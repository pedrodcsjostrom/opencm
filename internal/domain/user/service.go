package user

import (
	"context"

	"github.com/pedrodcsjostrom/opencm/internal/infrastructure/session"
)

//go:generate mockery --name=Service --case=underscore --inpackage
type Service interface {
	CreateUser(ctx context.Context, username, password, email string) (*UserResponse, error)
	GetUser(ctx context.Context, id string) (*UserResponse, error)
	Login(ctx context.Context, email, password string) (*session.Session, error)
	GetAllAppRoles(ctx context.Context) (*[]AppRole, error)
	GetUserAppRoles(ctx context.Context, userID string) ([]string, error)
	AssignAppRoleToUser(ctx context.Context, userID, roleID string) error
	RemoveAppRoleFromUser(ctx context.Context, userID, roleID string) error
	// Additional methods as needed
}

// Create a concrete implementation
type service struct {
	repo     Repository
	password PasswordHasher
	session  session.Manager
}

// Update the constructor to return the interface
func NewService(repo Repository, session session.Manager, passwordHasher PasswordHasher) Service {
	return &service{
		repo:     repo,
		password: passwordHasher,
		session:  session,
	}
}

func (s *service) CreateUser(ctx context.Context, username, password, email string) (*UserResponse, error) {
	existingUser, err := s.repo.FindByUsernameOrEmail(ctx, username, email)
	if err != nil {
		return &UserResponse{}, err
	}
	if existingUser != nil {
		return &UserResponse{}, ErrExistingUser
	}

	hashedPassword, salt, err := s.password.Hash(password)
	if err != nil {
		return &UserResponse{}, err
	}

	user, err := NewUser(username, hashedPassword, salt, email)
	if err != nil {
		return &UserResponse{}, err
	}
	uResponse, err := s.repo.Save(ctx, user)
	if err != nil {
		return &UserResponse{}, err
	}
	err = s.repo.AssignDefaultRoleToUser(ctx, user.ID)
	if err != nil {
		return &UserResponse{}, err
	}
	return uResponse, nil
}

func (s *service) GetUser(ctx context.Context, id string) (*UserResponse, error) {
	userResponse, err := s.repo.FindByIDWithRoles(ctx, id)
	if err != nil {
		return nil, err
	}
	if userResponse.ID == "" {
		return nil, ErrUserNotFound
	}
	return userResponse, nil
}

func (s *service) Login(ctx context.Context, email, password string) (*session.Session, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return &session.Session{}, err
	}
	if user == nil {
		return &session.Session{}, ErrUserNotFound
	}
	if !s.password.Validate(password, user.HashedPasword, user.Salt) {
		return &session.Session{}, ErrInvalidPassword
	}
	return s.session.CreateSession(ctx, user.ID)
}

func (s *service) Logout(ctx context.Context, sessionID string) error {
	// TODO: Implement the logout logic
	return nil
}

func (s *service) UpdatePassword(ctx context.Context, userID, password string) error {
	// TODO: Implement the update password logic
	return nil
}

func (s *service) UpdateEmail(ctx context.Context, userID, email string) error {
	// TODO: Implement the update email logic
	return nil
}

func (s *service) GetAllAppRoles(ctx context.Context) (*[]AppRole, error) {
	return s.repo.GetRoles(ctx)
}

func (s *service) GetUserAppRoles(ctx context.Context, userID string) ([]string, error) {
	return s.repo.GetUserRoles(ctx, userID)
}

func (s *service) AssignAppRoleToUser(ctx context.Context, userID, roleID string) error {
	return s.repo.AssignRoleToUser(ctx, userID, roleID)
}

func (s *service) RemoveAppRoleFromUser(ctx context.Context, userID, roleID string) error {
	return s.repo.RemoveRoleFromUser(ctx, userID, roleID)
}
