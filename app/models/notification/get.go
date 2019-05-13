package notification

import (
	"gin_bbs/database"
	gintime "gin_bbs/pkg/ginutils/time"
	"encoding/json"
)

// Get -
func Get(id int) (*Notification, error) {
	n := &Notification{}
	if err := database.DB.First(&n, id).Error; err != nil {
		return nil, err
	}

	return n, nil
}

// AllCount -
func AllCount() (count int, err error) {
	err = database.DB.Model(&Notification{}).Count(&count).Error
	return
}

// List -
func List(notifiableType string, notifiableID uint, offset, limit int) ([]interface{}, error) {
	ns := make([]*Notification, 0)
	result := make([]interface{}, 0)

	if err := database.DB.Where("notifiable_type = ? AND notifiable_id = ?",
		notifiableType,
		notifiableID,
	).Offset(offset).Limit(limit).Order("created_at").Find(&ns).Error; err != nil {
		return result, err
	}

	// 整理数据
	for _, n := range ns {
		j := make(map[string]interface{})
		json.Unmarshal([]byte(n.Data), &j)

		d := map[string]interface{}{
			"ID": n.ID,
			"Type": n.Type,
			"NotifiableType": n.NotifiableType,
			"NotifiableID": n.NotifiableID,
			"ReadAt": n.ReadAt,
			"CreatedAt": gintime.SinceForHuman(n.CreatedAt),
			"Data": j,
		}
		
		result = append(result, d)
	}

	return result, nil
}