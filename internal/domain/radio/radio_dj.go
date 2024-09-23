package radio

import (
	"errors"

	"github.com/google/uuid"
)

// 文字数ベースの制限を定数で定義(値に深い意味はない。大体の文字数を定義)
const MaxRadioDjNameChars = 20

type radioDj struct {
	id              string
	name            string
	playMusicIdList []string
	prevMusicId     string
	nowMusicId      string
	nextMusicId     string
	isPlaying       bool
}

// ファクトリーメソッド
// NewRadioDjは、指定された名前とプレイリストを持つ新しいradioDjインスタンスを作成します。
// 入力パラメータを検証し、検証に失敗した場合はエラーを返します。
//
// パラメータ:
//   - name: DJの名前。空であってはならず、MaxRadioDjNameChars文字を超えてはなりません。
//   - playList: DJが再生する音楽IDのリスト。空であってはなりません。
func NewRadioDj(name string, playList []string) (*radioDj, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	if len([]rune(name)) > MaxRadioDjNameChars {
		return nil, errors.New("name must not exceed the maximum number of characters")
	}

	if len(playList) == 0 {
		return nil, errors.New("playlist must not be empty")
	}

	// プレイリストが1曲の場合、前の楽曲と次の楽曲は同じになる
	nowMusicId := playList[0]
	prevMusicId := nowMusicId
	nextMusicId := nowMusicId

	// プレイリストが2曲以上の場合、前の楽曲と次の楽曲を設定する
	if len(playList) > 1 {
		nextMusicId = playList[1]
		prevMusicId = playList[len(playList)-1]
	}

	radioDj := &radioDj{
		id:              uuid.New().String(),
		name:            name,
		playMusicIdList: playList,
		prevMusicId:     prevMusicId,
		nowMusicId:      nowMusicId,
		nextMusicId:     nextMusicId,
		isPlaying:       false,
	}

	return radioDj, nil
}

// PlayListに音楽を追加するメソッド
func (r *radioDj) AddMusicIdList(musicId string) error {
	if musicId == "" {
		return errors.New("musicId must not be empty")
	}

	r.playMusicIdList = append(r.playMusicIdList, musicId)

	return nil
}

// 再生中の音楽をスキップして、変更するメソッド
func (r *radioDj) SkipMusic() error {
	if len(r.playMusicIdList) == 0 {
		return errors.New("music list is empty")
	}

	// 現在の楽曲がない場合はエラーを返す
	if r.nowMusicId == "" {
		return errors.New("now playing music is not found")
	}

	// 次の楽曲がない場合はエラーを返す
	if r.nextMusicId == "" {
		return errors.New("next music is not found")
	}

	// 前の楽曲がない場合はエラーを返す
	if r.prevMusicId == "" {
		return errors.New("prev music is not found")
	}

	r.prevMusicId = r.nowMusicId
	r.nowMusicId = r.nextMusicId

	// 前までの次の楽曲の次を次の楽曲にする
	for i, musicId := range r.playMusicIdList {
		if musicId == r.nextMusicId {
			if i == len(r.playMusicIdList)-1 {
				r.nextMusicId = r.playMusicIdList[0]
			} else {
				r.nextMusicId = r.playMusicIdList[i+1]
			}
			break
		}
	}

	return nil
}

// 音楽を再生するメソッド
func (r *radioDj) PlayMusic() error {
	if len(r.playMusicIdList) == 0 {
		return errors.New("music list is empty")
	}

	// 音楽が再生中の場合はエラーを返す
	if r.isPlaying {
		return errors.New("music is already playing")
	}

	// 現在の楽曲がない場合はエラーを返す
	if r.nowMusicId == "" {
		return errors.New("now playing music is not found")
	}

	r.isPlaying = true

	return nil
}

// 音楽を停止するメソッド
func (r *radioDj) PauseMusic() error {
	if len(r.playMusicIdList) == 0 {
		return errors.New("music list is empty")
	}

	// 音楽が再生中でない場合はエラーを返す
	if !r.isPlaying {
		return errors.New("music is already paused")
	}

	// 現在の楽曲がない場合はエラーを返す
	if r.nowMusicId == "" {
		return errors.New("now playing music is not found")
	}

	r.isPlaying = false

	return nil
}

// 　音楽を流すための初期設定を行うメソッド
func (r *radioDj) SetMusicList(musicIdList []string) error {
	if len(musicIdList) == 0 {
		return errors.New("music list is empty")
	}

	r.playMusicIdList = musicIdList
	r.nowMusicId = musicIdList[0]

	if len(musicIdList) > 1 {
		r.nextMusicId = musicIdList[1]
	}

	return nil
}
