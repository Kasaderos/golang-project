-- name: ReserveStock :exec
update stock 
set count = count - $1,
    total_reserved = total_reserved + $1
where sku = $2;

-- name: ReserveRemove :exec
update stock
set total_reserved = total_reserved - $1
where sku = $2;

-- name: ReserveCancel :exec
update stock 
set count = count + $1,
    total_reserved = total_reserved - $1
where sku = $2;

-- name: GetBySKU :one
select count - total_reserved from stock
where sku = $1;
