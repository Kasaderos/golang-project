-- +goose Up
CREATE TABLE order (
  id  text  default gen_random_uuid() not null primary key
);

-- +goose Down
DROP TABLE order;

