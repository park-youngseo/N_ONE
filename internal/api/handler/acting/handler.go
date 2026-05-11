package acting

import (
	"errors"
	"net/http"

	"github.com/insajin/autopus-adk/pkg/acting/service"
)

// Simplified HTTP router interface (e.g. mapping to echo or gin) for mock purposes
type SimpleRouter interface {
	GET(path string, handler func(w http.ResponseWriter, r *http.Request))
	POST(path string, handler func(w http.ResponseWriter, r *http.Request))
}

type Handler struct {
	svc service.ActingService
}

func NewHandler(svc service.ActingService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	// Mock: get studentID from query or context. Assuming "stu-1" for simplicity
	studentID := r.URL.Query().Get("student_id")
	if studentID == "" {
		studentID = "stu-1" 
	}

	summary, err := h.svc.GetDashboardSummary(studentID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// For a real implementation, write JSON response here.
	// For now, we assume success.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Mock JSON write
	// json.NewEncoder(w).Encode(summary)
	_ = summary
}
