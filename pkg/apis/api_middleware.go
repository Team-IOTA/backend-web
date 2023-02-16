package apis

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func EchoManager(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	default_ := e.Group("/classifie")
	defaultSecured_ := e.Group("/classifie")

	defaultSecured_.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	FN_Default(default_)

	//Api Documentation
	e.Use(middleware.Static("docs"))
}

func FN_Default(g *echo.Group) {

    g.POST("/api/CreateUser", CreateUserController)
    g.PUT("/api/UpdateUser", UpdateUserController)
    g.DELETE("/api/DeleteUser", DeleteUserController)
    g.GET("/api/FindUser", FindUserController)
    g.GET("/api/FindallUser", FindallUserController)
    g.POST("/api/CreateTimeStampData", CreateTimeStampDataController)
    g.PUT("/api/UpdateTimeStampData", UpdateTimeStampDataController)
    g.DELETE("/api/DeleteTimeStampData", DeleteTimeStampDataController)
    g.GET("/api/FindTimeStampData", FindTimeStampDataController)
    g.GET("/api/FindallTimeStampData", FindallTimeStampDataController)
    g.POST("/api/CreateTimeStampData", CreateTimeStampDataController)
    g.PUT("/api/UpdateTimeStampData", UpdateTimeStampDataController)
    g.DELETE("/api/DeleteTimeStampData", DeleteTimeStampDataController)
    g.GET("/api/FindTimeStampData", FindTimeStampDataController)
    g.GET("/api/FindallTimeStampData", FindallTimeStampDataController)
    g.POST("/api/CreateCustomNotes", CreateCustomNotesController)
    g.PUT("/api/UpdateCustomNotes", UpdateCustomNotesController)
    g.DELETE("/api/DeleteCustomNotes", DeleteCustomNotesController)
    g.GET("/api/FindCustomNotes", FindCustomNotesController)
    g.GET("/api/FindallCustomNotes", FindallCustomNotesController)
    g.POST("/api/CreateSummery", CreateSummeryController)
    g.PUT("/api/UpdateSummery", UpdateSummeryController)
    g.DELETE("/api/DeleteSummery", DeleteSummeryController)
    g.GET("/api/FindSummery", FindSummeryController)
    g.GET("/api/FindallSummery", FindallSummeryController)

	g.GET("/api/health-check", HealthCheck)
	g.GET("/swagger/*", echoSwagger.WrapHandler)

}