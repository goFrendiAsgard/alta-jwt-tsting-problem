package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"gotest.tools/assert"
)

func TestController(t *testing.T) {
	// userId is 73, because: https://bigbangtheory.fandom.com/wiki/73
	userId := 73
	token := CreateToken(userId)
	signedString, _ := CreateSignedString(userId)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", signedString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/sebuahPath")
	// This is the trick, somehow getting jwt.Token doesn't work
	context.Set("user", token)

	Controller(context)

	var response map[string]int
	json.Unmarshal(res.Body.Bytes(), &response)

	t.Run("GET /driver/orderlist", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, userId, response["userId"])
	})
}
