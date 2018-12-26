/* recify - db model */

-- users

CREATE TABLE recify_user (
  id SERIAL PRIMARY KEY,
  email VARCHAR(80) NOT NULL UNIQUE,
  password VARCHAR(80) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL
);

CREATE INDEX recify_user_email_idx ON recify_user (email);


-- groups

CREATE TABLE recify_group (
	id serial PRIMARY KEY,
	name VARCHAR(80) NOT NULL,
	description VARCHAR(250) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


-- users into a group

CREATE TABLE user_to_group (
	user_id SERIAL NOT NULL,
	group_id SERIAL NOT NULL,
	CONSTRAINT user_id_user_to_group_fk
	  FOREIGN KEY (user_id)
		REFERENCES recify_user(id)
		ON UPDATE CASCADE,
	CONSTRAINT group_id_user_to_group_fk
	  FOREIGN KEY (user_id)
		REFERENCES recify_group(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

CREATE INDEX user_id_user_to_group_idx ON user_to_group (user_id);
CREATE INDEX group_id_user_to_group_idx ON user_to_group (group_id);


-- recipes

CREATE TABLE recipe (
	id serial PRIMARY KEY,
	title VARCHAR(80) NOT NULL,
	description VARCHAR(1024)  NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


-- recipes contained into a group

CREATE TABLE group_to_recipe (
	group_id SERIAL NOT NULL,
	recipe_id SERIAL NOT NULL,
	CONSTRAINT group_id_group_to_recipe_fk
		FOREIGN KEY (group_id)
		REFERENCES recify_group(id)
		ON UPDATE CASCADE,
	CONSTRAINT recipe_id_group_to_recipe_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON UPDATE CASCADE
);

CREATE INDEX recipe_id_group_to_recipe_idx ON group_to_recipe (recipe_id);
CREATE INDEX group_id_group_to_recipe_idx ON group_to_recipe (group_id);


-- recipes liked by users

CREATE TABLE user_like_recipe (
	user_id SERIAL NOT NULL,
	recipe_id SERIAL NOT NULL,
	CONSTRAINT user_id_user_like_recipe_fk
		FOREIGN KEY (user_id)
		REFERENCES recify_user(id)
		ON UPDATE CASCADE,
	CONSTRAINT recipe_id_user_like_recipe_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

CREATE INDEX recipe_id_user_like_recipe_idx ON user_like_recipe (recipe_id);
CREATE INDEX user_id_user_like_recipe_idx ON user_like_recipe (user_id);


-- list of recipe categories

CREATE TABLE recipe_category (
	id SERIAL PRIMARY KEY,
	name VARCHAR(80) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


-- recipe's categories

CREATE TABLE recipe_to_category (
	recipe_id SERIAL NOT NULL,
	category_id SERIAL NOT NULL,
	CONSTRAINT recipe_id_recipe_to_category_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON UPDATE CASCADE,
	CONSTRAINT category_id_recipe_to_category_fk
		FOREIGN KEY (category_id)
		REFERENCES recipe_category(id)
		ON UPDATE CASCADE
);

CREATE INDEX recipe_id_recipe_to_category_idx ON recipe_to_category (recipe_id);
CREATE INDEX category_id_recipe_to_category_idx ON recipe_to_category (category_id);


-- ingredients

CREATE TABLE ingredient (
	id SERIAL PRIMARY KEY,
	name VARCHAR(80) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);

-- recipe's amount of ingredients

CREATE TABLE recipe_amount_ingredient (
	recipe_id SERIAL NOT NULL,
	ingredient_id SERIAL NOT NULL,
	amount REAL NOT NULL,
	measure CHAR(2) NOT NULL,
	CONSTRAINT recipe_id_recipe_amount_ingredient_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	CONSTRAINT ingredient_id_recipe_amount_ingreadient_fk
		FOREIGN KEY (ingredient_id)
		REFERENCES ingredient(id)
		ON UPDATE CASCADE
);

CREATE INDEX recipe_id_recipe_amount_ingredient_idx ON recipe_amount_ingredient (recipe_id);
CREATE INDEX ingredient_id_recipe_amount_ingredient_idx ON recipe_amount_ingredient (ingredient_id);


-- recipe's steps

CREATE TABLE recipe_step (
	recipe_id SERIAL NOT NULL,
	position SMALLSERIAL NOT NULL,
	description TEXT NOT NULL,
	PRIMARY KEY (recipe_id, position),
	CONSTRAINT recipe_id_recipe_step_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON DELETE CASCADE
);