create table Project(
    id bigserial primary key,
    workspace_id bigint not null,
    name varchar(100) not null,
    description varchar(200) not null,
    created_by bigint not null,
    created_at timestamp not null Default now(),
    updated_at timestamp not null Default now(),
    CONSTRAINT fk_Project_Workspace FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE,
    CONSTRAINT fk_Project_Creator FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);