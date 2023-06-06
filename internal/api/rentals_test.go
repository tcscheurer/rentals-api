package api

import (
	// "bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	// "strconv"
	"testing"

	// "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	// mockdb "github.com/tcscheurer/rentals/db/mock"
	"github.com/tcscheurer/rentals/db/sqlc"
	"github.com/tcscheurer/rentals/internal/models"
	"github.com/tcscheurer/rentals/utils"
)

func SetUpServer(t *testing.T) *Server {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		t.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		t.Fatal("cannot connect to db:", err)
	}
	q := sqlc.New(conn)
	return NewServer(q)
}

func TestGetRental(t *testing.T) {
	s := SetUpServer(t)
	
	tests := []struct {
		name string
		uriParam int32
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
	{
		name: "Get: 1",
		uriParam: 1,
		validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			data, err := ioutil.ReadAll(recorder.Body)
			require.NoError(t, err)
			var gotRental models.Rental
			err = json.Unmarshal(data, &gotRental)
			require.NoError(t, err)
			require.Equal(t, gotRental.ID, int32(1))
			require.Equal(t, gotRental.Name, "'Abaco' VW Bay Window: Westfalia Pop-top")
			require.Equal(t, gotRental.Type, "camper-van")
			require.Equal(t, gotRental.Sleeps, int32(4))
		},
	},
	{
		name: "Get: 2",
		uriParam: 2,
		validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			data, err := ioutil.ReadAll(recorder.Body)
			require.NoError(t, err)
			var gotRental models.Rental
			err = json.Unmarshal(data, &gotRental)
			require.NoError(t, err)
			require.Equal(t, gotRental.ID, int32(2))
			require.Equal(t, gotRental.Name, "Maupin: Vanagon Camper")
			require.Equal(t, gotRental.Type, "camper-van")
			require.Equal(t, gotRental.Sleeps, int32(4))
		},
	},
	{
		name: "Get: 24",
		uriParam: 24,
		validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			data, err := ioutil.ReadAll(recorder.Body)
			require.NoError(t, err)
			var gotRental models.Rental
			err = json.Unmarshal(data, &gotRental)
			require.NoError(t, err)
			require.Equal(t, gotRental.ID, int32(24))
			require.Equal(t, gotRental.Name, "2017 Ford Transit")
			require.Equal(t, gotRental.Type, "camper-van")
			require.Equal(t, gotRental.Sleeps, int32(1))
		},
	},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T){
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/rentals/%d", tc.uriParam)
			r, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			s.router.ServeHTTP(recorder, r)
			tc.validateResponse(t, recorder)
		})
	}
}

func TestGetRentals(t *testing.T) {
	s := SetUpServer(t)

	tests := []struct{
		name string
		queryParam string
		validateResponse func (t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "List: Empty Query",
			queryParam: "",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 11)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 10)
			},
		},
		{
			name: "List: With Offset",
			queryParam: "offset=2",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 11)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 10)
				require.Equal(t, gotRentals[0].ID, int32(3))
			},
		},
		{
			name: "List: With Limit",
			queryParam: "limit=2",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 11)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 2)
				require.Equal(t, gotRentals[0].ID, int32(1))
			},
		},
		{
			name: "List: With Offset and Limit",
			queryParam: "limit=3&offset=5",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 11)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 3)
				require.Equal(t, gotRentals[0].ID, int32(6))
				require.Equal(t, gotRentals[len(gotRentals)-1].ID, int32(8))
			},
		},
		{
			name: "List: With ID List Filter",
			queryParam: "ids=1,19,17",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 11)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 3)
				wl := map[int32]bool {
					1: true,
					19: true,
					17: true,
				}
				for _, r := range gotRentals {
					if !wl[r.ID] {
						t.Fatal("unexepected rental returned using id filter")
					}
				}
			},
		},
		{
			name: "List: With Price Range",
			queryParam: "price_min=17000&price_max=18500",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 11)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 3)
				for _, r := range gotRentals {
					if r.Price.Day < 1700 || r.Price.Day > 18500 {
						t.Fatal("Enexpted price range using price filtering: ", r.Price.Day)
					}
				}
			},
		},
		{
			name: "List: With Near Search",
			queryParam: "near=32.73,-117.24",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 15)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 5)
				for _, r := range gotRentals {
					if r.Location.State != "CA" {
						t.Fatal("Expected rentals to be in California using near filter. Actual: ", r.Location.State)
					}
				}
			},
		},
		{
			name: "List: With Limit and Near Search", 
			queryParam: "near=32.73,-117.24&limit=2",
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				data, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				gotRentals := make([]models.Rental, 0, 15)
				err = json.Unmarshal(data, &gotRentals)
				require.NoError(t, err)
				require.Equal(t, len(gotRentals), 2)
				for _, r := range gotRentals {
					if r.Location.State != "CA" {
						t.Fatal("Expected rentals to be in California using near filter. Actual: ", r.Location.State)
					}
				}
			},
		},
	}
	
	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T){
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/rentals?%s", tc.queryParam)
			r, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			s.router.ServeHTTP(recorder, r)
			tc.validateResponse(t, recorder)
		})
	}
}