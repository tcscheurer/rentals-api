package api

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tcscheurer/rentals/db/sqlc"
	"github.com/tcscheurer/rentals/internal/models"
)

var idListExp *regexp.Regexp = regexp.MustCompile(`^(\w+)(,\s*\w+)*$`)
var nearExp *regexp.Regexp = regexp.MustCompile(`^(\s*-?\d+(\.\d+)?)(\s*,\s*-?\d+(\.\d+)?)*$`)
var defaultListLimit int32 = 10
var customSortWhiteList []string = []string{"price"}

type getRentalRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listRentalsRequest struct {
	Offset   int32 `form:"offset"`
	Limit int32 `form:"limit"`
	IDs string `form:"ids"`
	PriceMin int32 `form:"price_min"`
	PriceMax int32 `form:"price_max"`
	Sort string `form:"sort"`
	Near string `form:"near"`
}

func (s *Server) GetRental(ctx *gin.Context) {
	var req getRentalRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	q := s.querier
	r, err := q.GetRentalByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, models.AdaptByID(r))
}

func (s *Server) GetRentals(ctx *gin.Context) {
var req listRentalsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shouldFilterIDs := len(req.IDs) > 0
	ppdLow := req.PriceMin
	var ppdHigh int32 = 25000
	if req.PriceMax != 0 {
		ppdHigh = req.PriceMax
	}

	var idList []int32
	if shouldFilterIDs {
		// validate id list
		if !idListExp.MatchString(req.IDs){
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid value for id filter")))
			return
		} else {
			for _, v := range strings.Split(req.IDs, ",") {
				i, err := strconv.Atoi(v)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				idList = append(idList, int32(i))
			}
		}

	}
	
	limit := req.Limit
	if limit == 0 {
		limit = defaultListLimit
	}

	priceSort := false
	if req.Sort != "" {
		//validate sort param, this would need to be improved to extend past price sorting
		for _, v := range customSortWhiteList {
			if v == req.Sort {
				priceSort = true
			}
		}
	} 

	findNear := false
	var nearLat float64 = 0
	var nearLng float64 = 0
	if len(req.Near) > 0 {
		if !nearExp.MatchString(req.Near){
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid value for near filter")))
			return
		} else {
			findNear = true
			for i, v := range strings.Split(req.Near, ",") {
				f, err := strconv.ParseFloat(v, 32)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				if i == 0 {
					nearLat = f
				} else {
					nearLng = f
				}
			}
		}
	}

	args := sqlc.GetRentalsParams{
		Limit: limit,
		Offset: req.Offset,
		FilterIds: shouldFilterIDs,
		IDList: idList,
		PricePerDayLow: ppdLow,
		PricePerDayHigh: ppdHigh,
		SortByPrice: priceSort,
		FindNear: findNear,
		NearLat: nearLat,
		NearLng: nearLng,
	}

	q := s.querier
	xr, err := q.GetRentals(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		models.AdaptSlice(xr),
	)
}