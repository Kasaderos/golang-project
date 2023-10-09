-- name: ReserveStock :exec
update stock 
set count = count - $1,
    reserved = reserved + $1
where count >= $1 and sku = $2;

-- name: ReserveRemove :exec
update stock
set reserved = reserved - $1
where sku = $2;

-- name: ReserveCancel :exec
update stock 
set count = count + $1,
    reserved = reserved - $1
where reserved >= $1 and sku = $2;

-- name: GetBySKU :one
select count - reserved from stock
where sku = $1;
