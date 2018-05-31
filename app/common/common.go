package common

type HttpResponse struct {
	Code int
	Message string
	Success bool
}

func GetInternalServerErrorResponse() HttpResponse {
	var httpResponse HttpResponse
	httpResponse.Code = 500
	httpResponse.Message = "Internal Server Error"
	httpResponse.Success = false
	return httpResponse
}
func GetNotFoundResponse() HttpResponse {
	var httpResponse HttpResponse
	httpResponse.Code = 404
	httpResponse.Message = "Not Found"
	httpResponse.Success = true
	return httpResponse
}
func GetSuccessResponse() HttpResponse {
	var httpResponse HttpResponse
	httpResponse.Code = 200
	httpResponse.Message = "Ok"
	httpResponse.Success = true
	return httpResponse
}