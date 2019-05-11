package factory

import (
	"fmt"
	"gin_bbs/app/models/category"
	"gin_bbs/app/models/topic"
	"gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/utils"

	"github.com/Pallinder/go-randomdata"
)

func topicFactory(uids, cids []uint) *topic.Topic {
	title := randomdata.Country(randomdata.FullCountry)
	paragraph := randomdata.Paragraph()
	excerpt := paragraph
	if len(excerpt) > 20 {
		excerpt = excerpt[:20]
	}
	ur := utils.RandInt(0, len(uids)-1)
	cr := utils.RandInt(0, len(cids)-1)
	randUID := uids[ur]
	randCID := cids[cr]

	return &topic.Topic{
		Title:      title,
		Body:       paragraph,
		Excerpt:    excerpt,
		UserID:     randUID,
		CategoryID: randCID,
	}
}

// TopicTableSeeder -
func TopicTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&topic.Topic{})
	}

	userIDs, _ := user.AllID()
	categoryIDs, _ := category.AllID()

	for i := 0; i < 30; i++ {
		topic := topicFactory(userIDs, categoryIDs)
		if err := topic.Create(); err != nil {
			fmt.Printf("mock topic error： %v\n", err)
		}
	}
}
