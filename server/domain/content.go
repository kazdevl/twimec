package domain

type Content struct {
	ContentConfig
	Chapters []Chapter `json:"chapters"`
}

func (c *Content) LoadChapters() error {
	// TODO impl
	return nil
}

type Chapter struct {
	Index int    `json:"name"`
	Icon  string `json:"icon"`
	Pages []int  `json:"pages"`
}
