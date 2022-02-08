package model

type Cards struct {
	ID            int    `gorm:"type:integer; primary key; not null" json:"id"`
	Name          string `gorm:"type: text; not null" json:"name"`
	NameEn        string `gorm:"type: text; not null" json:"name_en"`
	NameFr        string `gorm:"type: text; not null" json:"name_fr"`
	Type          string `gorm:"type: text; not null" json:"type"`
	Description   string `gorm:"type: text; not null" json:"description"`
	DescriptionEn string `gorm:"type: text; not null" json:"description_en"`
	DescriptionFr string `gorm:"type: text; not null" json:"description_fr"`
	Image         int    `gorm:"type: integer; not null" json:"image"`
	Attribute     string `gorm:"type: text; not null" json:"attribute"`
	Race          string `gorm:"type: text; not null" json:"race"`
	Archetype     string `gorm:"type: text; not null" json:"archetype"`
	Price         string `gorm:"type: float; not null" json:"price"`
	Atk           int    `gorm:"type: integer; not null" json:"atk"`
	Def           int    `gorm:"type: integer; not null" json:"def"`
	Level         int    `gorm:"type: integer; not null" json:"level"`
}