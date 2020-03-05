package core

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Gateway TODO
type Gateway struct {
	db *gorm.DB
}

// NewGateway TODO
func NewGateway(db *gorm.DB) Repository {
	return &Gateway{db}
}

//CreateRecipe TODO
func (r *Gateway) CreateRecipe(ctx context.Context, recipe Recipe) (Recipe, error) {
	err := r.db.Create(&recipe).Error
	return recipe, err
}

func (r *Gateway) ReadRecipes(ctx context.Context) ([]Recipe, error) {
	var recipes []Recipe
	err := r.db.Find(&recipes).Error
	return recipes, err
}

func (r *Gateway) ReadRecipe(ctx context.Context, id uint) (Recipe, error) {
	var recipe Recipe
	err := r.db.First(&recipe, id).Error
	return recipe, err
}

func (r *Gateway) UpdateRecipe(ctx context.Context, id uint, recipe Recipe) (Recipe, error) {
	var preRecipe Recipe
	if err := r.db.First(&preRecipe, id).Error; err != nil {
		return Recipe{}, err
	}
	preRecipe.Title = recipe.Title
	preRecipe.MakingTime = recipe.MakingTime
	preRecipe.Serves = recipe.Serves
	preRecipe.Ingredients = recipe.Ingredients
	preRecipe.Cost = recipe.Cost
	err := r.db.Save(&preRecipe).Error
	return recipe, err
}

func (r *Gateway) DeleteRecipe(ctx context.Context, recipe Recipe) (bool, error) {
	if r.db.First(&recipe).RecordNotFound() {
		return false, fmt.Errorf("gateway: record not found")
	}

	err := r.db.Delete(&recipe).Error
	return true, err
}
