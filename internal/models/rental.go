package models

import "github.com/tcscheurer/rentals/db/sqlc"

type Rental struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	Make string `json:"make"`
	Model string `json:"model"`
	Year int32 `json:"year"`
	Length string `json:"length"`
	Sleeps int32 `json:"sleeps"`
	PrimaryImageUrl string `json:"primary_image_url"`
	Price RentalPrice
	Location Location
	User User
}

type RentalPrice struct {
	Day int64 `json:"day"`
}

type Location struct {
	City string
	State string
	Zip string
	Country string
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}


// TODO: better null handling, and length's persistance model is incorrect
func AdaptByID(r sqlc.GetRentalByIDRow) Rental {
	rental := Rental{
		ID: r.ID,
		Name: r.Name.String,
		Description: r.Description.String,
		Type: r.Type.String,
		Make: r.VehicleMake.String,
		Model: r.VehicleModel.String,
		Year: r.VehicleYear.Int32,
		Length: r.VehicleLength.String,
		Sleeps: r.Sleeps.Int32,
		PrimaryImageUrl: r.PrimaryImageUrl.String,
		Price: RentalPrice{
			Day: r.PricePerDay.Int64,
		},
		Location: Location{
			City: r.HomeCity.String,
			State: r.HomeState.String,
			Zip: r.HomeZip.String,
			Country: r.HomeCountry.String,
			Latitude: r.Lat.Float64,
			Longitude: r.Lng.Float64,
		},
		User: User{
			ID: r.UserID.Int32,
			FirstName: r.FirstName.String,
			LastName: r.LastName.String,
		},
	}

	return rental
}

func AdaptSlice(xr []sqlc.GetRentalsRow ) []Rental {
	var out []Rental
	for _, r := range xr {
		out = append(out, Rental{
			ID: r.ID,
			Name: r.Name.String,
			Description: r.Description.String,
			Type: r.Type.String,
			Make: r.VehicleMake.String,
			Model: r.VehicleModel.String,
			Year: r.VehicleYear.Int32,
			Length: r.VehicleLength.String,
			Sleeps: r.Sleeps.Int32,
			PrimaryImageUrl: r.PrimaryImageUrl.String,
			Price: RentalPrice{
				Day: r.PricePerDay.Int64,
			},
			Location: Location{
				City: r.HomeCity.String,
				State: r.HomeState.String,
				Zip: r.HomeZip.String,
				Country: r.HomeCountry.String,
				Latitude: r.Lat.Float64,
				Longitude: r.Lng.Float64,
			},
			User: User{
				ID: r.UserID.Int32,
				FirstName: r.FirstName.String,
				LastName: r.LastName.String,
			},
		})
	}
	return out
}