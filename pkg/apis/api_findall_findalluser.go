package apis
import (
  "fmt"
  "classifie/pkg/functions"
  "classifie/pkg/util"
  "classifie/pkg/dbOps"
  "classifie/pkg/models"
  
  
  "github.com/labstack/echo/v4"
)


func FindallUserApi(c echo.Context) (*[]models.Model_User, error) {
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

returnValue, err := dbOps.FindallUser()
    if err != nil {
        return nil, err
    }
    DecryptedObjs,err :=  util.GetDecryptObjects(*returnValue)
        if err != nil {
            return nil,err
        }
        
        return &DecryptedObjs, nil
        
}
    return nil, echo.ErrForbidden
}
