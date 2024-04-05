package queryManager

import (
	"fmt"
	"net/url"
	"strconv"
)

type ParamType int

const (
	StringType ParamType = iota
	IntType
	FloatType
	BoolType
)

type QueryParam struct {
	Name         string
	Type         ParamType
	Required     bool
	DefaultValue any
}

type ParsedQueryParams map[string]any

func ParseUrlQuery(url url.URL, allowedParams []QueryParam) (ParsedQueryParams, error) {
	queryParams := url.Query()
	res := ParsedQueryParams{}

	for _, param := range allowedParams {
		values, exists := queryParams[param.Name]

		if !exists || len(values) == 0 {
			if param.Required {
				return nil, fmt.Errorf("%w: %s", ErrMissingRequiredParam, param.Name)
			} else {
				res[param.Name] = param.DefaultValue
			}
		} else {
			value, err := convertValueByType(values[0], param.Type)
			if err != nil {
				return nil, fmt.Errorf("%w %s: %w", ErrConvertingParam, param.Name, err)
			}
			res[param.Name] = value
		}
	}

	return res, nil

}

func convertValueByType(value string, paramType ParamType) (any, error) {
	switch paramType {
	case StringType:
		return value, nil
	case IntType:
		return strconv.Atoi(value)
	case FloatType:
		return strconv.ParseFloat(value, 64)
	case BoolType:
		return strconv.ParseBool(value)
	default:
		return nil, ErrUnsupportedParamType
	}
}
