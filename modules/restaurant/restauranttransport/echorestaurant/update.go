package echorestaurant

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restaurantbusiness"
	"simple-rest-api/modules/restaurant/restaurantmodel"
	"simple-rest-api/modules/restaurant/restaurantstorage"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		var data restaurantmodel.RestaurantUpdate

		if err = c.Bind(&data); err != nil {
			return err
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewUpdateRestaurantBusiness(store)

		if err = business.UpdateRestaurant(c.Request().Context(), id, &data); err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
