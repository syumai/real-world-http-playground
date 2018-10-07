package http

const (
	StatusOK        = 200
	StatusCreated   = 201
	StatusAccepted  = 202
	StatusNoContent = 204

	StatusMovedPermanently  = 301
	StatusNotModified       = 304
	StatusTemporaryRedirect = 307
	StatusPermanentRedirect = 308

	StatusBadRequest         = 400
	StatusUnauthorized       = 401
	StatusPaymentRequired    = 402
	StatusForbidden          = 403
	StatusNotFound           = 404
	StatusMethodNotAllowed   = 405
	StatusNotAcceptable      = 406
	StatusRequestTimeout     = 408
	StatusPreconditionFailed = 412

	StatusInternalServerError     = 500
	StatusNotImplemented          = 501
	StatusBadGateway              = 502
	StatusServiceUnavailable      = 503
	StatusGatewayTimeout          = 504
	StatusHTTPVersionNotSupported = 505
)

var statusText = map[int]string{
	StatusOK:        "OK",
	StatusCreated:   "Created",
	StatusAccepted:  "Accepted",
	StatusNoContent: "No Content",

	StatusMovedPermanently:  "Moved Permanently",
	StatusNotModified:       "Not Modified",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadRequest:         "Bad Request",
	StatusUnauthorized:       "Unauthorized",
	StatusPaymentRequired:    "Payment Required",
	StatusForbidden:          "Forbidden",
	StatusNotFound:           "Not Found",
	StatusMethodNotAllowed:   "Method Not Allowed",
	StatusNotAcceptable:      "Not Acceptable",
	StatusRequestTimeout:     "Request Timeout",
	StatusPreconditionFailed: "Precondition Failed",

	StatusInternalServerError:     "Internal Server Error",
	StatusNotImplemented:          "Not Implemented",
	StatusBadGateway:              "Bad Gateway",
	StatusServiceUnavailable:      "Service Unavailable",
	StatusGatewayTimeout:          "Gateway Timeout",
	StatusHTTPVersionNotSupported: "HTTP Version Not Supported",
}

func StatusText(code int) string {
	return statusText[code]
}
