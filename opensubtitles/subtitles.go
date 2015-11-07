package opensubtitles

type Subtitles []Subtitle

func (s Subtitles) First() *Subtitle {
	if len(s) == 0 {
		return nil
	}

	return &s[0]
}
