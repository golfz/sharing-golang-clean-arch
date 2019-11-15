package viewinterface

type ResponseSender interface {
	SendResponse(statusCode int, output interface{})
}
