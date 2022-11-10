package routes

import (
	"database/sql"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/cmd/server/handler"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/internal/storage"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}
func (r *router) buildSellerRoutes() {
	repo := storage.NewRepository(r.db)
	service := storage.NewService(repo)
	handler := handler.NewStorage(service)
	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:name", handler.Get())
	r.rg.POST("/products", handler.Create())
	r.rg.DELETE("/products/:id", handler.Delete())
	// r.rg.PATCH("/movies/:id", handler.Update())
}
