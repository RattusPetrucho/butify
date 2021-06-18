
-- +migrate Up
CREATE TABLE urls(
	id SERIAL PRIMARY KEY,
	small TEXT NOT NULL CHECK (small <> '') UNIQUE,
	origin TEXT NOT NULL CHECK (origin <> ''),
	created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX urls__small__index ON urls USING BTREE (small);

-- +migrate Down
DROP TABLE urls;
