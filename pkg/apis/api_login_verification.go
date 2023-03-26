package apis
import (
  "fmt"
  "classifie/pkg/functions"
  "classifie/pkg/util"
  "classifie/pkg/dbOps"
  "classifie/pkg/models"
  "github.com/labstack/echo/v4"
)


func LoginVerificationApi(c echo.Context) (*models.Model_Login_Response, error) {
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

	inputObj := models.Model_Login_Input{}
    if error := c.Bind(&inputObj); error != nil {
    		return nil, error
    	}
 
	returnValue, err := dbOps.FindUser(inputObj.UserName)
    if err != nil {
        return nil, err
    }
    DecryptedObj,err :=  util.GetDecryptObject(*returnValue)
        if err != nil {
            return nil,err
        }
		if DecryptedObj.Password == inputObj.Password{
			return &models.Model_Login_Response{Login: "True"},nil
		}else {
			return &models.Model_Login_Response{Login: "False"},nil
		}
        
        
}
    return nil, echo.ErrForbidden
}