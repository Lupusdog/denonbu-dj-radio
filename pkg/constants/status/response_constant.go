package status

// このファイルにレスポンス周りの定数を定義する

// レスポンスステータス(success, errorで応答が失敗か成功かを表す)
type ResponseStatus string

const (
	Success ResponseStatus = "success"
	Error   ResponseStatus = "error"
)

// レスポンスステータスコード(レスポンスの概要を表す)
type ResponseStatusCode int

const (
	// Success
	OK        ResponseStatusCode = iota + 1000 // 1000。以降連番
	Created   ResponseStatusCode = iota + 1000
	Accepted  ResponseStatusCode = iota + 1000
	NoContent ResponseStatusCode = iota + 1000

	// Error
	InvalidRequest       ResponseStatusCode = iota + 2000 - 4 // 2000。以降連番
	MissingRequiredParam ResponseStatusCode = iota + 2000 - 4
	InvalidParamValue    ResponseStatusCode = iota + 2000 - 4
	UnauthorizedAccess   ResponseStatusCode = iota + 2000 - 4
	Forbidden            ResponseStatusCode = iota + 2000 - 4
	ResourceNotFound     ResponseStatusCode = iota + 2000 - 4
	Conflict             ResponseStatusCode = iota + 2000 - 4
	UnprocessableEntity  ResponseStatusCode = iota + 2000 - 4
	TooManyRequests      ResponseStatusCode = iota + 2000 - 4
	InternalError        ResponseStatusCode = iota + 2000 - 4
	ServiceUnavailable   ResponseStatusCode = iota + 2000 - 4
	GatewayTimeout       ResponseStatusCode = iota + 2000 - 4
)

// レスポンスステータスメッセージ(レスポンスの概要を表す)
var ResponseStatusMsg = map[ResponseStatusCode]string{
	InvalidRequest:       "invalid request",
	MissingRequiredParam: "missing required parameter",
	InvalidParamValue:    "invalid parameter value",
	UnauthorizedAccess:   "unauthorized access",
	Forbidden:            "forbidden",
	ResourceNotFound:     "resource not found",
	Conflict:             "conflict",
	UnprocessableEntity:  "unprocessable entity",
	TooManyRequests:      "too many requests",
	InternalError:        "internal error",
	ServiceUnavailable:   "service unavailable",
	GatewayTimeout:       "gateway timeout",
}
