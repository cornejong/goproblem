package problem

import "net/http"

var (
	// ###################################################
	// #              400 range statuses
	// ###################################################
	BadRequestProblem = &Problem{
		Type:   "https://httpstatuscodes.io/400",
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: "The request could not be understood by the server due to malformed syntax.",
	}
	UnauthorizedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/401",
		Title:  "Unauthorized",
		Status: http.StatusUnauthorized,
		Detail: "The request requires user authentication.",
	}
	PaymentRequiredProblem = &Problem{
		Type:   "https://httpstatuscodes.io/402",
		Title:  "Payment Required",
		Status: http.StatusPaymentRequired,
		Detail: "Payment is required to access the requested resource.",
	}
	ForbiddenProblem = &Problem{
		Type:   "https://httpstatuscodes.io/403",
		Title:  "Forbidden",
		Status: http.StatusForbidden,
		Detail: "You do not have permission to access the requested resource.",
	}
	NotFoundProblem = &Problem{
		Type:   "https://httpstatuscodes.io/404",
		Title:  "Not Found",
		Status: http.StatusNotFound,
		Detail: "The requested resource could not be found.",
	}
	MethodNotAllowedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/405",
		Title:  "Method Not Allowed",
		Status: http.StatusMethodNotAllowed,
		Detail: "The requested method is not allowed for the requested resource.",
	}
	NotAcceptableProblem = &Problem{
		Type:   "https://httpstatuscodes.io/406",
		Title:  "Not Acceptable",
		Status: http.StatusNotAcceptable,
		Detail: "The requested resource is not available in a format acceptable to you.",
	}
	ProxyAuthenticationRequiredProblem = &Problem{
		Type:   "https://httpstatuscodes.io/407",
		Title:  "Proxy Authentication Required",
		Status: http.StatusProxyAuthRequired,
		Detail: "Authentication is required to access the requested resource through a proxy.",
	}
	RequestTimeoutProblem = &Problem{
		Type:   "https://httpstatuscodes.io/408",
		Title:  "Request Timeout",
		Status: http.StatusRequestTimeout,
		Detail: "The server timed out waiting for the request.",
	}
	ConflictProblem = &Problem{
		Type:   "https://httpstatuscodes.io/409",
		Title:  "Conflict",
		Status: http.StatusConflict,
		Detail: "The request could not be completed due to a conflict with the current state of the resource.",
	}
	GoneProblem = &Problem{
		Type:   "https://httpstatuscodes.io/410",
		Title:  "Gone",
		Status: http.StatusGone,
		Detail: "The requested resource is no longer available.",
	}
	LengthRequiredProblem = &Problem{
		Type:   "https://httpstatuscodes.io/411",
		Title:  "Length Required",
		Status: http.StatusLengthRequired,
		Detail: "A required length header was missing in the request.",
	}
	PreconditionFailedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/412",
		Title:  "Precondition Failed",
		Status: http.StatusPreconditionFailed,
		Detail: "The server does not meet one of the preconditions specified in the request.",
	}
	PayloadTooLargeProblem = &Problem{
		Type:   "https://httpstatuscodes.io/413",
		Title:  "Payload Too Large",
		Status: http.StatusRequestEntityTooLarge,
		Detail: "The request payload is too large.",
	}
	URITooLongProblem = &Problem{
		Type:   "https://httpstatuscodes.io/414",
		Title:  "URI Too Long",
		Status: http.StatusRequestURITooLong,
		Detail: "The requested URI is too long.",
	}
	UnsupportedMediaTypeProblem = &Problem{
		Type:   "https://httpstatuscodes.io/415",
		Title:  "Unsupported Media Type",
		Status: http.StatusUnsupportedMediaType,
		Detail: "The request entity has a media type which the server or resource does not support.",
	}
	RangeNotSatisfiableProblem = &Problem{
		Type:   "https://httpstatuscodes.io/416",
		Title:  "Range Not Satisfiable",
		Status: http.StatusRequestedRangeNotSatisfiable,
		Detail: "The request cannot be fulfilled as the range is not satisfiable.",
	}
	ExpectationFailedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/417",
		Title:  "Expectation Failed",
		Status: http.StatusExpectationFailed,
		Detail: "The server could not meet the expectation given in the request.",
	}
	TeapotProblem = &Problem{
		Type:   "https://httpstatuscodes.io/418",
		Title:  "I'm a teapot",
		Status: http.StatusTeapot,
		Detail: "I'm a teapot, not a coffee machine.",
	}
	MisdirectedRequestProblem = &Problem{
		Type:   "https://httpstatuscodes.io/421",
		Title:  "Misdirected Request",
		Status: http.StatusMisdirectedRequest,
		Detail: "The server cannot produce a response for this request.",
	}
	UnprocessableEntityProblem = &Problem{
		Type:   "https://httpstatuscodes.io/422",
		Title:  "Unprocessable Entity",
		Status: http.StatusUnprocessableEntity,
		Detail: "The request was well-formed but could not be followed due to semantic errors.",
	}
	LockedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/423",
		Title:  "Locked",
		Status: http.StatusLocked,
		Detail: "The requested resource is locked.",
	}
	FailedDependencyProblem = &Problem{
		Type:   "https://httpstatuscodes.io/424",
		Title:  "Failed Dependency",
		Status: http.StatusFailedDependency,
		Detail: "The request failed due to a failure of a previous request.",
	}
	TooEarlyProblem = &Problem{
		Type:   "https://httpstatuscodes.io/425",
		Title:  "Too Early",
		Status: http.StatusTooEarly,
		Detail: "The server is unwilling to risk processing a request that might be replayed.",
	}
	UpgradeRequiredProblem = &Problem{
		Type:   "https://httpstatuscodes.io/426",
		Title:  "Upgrade Required",
		Status: http.StatusUpgradeRequired,
		Detail: "The client should switch to a different protocol.",
	}
	PreconditionRequiredProblem = &Problem{
		Type:   "https://httpstatuscodes.io/428",
		Title:  "Precondition Required",
		Status: http.StatusPreconditionRequired,
		Detail: "The server requires a precondition in the request.",
	}
	TooManyRequestsProblem = &Problem{
		Type:   "https://httpstatuscodes.io/429",
		Title:  "Too Many Requests",
		Status: http.StatusTooManyRequests,
		Detail: "You have sent too many requests in a given amount of time.",
	}
	RequestHeaderFieldsTooLargeProblem = &Problem{
		Type:   "https://httpstatuscodes.io/431",
		Title:  "Request Header Fields Too Large",
		Status: http.StatusRequestHeaderFieldsTooLarge,
		Detail: "The server refuses to process the request because the header fields are too large.",
	}
	UnavailableForLegalReasonsProblem = &Problem{
		Type:   "https://httpstatuscodes.io/451",
		Title:  "Unavailable For Legal Reasons",
		Status: http.StatusUnavailableForLegalReasons,
		Detail: "The requested resource is unavailable due to legal reasons.",
	}

	// ###################################################
	// #              500 range statuses
	// ###################################################
	InternalServerErrorProblem = &Problem{
		Type:   "https://httpstatuscodes.io/500",
		Title:  "Internal Server Error",
		Status: http.StatusInternalServerError,
		Detail: "The server encountered an internal error and was unable to complete your request.",
	}
	NotImplementedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/501",
		Title:  "Not Implemented",
		Status: http.StatusNotImplemented,
		Detail: "The server does not support the functionality required to fulfill the request.",
	}
	BadGatewayProblem = &Problem{
		Type:   "https://httpstatuscodes.io/502",
		Title:  "Bad Gateway",
		Status: http.StatusBadGateway,
		Detail: "The server received an invalid response from an upstream server.",
	}
	ServiceUnavailableProblem = &Problem{
		Type:   "https://httpstatuscodes.io/503",
		Title:  "Service Unavailable",
		Status: http.StatusServiceUnavailable,
		Detail: "The server is currently unable to handle the request due to temporary overload or maintenance.",
	}
	GatewayTimeoutProblem = &Problem{
		Type:   "https://httpstatuscodes.io/504",
		Title:  "Gateway Timeout",
		Status: http.StatusGatewayTimeout,
		Detail: "The server did not receive a timely response from an upstream server.",
	}
	HTTPVersionNotSupportedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/505",
		Title:  "HTTP Version Not Supported",
		Status: http.StatusHTTPVersionNotSupported,
		Detail: "The server does not support the HTTP protocol version used in the request.",
	}
	VariantAlsoNegotiatesProblem = &Problem{
		Type:   "https://httpstatuscodes.io/506",
		Title:  "Variant Also Negotiates",
		Status: http.StatusVariantAlsoNegotiates,
		Detail: "Transparent content negotiation for the request results in a circular reference.",
	}
	InsufficientStorageProblem = &Problem{
		Type:   "https://httpstatuscodes.io/507",
		Title:  "Insufficient Storage",
		Status: http.StatusInsufficientStorage,
		Detail: "The server is unable to store the representation needed to complete the request.",
	}
	LoopDetectedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/508",
		Title:  "Loop Detected",
		Status: http.StatusLoopDetected,
		Detail: "The server detected an infinite loop while processing the request.",
	}
	NotExtendedProblem = &Problem{
		Type:   "https://httpstatuscodes.io/510",
		Title:  "Not Extended",
		Status: http.StatusNotExtended,
		Detail: "Further extensions to the request are required for the server to fulfill it.",
	}
	NetworkAuthenticationRequiredProblem = &Problem{
		Type:   "https://httpstatuscodes.io/511",
		Title:  "Network Authentication Required",
		Status: http.StatusNetworkAuthenticationRequired,
		Detail: "Network authentication is required to access the requested resource.",
	}
)
