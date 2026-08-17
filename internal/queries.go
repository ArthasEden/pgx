package internal

var (
	QAdd = `insert into users (id, name)
values ($1, $2)`

	QDel = `delete from users
where id = $1`

	QUp = `update from users
where id = $1
set name = $2`

	QCreateTable = `create table if not exists users(
id serial primary key,
name text not null
)`

	QAddMany = `insert into users(id, name)
values
	(1, 'Bob'),
	(2, 'Jack'),
	(3, 'Alice')`

	QGetAll = `select * from users
				order by id desc`
)
