package handlers

import "github.com/kataras/iris"

type Health struct {
	StatusCode int
	Message    string
}

func NewHealth() *Health {
	return &Health{
		StatusCode: 200,
		Message:    "ok",
	}
}

func (h *Health) Get(ctx iris.Context) {
	ctx.StatusCode(h.StatusCode)
	ctx.WriteString(h.Message)
}
