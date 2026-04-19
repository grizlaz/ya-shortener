package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleShortenBatch(t *testing.T) {
	t.Run("short batch urls", func(t *testing.T) {
		path := "/api/shorten/batch"
		URLCount := 5
		baseURL := "http://localhost:8080"
		baseForGenURL := "https://practicum.yandex.ru/?"
		requestData := make([]model.ShortenRequestBatch, 0, URLCount)
		for i := 0; i < URLCount; i++ {
			tURL := fmt.Sprintf("%s%d", baseForGenURL, i)
			requestData = append(requestData, model.ShortenRequestBatch{
				ID:  tURL,
				URL: tURL,
			})
		}

		bodyData, err := json.Marshal(requestData)
		require.NoError(t, err)

		body := bytes.NewReader(bodyData)
		shorten := service.NewService(context.Background(), repository.NewInMemory())
		handler := HandleAPIShortenBatch(shorten, baseURL)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, path, body)
		e := echo.New()
		c := e.NewContext(request, recorder)
		c.SetPath(path)

		require.NoError(t, handler(c))

		assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

		responseBody, err := io.ReadAll(recorder.Result().Body) //nolint:bodyclose
		recorder.Result().Body.Close()
		require.NoError(t, err)

		var response []batchResponse
		err = json.Unmarshal(responseBody, &response)
		require.NoError(t, err)

		assert.Equal(t, len(requestData), len(response))
		// for i := 0; i < URLCount; i++ {
		// 	assert.Equal(t, requestData[i].ID, response[i].ID)
		// 	assert.Contains(t, response[i].URL, baseURL)
		// }
	})
}
