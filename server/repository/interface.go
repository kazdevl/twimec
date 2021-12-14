package repository

import "github.com/kazdevl/twimec/domain"

type ConfigRepository interface {
	Get(contentID string) *domain.ConfigContentAcquisition
	All() []domain.ConfigContentAcquisition
	Store(config domain.ConfigContentAcquisition) error
	Update(config domain.ConfigContentAcquisition) error
}

type ContentInfoRepository interface {
	GetAuthorName(contentID string) string
	All() []domain.ContentInfo
	Store(content domain.ContentInfo) error
	Update(content domain.ContentInfo) error
}

type ChapterRepository interface {
	GetPages(contentID string, index int) domain.Pages
	GetList(contentID string) []domain.Chapter
	Store(chapter domain.Chapter) error
}
