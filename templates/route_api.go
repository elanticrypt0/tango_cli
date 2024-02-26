package templates

func (t *Templates) RouteAPI() string {
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
		f:=features.New$PC$Feature(ctx,tapp)
		return f.FindOne()
	})

	$PL$.GET("", func(ctx echo.Context) error {
		f:=features.New$PC$Feature(ctx,tapp)
		return f.FindAll()
	})

	$PL$.POST("create", func(ctx echo.Context) error {
		f:=features.New$PC$Feature(ctx,tapp)
		return f.Create()
	})

	$PL$.PUT("update/:id", func(ctx echo.Context) error {
		f:=features.New$PC$Feature(ctx,tapp)
		return f.Update()
	})

	$PL$.DELETE("delete/:id", func(ctx echo.Context) error {
		f:=features.New$PC$Feature(ctx,tapp)
		return f.Delete()
	})
}
	`
	return t.Replacements.Replace(template)
}
