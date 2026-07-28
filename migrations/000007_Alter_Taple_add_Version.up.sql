alter table task
add column version bigint not null default 1;
alter table project 
add column version bigint not null default 1;