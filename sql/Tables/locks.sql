create table if not exists public.locks
(
    id integer not null,
    last_lock_time timestamp without time zone,
    constraint locks_ID_uniq unique (id)
)