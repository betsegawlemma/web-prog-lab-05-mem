package service

import (
	"errors"

	"github.com/betsegawlemma/webprogmem/entity"
)

// CategoryCache provides an in-memory cache
type CategoryCache map[int]*entity.Category

// NewCategoryCache returns a new category cache
func NewCategoryCache() CategoryCache {
	return make(map[int]*entity.Category)
}

// Category returns a category for a given id from the cache
func (c CategoryCache) Category(id int) (*entity.Category, error) {
	if cat, ok := c[id]; ok {
		return cat, nil
	}
	return nil, errors.New("Category was not found")
}

// StoreCategory stores category data to the cache
func (c CategoryCache) StoreCategory(category *entity.Category) error {
	if _, ok := c[category.ID]; !ok {
		c[category.ID] = category
		return nil
	}
	return errors.New("Category already exists")
}
