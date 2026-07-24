create table workspaces(
    id bigserial primary key,
    name varchar(100) not null,
    description text,
    owner_id bigint not null,
    created_at timestamp not null Default now(),
    updated_at timestamp not null Default now(),
    CONSTRAINT fk_workspace_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);