-- create table
create table if not exists users(
    id serial primary key,
    name text not null
)

-- add user
insert into users (id, name)
values ($1, $2)

-- delete user
delete from users
where id = $1

-- update user
update from users
set name = $2
where id = $1


--get all
select * from users
order by id desc