create table TaskHistory(
    id bigserial primary key,
    task_id bigint not null,
    action varchar(20) not null,
    field_name varchar(20),
    old_value text,
    new_value text,
    changed_by bigint not null,
    created_at timestamp not null default now(),
    CONSTRAINT fk_task_history_task FOREIGN KEY (task_id) REFERENCES Task(id) ON DELETE CASCADE,
    CONSTRAINT fk_task_history_User FOREIGN KEY (changed_by) REFERENCES users(id) ON DELETE RESTRICT
)