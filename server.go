package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/atsushi-matsui/poke-search/db"
	"github.com/atsushi-matsui/poke-search/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	serverAddress := ":8081"

	r.LoadHTMLGlob("views/*")

	r.GET("/", Index)
	r.POST("/register", Register)
	r.POST("/register/poke/:no", RegisterPoke)
	r.GET("/search", Search)

	r.Run(serverAddress)

	log.Println("starting server at", serverAddress)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func Search(c *gin.Context) {
	fmt.Println("receive GET search")

	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, "not found param")
	}
	// プリプロセッサ実行
	terms := domain.ExecPrePro(keyword)
	res := make(map[int32]*domain.Res)
	for _, term := range terms {
		// Termごとのスコアを計算する
		domain.ScanScore(term, res)
	}
	// ポストプロセッサ実行
	postPro := domain.ExecPostPro(res)

	c.IndentedJSON(http.StatusOK, postPro)
}

func Register(c *gin.Context) {
	fmt.Println("receive POST register")

	// 登録処理
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "cast error")
	}
	term := c.PostForm("term")
	domain.RegisterTerm(int32(id), term)
}

func RegisterPoke(c *gin.Context) {
	fmt.Println("receive POST register poke")

	no, err := strconv.Atoi(c.Param("no"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "cast error")
	}
	no32 := int32(no)
	poke := c.Query("poke")
	if poke == "" {
		c.JSON(http.StatusBadRequest, "not found param")
	}

	pt := db.NewPokeTable()
	pt.AddPoke(no32, poke)
}
