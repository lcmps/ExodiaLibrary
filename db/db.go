package db

import (
	"fmt"

	"github.com/lcmps/ExodiaLibrary/app"
	"github.com/lcmps/ExodiaLibrary/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB     *gorm.DB
	Config *app.Config
}

func InitConnection() (Connection, error) {
	var conn Connection
	appData, err := app.InitConfig()
	if err != nil {
		return conn, err
	}
	conn.Config = appData
	connString := fmt.Sprintf(`host=localhost user=%s password=%s dbname=%s sslmode=disable`,
		conn.Config.DB_User,
		conn.Config.DB_Pass,
		conn.Config.DB_Name)
	conn.DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return conn, err
	}
	return conn, nil
}

func (conn *Connection) CreateTables() {
	conn.DB.AutoMigrate(model.Cards{})
}

func (conn *Connection) ImportCards() {
	conn.CreateTables()

	en, fr, pt := app.GetAllCardsLanguages()

	for _, card := range pt.Data {
		conn.DB.Exec(`
		INSERT INTO
			cards
		VALUES
		(?, ?, 'name_en', 'name_fr', ?, ?, 'desc_en', 'desc_fr', ?, ?, ?, ?, ?, ?, ?, ?)`, card.ID, card.Name, card.Type, card.Desc, card.CardImages[0].ID, card.Attribute, card.Race, card.Archetype, card.CardPrices[0].TcgplayerPrice, card.Atk, card.Def, card.Level)
	}

	for _, card := range en.Data {
		conn.DB.Exec(`
		UPDATE cards SET name_en = ?, description_en = ? WHERE id = ?;`, card.Name, card.Desc, card.ID)
	}

	for _, card := range fr.Data {
		conn.DB.Exec(`
		UPDATE cards SET name_fr = ?, description_fr = ? WHERE id = ?;`, card.Name, card.Desc, card.ID)
	}
}
