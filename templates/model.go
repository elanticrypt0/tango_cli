package templates

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
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"

}

type $SC$DTO struct {
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"
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
