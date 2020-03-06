package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	"github.com/tetsuzawa/voyageapi/containers/backend/cmd/server/controller"
	_ "github.com/tetsuzawa/voyageapi/containers/backend/cmd/server/docs"
	"github.com/tetsuzawa/voyageapi/containers/backend/pkg/env"
	"github.com/tetsuzawa/voyageapi/containers/backend/pkg/mysql"
)

// @title VOYAGE CTO CHALLENGE API
// @version 1.0
// @description This is a recipes API server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host voyageapi.tetsuzawa.com:80
// @BasePath /
func main() {
	e := createMux()
	apiCfg, err := env.ReadAPIEnv()
	if err != nil {
		log.Println(err)
		apiCfg.Host = "127.0.0.1"
		apiCfg.Port = "8080"
	}
	db := newDB()
	ctrls := InitializeControllers(db)
	handler := newHandler(e, ctrls)

	//log.Printf("Listening on %s:%s", apiCfg.Host, apiCfg.Port)
	//e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", apiCfg.Host, apiCfg.Port)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", apiCfg.Host, apiCfg.Port), handler))
}

func newDB() *gorm.DB {
	// Mysql
	mysqlCfg, err := env.ReadMysqlEnv()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := mysql.Connect(mysqlCfg)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	return e
}

func newHandler(e *echo.Echo, ctrls *controller.Controllers) http.Handler {
	recipes := e.Group("/recipes")
	recipes.POST("/", ctrls.Ctrl.HandleCreateRecipe)
	recipes.GET("/", ctrls.Ctrl.HandleReadRecipes)
	recipes.GET("/:id", ctrls.Ctrl.HandleReadRecipe)
	recipes.PATCH("/:id", ctrls.Ctrl.HandleUpdateRecipe)
	recipes.DELETE("/:id", ctrls.Ctrl.HandleDeleteRecipe)
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
