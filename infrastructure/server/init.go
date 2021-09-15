package server

// NewGinHTTPHandlerDefault ...
func NewGinHTTPHandlerDefault() GinHTTPHandler {
	httpHandler, err := NewGinHTTPHandler(":8080")
	if err != nil {
		panic(err.Error())
	}
	return httpHandler
}
