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

	$PL$.GET("", func(c echo.Context) error {
		return features.FindAllCategories(c, tapp)
	})

	$PL$.GET(":id", func(c echo.Context) error {
		return features.FindOne$SC$(c, tapp)
	})

	$PL$.GET("new", func(c echo.Context) error {
		return features.ShowForm$SC$(c, tapp, true)
	})

	$PL$.GET("edit/:id", func(c echo.Context) error {
		return features.ShowForm$SC$(c, tapp, false)
	})

	$PL$.POST("create", func(c echo.Context) error {
		return features.Create$SC$(c, tapp)
	})

	$PL$.POST("update/:id", func(c echo.Context) error {
		return features.Update$SC$(c, tapp)
	})

	$PL$.GET("delete/:id", func(c echo.Context) error {
		return features.Delete$SC$(c, tapp)
	})
}
	`
	return t.Replacements.Replace(template)
}
