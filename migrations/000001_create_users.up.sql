create table users(
    id bigserial primary key,
    name varchar(100) not null,
    email varchar(100) unique,
    Hashpassword text not null,
    created_at timestamp not null Default now(),
    updated_at timestamp not null Default now()
);