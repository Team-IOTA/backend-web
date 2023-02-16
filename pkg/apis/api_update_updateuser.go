package apis
import (
  "fmt"
  "classifie/pkg/functions"
  "classifie/pkg/util"
  "classifie/pkg/dbOps"
  "classifie/pkg/models"
  
  
  "net/http"
  "github.com/go-playground/validator"
  "github.com/labstack/echo/v4"
)


func UpdateUserApi(c echo.Context) (*models.Model_User, error) {
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


    inputObj := models.Model_User{}
    if error := c.Bind(&inputObj); error != nil {
    		return nil, error
    	}
 
if err :=functions.UniqueCheck_User(inputObj) ; err!=nil{
      return nil , c.String(http.StatusBadRequest , err.Error())
    }
validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return nil, validationErr
    }
    EncryptedObj,err :=  util.GetEncryptObject(inputObj)
    if err != nil {
       return nil,err
    }
    returnValue, err := dbOps.UpdateUser(&EncryptedObj)
    if err != nil {
        return nil, err
    }
return returnValue, nil
}
    return nil, echo.ErrForbidden
}
