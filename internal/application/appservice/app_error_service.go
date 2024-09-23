package appservice

import (
	"github.com/Lupusdog/denonbu-dj-radio/pkg/constants/status"
)

type appError struct {
	Status     string `json:"status"`      // エラーのステータス
	StatusCode int    `json:"status_code"` // エラーのステータスコード
	StatusMsg  string `json:"status_msg"`  // エラーの概要
	Detail     string `json:"-"`           // エラーの詳細(JSONに載せず、ログに出力する)
}

func NewAppError(statusCode status.ResponseStatusCode, originError error) appError {
	if statusCode < status.InvalidRequest || statusCode > status.GatewayTimeout {
		return appError{
			string(status.Error),
			int(status.InternalError),
			string(status.ResponseStatusMsg[status.InternalError]),
			"appError statusCode is invalid",
		}
	}

	if originError == nil {
		return appError{
			string(status.Error),
			int(status.InternalError),
			string(status.ResponseStatusMsg[statusCode]),
			"appError originError is nil",
		}
	}

	return appError{
		string(status.Error),
		int(statusCode),
		string(status.ResponseStatusMsg[statusCode]),
		originError.Error(),
	}
}

func (e *appError) Error() string {
	return e.Detail
}
