create table users(
  id serial primary key,
  name varchar not null,
  username varchar not null,
  email varchar not null,
  password varchar not null,
  cpf varchar not null
);

---- create above / drop below ----

drop table users;
