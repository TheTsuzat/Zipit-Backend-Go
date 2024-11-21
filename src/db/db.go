package db

import (
	"github.com/Tsuzat/zipit/src/config"
	"github.com/Tsuzat/zipit/src/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofiber/fiber/v2/log"
)

/*
InitDatabase - Initializes the database connection
*/
func InitDatabase() error {
	opt, err := pg.ParseURL(config.DB_URL)
	if err != nil {
		log.Error("Error parsing database URL: ", err)
		return err
	}
	config.DB = pg.Connect(opt)
	// Create the Schemas
	err = createSchema()
	if err != nil {
		log.Error("Error creating database schema: ", err)
		return err
	}
	return nil
}

func createSchema() error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Url)(nil),
	}
	for _, model := range models {
		err := config.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Error("Error creating table: ", err)
			return err
		}
	}
	return nil
}
