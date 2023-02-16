package apis
import (
  "fmt"
  "classifie/pkg/functions"
  "classifie/pkg/dbOps"
  "classifie/pkg/models"
  
  "github.com/go-playground/validator"
  "github.com/labstack/echo/v4"
)


func UpdateTimeStampDataApi(c echo.Context) (*models.Model_TimeStampData, error) {
claims, err := functions.DecodeJWT(c)
    	if err != nil {
    		return nil, err
    	}

    	role := fmt.Sprintf("%v", claims["role"])
    	accessApi, err := functions.ValidateRole(c, role)
    	if err != nil {
    		return nil, err
    	}

    	if accessApi {


    inputObj := models.Model_TimeStampData{}
    if error := c.Bind(&inputObj); error != nil {
    		return nil, error
    	}
 
validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return nil, validationErr
    }
    returnValue, err := dbOps.UpdateTimeStampData(&inputObj)
    if err != nil {
    	return nil, err
    }
return returnValue, nil
}
    return nil, echo.ErrForbidden
}
