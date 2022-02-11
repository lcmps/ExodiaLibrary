package db

import (
	"fmt"
	"reflect"

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

	for _, card := range en.Data {
		conn.DB.Exec(`
		INSERT INTO
			cards
		VALUES
		(?, ?, 'name_pt', 'name_fr', ?, ?, 'desc_pt', 'desc_fr', ?, ?, ?, ?, ?, ?, ?, ?)`, card.ID, card.Name,
			card.Type, card.Desc, card.CardImages[0].ID, card.Attribute, card.Race, card.Archetype,
			card.CardPrices[0].TcgplayerPrice, card.Atk, card.Def, card.Level)
	}

	for _, card := range pt.Data {
		conn.DB.Exec(`
		UPDATE cards SET name_pt = ?, description_pt = ? WHERE id = ? OR name = ?;`, card.Name, card.Desc,
			card.ID, card.NameEn)
	}

	for _, card := range fr.Data {
		conn.DB.Exec(`
		UPDATE cards SET name_fr = ?, description_fr = ? WHERE id = ? OR name = ?;`, card.Name, card.Desc,
			card.ID, card.NameEn)
	}
}

func getQueryMap(mod model.CardQuery) map[string]interface{} {
	w := make(map[string]interface{})

	if mod.Name != "" {
		w["name"] = mod.Name
	}
	if mod.Ctype != "" {
		w["type"] = mod.Ctype
	}
	if mod.Attribute != "" {
		w["attribute"] = mod.Attribute
	}
	if mod.Archetype != "" {
		w["archetype"] = mod.Archetype
	}
	if mod.Race != "" {
		w["race"] = mod.Race
	}
	if mod.Level != 0 {
		w["level"] = mod.Level
	}
	if mod.Atk != 0 {
		w["atk"] = mod.Atk
	}
	if mod.Def != 0 {
		w["def"] = mod.Def
	}
	if mod.Limit != 0 {
		w["limit"] = mod.Limit
	} else {
		w["limit"] = 10
	}
	if mod.Offset != 0 {
		w["offset"] = mod.Offset
	} else {
		w["offset"] = 0
	}

	return w
}

func (conn *Connection) GetCardsByFilter(mod model.CardQuery) []model.Cards {
	var res []model.Cards
	w := getQueryMap(mod)

	tx := conn.DB.Select(`id, name,	name_pt,
	name_fr , "type" , description, description_pt,
	description_fr , image,	"attribute" , race ,
	archetype, price, atk, def, "level"`).Table(`cards`)

	for k, v := range w {

		if k == "limit" {
			tx = tx.Limit(v.(int))
			continue
		}
		if k == "offset" {
			tx = tx.Offset(v.(int))
			continue
		}

		if reflect.TypeOf(v) == reflect.TypeOf("") {
			tx.Where("lower("+k+") LIKE lower(?)", "%"+v.(string)+"%")
		} else {
			tx.Where(k+" = ?", v)
		}

	}
	fmt.Println(tx)

	tx.Find(&res)

	return res
}
