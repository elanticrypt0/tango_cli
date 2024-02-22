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

	$PL$.GET("", func(ctx echo.Context) error {
		return features.FindAll$PC$(ctx, tapp)
	})

	$PL$.GET(":id", func(ctx echo.Context) error {
		return features.FindOne$SC$(ctx, tapp)
	})

	$PL$.POST("create", func(ctx echo.Context) error {
		return features.Create$SC$(ctx, tapp)
	})

	$PL$.PUT("update/:id", func(ctx echo.Context) error {
		return features.Update$SC$(ctx, tapp)
	})

	$PL$.DELETE("delete/:id", func(ctx echo.Context) error {
		return features.Delete$SC$(ctx, tapp)
	})
}
	`
	return t.Replacements.Replace(template)
}
