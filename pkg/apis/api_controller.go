package apis

import (
	"github.com/labstack/echo/v4"
	"net/http"
)
// @title          Quickstart for Evo-Codes REST API
// @version         1.0
// @description     Learn how to get started with evo-codes REST API.
// @contact.name   API Support
// @contact.url   https://evolza.org
// @contact.email  isuru@evolza.org
// @license.name  Evolza
// @license.url   http://evolza.org
// @BasePath  /classifie/api

func CreateUserController(c echo.Context) error {

		result, err := CreateUserApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func UpdateUserController(c echo.Context) error {

		result, err := UpdateUserApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func DeleteUserController(c echo.Context) error {

		result, err := DeleteUserApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindUserController(c echo.Context) error {

		result, err := FindUserApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindallUserController(c echo.Context) error {

		result, err := FindallUserApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func CreateTimeStampDataController(c echo.Context) error {

		result, err := CreateTimeStampDataApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func UpdateTimeStampDataController(c echo.Context) error {

		result, err := UpdateTimeStampDataApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func DeleteTimeStampDataController(c echo.Context) error {

		result, err := DeleteTimeStampDataApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindTimeStampDataController(c echo.Context) error {

		result, err := FindTimeStampDataApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindallTimeStampDataController(c echo.Context) error {

		result, err := FindallTimeStampDataApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}

func CreateCustomNotesController(c echo.Context) error {

		result, err := CreateCustomNotesApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func UpdateCustomNotesController(c echo.Context) error {

		result, err := UpdateCustomNotesApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func DeleteCustomNotesController(c echo.Context) error {

		result, err := DeleteCustomNotesApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindCustomNotesController(c echo.Context) error {

		result, err := FindCustomNotesApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindallCustomNotesController(c echo.Context) error {

		result, err := FindallCustomNotesApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func CreateSummeryController(c echo.Context) error {

		result, err := CreateSummeryApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func UpdateSummeryController(c echo.Context) error {

		result, err := UpdateSummeryApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func DeleteSummeryController(c echo.Context) error {

		result, err := DeleteSummeryApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindSummeryController(c echo.Context) error {

		result, err := FindSummeryApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}
func FindallSummeryController(c echo.Context) error {

		result, err := FindallSummeryApi(c)
		if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
	}


func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": "classifie is up and running",
	})
}
