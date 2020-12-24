package controller

import (
	"strings"

	"github.com/wmentor/langmod/model"
	"github.com/wmentor/langmod/tools"
	"github.com/wmentor/serv"
)

func init() {
	serv.Register("GET", "/generate", generate)
	serv.Register("POST", "/generate", generate)
	serv.Register("GET", "/loader", loader)
	serv.Register("POST", "/loader", loader)
}

func generate(c *serv.Context) {

	vars := tools.TemplateVars(c)

	vars["input"] = ""
	vars["output"] = ""

	if c.Method() == "POST" {
		val := c.FormValue("data")
		list := model.Generate(strings.NewReader(val))
		vars["input"] = val
		vars["output"] = strings.Join(list, " ")
	}

	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("generate.jet", vars)
}

func loader(c *serv.Context) {

	vars := tools.TemplateVars(c)

	if c.Method() == "POST" {
		val := c.FormValue("data")
		go model.Load(strings.NewReader(val))
	}

	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.WriteHeader(200)
	c.Render("load.jet", vars)
}
