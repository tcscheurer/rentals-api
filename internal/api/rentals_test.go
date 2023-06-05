package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/tcscheurer/rentals/db/mock"
	"github.com/tcscheurer/rentals/db/sqlc"
	"github.com/tcscheurer/rentals/internal/models"
)


func TestGetRentalsAPI(t *testing.T) {
	r := randomRental()
	xr := randomRentalSlice()

	testCases := []struct {
		name string
		pathExt string
		query string
		buildStubs func(querier *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Get by ID",
			pathExt: "/"+ strconv.Itoa(int(r.ID)),
			query: "",
			buildStubs: func(querier *mockdb.MockQuerier){
				querier.EXPECT().
				GetRentalByID(gomock.Any(), gomock.Any()).
				Times(1).
				Return(r, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchSingle(t, recorder.Body, r)
			},
		},
		{
			name: "Get List",
			pathExt: "",
			query: "",
			buildStubs: func(querier *mockdb.MockQuerier){
				querier.EXPECT().
				GetRentals(gomock.Any(), gomock.Any()).
				Times(1).
				Return(xr, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchSlice(t, recorder.Body, xr)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			q := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(q)

			s := NewServer(q)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/rentals%s", tc.pathExt)
			r, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			s.router.ServeHTTP(recorder, r)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomRental() sqlc.GetRentalByIDRow {
	return sqlc.GetRentalByIDRow{
		ID: 1, 
		UserID: sql.NullInt32{
			Valid: true,
			Int32: 23,
		},
		Name: sql.NullString{
			Valid: true,
			String: "Trevor",
		},
	}
}

func requireBodyMatchSingle(t *testing.T, b *bytes.Buffer, rentalRow sqlc.GetRentalByIDRow) {
	data, err := ioutil.ReadAll(b)
	require.NoError(t, err)
	var gotRental models.Rental
	err = json.Unmarshal(data, &gotRental)
	require.NoError(t, err)
	require.Equal(t, gotRental.ID, rentalRow.ID)
}

func randomRentalSlice() []sqlc.GetRentalsRow {
	return []sqlc.GetRentalsRow{
		{
			ID: 1, 
			UserID: sql.NullInt32{
				Valid: true,
				Int32: 23,
			},
			Name: sql.NullString{
				Valid: true,
				String: "Chad",
			},
		},
		{
			ID: 2, 
			UserID: sql.NullInt32{
				Valid: true,
				Int32: 23,
			},
			Name: sql.NullString{
				Valid: true,
				String: "Jess",
			},
		},
		{
			ID: 3, 
			UserID: sql.NullInt32{
				Valid: true,
				Int32: 23,
			},
			Name: sql.NullString{
				Valid: true,
				String: "Trevor",
			},
		},
	}
}

func requireBodyMatchSlice(t *testing.T, b *bytes.Buffer, rows []sqlc.GetRentalsRow) {
	data, err := ioutil.ReadAll(b)
	require.NoError(t, err)
	var gotRental []models.Rental
	err = json.Unmarshal(data, &gotRental)
	require.NoError(t, err)
	require.Equal(t, len(gotRental), len(rows))
	for i := range rows {
		require.Equal(t, gotRental[i].ID, rows[i].ID)
		require.Equal(t, gotRental[i].Name, rows[i].Name.String)
		require.Equal(t, gotRental[i].User.ID, rows[i].UserID.Int32)
	}
}