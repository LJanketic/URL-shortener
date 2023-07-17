package model

func GetAllMinfyrs() ([]Minifyr, error) {
	var minifyrs []Minifyr

	tx := db.Find(&minifyrs)
	if tx.Error != nil {
		return []Minifyr{}, tx.Error
	}

	return minifyrs, nil
}

func GetMinifyr(id uint64) (Minifyr, error){
	var minifyr Minifyr

	tx := db.Where("id = ?", id).First(&minifyr)
	if tx.Error != nil {
		return Minifyr{}, tx.Error
	}

	return minifyr, nil
}

func CreateMinifyr(minifyr Minifyr) error {
	tx := db.Create(&minifyr)

	return tx.Error
}

func UpdateMinifyr(minifyr Minifyr) error {
	tx := db.Save(&minifyr)

	return tx.Error
}

func DeleteMinifyr(id uint64) error {
	tx := db.Unscoped().Delete(&Minifyr{}, id)

	return tx.Error
}

func FindByMinifyrURL(url string) (Minifyr, error) {
	var minifyr Minifyr

	tx := db.Where("minifyr = ?", url).First(&minifyr)

	return minifyr, tx.Error
}
