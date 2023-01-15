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

func ListRestaurant(appCtx component.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		var filter restaurantmodel.Filter

		if err := c.Bind(&filter); err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		var paging common.Paging

		if err := c.Bind(&paging); err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewListRestaurantBusiness(store)

		result, err := business.ListRestaurant(c.Request().Context(), &filter, &paging)

		if err != nil {
			return c.JSON(401, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
