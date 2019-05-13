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

// Read -
func Read(notifiableType string, notifiableID uint) error {
	now := time.Now()

	return database.DB.Model(&Notification{}).Where("notifiable_type = ? AND notifiable_id = ?",
		notifiableType,
		notifiableID,
	).Updates(Notification{ReadAt: &now}).Error
}