package core

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

// MockGateway - MockDBのアダプターの構造体
type MockGateway struct {
	db *MockDB
}

// MockDB - テスト・開発用のDB
type MockDB struct {
	mu   sync.RWMutex
	data map[uint]Recipe
}

// NewMockDB - テスト・開発用のDBのコンストラクタ
func NewMockDB() *MockDB {
	return &MockDB{data: make(map[uint]Recipe)}
}

// NewMockGateway - MockDBのアダプターの構造体のコンストラクタ
func NewMockGateway(db *MockDB) Repository {
	return &MockGateway{db}
}

// CreateRecipe - レシピを作成
func (r *MockGateway) CreateRecipe(ctx context.Context, recipe Recipe) (Recipe, error) {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()

	// 割り当て可能なIDを探す
	var id int64
	for {
		id = src.Int63()
		_, ok := r.db.data[uint(id)]
		if !ok {
			break
		}
	}
	fmt.Println(recipe)
	fmt.Println(id)

	recipe.ID = uint(id)
	r.db.data[recipe.ID] = recipe

	return recipe, nil
}

// ReadRecipes - 全てのレシピを取得
func (r *MockGateway) ReadRecipes(ctx context.Context) ([]Recipe, error) {
	var recipes []Recipe
	for _, v := range r.db.data {
		recipes = append(recipes, v)
	}
	return recipes, nil
}

// ReadRecipe - 指定したIDのレシピを取得
func (r *MockGateway) ReadRecipe(ctx context.Context, id uint) (Recipe, error) {
	recipe := r.db.data[id]
	return recipe, nil
}

// UpdateRecipe - 指定したIDのレシピを更新
func (r *MockGateway) UpdateRecipe(ctx context.Context, id uint, recipe Recipe) (Recipe, error) {
	var preRecipe Recipe
	preRecipe = r.db.data[id]
	preRecipe.Title = recipe.Title
	preRecipe.MakingTime = recipe.MakingTime
	preRecipe.Serves = recipe.Serves
	preRecipe.Ingredients = recipe.Ingredients
	preRecipe.Cost = recipe.Cost
	r.db.data[id] = preRecipe
	return recipe, nil
}

// DeleteRecipe - 指定したIDのレシピを削除
func (r *MockGateway) DeleteRecipe(ctx context.Context, recipe Recipe) (bool, error) {
	delete(r.db.data, recipe.ID)
	return true, nil
}
