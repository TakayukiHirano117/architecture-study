
package models

import "time"

type CategoryDetailModel struct {
	CategoryID  string `db:"id"`
	CategoryName string `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}