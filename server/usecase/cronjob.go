package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kazdevl/twimec/domain"
	"github.com/sivchari/gotwtr"
)

type Cronjob struct {
	Client *gotwtr.Client
}

const (
	timeLayout = "2018-11-21T14:24:58.000Z"
)

func NewCronjob(c *gotwtr.Client) *Cronjob {
	return &Cronjob{
		Client: c,
	}
}

// テストがしやすいように関数分けを行う
func (c *Cronjob) FetchContents() {
	// TODO impl
	path, err := filepath.Abs("./../config/contents")
	if err != nil {
		log.Println(err)
		return
	}
	des, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}
	for _, de := range des {
		config := &domain.ContentConfig{}
		fileName := strings.Split(de.Name(), ".")
		config.LoadConfig(fileName[0])
		pagesList, latestTime, err := c.fetchContent(config.AuthorName, config.Keyword, config.LatestTime)
		if err != nil {
			log.Println(err)
			continue
		}

		path, err := filepath.Abs("./../assets")
		if err != nil {
			log.Println(err)
			continue
		}
		f, err := os.OpenFile(fmt.Sprintf("%s/%s", path, config.AuthorName), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()
		for _, pages := range pagesList {
			f.WriteString(strings.Join(pages, " ") + "\n")
		}

		config.LatestTime = latestTime
		config.LatestChapter = config.LatestChapter + len(pagesList)
		if err := config.StoreConfig(); err != nil {
			log.Println(err)
			continue
		}
	}
}

func (c *Cronjob) fetchContent(name, keyword string, previous time.Time) ([]domain.Pages, time.Time, error) {
	// TODO impl
	query := fmt.Sprintf("from:%s -is:retweet \"%s\"", name, keyword)
	res, err := c.Client.SearchRecentTweets(context.Background(), query, &gotwtr.SearchTweetsOption{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionAttachmentsMediaKeys},
		MediaFields: []gotwtr.MediaField{gotwtr.MediaFieldMediaKey, gotwtr.MediaFieldURL},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldAttachments, gotwtr.TweetFieldCreatedAt},
		MaxResults:  100,
		StartTime:   previous,
	})
	if err != nil {
		return nil, previous, err
	}
	if len(res.Tweets) == 0 {
		return nil, previous, errors.New("no tweets")
	}
	pagesList := make([]domain.Pages, len(res.Tweets))
	for index, tweet := range res.Tweets {
		pagesList[index] = getTweetImageLinks(tweet, res.Includes.Media)
	}
	latestTime, _ := time.Parse(timeLayout, res.Tweets[len(res.Tweets)-1].CreatedAt) // 昇順かどうかで変わる
	return pagesList, latestTime, nil
}

func getTweetImageLinks(tweet *gotwtr.Tweet, media []*gotwtr.Media) domain.Pages {
	links := make(domain.Pages, 0, 4)
	for _, key := range tweet.Attachments.MediaKeys {
		for _, m := range media {
			if key == m.MediaKey {
				links = append(links, m.URL)
			}
		}
	}
	return links
}

func (c *Cronjob) notify(mailAddress string) bool {
	// TODO impl
	return true
}
