CREATE TABLE IF NOT EXISTS "projects" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" text NOT NULL,
  "post_queue" uuid[] NOT NULL,
  "idea_queue" uuid[] NOT NULL,
  "created_by" uuid NOT NULL,
  "created_at" timestamp DEFAULT (NOW()),
  "updated_at" timestamp DEFAULT (NOW())
);

CREATE TABLE IF NOT EXISTS team_roles (
  "id" serial PRIMARY KEY,
  "role" varchar(20) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS team_members (
    project_id uuid NOT NULL,
    user_id uuid NOT NULL,
    default_user boolean DEFAULT false,
    added_at timestamp DEFAULT NOW(),
    PRIMARY KEY (project_id, user_id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS team_members_roles (
    project_id uuid NOT NULL,
    team_role_id int NOT NULL,
    user_id uuid NOT NULL,
    PRIMARY KEY (project_id, team_role_id, user_id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (team_role_id) REFERENCES team_roles(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id, user_id) REFERENCES team_members(project_id, user_id) ON DELETE CASCADE
);



ALTER TABLE "projects" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON DELETE CASCADE;
INSERT INTO team_roles (role) VALUES ('member'), ('manager'), ('owner') ON CONFLICT DO NOTHING;