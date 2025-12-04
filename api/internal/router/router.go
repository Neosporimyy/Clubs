package router

import (
	"SportsClubs/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(clubHandler *handlers.ClubHandler, townHandler *handlers.TownHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	api := r.Group("/api")
	{
		clubs := api.Group("/clubs")
		{
			clubs.GET("", clubHandler.GetAllClubs)
			clubs.GET("/search", clubHandler.SearchClubs)
			clubs.GET("/:id", clubHandler.GetClubByID)
			clubs.POST("", clubHandler.CreateClub)
			clubs.PUT("/:id", clubHandler.UpdateClub)
			clubs.DELETE("/:id", clubHandler.DeleteClub)
		}

		towns := api.Group("/towns")
		{
			towns.GET("", townHandler.GetTowns)
			towns.GET("/search", townHandler.GetClubsByTownName)
			towns.GET("/:id", townHandler.GetTownByID)
			towns.POST("", townHandler.CreateTown)
			towns.PUT("/:id", townHandler.UpdateTown)
			towns.DELETE("/:id", townHandler.DeleteTown)
		}
	}

	return r
}
