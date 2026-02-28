package service_test

import (
	"fmt"
	"testing"

	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShorten(t *testing.T) {
	var id uint32 = 1024
	resForId := "sw"
	t.Run("returns short identifier", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}

		testCases := []testCase{
			{
				id:       id,
				expected: resForId,
			},
			{
				id:       0,
				expected: "",
			},
		}

		for _, tc := range testCases {
			actual := service.Shorten(tc.id)
			assert.Equal(t, tc.expected, actual)
		}
	})

	t.Run("is idempotent", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			assert.Equal(t, resForId, service.Shorten(id))
		}
	})

	t.Run("return prepared base url", func(t *testing.T) {
		type testCase struct {
			baseURL    string
			identifier string
			expected   string
			noError    bool
		}

		testCases := []testCase{
			{
				baseURL:    "http://localhost",
				identifier: "asd",
				expected:   fmt.Sprintf("%s/%s", "http://localhost", "asd"),
				noError:    true,
			},
			{
				baseURL:    string('\n'),
				identifier: "err",
				noError:    false,
			},
		}

		for _, tc := range testCases {
			actual, err := service.PrependBaseURL(tc.baseURL, tc.identifier)
			if tc.noError {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			} else {
				assert.Error(t, err)
			}
		}
	})
}
