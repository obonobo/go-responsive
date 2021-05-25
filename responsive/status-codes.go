package responsive

// 1xx
const (
	// This interim response indicates that everything so far is OK and that the
	// client should continue the request, or ignore the response if the request
	// is already finished.
	CONTINUE int = 100

	// This code is sent in response to an Upgrade request header from the
	// client, and indicates the protocol the server is switching to.
	SWITCHING_PROTOCOL int = 101

	// This code indicates that the server has received and is processing the
	// request, but no response is available yet.
	PROCESSING int = 102

	// This status code is primarily intended to be used with the Link header,
	// letting the user agent start preloading resources while the server
	// prepares a response.
	EARLY_HINTS int = 103
)

// 2xx
const (
	OK                            int = 200
	CREATED                       int = 201
	ACCEPTED                      int = 202
	NON_AUTHORITATIVE_INFORMATION int = 203
	NO_CONTENT                    int = 204
	RESET_CONTENT                 int = 205
	PARTIAL_CONTENT               int = 206
	MULTI_STATUS                  int = 207
	ALREADY_REPORTED              int = 208
	IM_USED                       int = 226
)

// 3xx
const (
	MULTIPLE_CHOICES   int = 300
	MOVED_PERMANENTLY  int = 301
	FOUND              int = 302
	SEE_OTHER          int = 303
	NOT_MODIFIED       int = 304
	TEMPORARY_REDIRECT int = 307
	PERMANENT_REDIRECT int = 308
)

// 4xx
const (

	// Client Error
	CLIENT_ERROR                    int = 400
	UNAUTHORIZED                    int = 401
	PAYMENT_REQUIRED                int = 402
	FORBIDDEN                       int = 403
	NOT_FOUND                       int = 404
	METHOD_NOT_ALLOWED              int = 405
	NOT_ACCEPTABLE                  int = 406
	PROXY_AUTHENTICATION_REQUIRED   int = 407
	REQUEST_TIMEOUT                 int = 408
	CONFLICT                        int = 409
	GONE                            int = 410
	LENGTH_REQUIRED                 int = 411
	PRECONDITION_FAILED             int = 412
	PAYLOAD_TOO_LARGE               int = 413
	URI_TOO_LONG                    int = 414
	UNSUPPORTED_MEDIA_TYPE          int = 415
	RANGE_NOT_SATISFIABLE           int = 416
	EXPECTATION_FAILED              int = 417
	IM_A_TEAPOT                     int = 418
	MISDIRECTED_REQUEST             int = 421
	UNPROCESSABLE_ENTITY            int = 422
	LOCKED                          int = 423
	FAILED_DEPENDENCY               int = 424
	TOO_EARLY                       int = 425
	UPGRADE_REQUIRED                int = 426
	PRECONDITION_REQUIRED           int = 428
	TOO_MANY_REQUESTS               int = 429
	REQUEST_HEADER_FIELDS_TOO_LARGE int = 431
	UNAVAILABLE_FOR_LEGAL_REASONS   int = 451
)

// 5xx
const (
	INTERNAL_SERVER_ERROR           int = 500
	NOT_IMPLEMENTED                 int = 501
	BAD_GATEWAY                     int = 502
	SERVICE_UNAVAILABLE             int = 503
	GATEWAY_TIMEOUT                 int = 504
	HTTP_VERSION_NOT_SUPPORTED      int = 505
	VARIANT_ALSO_NEGOTIATES         int = 506
	INSUFFICIENT_STORAGE            int = 507
	LOOP_DETECTED                   int = 508
	NOT_EXTENDED                    int = 510
	NETWORK_AUTHENTICATION_REQUIRED int = 511
)
