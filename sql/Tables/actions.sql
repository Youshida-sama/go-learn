create table if not exists public.actions
(
    id integer not null,
    name text not null,
    constraint actions_ID_uniq unique (id)
)