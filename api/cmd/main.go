package main

import (
	"SportsClubs/internal/db"
	"SportsClubs/internal/handlers"
	"SportsClubs/internal/repositories"
	"SportsClubs/internal/router"
	"SportsClubs/internal/services"
)

func main() {
	database := db.InitDB()
	defer func() {
		pgsql, _ := database.DB()
		pgsql.Close()
	}()

	db.SeedDB(database)

	clubRepo := repositories.NewClubRepository(database)
	clubService := services.NewClubService(clubRepo)
	clubHandler := handlers.NewClubHandler(clubService)

	townRepo := repositories.NewTownRepository(database)
	townService := services.NewTownService(townRepo)
	townHandler := handlers.NewTownHandler(townService)

	r := router.SetupRouter(clubHandler, townHandler)
	r.Run(":8080")
}
