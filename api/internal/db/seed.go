package db

import (
	"SportsClubs/internal/models"
	"log"

	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	var townCount int64
	db.Model(&models.Town{}).Count(&townCount)
	if townCount == 0 {
		towns := []models.Town{
			{Name: "Москва"},
			{Name: "Санкт-Петербург"},
			{Name: "Новосибирск"},
			{Name: "Екатеринбург"},
			{Name: "Казань"},
			{Name: "Орел"},
			{Name: "Владивосток"},
			{Name: "Челябинск"},
			{Name: "Краснодар"},
			{Name: "Тюмень"},
			{Name: "Уфа"},
			{Name: "Красноярск"},
			{Name: "Омск"},
			{Name: "Томск"},
			{Name: "Самара"},
			{Name: "Кемерово"},
			{Name: "Нижний Новгород"},
			{Name: "Иркутск"},
			{Name: "Барнаул"},
			{Name: "Волгоград"},
		}
		for _, town := range towns {
			if err := db.Create(&town).Error; err != nil {
				log.Printf("Не удалось заполнить город: %v", err)
			}
		}
		log.Println("Таблица городов заполнена")
	}

	var clubCount int64
	db.Model(&models.Club{}).Count(&clubCount)
	if clubCount == 0 {
		clubs := []models.Club{
			{Name: "Динамо", QuantityTituls: 21, AverageAgePlayers: 22.5, TownID: 1},
			{Name: "Спартак", QuantityTituls: 30, AverageAgePlayers: 24.0, TownID: 1},
			{Name: "Зенит", QuantityTituls: 18, AverageAgePlayers: 21.0, TownID: 2},
			{Name: "Локомотив", QuantityTituls: 19, AverageAgePlayers: 23.7, TownID: 3},
			{Name: "Торпеда", QuantityTituls: 12, AverageAgePlayers: 18.1, TownID: 4},
			{Name: "Ак Барс", QuantityTituls: 14, AverageAgePlayers: 23.2, TownID: 1},
			{Name: "Металург", QuantityTituls: 11, AverageAgePlayers: 29.2, TownID: 10},
			{Name: "Сибирь", QuantityTituls: 17, AverageAgePlayers: 21.2, TownID: 3},
			{Name: "Витязи", QuantityTituls: 18, AverageAgePlayers: 31.1, TownID: 5},
			{Name: "Соколы", QuantityTituls: 21, AverageAgePlayers: 24.4, TownID: 6},
			{Name: "Уралмаш", QuantityTituls: 15, AverageAgePlayers: 26.2, TownID: 8},
			{Name: "Спартак", QuantityTituls: 13, AverageAgePlayers: 23.1, TownID: 7},
			{Name: "Балтика", QuantityTituls: 11, AverageAgePlayers: 25.5, TownID: 18},
			{Name: "Алмаз", QuantityTituls: 22, AverageAgePlayers: 22.0, TownID: 20},
			{Name: "Медведи", QuantityTituls: 6, AverageAgePlayers: 28.1, TownID: 19},
			{Name: "Арсенал", QuantityTituls: 15, AverageAgePlayers: 23.8, TownID: 17},
			{Name: "Титаны", QuantityTituls: 12, AverageAgePlayers: 24.0, TownID: 16},
			{Name: "Метеор", QuantityTituls: 20, AverageAgePlayers: 26.0, TownID: 10},
			{Name: "Юпитер", QuantityTituls: 4, AverageAgePlayers: 22.1, TownID: 12},
			{Name: "Нептун", QuantityTituls: 14, AverageAgePlayers: 20.5, TownID: 9},
			{Name: "Спарта", QuantityTituls: 24, AverageAgePlayers: 20.9, TownID: 1},
			{Name: "Империя", QuantityTituls: 29, AverageAgePlayers: 28.5, TownID: 11},
			{Name: "Вулкан", QuantityTituls: 12, AverageAgePlayers: 21.4, TownID: 2},
			{Name: "Торнадо", QuantityTituls: 32, AverageAgePlayers: 26.6, TownID: 1},
			{Name: "Дрим тим", QuantityTituls: 12, AverageAgePlayers: 22.4, TownID: 3},
			{Name: "Орион", QuantityTituls: 7, AverageAgePlayers: 23.0, TownID: 15},
			{Name: "Арсенал", QuantityTituls: 3, AverageAgePlayers: 18.4, TownID: 12},
			{Name: "Зенит", QuantityTituls: 13, AverageAgePlayers: 29.5, TownID: 14},
			{Name: "Спартак", QuantityTituls: 19, AverageAgePlayers: 26.1, TownID: 9},
			{Name: "ЦСКА", QuantityTituls: 17, AverageAgePlayers: 26.9, TownID: 3},
			{Name: "Викинги", QuantityTituls: 23, AverageAgePlayers: 23.1, TownID: 13},
			{Name: "Энергия", QuantityTituls: 16, AverageAgePlayers: 22.1, TownID: 14},
			{Name: "Легион", QuantityTituls: 11, AverageAgePlayers: 19.9, TownID: 16},
			{Name: "Факел", QuantityTituls: 15, AverageAgePlayers: 20.0, TownID: 11},
			{Name: "Импульс", QuantityTituls: 25, AverageAgePlayers: 18.4, TownID: 9},
			{Name: "Шторм", QuantityTituls: 22, AverageAgePlayers: 19.8, TownID: 8},
			{Name: "Ледяной щит", QuantityTituls: 20, AverageAgePlayers: 20.8, TownID: 10},
			{Name: "Балтика", QuantityTituls: 18, AverageAgePlayers: 24.3, TownID: 6},
			{Name: "Торнадо", QuantityTituls: 14, AverageAgePlayers: 25.3, TownID: 5},
			{Name: "Витязи", QuantityTituls: 19, AverageAgePlayers: 21.9, TownID: 17},
			{Name: "Локомотив", QuantityTituls: 11, AverageAgePlayers: 24.9, TownID: 9},
			{Name: "Дрим тим", QuantityTituls: 31, AverageAgePlayers: 28.1, TownID: 1},
		}
		for _, club := range clubs {
			if err := db.Create(&club).Error; err != nil {
				log.Printf("Не удалось заполнить клуб: %v", err)
			}
		}
		log.Println("Таблица клубов заполнена")
	}
}
