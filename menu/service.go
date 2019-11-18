package menu

import "github.com/betsegawlemma/webprogmem/entity"

// CategoryService specifies food menu category related services
type CategoryService interface {
	Category(id int) (*entity.Category, error)
	StoreCategory(category *entity.Category) error
}
