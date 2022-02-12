package model

// CardList is a struct that holds query results from user input
type CardQuery struct {
	Name      string `form:"name"`
	Ctype     string `form:"type"`
	Attribute string `form:"attribute"`
	Race      string `form:"race"`
	Archetype string `form:"archetype"`
	Level     int    `form:"level"`
	Atk       int    `form:"atk"`
	Def       int    `form:"def"`
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
}

type CardResponse struct {
	Total int64   `json:"total"`
	Cards []Cards `json:"cards"`
}
