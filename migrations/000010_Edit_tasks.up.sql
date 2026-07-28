ALTER TABLE task
DROP CONSTRAINT check_task_stauts;
ALTER TABLE task
add CONSTRAINT check_task_stauts CHECK(status IN('Todo','InProgress','Done','Overdue'));