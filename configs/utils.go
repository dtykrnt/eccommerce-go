package configs

import (
	"gorm.io/gorm"
)

type paginate struct {
	limit int
	page  int
}

func Pagination(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) Result(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit
	return db.Offset(offset).Limit(p.limit)
}
