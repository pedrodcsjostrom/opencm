package user

import "context"

//go:generate mockery --name=Repository --case=underscore --inpackage
type Repository interface {
	Save(ctx context.Context, usr *User) error
	FindByID(ctx context.Context, id string) (*UserResponse, error)
	FindByUsernameOrEmail(ctx context.Context, username, email string) (*UserResponse, error)
	FindByEmail(ctx context.Context, email string) (*FullUserResponse, error)
	// Additional methods as needed
}
