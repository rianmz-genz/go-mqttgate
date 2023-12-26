package exception

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	} else if forbiddenError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func forbiddenError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(ForbiddenError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusForbidden)

		webResponse := web.WebResponse{
			Status:  "Fail",
			Code:    http.StatusForbidden,
			Message: "User doesn't have match permission",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Status:  "Fail",
			Code:    404,
			Message: "Item not found",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Status:  "Fail",
		Code:    500,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
