package acting_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/insajin/autopus-adk/internal/api/handler/acting"
	"github.com/insajin/autopus-adk/pkg/acting/service"
)

func TestHandler_GetDashboard(t *testing.T) {
	svc := service.NewInMemoryActingService()
	handler := acting.NewHandler(svc)

	t.Run("Returns 200 OK for valid student", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/dashboard?student_id=stu-1", nil)
		rr := httptest.NewRecorder()

		handler.GetDashboard(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	})

	t.Run("Returns 404 Not Found for invalid student", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/dashboard?student_id=unknown", nil)
		rr := httptest.NewRecorder()

		handler.GetDashboard(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
