package tools

import (
	"time"

	"github.com/wmentor/serv"
)

var (
	version int64 = time.Now().Unix()
)

func TemplateVars(c *serv.Context) map[string]interface{} {

	path := c.RequestURI()

	vars := map[string]interface{}{
		"path":        path,
		"requestpath": c.RequestPath(),
		"version":     version,
		"epoch":       time.Now().UnixNano(),
		"title":       "langmod",
	}

	return vars
}
