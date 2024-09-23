package music

// MusicRepositoryは、音楽に関する永続化を行うためのリポジトリです。
type MusicRepository interface {
	Save(music *music) error
	FindByID(id string) (*music, error)
	Update(music *music) error
	Delete(id string) error
}
