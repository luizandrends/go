alter table users drop column id;
alter table users add column id uuid primary key default gen_random_uuid();

---- create above / drop below ----

alter table users drop column id;
alter table users add column id serial primary key;
