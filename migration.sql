CREATE TABLE IF NOT EXISTS table_users (
   user_id serial PRIMARY KEY,
   username VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT exists table_notes (
	user_note_id serial primary key,
	user_id INT references table_users(user_id),
  	note VARCHAR(50) UNIQUE NOT null,
   	note_day TIMESTAMP NOT NULL
);

insert into table_users(username)values('Ella');
