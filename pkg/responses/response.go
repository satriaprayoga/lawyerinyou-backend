package responses

import (
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/utils"

	"github.com/labstack/echo/v4"
)

// Res :
type Res struct {
	R echo.Context
}

// ResponseModel :
type ResponseModel struct {
	// Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// Response :
func (e Res) Response(httpCode int, errMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{

		Msg:  errMsg,
		Data: data,
	}
	logger.Info(string(utils.Stringify(response)))
	return e.R.JSON(httpCode, response)
}

// ResponseError :
func (e Res) ResponseError(httpCode int, errMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{

		Msg:  errMsg,
		Data: data,
	}
	logger.Error(string(utils.Stringify(response)))
	return e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}

// ResponseErrorList :
func (e Res) ResponseErrorList(httpCode int, errMsg string, data models.ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Error(string(utils.Stringify(data)))
	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}

// ResponseList :
func (e Res) ResponseList(httpCode int, errMsg string, data models.ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Info(string(utils.Stringify(data)))
	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}
