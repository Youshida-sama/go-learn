create table if not exists public.user_actions
(
    id integer not null,
    name text,
    surname text,
    time timestamp without time zone,
    constraint user_actions_ID_uniq unique (id)
)