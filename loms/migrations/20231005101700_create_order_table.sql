-- +goose Up 
CREATE TABLE "order"
(
  id text default gen_random_uuid() not null,
  user_id text,
  status text,
  primary key(id)
);

CREATE TABLE "cart"

-- +goose Down
DROP TABLE "order";

