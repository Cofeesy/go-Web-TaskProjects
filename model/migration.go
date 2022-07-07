package model

func migration() {

	//实现自动迁移Task和User表
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Task{})
	DB.Model(&Task{})

	if err != nil {
		return
	}
}
