package templates

func (t *Templates) Route() string {
	t.setReplacements()

	template := `
package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func $PL$Routes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	$PL$ := rootPath.Group("/$PL$/")

	$PL$.GET(":id", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return f.FindOne()
	})

	$PL$.GET("", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return f.FindAll()
	})

	$PL$.GET("new", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return f.ShowForm(, true)
	})

	$PL$.GET("edit/:id", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return f.ShowForm(, false)
	})

	$PL$.POST("create", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return featufres.Create()
	})

	$PL$.POST("update/:id", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return f.Update()
	})

	$PL$.GET("delete/:id", func(ctx echo.Context) error {
		f:=features.New$SC$Feature(ctx,tapp)
		return f.Delete()
	})
}
	`
	return t.Replacements.Replace(template)
}
