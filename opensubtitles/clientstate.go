package opensubtitles

type ClientState interface {
	Authenticate(user, password string) (ClientState, error)
	Search(hash string, languages []string) (Subtitles, error)
	Download(subtitle Subtitle) (SubtitleFile, error)
}
