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

	// Assets
	r.Static("/css", "./pages/assets/css")
	r.Static("/js", "./pages/assets/js")
	r.Static("/fvc", "./pages/assets/fvc")
	r.Static("/img", "./pages/assets/img")
	r.Static("/card-img", "./pages/card-img")
	r.Static("/fonts", "./pages/assets/fonts")

	// API
	r.GET("/card", getAllCards)
	r.GET("/random", GetRandomCards)

	// Pages
	r.GET("/", home)

	// HTML
	r.LoadHTMLGlob("./pages/html/*.html")

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
			"title": "𓂀 Exodia Library 𓂀",
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

func GetRandomCards(ctx *gin.Context) {
	cn, err := db.InitConnection()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := cn.GetTenRandomCards()
	ctx.JSON(http.StatusOK, m)
}
