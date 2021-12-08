package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kazdevl/twimec/domain"
	"github.com/sivchari/gotwtr"
)

type Cronjob struct {
	Client *gotwtr.Client
}

func (c *Cronjob) FetchContents() {
	// TODO impl
}

func (c *Cronjob) fetchContent(name, keyword string, previous time.Time) (pagesList []domain.Pages) {
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
		log.Println(err)
		return
	}
	if len(res.Tweets) == 0 {
		log.Println("no tweets")
		return
	}
	pagesList = make([]domain.Pages, len(res.Tweets))
	for index, tweet := range res.Tweets {
		pagesList[index] = getTweetImageLinks(tweet, res.Includes.Media)
	}
	return pagesList
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
