CREATE TABLE recify_user (
  id SERIAL PRIMARY KEY,
  email VARCHAR(80) NOT NULL UNIQUE,
  password VARCHAR(80) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL
);

CREATE INDEX recify_user_email_idx ON recify_user (email)

CREATE TABLE recify_group (
	id serial PRIMARY KEY,
	name VARCHAR(80) NOT NULL,
	description VARCHAR(250) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


CREATE TABLE user2group (
	user_id SERIAL NOT NULL,
	group_id SERIAL NOT NULL,
	CONSTRAINT user_user2group_fk
	  FOREIGN KEY (user_id)
		REFERENCES recify_user(id)
		ON UPDATE CASCADE,
	CONSTRAINT group_user2group_fk
	  FOREIGN KEY (user_id)
		REFERENCES recify_group(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

CREATE INDEX user_id_user2group_idx ON user2group (user_id)
CREATE INDEX group_id_user2group_idx ON user2group (group_id)

/* example to test group cascade deletion:
INSERT INTO recify_user VALUES (1, 'raul@email.com', '1234', current_timestamp, NULL)
INSERT INTO recify_group VALUES (1, 'family','family recipes', current_timestamp, null)
INSERT INTO user2group VALUES (1,1)
DELETE FROM recify_group WHERE id = 1
*/

CREATE TABLE recipe (
	id serial PRIMARY KEY,
	title VARCHAR(80) NOT NULL,
	description VARCHAR(1024)  NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
)


CREATE TABLE group2recipe (
	group_id SERIAL NOT NULL,
	recipe_id SERIAL NOT NULL,
	CONSTRAINT group_group2recipe_fk
		FOREIGN KEY (group_id)
		REFERENCES recify_group(id)
		ON UPDATE CASCADE,
	CONSTRAINT recipe_group2recipe_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON UPDATE CASCADE
)

CREATE INDEX recipe_id_group2recipe_idx ON group2recipe (recipe_id)
CREATE INDEX group_id_group2recipe_idx ON group2recipe (group_id)


/* WIP */