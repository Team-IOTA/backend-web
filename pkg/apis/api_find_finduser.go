package apis
import (
  "fmt"
  "classifie/pkg/functions"
  "classifie/pkg/util"
  "classifie/pkg/dbOps"
  "classifie/pkg/models"
  
  
  "github.com/labstack/echo/v4"
)


func FindUserApi(c echo.Context) (*models.Model_User, error) {
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

objectid := c.QueryParam("objectid")
returnValue, err := dbOps.FindUser(objectid)
    if err != nil {
        return nil, err
    }
    DecryptedObj,err :=  util.GetDecryptObject(*returnValue)
        if err != nil {
            return nil,err
        }
        
        return &DecryptedObj, nil
        
}
    return nil, echo.ErrForbidden
}
