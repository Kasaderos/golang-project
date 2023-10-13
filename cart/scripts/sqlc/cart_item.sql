-- name: AddCartItem :exec
insert into cart_item (
    user_id,
    sku,
    price,
    amount
) VALUES ($1, $2, $3, $4)
on conflict on constraint id
do update
set amount = amount + $4,
    price = $3
where user_id = $1 and sku = $2;

-- name: DeleteItem :exec
delete from cart_item
where user_id = $1 and sku = $2;

-- name: DeleteItemByUser :exec
delete from cart_item
where user_id = $1;

-- name: GetItemsByUserID :many
select * from cart_item
where user_id = $1;
