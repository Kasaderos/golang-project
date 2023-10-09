-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
  id      bigserial PRIMARY KEY,
  user_id bigint,
  status  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders CASCADE;
-- +goose StatementEnd
