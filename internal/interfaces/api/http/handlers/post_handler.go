package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pedrodcsjostrom/opencm/internal/domain/post"
	e "github.com/pedrodcsjostrom/opencm/internal/utils/errors"
)

type PostHandler struct {
	Service post.Service
}

func NewPostHandler(service post.Service) *PostHandler {
	return &PostHandler{Service: service}
}

type createPostRequest struct {
	Title       string    `json:"title"`
	TextContent string    `json:"text_content"`
	ImageLinks  []string  `json:"image_links"`
	VideoLinks  []string  `json:"video_links"`
	IsIdea      bool      `json:"is_idea"`
	ScheduledAt time.Time `json:"scheduled_at"`
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with the given title, text content, image links, video links, is idea and scheduled at
// @Tags posts
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Param post body createPostRequest true "Post creation request"
// @Success 201 {object} post.Post
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 409 {object} errors.APIError "Post already exists"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/add [post]
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req createPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e.WriteHttpError(w, e.NewValidationError("Invalid request payload", nil))
		return
	}

	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteHttpError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}))
		return
	}

	if req.Title == "" {
		e.WriteHttpError(w, e.NewValidationError("Title is required", map[string]string{
			"title": "required",
		}))
		return
	}

	post, err := h.Service.CreatePost(
		r.Context(),
		projectID,
		req.Title,
		req.TextContent,
		req.ImageLinks,
		req.VideoLinks,
		req.IsIdea,
		req.ScheduledAt,
	)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

// GetPost godoc
// @Summary Get a post by id
// @Description Get a post by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} post.Post
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router  /posts/{project_id}/{post_id} [get]
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("post_id")
	if postID == "" {
		e.WriteHttpError(w, e.NewValidationError("Post id is required", map[string]string{
			"post_id": "required",
		}))
		return
	}

	post, err := h.Service.GetPost(r.Context(), postID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

// ListProjectPosts godoc
// @Summary List all posts of a project
// @Description List all posts of a project by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Success 200 {array} post.Post
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Project not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id} [get]
func (h *PostHandler) ListProjectPosts(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteHttpError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}))
		return
	}

	posts, err := h.Service.ListProjectPosts(r.Context(), projectID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

// ArchivePost godoc
// @Summary Archive a post
// @Description Archive a post by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 204 "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/{post_id}/archive [patch]
func (h *PostHandler) ArchivePost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("post_id")
	if postID == "" {
		e.WriteHttpError(w, e.NewValidationError("Post id is required", map[string]string{
			"post_id": "required",
		}))
		return
	}

	err := h.Service.ArchivePost(r.Context(), postID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by its id. We might or might not want to implement pagination and filtering. For the time being, we will keep it simple.
// @Tags posts
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 204 "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/{post_id} [delete]
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("post_id")
	if postID == "" {
		e.WriteHttpError(w, e.NewValidationError("Post id is required", map[string]string{
			"post_id": "required",
		}))
		return
	}

	err := h.Service.DeletePost(r.Context(), postID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// AddSocialMediaPublisherPlatform godoc
// @Summary Add a social media publisher platform to a post
// @Description Add a social media publisher platform to a post by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Param post_id path string true "Post ID"
// @Param platform_id path string true "Platform ID"
// @Success 204 "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/{post_id}/platforms/{platform_id} [post]
func (h *PostHandler) AddSocialMediaPublisherPlatform(w http.ResponseWriter, r *http.Request) {
	platformID := r.PathValue("platform_id")
	if platformID == "" {
		e.WriteHttpError(w, e.NewValidationError("Publisher ID is required", map[string]string{
			"platform_id": "required",
		}))
		return
	}

	postID := r.PathValue("post_id")
	if postID == "" {
		e.WriteHttpError(w, e.NewValidationError("Post id is required", map[string]string{
			"post_id": "required",
		}))
		return
	}

	projectID := r.PathValue("project_id")

	err := h.Service.AddSocialMediaPublisher(r.Context(), projectID, postID, platformID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

type schedulePostRequest struct {
	ScheduledAt time.Time `json:"scheduled_at"`
}

// SchedulePost godoc
// @Summary Schedule a post
// @Description Schedule a post by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Param project_id path string true "Project ID"
// @Param scheduled_at body schedulePostRequest true "Scheduled at"
// @Success 204 "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/{post_id}/schedule [patch]
func (h* PostHandler) SchedulePost(w http.ResponseWriter, r *http.Request) {
	var req schedulePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e.WriteHttpError(w, e.NewValidationError("Invalid request payload", nil))
		return
	}

	if req.ScheduledAt.IsZero() {
		e.WriteHttpError(w, e.NewValidationError("Scheduled at is required", map[string]string{
			"scheduled_at": "required",
		}))
		return
	}

	postID := r.PathValue("post_id")
	if postID == "" {
		e.WriteHttpError(w, e.NewValidationError("Post id is required", map[string]string{
			"post_id": "required",
		}))
		return
	}

	err := h.Service.SchedulePost(r.Context(), postID, req.ScheduledAt)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// AddPostToProjectQueue godoc
// @Summary Add a post to a project queue
// @Description Add a post to a project queue by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Param project_id path string true "Project ID"
// @Success 204 "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/{post_id}/enqueue [patch]
func (h * PostHandler) AddPostToProjectQueue(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("post_id")
	if postID == "" {
		e.WriteHttpError(w, e.NewValidationError("Post id is required", map[string]string{
			"post_id": "required",
		}))
		return
	}

	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteHttpError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}))
		return
	}

	err := h.Service.AddToProjectQueue(r.Context(), projectID, postID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// GetProjectQueuedPosts godoc
// @Summary Get all queued posts of a project
// @Description Get all queued posts of a project by its id
// @Tags posts
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Success 200 {array} post.Post
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Project not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/queue [get]
func (h * PostHandler) GetProjectQueuedPosts(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteHttpError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}))
		return
	}

	posts, err := h.Service.GetProjectQueuedPosts(r.Context(), projectID)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		e.WriteHttpError(w, e.NewInternalError("Failed to encode response"))
	}
}

type movePostRequest struct {
	CurrentIndex int `json:"current_index"`
	NewIndex     int `json:"new_index"`
}

// MovePostInQueue godoc
// @Summary Move a post in the project queue
// @Description Move a post in the project queue by its current and new index
// @Tags posts
// @Accept json
// @Produce json
// @Param project_id path string true "Project ID"
// @Param post_id path string true "Post ID"
// @Param move body movePostRequest true "Move post request"
// @Success 204 "No content"
// @Failure 400 {object} errors.APIError "Validation error"
// @Failure 401 {object} errors.APIError "Unauthorized"
// @Failure 410 {object} errors.APIError "Post not found"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Security ApiKeyAuth
// @Router /posts/{project_id}/queue/move [patch]
func (h *PostHandler) MovePostInQueue(w http.ResponseWriter, r *http.Request) {
	var req movePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		e.WriteHttpError(w, e.NewValidationError("Invalid request payload", nil))
		return
	}

	projectID := r.PathValue("project_id")
	if projectID == "" {
		e.WriteHttpError(w, e.NewValidationError("Project id is required", map[string]string{
			"project_id": "required",
		}))
		return
	}

	err := h.Service.MovePostInQueue(r.Context(), projectID, req.CurrentIndex, req.NewIndex)
	if err != nil {
		e.WriteBusinessError(w, err, mapPostErrorToAPIError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}