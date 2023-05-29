package view

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func SuccessCreated(data interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: "success create",
		Data:    data,
	}
}

func SuccessAction(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success action",
		Data:    data,
	}
}

func SuccessUpload(data interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: "success upload",
		Data:    data,
	}
}

func SuccessLogin(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success login",
		Data:    data,
	}
}

func SuccessRegister(data interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: "success register",
		Data:    data,
	}
}

func SuccessUpdated(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success updated",
		Data:    data,
	}
}

func SuccessDeleted(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success delete",
		Data:    data,
	}
}

func SuccessFind(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success get data",
		Data:    data,
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Error:   err,
	}
}

func ErrInternalServer(err interface{}) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   err,
	}
}
func ErrNotFound(err ...interface{}) *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "Data not found",
		Error:   err,
	}
}
func ErrUnauthorized() *Response {
	return &Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   "UNAUTHORIZED",
	}
}
