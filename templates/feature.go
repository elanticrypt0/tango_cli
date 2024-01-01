package templates

func (t *Templates) Feature() string {

	t.setReplacements()

	template := `
package features

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/views"
	"github.com/k23dev/tango/pkg/pagination"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

var itemsPerPage = 15

func FindOne$SC$(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, _ := $FL$.FindOne(tapp.App.DB.Primary, id)
	if $SL$ != nil {
		return utils.Render(c, views.$PC$ShowOne(tapp.GetTitleAndVersion(), *$SL$))
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAll$PC$(c echo.Context, tapp *webcore.TangoApp) error {
	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	$FL$ := models.New$SC$()
	counter, _ := $FL$.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	$FL$, _ := $FL$.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return utils.Render(c, views.$PCShowList(tapp.GetTitleAndVersion(), *$FL$, *pagination))
}

func ShowForm$SC$(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	$FL$ := models.New$SC$()

	if is_new {
		return utils.Render(c, views.$PCFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		$FL$, _ := $FL$.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.$PCFormUpdate(tapp.GetTitleAndVersion(), $FL$))
	}
}

func Create$SC$(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := c.Bind(&$FL$DTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	$FL$ := models.New$SC$()
	$FL$.Create(tapp.App.DB.Primary, $FL$DTO.Name)

	return c.Redirect(http.StatusMovedPermanently, "/$PL$/")
}

func Update$SC$(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := c.Bind(&$FL$DTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	$FL$ := models.New$SC$()
	$FL$.Name = strings.ToLower($FL$DTO.Name)

	$FL$.Update(tapp.App.DB.Primary, id, $FL$.Name)

	return c.Redirect(http.StatusMovedPermanently, "/$PL$/")
}

func Delete$SC$(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	$FL$ := models.New$SC$()
	$FL$.Delete(tapp.App.DB.Primary, id)

	return c.Redirect(http.StatusMovedPermanently, "/$PL$/")
}
	`
	return t.Replacements.Replace(template)

}
