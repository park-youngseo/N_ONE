package service

import (
	"errors"
	"github.com/insajin/autopus-adk/pkg/acting/domain"
)

var ErrNotFound = errors.New("not found")

// ActingService provides business logic for the acting school dashboard
type ActingService interface {
	GetDashboardSummary(studentID string) (*DashboardSummary, error)
	GetScript(scriptID string) (*domain.Script, error)
	UpdateSceneProgress(scriptID string, sceneID string, progress UpdateProgressReq) error
}

type DashboardSummary struct {
	Student          domain.Student      `json:"student"`
	AttendanceRate   int                 `json:"attendance_rate"`
	RecentScript     *domain.Script      `json:"recent_script,omitempty"`
}

type UpdateProgressReq struct {
	ReadDone      bool `json:"read_done"`
	MemorizeDone  bool `json:"memorize_done"`
	AnalyzeDone   bool `json:"analyze_done"`
	RehearsalDone bool `json:"rehearsal_done"`
}

// InMemoryActingService is a mock implementation for early development
type InMemoryActingService struct {
	students    map[string]domain.Student
	attendances map[string][]domain.Attendance // studentID -> attendances
	scripts     map[string]domain.Script
}

func NewInMemoryActingService() *InMemoryActingService {
	return &InMemoryActingService{
		students: map[string]domain.Student{
			"stu-1": {ID: "stu-1", Name: "홍길동"},
		},
		attendances: map[string][]domain.Attendance{
			"stu-1": {
				{ID: "att-1", StudentID: "stu-1", Date: "2026-05-01", Status: "present"},
				{ID: "att-2", StudentID: "stu-1", Date: "2026-05-08", Status: "present"},
				{ID: "att-3", StudentID: "stu-1", Date: "2026-05-15", Status: "late"},
			},
		},
		scripts: map[string]domain.Script{
			"script-1": {
				ID:          "script-1",
				Title:       "로미오와 줄리엣",
				TotalScenes: 2,
				Scenes: []domain.Scene{
					{ID: "scene-1", ScriptID: "script-1", Number: 1, ReadDone: true, MemorizeDone: true, AnalyzeDone: false, RehearsalDone: false},
					{ID: "scene-2", ScriptID: "script-1", Number: 2, ReadDone: false, MemorizeDone: false, AnalyzeDone: false, RehearsalDone: false},
				},
			},
		},
	}
}

func (s *InMemoryActingService) GetDashboardSummary(studentID string) (*DashboardSummary, error) {
	student, ok := s.students[studentID]
	if !ok {
		return nil, ErrNotFound
	}

	atts := s.attendances[studentID]
	total := len(atts)
	presentCount := 0
	for _, a := range atts {
		if a.Status == "present" || a.Status == "late" {
			presentCount++
		}
	}

	rate := 0
	if total > 0 {
		rate = (presentCount * 100) / total
	}

	// Just return the first script as 'recent' for the mock
	var recent *domain.Script
	for _, script := range s.scripts {
		scriptCopy := script
		recent = &scriptCopy
		break
	}

	return &DashboardSummary{
		Student:        student,
		AttendanceRate: rate,
		RecentScript:   recent,
	}, nil
}

func (s *InMemoryActingService) GetScript(scriptID string) (*domain.Script, error) {
	script, ok := s.scripts[scriptID]
	if !ok {
		return nil, ErrNotFound
	}
	return &script, nil
}

func (s *InMemoryActingService) UpdateSceneProgress(scriptID string, sceneID string, req UpdateProgressReq) error {
	script, ok := s.scripts[scriptID]
	if !ok {
		return ErrNotFound
	}

	sceneFound := false
	for i, scene := range script.Scenes {
		if scene.ID == sceneID {
			script.Scenes[i].ReadDone = req.ReadDone
			script.Scenes[i].MemorizeDone = req.MemorizeDone
			script.Scenes[i].AnalyzeDone = req.AnalyzeDone
			script.Scenes[i].RehearsalDone = req.RehearsalDone
			sceneFound = true
			break
		}
	}

	if !sceneFound {
		return ErrNotFound
	}

	s.scripts[scriptID] = script // Save back to map
	return nil
}
