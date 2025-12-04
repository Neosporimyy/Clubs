package models

type Club struct {
	ID                uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name              string  `json:"name"`
	QuantityTituls    int     `json:"quantity_tituls"`
	AverageAgePlayers float64 `json:"average_age_players"`
	TownID            uint    `json:"town_id"`
	Town              Town    `json:"town" gorm:"foreignKey:TownID"`
}
