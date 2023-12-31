package templates

import (
	"strings"
	"tango_cli/parser"
)

// README
// Para reemplazar los archivos tienen una connotación especial
// $[TIPO: Singular | Plural][Capitalized | Lowercase]$
// EJ: $PC$ (Plural Capitalized)

type Templates struct {
	Parser       *parser.Parser
	Replacements *strings.Replacer
}

func New(p *parser.Parser) *Templates {
	return &Templates{
		Parser: p,
	}
}

func (t *Templates) setReplacements() {

	pc := t.Parser.ConvertToTitle(t.Parser.NamePlural)
	pl := t.Parser.NamePlural
	sc := t.Parser.ConvertToTitle(t.Parser.NameSingular)
	sl := t.Parser.NameSingular
	fl := t.Parser.FirstChar

	t.Replacements = strings.NewReplacer("$PC$", pc, "$PL$", pl, "$SC$", sc, "$SL$", sl, "$FL$", fl)

}

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
		return utils.Render(c, views.$PCShowOne(tapp.GetTitleAndVersion(), *$SL$))
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
func (t *Templates) Model() string {

	t.setReplacements()

	template := `
package models

import (
	"fmt"

	"github.com/k23dev/tango/pkg/tango_errors"
	"gorm.io/gorm"
)

type $SC$ struct {
	gorm.Model
	Name string
}

type $SC$DTO struct {
	Name string
}

type $SC$Counter struct {
	Total int
}

func New$SC$() *$SC$ {
	return &$SC${}
}

func (c *$SC$) Count(db *gorm.DB) (int, error) {
	counter := &$SC$Counter{}
	db.Model(&$SC${}).Select("count(ID) as total").Where("delete = ? ", "").Find(&counter)
	return counter.Total, nil
}

func (c *$SC$) FindOne(db *gorm.DB, id int) (*$SC$, error) {
	var $SL$ $SC$
	db.First(&$SL$, id)
	if $SL$.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "$SC$",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &$SL$, nil
}

func (c *$SC$) FindAll(db *gorm.DB) ([]$SC$, error) {
	var $PL$ []$SC$
	db.Order("created_at ASC").Find(&$PL$)
	if len($PL$) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "$SC$",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return $PL$, nil
}

func (c *$SC$) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]$SC$, error) {
	$PL$ := []$SC${}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&$PL$)
	if len($PL$) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "$SC$",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &$PL$, nil
}

func (c *$SC$) Create(db *gorm.DB, name string) (*$SC$, error) {
	$SL$ := $SC${
		Name: name,
	}
	db.Create(&$SL$)
	return &$SL$, nil
}

func (c *$SC$) Update(db *gorm.DB, id int, name string) (*$SC$, error) {
	db.Model(&$SC${}).Where("ID =?", id).Update("name", name)
	return c, nil
}

func (c *$SC$) Delete(db *gorm.DB, id int) (*$SC$, error) {
	$SL$, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&$SL$)
	return $SL$, nil
}

func (c *$SC$) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}	
	`
	return t.Replacements.Replace(template)

}
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
func (t *Templates) Views() string {
	t.setReplacements()

	template := `
package views

import (
    "github.com/k23dev/tango/app/views/layouts"
    "github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/views/components"
	"github.com/k23dev/tango/app/views/forms"
    "github.com/k23dev/tango/pkg/pagination"
)

templ $PC$ShowList(appTitle string,$PL$ []models.$SC$,pagination pagination.Pagination){
    @layouts.Default(appTitle){
        @components.Table($PL$,pagination)
    }

}

templ $PC$ShowOne(appTitle string,$FL$ models.$SC$){
    @layouts.Default(appTitle){
        <h1>$SC$</h1>
        <h2>{$FL$.Name}</h2>
    }
}

templ $PC$FormCreate(appTitle string){
    @layouts.Default(appTitle){
        @forms.$SC$("/$PL$/create","")
    }
}

templ $PC$FormUpdate(appTitle string,$FL$ *models.$SC$){
    @layouts.Default(appTitle){
        @forms.$SC$("/$PL$/update/"+$FL$.GetIDAsString(),$FL$.Name)
    }
}

templ $PC$Delete(appTitle string,$FL$ *models.$SC$){
    @layouts.Default(appTitle){
        <h1>Borrar $SL$</h1>
    }
}
	`
	return t.Replacements.Replace(template)
}
func (t *Templates) ViewsFull() string {
	t.setReplacements()

	template := `
	//
	`
	return t.Replacements.Replace(template)
}
