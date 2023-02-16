package functions

import (
	"encoding/csv"
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"strings"
)

func ValidateRole(c echo.Context, accessRole string) (bool, error) {
	file, err := os.Open("pkg/localStorage/policy.csv")
	if err != nil {
		return false, err
	}

	csvReader := csv.NewReader(file)
	var accessibleRole string
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
		if rec[0] == "*" {
       		accessibleRole = rec[1]
        	break
   		} else if rec[0] == c.Request().URL.Path {
    		accessibleRole = rec[1]
        	break
       	}
	}
	if strings.EqualFold(accessRole, accessibleRole) {
		return true, nil
	}
	return false, nil
}