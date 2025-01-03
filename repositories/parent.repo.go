package repositories

import (
	"github.com/MarcelArt/family_tree/models"
	"gorm.io/gorm"
)

const parentPageQuery = `
	select 
		p.id id,
		p.husband_id husband_id,
		h.name husband,
		p.wife_id wife_id,
		w.name wife
	from parents p
	join people h on p.husband_id = h.id
	join people w on p.wife_id = w.id
`

type IParentRepo interface {
	IBaseCrudRepo[models.Parent, models.ParentDTO, models.ParentPage]
}

type ParentRepo struct {
	BaseCrudRepo[models.Parent, models.ParentDTO, models.ParentPage]
}

func NewParentRepo(db *gorm.DB) *ParentRepo {
	return &ParentRepo{
		BaseCrudRepo: BaseCrudRepo[models.Parent, models.ParentDTO, models.ParentPage]{
			db:        db,
			pageQuery: parentPageQuery,
		},
	}
}
