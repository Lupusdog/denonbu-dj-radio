package appservice

import (
	"github.com/Lupusdog/denonbu-dj-radio/pkg/constants/status"
)

type appSuccess struct {
	Status     string      `json:"status"`      // 成功のステータス
	StatusCode int         `json:"status_code"` // 成功のステータスコード
	StatusMsg  string      `json:"status_msg"`  // 成功の概要
	Data       interface{} `json:"data"`        // 成功時のデータ
}

// 成功用の返答の構造体を生成する
func NewAppSuccess(statusCode status.ResponseStatusCode, data interface{}) appSuccess {
	if statusCode < status.OK || statusCode > status.NoContent {
		return appSuccess{
			string(status.Error),
			int(status.InternalError),
			string(status.ResponseStatusMsg[status.InternalError]),
			"appSuccess statusCode is invalid",
		}
	}

	if data == nil {
		return appSuccess{
			string(status.Error),
			int(status.InternalError),
			string(status.ResponseStatusMsg[statusCode]),
			"appSuccess data is nil",
		}
	}

	return appSuccess{
		string(status.Success),
		int(statusCode),
		string(status.ResponseStatusMsg[statusCode]),
		data,
	}
}
