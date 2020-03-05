package core

import (
	"context"
)

// Provider TODO
type Provider struct {
	r Repository
}

// NewProvider TODO
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

// CreateRecipe TODO
func (p *Provider) CreateRecipe(ctx context.Context, title, makingTime, serves, ingredients string, cost int) (Recipe, error) {
	recipe := Recipe{
		Title:       title,
		MakingTime:  makingTime,
		Serves:      serves,
		Ingredients: ingredients,
		Cost:        cost,
	}

	recipe, err := p.r.CreateRecipe(ctx, recipe)
	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func (p *Provider) ReadRecipes(ctx context.Context) ([]Recipe, error) {
	recipes, err := p.r.ReadRecipes(ctx)
	if err != nil {
		return []Recipe{}, err
	}
	return recipes, nil
}

func (p *Provider) ReadRecipe(ctx context.Context, id uint) (Recipe, error) {
	recipe, err := p.r.ReadRecipe(ctx, id)
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}

func (p *Provider) UpdateRecipe(ctx context.Context, id uint, title, makingTime, serves, ingredients string, cost int) (Recipe, error) {
	recipe := Recipe{
		Title:       title,
		MakingTime:  makingTime,
		Serves:      serves,
		Ingredients: ingredients,
		Cost:        cost,
	}
	recipe, err := p.r.UpdateRecipe(ctx, id, recipe)
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}

func (p *Provider) DeleteRecipe(ctx context.Context, id uint) (bool, error) {
	recipe := Recipe{
		ID: id,
	}
	return p.r.DeleteRecipe(ctx, recipe)
}
