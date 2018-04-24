-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (username varchar, password varchar);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;