package music

import (
	"errors"

	"github.com/google/uuid"
)

// 文字数ベースの制限を定数で定義(値に深い意味はない。大体の文字数を定義)
const maxMusicTitleChars = 20
const maxArtistNameChars = 20
const maxDiscriptionChars = 100
const maxMusicSourceChars = 100

type music struct {
	musicId          string
	musicTitle       string
	musicArtist      string
	musicDiscription string
	musicSource      string
}

// music構造体のファクトリーメソッド。引数にタイトル、アーティスト、説明を取り、それぞれの文字数が制限を超えていないかチェック
func NewMusic(title string, artist string, discription string, musicSource string) (*music, error) {
	if title == "" || artist == "" || discription == "" || musicSource == "" {
		return nil, errors.New("title, artist, and description must not be empty")
	}

	if len([]rune(title)) > maxMusicTitleChars || len([]rune(artist)) > maxArtistNameChars || len([]rune(discription)) > maxDiscriptionChars || len([]rune(musicSource)) > maxMusicSourceChars {
		return nil, errors.New("title, artist, and discription must not exceed the maximum number of characters")
	}

	return &music{
		musicId:          uuid.New().String(),
		musicTitle:       title,
		musicArtist:      artist,
		musicDiscription: discription,
		musicSource:      musicSource,
	}, nil
}

// タイトルの変更を行うメソッド
func (m *music) ChangeTitle(newTitle string) error {
	if newTitle == "" {
		return errors.New("title must not be empty")
	}

	if len([]rune(newTitle)) > maxMusicTitleChars {
		return errors.New("title must not exceed the maximum number of characters")
	}

	m.musicTitle = newTitle
	return nil
}

// アーティスト名の変更を行うメソッド
func (m *music) ChangeArtist(newArtist string) error {
	if newArtist == "" {
		return errors.New("artist must not be empty")
	}

	if len([]rune(newArtist)) > maxArtistNameChars {
		return errors.New("artist must not exceed the maximum number of characters")
	}

	m.musicArtist = newArtist
	return nil
}

// 説明の変更を行うメソッド
func (m *music) ChangeDiscription(newDiscription string) error {
	if newDiscription == "" {
		return errors.New("discription must not be empty")
	}

	if len([]rune(newDiscription)) > maxDiscriptionChars {
		return errors.New("discription must not exceed the maximum number of characters")
	}

	m.musicDiscription = newDiscription
	return nil
}

// 音楽ソースの変更を行うメソッド
func (m *music) ChangeMusicSource(newMusicSource string) error {
	if newMusicSource == "" {
		return errors.New("music source must not be empty")
	}

	if len([]rune(newMusicSource)) > maxMusicSourceChars {
		return errors.New("music source must not exceed the maximum number of characters")
	}

	m.musicSource = newMusicSource
	return nil
}
