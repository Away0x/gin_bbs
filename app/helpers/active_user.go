package helpers

import (
	"gin_bbs/app/cache"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/database"
	"sort"
	"time"

	"github.com/lexkong/log"
)

var (
	// 缓存相关配置
	activeUserCacheKey   = "ginbbs_active_users"
	cacheExpireInMinutes = 65 * time.Minute
)

// ActiveUser 活跃用户
type ActiveUser struct {
	topicWeight int           // 话题权重
	replyWeight int           // 回复权重
	passDays    time.Duration // 多少天内发表过内容
	userNumber  int           // 取出来多少用户
	tempUsers   map[uint]int  // 存储临时用户数据
}

// NewActiveUser -
func NewActiveUser() *ActiveUser {
	return &ActiveUser{
		topicWeight: 4,
		replyWeight: 1,
		passDays:    7 * 24 * time.Hour,
		userNumber:  6,
		tempUsers:   make(map[uint]int),
	}
}

// GetActiveUsers 获取活跃用户
func (a *ActiveUser) GetActiveUsers() []*userModel.User {
	if cachedData, ok := cache.Get(activeUserCacheKey); ok {
		if result, ok := cachedData.([]*userModel.User); ok {
			return result
		}
	}

	return a.CalculateAndCacheActiveUsers()
}

func (a *ActiveUser) CalculateAndCacheActiveUsers() []*userModel.User {
	// 取得活跃用户列表
	activeUsers := a.calculateActiveUsers()
	// 缓存
	a.cacheActiveUsers(activeUsers)

	return activeUsers
}

// 计算出活跃用户列表
func (a *ActiveUser) calculateActiveUsers() []*userModel.User {
	a.calculateTopicScore()
	a.calculateReplyScore()

	// 数组按得分排序 (倒序，高分靠前)
	scoreUsersIDs := scoreSort(a.tempUsers)
  if len(scoreUsersIDs) > a.userNumber {
	  scoreUsersIDs = scoreUsersIDs[:a.userNumber]
  }

	activeUsers := make([]*userModel.User, 0)
	for _, v := range scoreUsersIDs {
		u, _ := userModel.Get(int(v))
		if u != nil {
			activeUsers = append(activeUsers, u)
		}
	}

	return activeUsers
}

// 将数据放入缓存中
func (a *ActiveUser) cacheActiveUsers(users []*userModel.User) {
	cache.Put(activeUserCacheKey, users, cacheExpireInMinutes)
}

func (a *ActiveUser) calculateTopicScore() {
	// 从话题数据表里取出限定时间范围（a.PassDay）内，有发表过话题的用户
	// 并且同时取出用户此段时间内发布话题的数量
	data := []struct {
		TopicCount int
		UserID     uint
	}{}
	d := database.DB.Raw(`SELECT user_id, count(*) as topic_count from topics
    WHERE created_at >= ?
    GROUP BY user_id`, time.Now().Add(-a.passDays)).Scan(&data)
	if err := d.Error; err != nil {
		log.Infof("active user calculateTopicScore error: %s", err.Error())
		return
	}

	for _, v := range data {
		a.tempUsers[v.UserID] = v.TopicCount * a.topicWeight
	}
}

func (a *ActiveUser) calculateReplyScore() {
	// 从回复数据表里取出限定时间范围（a.PassDay）内，有发表过回复的用户
	// 并且同时取出用户此段时间内发布回复的数量
	data := []struct {
		ReplyCount int
		UserID     uint
	}{}
	d := database.DB.Raw(`SELECT user_id, count(*) as reply_count from replies
    WHERE created_at >= ?
    GROUP BY user_id`, time.Now().Add(-a.passDays)).Scan(&data)
	if err := d.Error; err != nil {
		log.Infof("active user calculateReplyScore error: %s", err.Error())
		return
	}

	for _, v := range data {
		replyScore := v.ReplyCount * a.replyWeight
		if _, ok := a.tempUsers[v.UserID]; ok {
			a.tempUsers[v.UserID] += replyScore
		} else {
			a.tempUsers[v.UserID] = replyScore
		}
	}
}

func scoreSort(mp map[uint]int) []uint {
	newMp := make([]int, 0)
	newMpKey := make([]uint, 0)
	for k, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(newMp)))
	for k, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, k)
	}

	return newMpKey
}
