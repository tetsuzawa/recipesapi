package core

import (
	"context"
)

// Provider - アプリケーションコアの構造体
type Provider struct {
	r Repository
}

// NewProvider - アプリケーションコアの構造体のコンストラクタ
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

// CreateRecipe - レシピを作成
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

// ReadRecipes - 全てのレシピを取得
func (p *Provider) ReadRecipes(ctx context.Context) ([]Recipe, error) {
	recipes, err := p.r.ReadRecipes(ctx)
	if err != nil {
		return []Recipe{}, err
	}
	return recipes, nil
}

// ReadRecipe - 指定したIDのレシピを取得
func (p *Provider) ReadRecipe(ctx context.Context, id uint) (Recipe, error) {
	recipe, err := p.r.ReadRecipe(ctx, id)
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}

// UpdateRecipe - 指定したIDのレシピを更新
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

// DeleteRecipe - 指定したIDのレシピを削除
func (p *Provider) DeleteRecipe(ctx context.Context, id uint) (bool, error) {
	recipe := Recipe{
		ID: id,
	}
	return p.r.DeleteRecipe(ctx, recipe)
}
