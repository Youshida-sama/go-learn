
create procedure create_user_action(i integer, n text, s text, a text, ts timestamp without time zone)
language sql
as $$
	insert into locks (id) 
		values (i)
		on conflict (id) 
		do update set last_lock_time = clock_timestamp();

	select * 
		from locks 
		where "id" = i 
		for update;

	--Проверка блокировок
	--select pg_sleep(1);

	insert into user_actions
	select
		i,
		coalesce(usrs.id, -1),
		acts.id,
		ts
	from (select 1)
	left join
		users as usrs
		on usrs.name = n
		and usrs.surname = s
	left join
		actions as acts
		on acts.name = a;
	
	update locks
		set last_lock_time = clock_timestamp()
		where id = i;
$$;