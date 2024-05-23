create procedure create_user_action(i integer, n text, s text, ts timestamp without time zone)
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

	insert into user_actions (id, name, surname, time)
		values (i, n, s, ts);

	
	update locks
		set last_lock_time = clock_timestamp()
		where id = i;
$$;