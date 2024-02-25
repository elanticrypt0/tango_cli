package templates

func (t *Templates) FeatureAPI() string {

	t.setReplacements()

	template := `
package features

import (
	"net/http"
	"strconv"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

const $PL$Pagination = false
const $PL$PaginationItemsPerPage = 15

func FindOne$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, err := $FL$.FindOne(tapp.App.DB.Primary, id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK,$SL$.ConvertToDTO())
}

func FindAll$PC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	var $FL$Buf *[]models.$SC$
	$FL$ := models.New$SC$()

	if $PL$Pagination==true{
		queryPage := ctx.Param("page")
		currentPage:= 0
		if queryPage != "" {
			currentPage, _ = strconv.Atoi(queryPage)
		}
	
		// total de registros en la db
		// counter, _ := c.Count(tapp.App.DB.Primary)
		// pagination := pagination.NewPagination(currentPage,$PL$PaginationItemsPerPage,counter)
	
		$FL$Buf, _ = $FL$.FindAllPagination(tapp.App.DB.Primary, $PL$PaginationItemsPerPage, currentPage)
	}else{
		$FL$Buf, _ = $FL$.FindAll(tapp.App.DB.Primary)
	}

	return ctx.JSON(http.StatusOK,$FL$Buf)

}

func Create$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := ctx.Bind(&$FL$DTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf,err:= $FL$.Create(tapp.App.DB.Primary, $FL$DTO)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, $FL$Buf.ConvertToDTO())
}

func Update$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := ctx.Bind(&$FL$DTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf, err:=$FL$.Update(tapp.App.DB.Primary, id, $FL$DTO)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, $FL$Buf.ConvertToDTO())
}

func Delete$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$Buf,err:=$FL$.Delete(tapp.App.DB.Primary, id)
	
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, $FL$Buf.ConvertToDTO())
}
	`
	return t.Replacements.Replace(template)

}
