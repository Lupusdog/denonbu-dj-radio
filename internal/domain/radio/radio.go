package radio

import "github.com/google/uuid"

type radio struct {
	id               string
	radioDj          string
	radioStationName string
}

// NewRadioは、指定されたDJ名とラジオ局名を持つ新しいRadioインスタンスを作成します。
func NewRadio(radioDj string, radioStationName string) radio {
	return radio{
		id:               uuid.New().String(),
		radioDj:          radioDj,
		radioStationName: radioStationName,
	}
}
