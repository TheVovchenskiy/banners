package queryManager

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUrlQuery(t *testing.T) {
	testCases := []struct {
		name          string
		url           string
		allowedParams []QueryParam
		expected      ParsedQueryParams
		expectedErr   error
	}{
		{
			"int param",
			"http://example.com?param=1",
			[]QueryParam{
				{
					"param",
					IntType,
					false,
					0,
				},
			},
			ParsedQueryParams{
				"param": 1,
			},
			nil,
		},
		{
			"bool param",
			"http://example.com?param=false",
			[]QueryParam{
				{
					"param",
					BoolType,
					false,
					true,
				},
			},
			ParsedQueryParams{
				"param": false,
			},
			nil,
		},
		{
			"string param",
			"http://example.com?param=string",
			[]QueryParam{
				{
					"param",
					StringType,
					false,
					"string",
				},
			},
			ParsedQueryParams{
				"param": "string",
			},
			nil,
		},
		{
			"float param",
			"http://example.com?param=3.14",
			[]QueryParam{
				{
					"param",
					FloatType,
					false,
					0.0,
				},
			},
			ParsedQueryParams{
				"param": 3.14,
			},
			nil,
		},
		{
			"default value",
			"http://example.com",
			[]QueryParam{
				{
					"param",
					IntType,
					false,
					37,
				},
			},
			ParsedQueryParams{
				"param": 37,
			},
			nil,
		},
		{
			"empty required param error",
			"http://example.com?param1=1",
			[]QueryParam{
				{
					"param",
					IntType,
					true,
					0,
				},
			},
			nil,
			ErrMissingRequiredParam,
		},
		{
			"invalid conversion error",
			"http://example.com?param=string",
			[]QueryParam{
				{
					"param",
					IntType,
					true,
					0,
				},
			},
			nil,
			ErrConvertingParam,
		},
		{
			"unsupported param type error",
			"http://example.com?param=string",
			[]QueryParam{
				{
					"param",
					33,
					true,
					0,
				},
			},
			nil,
			ErrUnsupportedParamType,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			url, _ := url.Parse(tt.url)
			actual, actualErr := ParseUrlQuery(*url, tt.allowedParams)

			assert.Equal(t, tt.expected, actual)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}
