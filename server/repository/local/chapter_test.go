package local

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kazdevl/twimec/domain"
	"github.com/stretchr/testify/assert"
)

func Test__ChapterGetPages(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata")
	contentID := "sample"
	index := 0
	repo := NewChapterRepository(storagePath)
	pages := repo.GetPages(contentID, index)
	want := domain.Pages([]string{"sample", "sample", "sample"})
	assert.Equal(t, want, pages)
}
func Test__ChapterGetList(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata")
	contentID := "sample"
	repo := NewChapterRepository(storagePath)
	chapters := repo.GetList(contentID)
	want := []domain.Chapter{
		{
			ContentID: contentID,
			Index:     0,
			Icon:      "sample",
			Pages:     "sample,sample,sample",
		},
		{
			ContentID: contentID,
			Index:     1,
			Icon:      "sample",
			Pages:     "sample,sample,sample",
		},
	}
	assert.Equal(t, want, chapters)
}

func Test_ChapterStore(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata")
	contentID := "sample"
	pages := "sample,sample,sample,sample"
	repo := NewChapterRepository(storagePath)
	err := repo.Store(domain.Chapter{ContentID: contentID, Pages: pages})
	assert.NoError(t, err)

	filename := fmt.Sprintf("%s/2.txt", filepath.Join(storagePath, contentID))
	wantByte, err := os.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, string(wantByte), pages)

	err = os.Remove(filename)
	assert.NoError(t, err)
}
