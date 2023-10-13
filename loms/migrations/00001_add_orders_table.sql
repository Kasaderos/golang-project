-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
  id          bigserial PRIMARY KEY,
  user_id     bigint,
  status_id   int,
  created_at  timestamptz default now() not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders CASCADE;
-- +goose StatementEnd
