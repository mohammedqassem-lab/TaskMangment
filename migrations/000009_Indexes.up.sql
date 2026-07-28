CREATE UNIQUE INDEX idx_users_email
ON users (email);
CREATE INDEX idx_workspace_id
on workspaces(id);
CREATE INDEX idx_workspace_id_version
on workspaces(id,version);
CREATE INDEX idx_workspaces_member_id_user_id
on workspaces_member(id,user_id);
CREATE INDEX idx_workspaces_member_workspace_id_user_id
on workspaces_member(workspace_id,user_id);
CREATE INDEX idx_task_project_id_priority
on task(project_id,priority);
CREATE INDEX idx_task_project_id_status
on task(project_id,status);
CREATE INDEX idx_task_status_priority
on task(status,priority);
CREATE INDEX idx_task_id_version
on task(id,version);
CREATE INDEX idx_Project_id
on Project(id);
CREATE INDEX idx_Project_id_version
on Project(id,version);