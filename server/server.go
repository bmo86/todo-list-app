package server

import (
	"errors"
	"log"
	"net/http"
	"todo-api/cache"
	"todo-api/database"
	repoapp "todo-api/repository/repo-app"
	repocache "todo-api/repository/repo-cache"
	repousr "todo-api/repository/repo-usr"
	"todo-api/websocket"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type Config struct {
	PORT       string
	JWT_SECRET string
	DATABASE   string
	REDIS_URL  string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *gin.Engine
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(config *Config) (*Broker, error) {
	if config.PORT == "" {
		return nil, errors.New("port is required")
	}

	if config.DATABASE == "" {
		return nil, errors.New("database is required")
	}

	if config.JWT_SECRET == "" {
		return nil, errors.New("JWT is required")
	}

	if config.REDIS_URL == "" {
		return nil, errors.New("REDIS id required")
	}

	return &Broker{
		config: config,
		router: gin.New(),
		hub:    websocket.NewHub(),
	}, nil

}

func (b *Broker) Strat(binder func(s Server, r *gin.Engine)) {
	b.router = gin.New()
	binder(b, b.router)

	handler := cors.Default().Handler(b.router)

	repo, err := database.NewConnectionDB(b.config.DATABASE)
	if err != nil {
		log.Fatal(err)
	}

	rdb := cache.NewRedis(b.config.REDIS_URL)
	repocache.SetRepo(rdb)

	go b.Hub().Run()

	repousr.SetRepo(repo)
	repoapp.SetRepoApp(repo)

	log.Println("Loading Server on port ", b.config.PORT)
	if err := http.ListenAndServe(b.config.PORT, handler); err != nil {
		log.Fatal("error in ListandServe : ", err)
	}

}
