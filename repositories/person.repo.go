package repositories

import (
	"github.com/MarcelArt/family_tree/models"
	"gorm.io/gorm"
)

const personPageQuery = `
	select 
		p.id id,
		p."name" "name",
		p.parent_id parent_id,
		f.id father_id,
		f.name father,
		m.id mother_id,
		m."name" mother
	from people p 
	left join parents p2 on p.parent_id = p2.id 
	left join people f on p2.husband_id = f.id 
	left join people m on p2.wife_id = m.id 
`

type IPersonRepo interface {
	IBaseCrudRepo[models.Person, models.PersonDTO, models.PersonPage]
}

type PersonRepo struct {
	BaseCrudRepo[models.Person, models.PersonDTO, models.PersonPage]
}

func NewPersonRepo(db *gorm.DB) *PersonRepo {
	return &PersonRepo{
		BaseCrudRepo: BaseCrudRepo[models.Person, models.PersonDTO, models.PersonPage]{
			db:        db,
			pageQuery: personPageQuery,
		},
	}
}
