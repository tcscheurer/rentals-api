## Rentals API
An API for returning rental and user information!
_Definitely a work in progress, and needs sufficient testing_

Supports two endpoints:
- GET /rentals/:id
- GET /rentals
    - Supported params:
        - price_min
        - price_max
        - limit
        - offset
        - ids
        - near
        - sort 
    - Examples:
        - `rentals?price_min=9000&price_max=75000`
        - `rentals?limit=3&offset=6`
        - `rentals?ids=3,4,5`
        - `rentals?near=33.64,-117.93` // finds within 100 miles
        - `rentals?sort=price`
        - `rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`

Example rental response:

```json
{
  "id": "int",
  "name": "string",
  "description": "string",
  "type": "string",
  "make": "string",
  "model": "string",
  "year": "int",
  "length": "decimal",
  "sleeps": "int",
  "primary_image_url": "string",
  "price": {
    "day": "int"
  },
  "location": {
    "city": "string",
    "state": "string",
    "zip": "string",
    "country": "string",
    "lat": "decimal",
    "lng": "decimal"
  },
  "user": {
    "id": "int",
    "first_name": "string",
    "last_name": "string"
  }
}
```

## Running Locally 
`make local-up`
This is a helper for running `docker-compose up -d && go run ./cmd/api/main.go`

Other commands are provided for managing the runtime of the db and service independently


## Testing
Test are set up to connect to the local containerized db and assert based on the APIs interactions with the seeded data
`make db-up` (wrapper for `docker-compose up -d`)
and then,
`make test` (wrapper for `go test ./... -v`)
