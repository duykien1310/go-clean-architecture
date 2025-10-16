package main

import (
	authHandler "auth_service/api/handler/auth"
	followingHandler "auth_service/api/handler/following"
	postHandler "auth_service/api/handler/post"
	userHandler "auth_service/api/handler/user"
	"fmt"

	"auth_service/config"
	"auth_service/docs"
	"auth_service/infrastucture/repository"
	"auth_service/usecase/auth"
	"auth_service/usecase/following"
	"auth_service/usecase/post"
	"auth_service/usecase/user"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.SetConfigFile("config")
	os.Setenv("TZ", "Etc/GMT")
}

func main() {
	envConfig := getConfig()

	// Database
	db, err := repository.ConnectDatabase(envConfig.Postgres)
	if err != nil {
		log.Println(err)
		return
	}

	// Redis
	redisClient, err := repository.ConnectRedis(envConfig.Redis)
	if err != nil {
		log.Println(err)
		return
	}

	// App
	app := gin.New()
	// Cors
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization", "Time-Offset"},
	})
	app.Use(crs)
	// authHandler.MakeHandlers(app)

	if err != nil {
		log.Println(err)
		return
	}

	// Verifier
	// verifier := jwtUtil.NewVerifier(redisClient)

	// Define Repository
	caching := repository.NewCaching(redisClient)
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	followingRepo := repository.NewFollowingRepository(db)

	// Define Service
	authService := auth.NewService(userRepo)
	userService := user.NewService(userRepo)
	postService := post.NewService(postRepo, userRepo, caching)
	followingService := following.NewService(followingRepo)

	// Handler
	authHandler.MakeHandlers(app, authService)
	userHandler.MakeHandlers(app, userService)
	postHandler.MakeHandlers(app, postService)
	followingHandler.MakeHandlers(app, followingService)

	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run(fmt.Sprintf("%s%s%v", envConfig.Host, ":", envConfig.Port))
}

func getConfig() config.EnvConfig {
	return config.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Postgres: config.PostgresConfig{
			Timeout:  config.GetInt("database.postgres.timeout"),
			DBname:   config.GetString("database.postgres.dbname"),
			Username: config.GetString("database.postgres.user"),
			Password: config.GetString("database.postgres.password"),
			Host:     config.GetString("database.postgres.host"),
			Port:     config.GetString("database.postgres.port"),
		},
		Redis: config.RedisConfig{
			Host: config.GetString("redis.host"),
			Port: config.GetString("redis.port"),
		},
	}
}
