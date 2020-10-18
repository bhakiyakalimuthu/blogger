begin;

create table users (
    id UUID not null unique primary key,
    name varchar not null,
    email_id varchar,
    phone_number varchar
);

commit;