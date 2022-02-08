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
	cards, err := app.GetAllCards(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, card := range cards.Data {
		conn.DB.Exec(`
		INSERT INTO
			cards
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, card.ID, card.Name, card.Type, card.Desc, card.CardImages[0].ID, card.Attribute, card.Race, card.Archetype, card.CardPrices[0].TcgplayerPrice, card.Atk, card.Def, card.Level)
	}
}
