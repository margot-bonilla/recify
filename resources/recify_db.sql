/* recify - db model */


/* remove all tables */
DROP TABLE IF EXISTS "user_group";
DROP TABLE IF EXISTS "group_recipe";
DROP TABLE IF EXISTS "user_likes_recipe";
DROP TABLE IF EXISTS "recipe_category";
DROP TABLE IF EXISTS "recipe_ingredient";
DROP TABLE IF EXISTS "recipe_category";
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "group";
DROP TABLE IF EXISTS "ingredient";
DROP TABLE IF EXISTS "step";
DROP TABLE IF EXISTS "recipe";


-- users

CREATE TABLE "user" (
  id SERIAL PRIMARY KEY,
  email VARCHAR(80) NOT NULL UNIQUE,
	token varchar(100) DEFAULT NULL,
  password VARCHAR(80) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL
);

CREATE INDEX user_email_idx ON "user" (email);


-- groups

CREATE TABLE "group" (
	id serial PRIMARY KEY,
	name VARCHAR(80) NOT NULL,
	description VARCHAR(250) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


-- users into a group

CREATE TABLE user_group (
	user_id SERIAL NOT NULL,
	group_id SERIAL NOT NULL,
	CONSTRAINT user_id_user_group_fk
	  FOREIGN KEY (user_id)
		REFERENCES "user"(id)
		ON UPDATE CASCADE
);
ALTER TABLE user_group
	ADD CONSTRAINT group_id_user_group_fk
	  FOREIGN KEY (user_id)
		REFERENCES "group"(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
;

CREATE INDEX user_id_user_group_idx ON user_group (user_id);
CREATE INDEX group_id_user_group_idx ON user_group (group_id);


-- recipes

CREATE TABLE recipe (
	id serial PRIMARY KEY,
	title VARCHAR(80) NOT NULL,
	description VARCHAR(1024)  NULL,
	rating DECIMAL(3,2) DEFAULT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


-- recipes contained into a group

CREATE TABLE group_recipe (
	group_id SERIAL NOT NULL,
	recipe_id SERIAL NOT NULL,
	CONSTRAINT group_id_group_recipe_fk
		FOREIGN KEY (group_id)
		REFERENCES "group"(id)
		ON UPDATE CASCADE
);
ALTER TABLE group_recipe
	ADD CONSTRAINT recipe_id_group_recipe_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON UPDATE CASCADE
;

CREATE INDEX recipe_id_group_recipe_idx ON group_recipe (recipe_id);
CREATE INDEX group_id_group_recipe_idx ON group_recipe (group_id);


-- recipes liked by users

CREATE TABLE user_likes_recipe (
	user_id SERIAL NOT NULL,
	recipe_id SERIAL NOT NULL,
	CONSTRAINT user_id_user_likes_recipe_fk
		FOREIGN KEY (user_id)
		REFERENCES "user"(id)
		ON UPDATE CASCADE
);
ALTER TABLE user_likes_recipe
	ADD CONSTRAINT recipe_id_user_likes_recipe_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
;

CREATE INDEX recipe_id_user_likes_recipe_idx ON user_likes_recipe (recipe_id);
CREATE INDEX user_id_user_likes_recipe_idx ON user_likes_recipe (user_id);


-- list of recipe categories

CREATE TABLE "category" (
	id SERIAL PRIMARY KEY,
	name VARCHAR(80) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL
);


-- recipe's categories

CREATE TABLE recipe_category (
	recipe_id SERIAL NOT NULL,
	category_id SERIAL NOT NULL,
	CONSTRAINT recipe_id_recipe_category_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON UPDATE CASCADE
);
ALTER TABLE recipe_category
	ADD CONSTRAINT category_id_recipe_category_fk
		FOREIGN KEY (category_id)
		REFERENCES "category"(id)
		ON UPDATE CASCADE
;

CREATE INDEX recipe_id_recipe_category_idx ON recipe_category (recipe_id);
CREATE INDEX category_id_recipe_category_idx ON recipe_category (category_id);

-- recipe's amount of ingredients

CREATE TABLE "recipe_ingredient" (
	recipe_id SERIAL NOT NULL,
	ingredient_id SERIAL NOT NULL,
	name VARCHAR(100) NOT NULL,
	amount REAL NOT NULL,
	measure VARCHAR(20) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT recipe_id_recipe_amount_ingredient_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

CREATE INDEX recipe_id_recipe_ingredient_idx ON recipe_ingredient (recipe_id);
CREATE INDEX ingredient_id_recipe_ingredient_idx ON recipe_ingredient (ingredient_id);


-- recipe's steps

CREATE TABLE "step" (
	recipe_id SERIAL NOT NULL,
	position SMALLSERIAL NOT NULL,
	description TEXT NOT NULL,
	PRIMARY KEY (recipe_id, position),
	CONSTRAINT recipe_id_step_fk
		FOREIGN KEY (recipe_id)
		REFERENCES recipe(id)
		ON DELETE CASCADE
);