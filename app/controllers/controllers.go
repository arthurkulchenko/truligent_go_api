package controllers

import(
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/app/services"
)

type ResultObject struct {
	StringValue string
	IntValue int
	BoolValue bool
	Error error
}

type ProccessObject struct {
	RaiseError bool
	State string
}

func (o ResultObject) value() {
	if o.Error != nil {
		return o.Error
	}
	if o.StringValue != nil {
		return o.StringValue
	}
	if o.IntValue != nil {
		return o.IntValue
	}
	if o.BoolValue != nil {
		return o.BoolValue
	}
}

func RotateToken(c echo.Context) error {
	defer c.Request().Body.Close()
	var err error
	var body []byte
	// body := readRequest(request, logAndReturn, "error message")
	body, err = ioutil.ReadAll(c.Request().Body)
	// logAndReturn
	if err != nil {
		log.Println("failed to read rotate_token request body: %s", err)
		return c.String(http.StatusInternalServerError, "failed")
	}
	// requestBody := unmarshalRequest(body, logAndReturn, "error message")
	// getToken instrad of company_id, 
	requestBody := struct { CompanyId string `json:"company_id"` }{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		log.Println("failed to unmatshal rotate_token request body: %s", err)
		return c.String(http.StatusInternalServerError, "failed")
	}
	// tokenString := rotateToken(requestBody, logAndReturn, "error message")
	tokenString, serviceErr := services.RotateTokenService(requestBody.CompanyId)
	if serviceErr != nil {
		return c.String(http.StatusInternalServerError, "failed")
	}
	return c.JSON(http.StatusOK, echo.Map{ "data": tokenString, })
}

// func logAndReturn() {
	// failFunction = func(errMsg string) { return c.String(http.StatusInternalServerError, errMsg) }
	// logFunction = func(msg string) { return log.Println(msg) }
// }

// func RotateToken(c echo.Context) error {
	// body := readRequest(request, logAndReturn, "error message")
	// requestBody := unmarshalRequest(body, logAndReturn, "error message")
	// tokenString := rotateToken(requestBody, logAndReturn, "error message")
	// return c.JSON(http.StatusOK, echo.Map{ "data": tokenString, })
// }
