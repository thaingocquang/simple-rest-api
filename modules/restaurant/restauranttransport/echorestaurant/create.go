package echorestaurant

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/modules/restaurant/restaurantbusiness"
	"simple-rest-api/modules/restaurant/restaurantmodel"
	"simple-rest-api/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(appCtx component.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data restaurantmodel.RestaurantCreate

		if err := c.Bind(&data); err != nil {
			return err
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)

		if err := business.CreateRestaurant(c.Request().Context(), &data); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
