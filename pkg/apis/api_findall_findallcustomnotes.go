package apis
import (
  "fmt"
  "classifie/pkg/functions"
  "classifie/pkg/dbOps"
  "classifie/pkg/models"
  
  
  "github.com/labstack/echo/v4"
)


func FindallCustomNotesApi(c echo.Context) (*[]models.Model_CustomNotes, error) {
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

returnValue, err := dbOps.FindallCustomNotes()
    if err != nil {
        return nil, err
    }
    
         return returnValue, nil
    
}
    return nil, echo.ErrForbidden
}
