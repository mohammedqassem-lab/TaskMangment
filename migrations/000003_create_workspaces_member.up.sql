create table Workspaces_member(
    id bigserial primary key,
    workspace_id bigint not null,
    user_id bigint not null,
    role varchar(50) not null,
    created_at timestamp not null Default now(),
    updated_at timestamp not null Default now(),
    CONSTRAINT fk_workspace_owner FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_owner FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);