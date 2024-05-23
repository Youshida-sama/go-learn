create table if not exists public.user_actions
(
    id integer not null,
    user_id integer,
    action_id integer not null,
    time timestamp without time zone,
    constraint user_actions_ID_uniq unique (id)
)