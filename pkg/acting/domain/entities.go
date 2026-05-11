package domain

// Student represents an acting school student
type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Attendance represents a student's daily attendance
type Attendance struct {
	ID        string `json:"id"`
	StudentID string `json:"student_id"`
	Date      string `json:"date"`   // Format: YYYY-MM-DD
	Status    string `json:"status"` // "present", "late", "absent"
}

// Script represents a practice script
type Script struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	TotalScenes int     `json:"total_scenes"`
	Scenes      []Scene `json:"scenes"`
}

// Scene represents a specific scene within a script and its practice progress
type Scene struct {
	ID              string `json:"id"`
	ScriptID        string `json:"script_id"`
	Number          int    `json:"number"`
	ReadDone        bool   `json:"read_done"`
	MemorizeDone    bool   `json:"memorize_done"`
	AnalyzeDone     bool   `json:"analyze_done"`
	RehearsalDone   bool   `json:"rehearsal_done"`
}

// CalculateProgress returns the completion percentage of the script (0-100)
func (s *Script) CalculateProgress() int {
	if s.TotalScenes == 0 || len(s.Scenes) == 0 {
		return 0
	}
	
	totalChecks := s.TotalScenes * 4 // 4 checks per scene
	completedChecks := 0

	for _, scene := range s.Scenes {
		if scene.ReadDone {
			completedChecks++
		}
		if scene.MemorizeDone {
			completedChecks++
		}
		if scene.AnalyzeDone {
			completedChecks++
		}
		if scene.RehearsalDone {
			completedChecks++
		}
	}

	return (completedChecks * 100) / totalChecks
}
