package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pedrodcsjostrom/opencm/internal/domain/project"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

type TableNames string

const (
	Projects         TableNames = "projects"
	TeamMembers      TableNames = "team_members"
	TeamMembersRoles TableNames = "team_members_roles"
	TeamRoles        TableNames = "team_roles"
)

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Save(ctx context.Context, p *project.Project) (*project.Project, error) {
	_, err := r.db.Exec(ctx, fmt.Sprintf(`
		INSERT INTO %s (id, name, description, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, Projects), p.ID, p.Name, p.Description, p.CreatedBy, time.Now(), time.Now())
	if err != nil {
		return &project.Project{}, err
	}

	return p, nil
}

func (r *ProjectRepository) ListByUserID(ctx context.Context, userID string) ([]*project.Project, error) {
	rows, err := r.db.Query(ctx, fmt.Sprintf(`
		SELECT p.id, p.name, p.description, p.created_by, p.created_at, p.updated_at
		FROM %s p
		INNER JOIN %s tm ON p.id = tm.project_id
		WHERE tm.user_id = $1
	`, Projects, TeamMembers), userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*project.Project
	for rows.Next() {
		p := &project.Project{}
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectRepository) AssignProjectOwner(ctx context.Context, projectID, userID string) error {
	// Begin a transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	// Insert into team_members
	insertMember := fmt.Sprintf(`
        INSERT INTO %s (project_id, user_id, added_at)
        VALUES ($1, $2, $3)
    `, TeamMembers)

	_, err = tx.Exec(ctx, insertMember, projectID, userID, time.Now())
	if err != nil {
		return err
	}

	// Prepare the INSERT statement for team_members_roles
	insertRoleStmt := fmt.Sprintf(`
        INSERT INTO %s (project_id,team_role_id, user_id)
        VALUES ($1, $2, $3)
    `, TeamMembersRoles)

	// Batch insert roleIDs
	roleIDs := []project.TeamRoleIDs{project.OwnerRoleID, project.ManagerRoleID, project.MemberRoleID}
	for _, roleID := range roleIDs {
		_, err = tx.Exec(ctx, insertRoleStmt, projectID, roleID, userID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ProjectRepository) GetUserRoles(ctx context.Context, userID, projectID string) ([]string, error) {
	rows, err := r.db.Query(ctx, fmt.Sprintf(`
		SELECT tr.role
		FROM %s tmr
		INNER JOIN %s tr ON tmr.team_role_id = tr.id
		WHERE tmr.user_id = $1 AND tmr.project_id = $2
	`, TeamMembersRoles, TeamRoles), userID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		err = rows.Scan(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *ProjectRepository) GetProject(ctx context.Context, projectID string) (*project.Project, error) {
	row := r.db.QueryRow(ctx, fmt.Sprintf(`
		SELECT id, name, description, created_by, created_at, updated_at
		FROM %s
		WHERE id = $1
	`, Projects), projectID)

	p := &project.Project{}
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProjectRepository) GetProjectUsers(ctx context.Context, projectID string) ([]*project.TeamMember, error) {
	rows, err := r.db.Query(ctx, fmt.Sprintf(`
		SELECT tm.user_id, tm.added_at
		FROM %s tm
		WHERE tm.project_id = $1
	`, TeamMembers), projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*project.TeamMember
	for rows.Next() {
		tm := &project.TeamMember{}
		err = rows.Scan(&tm.ID, &tm.AddedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, tm)
	}

	return users, nil
}