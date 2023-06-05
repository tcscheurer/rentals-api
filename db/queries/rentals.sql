-- name: GetRentalByID :one
select * from rentals
join users on rentals.user_id = users.id
where rentals.id = $1;

-- name: GetRentals :many
select * from rentals 
join users on rentals.user_id = users.id
where

CASE 
WHEN @filter_ids::bool
THEN 
    rentals.id = ANY(@id_list::int[]) and
    price_per_day >= @price_per_day_low::integer and
    price_per_day <= @price_per_day_high::integer
ELSE
    price_per_day >= @price_per_day_low::integer and
    price_per_day <= @price_per_day_high::integer
END

and
CASE
WHEN @find_near::bool
THEN
    ST_DistanceSphere(ST_MakePoint(rentals.lng,rentals.lat), ST_MakePoint(@near_lng::double precision,@near_lat::double precision)) <= 100 * 1609.34
ELSE
    rentals.id is not null
END

order by
CASE 
WHEN @sort_by_price::bool
THEN
    rentals.price_per_day
ELSE
    rentals.id
END
LIMIT $1
OFFSET $2;