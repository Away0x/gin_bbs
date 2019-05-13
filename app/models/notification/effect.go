package notification

import (
	"encoding/json"
	"gin_bbs/database"
	"time"
)

// Create -
func (n *Notification) Create() error {
	return database.DB.Create(&n).Error
}

// Read -
func Read(id int) error {
	n, err := Get(id)
	if err != nil {
		return err
	}

	now := time.Now()
	n.ReadAt = &now
	return database.DB.Save(&n).Error
}

// Notify -
func Notify(typeName string, notifiableType string, notifiableID uint, data map[string]interface{}) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	n := &Notification{
		Type:           typeName,
		NotifiableType: notifiableType,
		NotifiableID:   notifiableID,
		Data:           string(jsonStr),
	}

	return n.Create()
}
