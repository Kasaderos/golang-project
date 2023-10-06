-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_order
(
  id      bigserial PRIMARY KEY,
  user_id bigint,
  status  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_order CASCADE;
-- +goose StatementEnd
