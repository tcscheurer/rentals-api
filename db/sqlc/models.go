// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlc

import (
	"database/sql"
)

type Rental struct {
	ID              int32
	UserID          sql.NullInt32
	Name            sql.NullString
	Type            sql.NullString
	Description     sql.NullString
	Sleeps          sql.NullInt32
	PricePerDay     sql.NullInt64
	HomeCity        sql.NullString
	HomeState       sql.NullString
	HomeZip         sql.NullString
	HomeCountry     sql.NullString
	VehicleMake     sql.NullString
	VehicleModel    sql.NullString
	VehicleYear     sql.NullInt32
	VehicleLength   sql.NullString
	Created         sql.NullTime
	Updated         sql.NullTime
	Lat             sql.NullFloat64
	Lng             sql.NullFloat64
	PrimaryImageUrl sql.NullString
}

type User struct {
	ID        int32
	FirstName sql.NullString
	LastName  sql.NullString
}