package postgres

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pedrodcsjostrom/opencm/internal/domain/post"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Save(p *post.Post) error {
	// Implementation...
	return nil
}

func (r *PostRepository) FindByID(id post.PostID) (*post.Post, error) {
	// Implementation...
	return nil, nil
}

func (r *PostRepository) FindByTeamID(teamID uuid.UUID) ([]*post.Post, error) {
	// Implementation...
	return nil, nil
}
