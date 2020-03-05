package core

import "context"

// Repository todo
type Repository interface {
	CreateRecipe(ctx context.Context, recipe Recipe) (Recipe, error)
	ReadRecipes(ctx context.Context) ([]Recipe, error)
	ReadRecipe(ctx context.Context, id uint) (Recipe, error)
	UpdateRecipe(ctx context.Context, id uint, recipe Recipe) (Recipe, error)
	DeleteRecipe(ctx context.Context, recipe Recipe) (bool, error)
}
