package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	godotenv.Load(".env")
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_DATABASE")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(os.Getenv("DB_CONNECTION"), URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Register all models
func Config() {
	db := SetupDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(About{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(ArticleComment{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(ArticleCategory{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Article{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(CategoryArticle{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(CategoryFaq{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(CategoryProduct{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(CategoryProject{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Contact{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Content{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Country{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(EmailVerifiaction{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Faq{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Feature{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Message{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Notification{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Product{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Project{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Team{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(User{})
}
