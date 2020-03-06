package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	_ "github.com/tetsuzawa/voyageapi/containers/backend/cmd/server/docs"
	"github.com/tetsuzawa/voyageapi/containers/backend/pkg/env"
	"github.com/tetsuzawa/voyageapi/containers/backend/pkg/mysql"
)

var e = createMux()

// @title VOYAGE CTO CHALLENGE API
// @version 1.0
// @description This is a recipes API server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host voyageapi.tetsuzawa.com:80
// @BasePath /
func main() {
	apiCfg, err := env.ReadAPIEnv()
	if err != nil {
		log.Println(err)
		apiCfg.Host = "127.0.0.1"
		apiCfg.Port = "8080"
	}

	log.Printf("Listening on %s:%s", apiCfg.Host, apiCfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", apiCfg.Host, apiCfg.Port)))
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Mysql
	mysqlCfg, err := env.ReadMysqlEnv()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := mysql.Connect(mysqlCfg)
	if err != nil {
		log.Fatalln(err)
	}

	// MockDB
	//db := core.NewMockDB()

	ctrls := InitializeControllers(db)

	recipes := e.Group("/recipes")
	recipes.POST("/", ctrls.Ctrl.HandleCreateRecipe)
	recipes.GET("/", ctrls.Ctrl.HandleReadRecipes)
	recipes.GET("/:id", ctrls.Ctrl.HandleReadRecipe)
	recipes.PATCH("/:id", ctrls.Ctrl.HandleUpdateRecipe)
	recipes.DELETE("/:id", ctrls.Ctrl.HandleDeleteRecipe)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	return e
}
