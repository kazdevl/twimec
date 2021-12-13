package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kazdevl/twimec/domain"
	"github.com/stretchr/testify/assert"
)

func Test_ContentinfoGetAuthorName(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	contentID := "sample"
	repo := NewContentInfoRepository(storagePath)
	name := repo.GetAuthorName(contentID)

	want := "sample"
	assert.Equal(t, want, name)
}

func Test_ContentinfoAll(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	repo := NewContentInfoRepository(storagePath)
	contentinfoList := repo.All()

	wants := []domain.ContentInfo{
		{
			ID:         "sample",
			AuthorName: "sample",
			Header:     "",
			Cover:      "",
			Title:      "",
		},
		{
			ID:         "sample1",
			AuthorName: "sample1",
			Header:     "",
			Cover:      "",
			Title:      "",
		},
	}
	assert.Equal(t, wants, contentinfoList)
}

func Test_ContentinfoStore(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	contentID := "sample3"
	contentDir := filepath.Join(storagePath, contentID)
	err := os.Mkdir(contentDir, 0777)
	assert.NoError(t, err)

	want := domain.ContentInfo{
		ID:         "sample3",
		AuthorName: "sample3",
		Header:     "",
		Cover:      "",
		Title:      "",
	}

	repo := NewContentInfoRepository(storagePath)
	err = repo.Store(want)
	assert.NoError(t, err)

	filename := fmt.Sprintf("%s/info.json", contentDir)
	wantByte, err := os.ReadFile(filename)

	var result domain.ContentInfo
	err = json.Unmarshal(wantByte, &result)
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = os.RemoveAll(filepath.Join(storagePath, contentID))
	assert.NoError(t, err)
}

func Test_ContentinfoUpdate(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	contentID := "sample3"
	contentDir := filepath.Join(storagePath, contentID)
	err := os.Mkdir(contentDir, 0777)
	assert.NoError(t, err)

	origin := domain.ContentInfo{
		ID:         "sample3",
		AuthorName: "sample3",
		Header:     "",
		Cover:      "",
		Title:      "",
	}
	f, err := os.Create(fmt.Sprintf("%s/info.json", contentDir))
	assert.NoError(t, err)

	err = json.NewEncoder(f).Encode(origin)
	f.Close()
	assert.NoError(t, err)

	repo := NewContentInfoRepository(storagePath)
	want := domain.ContentInfo{
		ID:         "sample3",
		AuthorName: "sample4",
		Header:     "",
		Cover:      "",
		Title:      "",
	}
	err = repo.Store(want)
	assert.NoError(t, err)

	filename := fmt.Sprintf("%s/info.json", contentDir)
	wantByte, err := os.ReadFile(filename)

	var result domain.ContentInfo
	err = json.Unmarshal(wantByte, &result)
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = os.RemoveAll(filepath.Join(storagePath, contentID))
	assert.NoError(t, err)
}
