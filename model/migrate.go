package model

func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
