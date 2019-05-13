package notification

import "gin_bbs/database"

// Get -
func Get(id int) (*Notification, error) {
	n := &Notification{}
	if err := database.DB.First(&n, id).Error; err != nil {
		return nil, err
	}

	return n, nil
}
