package responsive

import (
	"net/http"

	"github.com/obonobo/go-responsive/responsive/json"
)

var (
	defaultHeaders = map[string]string{
		"Content-Type": "application/json",
	}
)

func Empty() Response {
	return Ok(nil, nil)
}

func Ok(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(OK, body, headers)
}

func Redirect(to string) Response {
	return CreateResponse(SEE_OTHER, nil, map[string]string{
		"Location": to,
	})
}

func BasicMessage(body string) Response {
	return CreateResponse(OK, map[string]interface{}{
		"message": body,
	}, nil)
}

func ClientError(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(CLIENT_ERROR, body, headers)
}

func NotFound() Response {
	return NotFoundWithHeaders(nil)
}

func NotFoundWithHeaders(headers map[string]string) Response {
	return CreateResponse(NOT_FOUND, nil, headers)
}

func UnprocessableEntity() Response {
	return UnprocessableEntityWithHeaders(nil)
}

func UnprocessableEntityWithHeaders(headers map[string]string) Response {
	return CreateResponse(UNPROCESSABLE_ENTITY, nil, headers)
}

func InternalServerError(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(INTERNAL_SERVER_ERROR, body, headers)
}

func ServiceUnavailable(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(SERVICE_UNAVAILABLE, body, headers)
}

func CreateResponse(status int, body map[string]interface{}, headers map[string]string) Response {
	return Response{
		IsBase64Encoded: false,
		StatusCode:      status,
		Headers:         CombineHeaders(defaultHeaders, headers),
		Body:            json.Stringify(body),
	}
}

func Panic(err error) (response *Response, ok bool) {
	if err == nil {
		return nil, true
	}
	errorResponse := ClientError(nil, nil)
	return &errorResponse, false
}

func ResponseIsBad(resp *http.Response) bool {
	return resp.StatusCode > 400
}
