package echorestaurant

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restaurantbusiness"
	"simple-rest-api/modules/restaurant/restaurantstorage"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewGetRestaurantBusiness(store)

		data, err := business.GetRestaurant(c.Request().Context(), id)

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
