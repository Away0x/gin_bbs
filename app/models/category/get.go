package category

import (
	"gin_bbs/database"
)

// Get -
func Get(id int) (*Category, error) {
	cat := &Category{}
	if err := database.DB.First(&cat, id).Error; err != nil {
		return cat, nil
	}

	return cat, nil
}

// All -
func All() (cats []*Category, err error) {
	cats = make([]*Category, 0)

	if err := database.DB.Order("id").Find(&cats).Error; err != nil {
		return cats, err
	}

	return cats, nil
}

// AllID -
func AllID() (ids []uint, err error) {
	ids = make([]uint, 0)
	cats, err := All()
	if err != nil {
		return ids, err
	}

	for _, u := range cats {
		ids = append(ids, u.ID)
	}

	return ids, nil
}

// AllName -
func AllName() (names []string, err error) {
	names = make([]string, 0)
	cats, err := All()
	if err != nil {
		return names, err
	}

	for _, u := range cats {
		names = append(names, u.Name)
	}

	return names, nil
}