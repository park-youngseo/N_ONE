package setup

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func scanDirectoryTree(dir string, depth int) []DirEntry {
	if depth >= maxDepth {
		return nil
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var tree []DirEntry
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		// Skip hidden dirs and common non-essential dirs
		if strings.HasPrefix(name, ".") || isIgnoredDir(name) {
			continue
		}

		rel, _ := filepath.Rel(filepath.Dir(dir), filepath.Join(dir, name))
		if depth == 0 {
			rel = name
		}

		entry := DirEntry{
			Name:        name,
			Path:        rel,
			Description: inferDirDescription(name),
			Children:    scanDirectoryTree(filepath.Join(dir, name), depth+1),
		}
		tree = append(tree, entry)
	}
	return tree
}

// parseMakefileTargets extracts all target names and their first recipe line from a Makefile.
func parseMakefileTargets(path string) map[string]string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	targets := make(map[string]string)
	scanner := bufio.NewScanner(f)
	var currentTarget string
	var currentRecipe []string

	flushTarget := func() {
		if currentTarget != "" && len(currentRecipe) > 0 {
			recipe := strings.Join(currentRecipe, " && ")
			recipe = strings.TrimPrefix(recipe, "@")
			targets[currentTarget] = recipe
		} else if currentTarget != "" {
			targets[currentTarget] = "make " + currentTarget
		}
		currentTarget = ""
		currentRecipe = nil
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "\t") {
			recipe := strings.TrimSpace(line)
			if currentTarget != "" && recipe != "" && !strings.HasPrefix(recipe, "#") {
				if !strings.HasPrefix(recipe, "@echo") && !strings.HasPrefix(recipe, "echo ") {
					currentRecipe = append(currentRecipe, strings.TrimPrefix(recipe, "@"))
				}
			}
			continue
		}
		flushTarget()
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") || (strings.Contains(trimmed, "=") && !strings.Contains(trimmed, ":")) {
			continue
		}
		if strings.Contains(trimmed, ":") {
			parts := strings.SplitN(trimmed, ":", 2)
			name := strings.TrimSpace(parts[0])
			if name != "" && !strings.HasPrefix(name, ".") && !strings.Contains(name, " ") && !strings.Contains(name, "/") && !strings.Contains(name, "$") {
				currentTarget = name
			}
		}
	}
	flushTarget()
	return targets
}

func parsePyprojectScripts(path string) map[string]string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	cmds := make(map[string]string)
	scanner := bufio.NewScanner(f)
	inScripts := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "[tool.poetry.scripts]" || line == "[project.scripts]" {
			inScripts = true
			continue
		}
		if strings.HasPrefix(line, "[") {
			inScripts = false
			continue
		}
		if inScripts && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			name := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
			cmds[name] = value
		}
	}
	return cmds
}

func findDirsWithSuffix(dir, suffix string) []string {
	seen := make(map[string]bool)
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), suffix) {
			rel, _ := filepath.Rel(dir, filepath.Dir(path))
			if !seen[rel] {
				seen[rel] = true
			}
		}
		return nil
	})
	var dirs []string
	for d := range seen {
		dirs = append(dirs, d)
	}
	return dirs
}

func inferDirDescription(name string) string {
	descriptions := map[string]string{
		"cmd": "CLI entry points", "pkg": "Public reusable libraries", "internal": "Private implementation packages",
		"api": "API definitions and handlers", "web": "Web server and routes", "src": "Source code",
		"lib": "Library code", "test": "Test files", "docs": "Documentation", "scripts": "Scripts",
	}
	return descriptions[name]
}

func isIgnoredDir(name string) bool {
	ignored := map[string]bool{
		"node_modules": true, "vendor": true, ".git": true, "dist": true, "build": true, "target": true,
	}
	return ignored[name]
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func mergeMaps(a, b map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range a { result[k] = v }
	for k, v := range b { result[k] = v }
	return result
}
