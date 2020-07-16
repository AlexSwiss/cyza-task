package seed

import (
	"log"

	"github.com/AlexSwiss/cyza-task/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Alexander Swiss",
		Email:    "alexyswiss@gmail.com",
		Password: "@kingkunta",
	},
	models.User{
		Nickname: "John Doe",
		Email:    "johndoe@gmail.com",
		Password: "@john123",
	},
}

// Load dummy data
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	//err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	//if err != nil {
	//	log.Fatalf("attaching foreign key error: %v", err)
	//}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
