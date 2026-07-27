create table Task(
    id bigserial primary key,
    project_id bigint not null,
    title varchar(50) not null,
    description varchar(200) not null,
    status varchar(15) not null default 'Todo',
    priority varchar(10) not null default 'Medium',
    parent_task_id bigint,
    assignee_id bigint not null,
    created_by bigint not null,
    due_date timestamp,
    created_at timestamp not null Default now(),
    updated_at timestamp not null Default now(),
    CONSTRAINT fk_task_project FOREIGN KEY (project_id) REFERENCES Project(id) ON DELETE CASCADE,
    CONSTRAINT fk_task_CreateUser FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_task_assigned FOREIGN KEY (assignee_id) REFERENCES users(id) ON DELETE set null,
    CONSTRAINT check_task_stauts CHECK(status IN('Todo','InProgress','Done')),
    CONSTRAINT check_task_priority CHECK(priority IN('Low','Medium','Hihg'))
)