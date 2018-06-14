package main

import (
	"github.com/kataras/iris"
)

func firstMiddleware(ctx iris.Context) {
	ctx.Writef("1. This is the first middleware, before any of route handlers \n")
	ctx.Next()
}

func secondMiddleware(ctx iris.Context) {
	ctx.Writef("2. This is the second middleware, before the '/' route handler \n")
	ctx.Next()
}

func thirdMiddleware(ctx iris.Context) {
	ctx.Writef("3. This is the 3rd middleware, after the '/' route handler \n")
	ctx.Next()
}

func lastAlwaysMiddleware(ctx iris.Context) {
	ctx.Writef("4. This is ALWAYS the last Handler \n")
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))

	// with parties:
	myParty := app.Party("/myparty", firstMiddleware)
	myParty.Get("/", secondMiddleware, func(ctx iris.Context) {
		ctx.Writef("Hello from /myparty/ \n")
		ctx.Next() // .Next because we 're using the third middleware after that, and lastAlwaysMiddleware also
	}, thirdMiddleware)


	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/user/{id}", userById)
	app.Get("/profile/{name:string regexp(^[A-Za-z]+)}", profileByUsername)
	app.Run(iris.Addr(":8080"))
}


func profileByUsername(ctx iris.Context) {
	name := ctx.Params().Get("name")

	ctx.Writef("name", name)
}

func userById(ctx iris.Context) {
	userID, err := ctx.Params().GetInt("id")
	if err != nil {
		//
	}
	//ctx.Writef("User %d", userID)
	ctx.ViewData("message", "User")
	ctx.ViewData("userID", userID)
	ctx.View("hello.html")
}
