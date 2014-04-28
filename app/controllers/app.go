package controllers

import "github.com/revel/revel"

type Application struct {
	GorpController
}

func (c Application) Wildcard() revel.Result {
	return c.RenderJson(struct{}{})
}
