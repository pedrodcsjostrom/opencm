package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrodcsjostrom/opencm/internal/domain/project"
	e "github.com/pedrodcsjostrom/opencm/internal/utils/errors"
)

type ProjectHandler struct {
	Service project.Service
}

func NewProjectHandler(service project.Service) *ProjectHandler {
	return &ProjectHandler{Service: service}
}

type createProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project with the given name and description
// @Tags projects
// @Accept json
// @Produce json
// @Param project body createProjectRequest true "Project creation request"
// @Success 201 {object} project.Project
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 409 {object} errors.APIError "Project already exists"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /projects [post]
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req createProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e.WriteBusinessError(w, e.NewValidationError("Invalid request payload", nil), nil)
		return
	}

	if req.Name == "" {
		e.WriteBusinessError(w, e.NewValidationError("Name is required", map[string]string{
			"name": "required",
		}), nil)
		return
	}

	p, err := h.Service.CreateProject(r.Context(), req.Name, req.Description)
	if err != nil {
		e.WriteBusinessError(w, err, mapProjectErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

// ListProjects godoc
// @Summary List all projects
// @Description List all projects that the user is a member of
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {array} project.Project
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /projects [get]
func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projects, err := h.Service.ListProjects(ctx)
	if err != nil {
		e.WriteBusinessError(w, err, mapProjectErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

// GetProject godoc
// @Summary Get a project
// @Description Get a project by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Success 200 {object} project.Project
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Project not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /projects/{project_id} [get]
func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteBusinessError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}), nil)
		return
	}

	p, err := h.Service.GetProject(ctx, projectID)
	if err != nil {
		e.WriteBusinessError(w, err, mapProjectErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

type addUserRequest struct {
	Email string `json:"email"`
}

// AddUserToProject godoc
// @Summary Add a user to a project
// @Description Add a user to a project by their email
// @Tags projects
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Param user_id path string true "User ID"
// @Success 204 {string} string "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Project not found"
// @Failure 409 {object} errors.APIError "User already exists"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /projects/{project_id}/add [post]
func (h *ProjectHandler) AddUserToProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteBusinessError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}), nil)
		return
	}

	var req addUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e.WriteBusinessError(w, e.NewValidationError("Invalid request payload", nil), nil)
		return
	}
	if req.Email == "" {
		e.WriteBusinessError(w, e.NewValidationError("Email is required", map[string]string{
			"email": "required",
		}), nil)
		return
	}

	err := h.Service.AddUserToProject(ctx, projectID, req.Email)
	if err != nil {
		e.WriteBusinessError(w, err, mapProjectErrorToAPIError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func mapProjectErrorToAPIError(err error) *e.APIError {
	switch {
	case e.MatchError(
		err,
		project.ErrProjectExists,
	):
		return &e.APIError{
			Status:  http.StatusConflict,
			Code:    e.ErrCodeConflict,
			Message: err.Error(),
		}
	case e.MatchError(
		err,
		project.ErrProjectNotFound,
	):
		return &e.APIError{
			Status:  http.StatusGone,
			Code:    e.ErrCodeNotFound,
			Message: err.Error(),
		}
	case e.MatchError(
		err,
		project.ErrUserAlreadyInProject,
	):
		return &e.APIError{
			Status:  http.StatusConflict,
			Code:    e.ErrCodeConflict,
			Message: err.Error(),
		}
	default:
		return &e.APIError{
			Status:  http.StatusInternalServerError,
			Code:    e.ErrCodeInternal,
			Message: "Internal server error",
		}
	}
}
