package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

type ContentConfig struct {
	AuthorName    string    `json:"author_name"`
	Keyword       string    `json:"keyword"`
	Header        string    `json:"header"`
	Cover         string    `json:"cover"`
	Title         string    `json:"title"`
	Desciption    string    `json:"description"`
	LatestChapter int       `json:"latest_chapter"`
	LatestTime    time.Time `json:"latest_time"`
}

func (c *ContentConfig) LoadConfig(authorName string) error {
	// TODO impl
	path, err := filepath.Abs("./../config/contents")
	if err != nil {
		return err
	}
	data, err := os.ReadFile(fmt.Sprintf("%s/%s.json", path, authorName))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}

func (c *ContentConfig) StoreConfig() error { // 新規作成と更新
	// TODO impl
	return nil
}
