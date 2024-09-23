package radio

import "testing"

func TestCreateDj(t *testing.T) {
	type args struct {
		name     string
		playList []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "DJの名前が空の場合、エラーが返る",
			args: args{
				name:     "",
				playList: []string{"1", "2"},
			},
			wantErr: true,
		},
		{
			name: "DJの名前がMaxRadioDjNameChars文字を超える場合、エラーが返る",
			args: args{
				name:     "123456789012345678901",
				playList: []string{"1", "2"},
			},
			wantErr: true,
		},
		{
			name: "プレイリストが空の場合、エラーが返る",
			args: args{
				name:     "test",
				playList: []string{},
			},
			wantErr: true,
		},
		{
			name: "正常な入力の場合、エラーが返らない",
			args: args{
				name:     "test",
				playList: []string{"1", "2"},
			},
			wantErr: false,
		},
		{
			name: "正常な入力の場合、エラーが返らない",
			args: args{
				name:     "test",
				playList: []string{"1", "2", "3"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRadioDj(tt.args.name, tt.args.playList)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRadioDj() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSkipMusic(t *testing.T) {
	tests := []struct {
		name     string
		playList []string
		wantErr  bool
	}{
		{
			name:     "プレイリストが1曲の場合、エラーが返らない",
			playList: []string{"1"},
			wantErr:  false,
		},
		{
			name:     "プレイリストが複数曲の場合、エラーが返らない",
			playList: []string{"1", "2", "3"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radioDj, err := NewRadioDj("test", tt.playList)
			if err != nil {
				t.Fatalf("NewRadioDj() error = %v", err)
			}
			// 数回スキップしてもエラーが返らないことを確認
			_ = radioDj.SkipMusic()
			_ = radioDj.SkipMusic()
			err = radioDj.SkipMusic()
			if (err != nil) != tt.wantErr {
				t.Errorf("SkipMusic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
