package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo"

	"go-store/templates"
)


func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return Render(ctx, http.StatusOK, templates.Base())
	})

	e.Logger.Fatal(e.Start(":8000"))
}

// INFO: This is a simplified render method that replaces `echo`'s with a custom
// one. This should simplify rendering out of an echo route.
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}