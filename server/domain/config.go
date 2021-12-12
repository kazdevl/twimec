package domain

import (
	"time"
)

type Auth struct {
	Key string `json:"key"`
}

func (a *Auth) Use() string {
	// TODO impl
	return ""
}

func (a *Auth) Store(string) error {
	// TODO impl
	return nil
}

type ConfigContentAcquisition struct {
	ContentID     string    `json:"content_id"`
	Keyword       string    `json:"keyword"`
	LatestChapter int       `json:"latest_chapter"`
	LatestTime    time.Time `json:"latest_time"`
}
