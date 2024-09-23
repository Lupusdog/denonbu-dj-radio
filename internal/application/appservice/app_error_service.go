package appservice

import (
	"github.com/Lupusdog/denonbu-dj-radio/pkg/constants/status"
)

type AppError struct {
	Status     string `json:"status"`      // エラーのステータス
	StatusCode int    `json:"status_code"` // エラーのステータスコード
	StatusMsg  string `json:"status_msg"`  // エラーの概要
	Detail     string `json:"-"`           // エラーの詳細(JSONに載せず、ログに出力する)
}

func NewAppError(statusCode status.ResponseStatusCode, originError error) AppError {
	if statusCode < status.InvalidRequest || statusCode > status.GatewayTimeout {
		return AppError{
			string(status.Error),
			int(status.InternalError),
			string(status.ResponseStatusMsg[status.InternalError]),
			"AppError statusCode is invalid",
		}
	}

	if originError == nil {
		return AppError{
			string(status.Error),
			int(status.InternalError),
			string(status.ResponseStatusMsg[statusCode]),
			"AppError originError is nil",
		}
	}

	return AppError{
		string(status.Error),
		int(statusCode),
		string(status.ResponseStatusMsg[statusCode]),
		originError.Error(),
	}
}

func (e *AppError) Error() string {
	return e.Detail
}
