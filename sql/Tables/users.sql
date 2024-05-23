create table if not exists public.users
(
    id integer not null,
    name text not null,
    surname text not null,
    constraint users_ID_uniq unique (id)
)