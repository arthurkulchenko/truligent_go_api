package controllers

import(
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/app/services"
)

func RotateToken(c echo.Context) error {
	defer c.Request().Body.Close()
	var err error
	var body []byte
	body, err = ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("failed to read rotate_token request body: %s", err)
		return c.String(http.StatusInternalServerError, "failed")
	}
	requestBody := struct { CompanyId string `json:"company_id"` }{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		log.Println("failed to unmatshal rotate_token request body: %s", err)
		return c.String(http.StatusInternalServerError, "failed")
	}
	tokenString, serviceErr := services.RotateTokenService(requestBody.CompanyId)
	if serviceErr != nil {
		return c.String(http.StatusInternalServerError, "failed")
	}
	return c.JSON(http.StatusOK, echo.Map{ "data": tokenString, })
}
