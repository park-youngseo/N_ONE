package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/insajin/autopus-adk/pkg/acting/service"
)

func TestInMemoryActingService_GetDashboardSummary(t *testing.T) {
	svc := service.NewInMemoryActingService()

	t.Run("Returns correct summary for existing student", func(t *testing.T) {
		summary, err := svc.GetDashboardSummary("stu-1")
		assert.NoError(t, err)
		assert.NotNil(t, summary)
		assert.Equal(t, "홍길동", summary.Student.Name)
		assert.Equal(t, 100, summary.AttendanceRate) // 3/3 presents or lates
		assert.NotNil(t, summary.RecentScript)
		assert.Equal(t, "로미오와 줄리엣", summary.RecentScript.Title)
	})

	t.Run("Returns ErrNotFound for unknown student", func(t *testing.T) {
		summary, err := svc.GetDashboardSummary("unknown")
		assert.ErrorIs(t, err, service.ErrNotFound)
		assert.Nil(t, summary)
	})
}

func TestInMemoryActingService_UpdateSceneProgress(t *testing.T) {
	svc := service.NewInMemoryActingService()

	t.Run("Successfully updates scene progress and script progress changes", func(t *testing.T) {
		// Initial progress check
		script, _ := svc.GetScript("script-1")
		initialProgress := script.CalculateProgress()
		assert.Equal(t, 25, initialProgress) // 2/8 checks initially

		// Update
		req := service.UpdateProgressReq{
			ReadDone:      true,
			MemorizeDone:  true,
			AnalyzeDone:   true,
			RehearsalDone: true,
		}
		err := svc.UpdateSceneProgress("script-1", "scene-1", req)
		assert.NoError(t, err)

		// Re-fetch and check
		script, _ = svc.GetScript("script-1")
		assert.True(t, script.Scenes[0].AnalyzeDone)
		assert.Equal(t, 50, script.CalculateProgress()) // Now 4/8 checks
	})

	t.Run("Returns ErrNotFound for unknown script", func(t *testing.T) {
		err := svc.UpdateSceneProgress("unknown", "scene-1", service.UpdateProgressReq{})
		assert.ErrorIs(t, err, service.ErrNotFound)
	})
	
	t.Run("Returns ErrNotFound for unknown scene", func(t *testing.T) {
		err := svc.UpdateSceneProgress("script-1", "unknown", service.UpdateProgressReq{})
		assert.ErrorIs(t, err, service.ErrNotFound)
	})
}
