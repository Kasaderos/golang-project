-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_item
(
  user_id bigint,
  sku bigint,
  name text,
  price int,
  count int
);

CREATE INDEX idx_user_id ON cart_item(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_item CASCADE;
-- +goose StatementEnd