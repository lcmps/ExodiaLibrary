package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcmps/ExodiaLibrary/app"
	"github.com/lcmps/ExodiaLibrary/db"
	"github.com/lcmps/ExodiaLibrary/model"
	"github.com/sirupsen/logrus"
)

type WebApp struct {
	Config *app.Config
	Logger *logrus.Logger
}

func New() (*WebApp, error) {
	var App WebApp
	l := logrus.New()

	appData, err := app.InitConfig()
	if err != nil {
		return &App, err
	}

	App.Config = appData
	App.Logger = l

	return &App, nil
}

func (App *WebApp) Host() {
	r := gin.Default()

	r.GET("/", home)

	// Assets
	r.Static("/css", "./pages/assets/css")
	r.Static("/js", "./pages/assets/js")
	r.Static("/card-img", "./pages/assets/card-img")

	// HTML
	r.LoadHTMLGlob("./pages/html/*.html")

	// API
	r.GET("/card", getAllCards)

	gin.SetMode(App.Config.Env)
	err := r.Run(":" + App.Config.Web_Port)
	if err != nil {
		App.Logger.Error(err.Error())
	}
}

func home(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Exodia Library",
		},
	)
}

func getAllCards(ctx *gin.Context) {
	var q model.CardQuery

	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cn, err := db.InitConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	m := cn.GetCardsByFilter(q)
	ctx.JSON(http.StatusOK, m)
}
