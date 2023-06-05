CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name text,
    last_name text
);

CREATE TABLE IF NOT EXISTS rentals (
    id SERIAL PRIMARY KEY,
    user_id integer,
    name text,
    type text,
    description text,
    sleeps integer,
    price_per_day bigint,
    home_city text,
    home_state text,
    home_zip text,
    home_country text,
    vehicle_make text,
    vehicle_model text,
    vehicle_year integer,
    vehicle_length numeric(4,2),
    created timestamp with time zone,
    updated timestamp with time zone,
    lat double precision,
    lng double precision,
    primary_image_url text
);