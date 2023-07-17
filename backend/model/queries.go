package model

func getAllMinfyers() ([]Minifyr, error) {
	var minifyers []Minifyr

	tx := db.Find(&minifyers)
	if tx.Error != nil {
		return []Minifyr{}, tx.Error
	}

	return minifyers, nil
}

func getMinifyr(id uint64) (Minifyr, error){
	var minifyr Minifyr

	tx := db.Where("id = ?", id).First(&minifyr)
	if tx.Error != nil {
		return Minifyr{}, tx.Error
	}

	return minifyr, nil
}

func createMinifyr(minifyr Minifyr) error {
	tx := db.Create(&minifyr)

	return tx.Error
}

func updateMinifyr(minifyr Minifyr) error {
	tx := db.Save(&minifyr)

	return tx.Error
}

func deleteMinifyr(id uint64) error {
	tx := db.Unscoped().Delete(&Minifyr{}, id)

	return tx.Error
}
