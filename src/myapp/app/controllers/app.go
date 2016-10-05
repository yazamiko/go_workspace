package controllers

import (
	_ "fmt"
	"github.com/revel/revel"
	"log"
	"myapp/app"
	"myapp/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Hello, Gabe!"
	return c.Render(greeting)
}

func (c App) Save() revel.Result {
	var myName string
	var file string

	c.Params.Bind(&myName, "title")
	c.Params.Bind(&file, "file")

	fasta := models.NewFasta(myName, file)
	err := app.Dbmap.Insert(&fasta)

	if err != nil {
		msg := "Insert failed!"
		log.Println(err, msg)
		c.RenderArgs["myName"] = msg
		return c.RenderTemplate("App/savefasta.html")
	}

	c.RenderArgs["myName"] = myName
	return c.RenderTemplate("App/savefasta.html")
}

func (c App) NotFound() revel.Result {
	return c.RenderTemplate("bad_request.html")
}
