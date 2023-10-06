-- name: AddCartItem :exec
insert into cart_item (
    user_id,
    sku,
    name,
    price,
    count
) VALUES ($1, $2, $3, $4, $5);