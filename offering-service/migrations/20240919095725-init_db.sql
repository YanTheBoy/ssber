-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE offering
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(50) NOT NULL,
    description VARCHAR(50),
    flat_type varchar,
    price INTEGER,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(10)
);


-- +migrate Down
drop table if exists offering;
