package core

import "time"

// Recipe - Recipeのモデル. 外部に公開する.
type Recipe struct {
	ID          uint       `json:"id,omitempty"`
	Title       string     `json:"title" validate:"required"`
	MakingTime  string     `json:"making_time" validate:"required"`
	Serves      string     `json:"serves" validate:"required"`
	Ingredients string     `json:"ingredients" validate:"required"`
	Cost        int        `json:"cost,string" validate:"required"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

