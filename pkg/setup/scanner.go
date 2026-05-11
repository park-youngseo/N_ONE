package setup

import (
	"os"
	"path/filepath"
)

const maxDepth = 3

// Scan analyzes a project directory and returns ProjectInfo.
func Scan(projectDir string) (*ProjectInfo, error) {
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return nil, err
	}

	info := &ProjectInfo{
		Name:    filepath.Base(absDir),
		RootDir: absDir,
	}

	info.Languages = detectLanguages(absDir)
	info.Frameworks = detectFrameworks(absDir)
	info.BuildFiles = detectBuildFiles(absDir)
	info.EntryPoints = detectEntryPoints(absDir, info.Languages)
	info.TestConfig = detectTestConfig(absDir, info.Languages, info.BuildFiles)
	info.Structure = scanDirectoryTree(absDir, 0)
	info.Conventions = AnalyzeConventions(absDir, info.Languages)
	info.Workspaces = DetectWorkspaces(absDir)

	return info, nil
}

func detectEntryPoints(dir string, langs []Language) []EntryPoint {
	var eps []EntryPoint

	for _, lang := range langs {
		switch lang.Name {
		case "Go":
			// cmd/ directories
			cmdDir := filepath.Join(dir, "cmd")
			if entries, err := os.ReadDir(cmdDir); err == nil {
				for _, e := range entries {
					if e.IsDir() {
						mainFile := filepath.Join("cmd", e.Name(), "main.go")
						if fileExists(filepath.Join(dir, mainFile)) {
							eps = append(eps, EntryPoint{
								Path:        mainFile,
								Description: e.Name() + " CLI entry point",
							})
						}
					}
				}
			}
			// root main.go
			if fileExists(filepath.Join(dir, "main.go")) {
				eps = append(eps, EntryPoint{
					Path:        "main.go",
					Description: "Main entry point",
				})
			}
		case "TypeScript", "JavaScript":
			for _, f := range []string{"src/index.ts", "src/index.js", "src/main.ts", "src/main.js", "index.ts", "index.js"} {
				if fileExists(filepath.Join(dir, f)) {
					eps = append(eps, EntryPoint{
						Path:        f,
						Description: "Application entry point",
					})
				}
			}
		case "Python":
			for _, f := range []string{"main.py", "app.py", "src/main.py", "manage.py"} {
				if fileExists(filepath.Join(dir, f)) {
					eps = append(eps, EntryPoint{
						Path:        f,
						Description: "Application entry point",
					})
				}
			}
		}
	}
	return eps
}
