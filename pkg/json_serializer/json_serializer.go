package json_serializer

import (
	"encoding/json"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	echo "github.com/labstack/echo/v5"
)

type JsonSerializer struct{}

// Serialize
// @param c
// @param i
// @param indent
// @date 2022-09-10 17:36:02
func (j *JsonSerializer) Serialize(c echo.Context, i interface{}, indent string) error {

	enc := jsoniter.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize
// @param c
// @param i
// @date 2022-09-10 17:36:01
func (j *JsonSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := jsoniter.NewDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).WithInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).WithInternal(err)
	}
	return err
}
