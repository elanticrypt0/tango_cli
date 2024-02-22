package templates

func (t *Templates) FeatureAPI() string {

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


func FindOne$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, _ := $FL$.FindOne(tapp.App.DB.Primary, id)
	return ctx.JSON(http.StatusOK,$SL$.ConvertToDTO)
}

func FindAll$PC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	$FL$ := models.New$SC$()
	counter, _ := $FL$.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	$FL$Buf, _ := $FL$.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return ctx.JSON(http.StatusOK,$FL$Buf)

}

func ShowForm$SC$(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	$FL$ := models.New$SC$()

	if is_new {
		return ctx.JSON(http.StatusOk,interface{})
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		$FL$, _ := $FL$.FindOne(tapp.App.DB.Primary, id)
		return ctx.JSON(http.StatusOK,$FL$Buf)
	}
}

func Create$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := ctx.Bind(&$FL$DTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, interface{})
	}

	$FL$ := models.New$SC$()
	$FL$.Create(tapp.App.DB.Primary, $FL$DTO)

	return ctx.JSON(http.StatusOK, $FL$.ConvertToDTO)
}

func Update$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := ctx.Bind(&$FL$DTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, interface{})
	}

	$FL$ := models.New$SC$()
	$FL$.Update(tapp.App.DB.Primary, id, $FL$DTO)

	return ctx.JSON(http.StatusOK, $FL$DTO)
}

func Delete$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$.Delete(tapp.App.DB.Primary, id)

	return ctx.JSON(http.StatusOK, $FL$.ConvertToDTO)
}
	`
	return t.Replacements.Replace(template)

}
