package domain_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/insajin/autopus-adk/pkg/acting/domain"
)

func TestScript_CalculateProgress(t *testing.T) {
	t.Run("Empty script returns 0 progress", func(t *testing.T) {
		s := domain.Script{TotalScenes: 0}
		assert.Equal(t, 0, s.CalculateProgress())
	})

	t.Run("Script with no scenes returns 0 progress", func(t *testing.T) {
		s := domain.Script{TotalScenes: 2, Scenes: []domain.Scene{}}
		assert.Equal(t, 0, s.CalculateProgress())
	})

	t.Run("Calculates correct progress for partial completion", func(t *testing.T) {
		s := domain.Script{
			TotalScenes: 1,
			Scenes: []domain.Scene{
				{ReadDone: true, MemorizeDone: true, AnalyzeDone: false, RehearsalDone: false}, // 2/4 checks done
			},
		}
		assert.Equal(t, 50, s.CalculateProgress())
	})

	t.Run("Calculates correct progress for full completion across multiple scenes", func(t *testing.T) {
		s := domain.Script{
			TotalScenes: 2,
			Scenes: []domain.Scene{
				{ReadDone: true, MemorizeDone: true, AnalyzeDone: true, RehearsalDone: true}, // 4/4
				{ReadDone: true, MemorizeDone: true, AnalyzeDone: true, RehearsalDone: true}, // 4/4
			},
		}
		assert.Equal(t, 100, s.CalculateProgress()) // 8/8 checks
	})
	
	t.Run("Calculates correct progress for complex partial completion", func(t *testing.T) {
		s := domain.Script{
			TotalScenes: 2,
			Scenes: []domain.Scene{
				{ReadDone: true, MemorizeDone: true, AnalyzeDone: true, RehearsalDone: true}, // 4/4
				{ReadDone: true, MemorizeDone: true, AnalyzeDone: false, RehearsalDone: false}, // 2/4
			},
		}
		assert.Equal(t, 75, s.CalculateProgress()) // 6/8 checks = 75%
	})
}
