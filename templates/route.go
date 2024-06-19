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

func $PL$Routes(tapp *tangoapp.TangoApp, rootPath *echo.Group) {
	$PL$ := rootPath.Group("/$PL$/")
	feat:=features.New$PC$Feature(tapp)

	$PL$.GET(":id", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return feat.FindOne()
	})

	$PL$.GET("", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return feat.FindAll()
	})

	$PL$.GET("new", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return feat.ShowForm(, true)
	})

	$PL$.GET("edit/:id", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return feat.ShowForm(, false)
	})

	$PL$.POST("create", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return featufres.Create()
	})

	$PL$.POST("update/:id", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return feat.Update()
	})

	$PL$.GET("delete/:id", func(ctx echo.Context) error {
		feat.setCtx(ctx)
		return feat.Delete()
	})
}
	`
	return t.Replacements.Replace(template)
}
