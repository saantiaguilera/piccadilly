package random

import (
	"fmt"
	"github.com/juanimoli/piccadilly/api/controller"
	"github.com/juanimoli/piccadilly/api/http"
)

func CreatePostController() controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/slack/random",
		Body:   CreatePostBody(),
	}
}

func CreatePostBody() http.Handler {
	return func(ctx *http.Context) {
		fmt.Println(ctx.GetHeaders())
		var body map[string]interface{}
		_ = ctx.ReadBody(&body)
		fmt.Println(body)
	}
}