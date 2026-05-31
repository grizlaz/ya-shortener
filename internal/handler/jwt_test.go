package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestJwt(t *testing.T) {
	t.Run("get userID from request", func(t *testing.T) {
		path := "/ping"
		request := httptest.NewRequest(http.MethodGet, path, nil)
		origUserID := uuid.New()
		token, err := makeJWT(origUserID)

		require.NoError(t, err, "error create JWT token")

		cookie := &http.Cookie{
			Name:  "Authorization",
			Value: token,
		}
		request.AddCookie(cookie)
		recorder := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(request, recorder)
		c.SetPath(path)

		userID, err := getUserID(c)
		require.NoError(t, err, "error get userID")

		require.Equal(t, origUserID, userID, "userID not equal: %s != %s", origUserID, userID)
	})
}
