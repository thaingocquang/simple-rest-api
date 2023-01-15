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

func DeleteRestaurant(appCtx component.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err = business.DeleteRestaurant(c.Request().Context(), id); err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
