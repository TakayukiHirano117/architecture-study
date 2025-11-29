package categorydm

import (
	"time"
)

type Category struct {
	createdAt time.Time
	updatedAt time.Time
	id        CategoryID
	name      CategoryName
}

func NewCategory(id CategoryID, name CategoryName) (*Category, error) {
	return &Category{
		id:        id,
		name:      name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func NewCategoryByVal(id CategoryID, name CategoryName) (*Category, error) {
	return &Category{
		id:        id,
		name:      name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (c *Category) ID() CategoryID {
	return c.id
}

func (c *Category) Name() CategoryName {
	return c.name
}

func (c *Category) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) UpdatedAt() time.Time {
	return c.updatedAt
}
