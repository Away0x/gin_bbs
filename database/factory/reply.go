package factory

import (
	"fmt"
	"gin_bbs/app/models/reply"
	"gin_bbs/app/models/topic"
	"gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/utils"

	"github.com/Pallinder/go-randomdata"
)

func replyFactory(uids, tids []uint) *reply.Reply {
	paragraph := randomdata.Paragraph()

	ur := utils.RandInt(0, len(uids)-1)
	tr := utils.RandInt(0, len(tids)-1)
	randUID := uids[ur]
	randTID := tids[tr]

	return &reply.Reply{
		Content: paragraph,
		UserID:  randUID,
		TopicID: randTID,
	}
}

// ReplyTableSeeder -
func ReplyTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&reply.Reply{})
	}

	userIDs, _ := user.AllID()
	topicIDs, _ := topic.AllID()

	for i := 0; i < 1000; i++ {
		reply := replyFactory(userIDs, topicIDs)
		if err := reply.Create(); err != nil {
			fmt.Printf("mock reply error： %v\n", err)
		}
	}
}
