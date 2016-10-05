package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"memory_game/app"
	"memory_game/app/models"
	"strconv"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Save() revel.Result {
	var score string
	var nickname string

	c.Params.Bind(&score, "score")
	c.Params.Bind(&nickname, "nickname")

	value, _ := strconv.Atoi(score)
	s := models.NewScore(int64(value), nickname)

	if err := app.Dbmap.Insert(&s); err != nil {
		c.RenderError(err)
	}

	var scores []models.Score
	_, err := app.Dbmap.Select(&scores, "select value, nickname from scores order by Value desc limit 5")

	fmt.Println(scores)
	if err != nil {
		fmt.Println(err)
		c.RenderError(err)
	}

	return c.RenderJson(scores)
}

func (c App) Scoreboard() revel.Result {
	var scores []models.Score
	_, err := app.Dbmap.Select(&scores, "select value, nickname, inserted from scores order by Value desc limit 5")

	if err != nil {
		fmt.Println(err)
		c.RenderError(err)
	}

	return c.RenderJson(scores)
}
