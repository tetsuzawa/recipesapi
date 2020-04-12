package controller

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"

	"github.com/tetsuzawa/recipesapi/containers/backend/internal/core"
)

// Controller - リクエストを処理しアプリケーションコアに渡す
type Controller struct {
	p *core.Provider
}

// NewController - Controllerのコンストラクタ
func NewController(p *core.Provider) *Controller {
	return &Controller{p}
}

// Response - Controllerのレスポンスを定義した構造体
type Response struct {
	Message  string        `json:"message,omitempty"`
	Required string        `json:"required,omitempty"`
	Recipe   *core.Recipe  `json:"recipe,omitempty"`
	Recipes  []core.Recipe `json:"recipes,omitempty"`
}

// HandleCreateRecipes - レシピを作成する.
// @Summary レシピを作成する
// @Description title, making_tike, serves, ingredients, costからレシピを作成する
// @Accept json
// @Produce json
// @Param title query string true "Title"
// @Param making_time query string true "Making Time"
// @Param serves query string true "Serves"
// @Param ingredients query string true "Ingredients"
// @Param cost query string true "Cost"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /recipes/ [post]
func (ctrl *Controller) HandleCreateRecipe(c echo.Context) error {

	resp := Response{
		Message:  "Recipe creation failed!",
		Required: "title, making_time, serves, ingredients, cost",
	}

	var in core.Recipe
	if err := c.Bind(&in); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err := core.Validate.Struct(&in)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	ctx := c.Request().Context()
	recipe, err := ctrl.p.CreateRecipe(ctx, in.Title, in.MakingTime, in.Serves, in.Ingredients, in.Cost)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp = Response{
		Message: "Recipe successfully created!",
		Recipe: &core.Recipe{
			Title:       recipe.Title,
			MakingTime:  recipe.MakingTime,
			Serves:      recipe.Serves,
			Ingredients: recipe.Ingredients,
			Cost:        recipe.Cost,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

// HandleReadRecipes - レシピを全て取得.
// @Summary レシピを全て取得
// @Description レシピを全て取得し、配列にする
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /recipes/ [get]
func (ctrl *Controller) HandleReadRecipes(c echo.Context) error {
	ctx := c.Request().Context()
	recipes, err := ctrl.p.ReadRecipes(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}
	resp := Response{
		Recipes: []core.Recipe{},
	}
	for _, recipe := range recipes {
		resp.Recipes = append(resp.Recipes, core.Recipe{
			ID:          recipe.ID,
			Title:       recipe.Title,
			MakingTime:  recipe.MakingTime,
			Serves:      recipe.Serves,
			Ingredients: recipe.Ingredients,
			Cost:        recipe.Cost,
		})
	}
	return c.JSON(http.StatusOK, resp)
}

// HandleReadRecipe - 指定したIDのレシピを取得.
// @Summary 指定したIDのレシピを取得
// @Description 指定したIDのレシピを取得
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /recipes/{id} [get]
func (ctrl *Controller) HandleReadRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	ctx := c.Request().Context()
	recipe, err := ctrl.p.ReadRecipe(ctx, uint(id))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, echo.ErrNotFound)
	}
	res := struct {
		Message string      `json:"message"`
		Recipe  core.Recipe `json:"recipe"`
	}{
		Message: "Recipe details by id",
		Recipe: core.Recipe{
			Title:       recipe.Title,
			MakingTime:  recipe.MakingTime,
			Serves:      recipe.Serves,
			Ingredients: recipe.Ingredients,
			Cost:        recipe.Cost,
		},
	}
	return c.JSON(http.StatusOK, res)
}

// HandleUpdateRecipe - 指定したIDのレシピを更新.
// @Summary 指定したIDのレシピを更新
// @Description 指定したIDのレシピを更新
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Param title query string true "Title"
// @Param making_time query string true "Making Time"
// @Param serves query string true "Serves"
// @Param ingredients query string true "Ingredients"
// @Param cost query string true "Cost"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /recipes/{id} [patch]
func (ctrl *Controller) HandleUpdateRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	var in core.Recipe
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}
	err = core.Validate.Struct(&in)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	ctx := c.Request().Context()
	recipe, err := ctrl.p.UpdateRecipe(ctx, uint(id), in.Title, in.MakingTime, in.Serves, in.Ingredients, in.Cost)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}
	resp := Response{
		Message: "Recipe successfully updated!",
		Recipe: &core.Recipe{
			Title:       recipe.Title,
			MakingTime:  recipe.MakingTime,
			Serves:      recipe.Serves,
			Ingredients: recipe.Ingredients,
			Cost:        recipe.Cost,
		},
	}
	return c.JSON(http.StatusOK, resp)

}

// HandleUpdateRecipe - 指定したIDのレシピを削除.
// @Summary 指定したIDのレシピを削除
// @Description 指定したIDのレシピを削除
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /recipes/{id} [delete]
func (ctrl *Controller) HandleDeleteRecipe(c echo.Context) error {
	sid := c.Param("id")
	if sid == "" {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}
	ctx := c.Request().Context()
	existed, err := ctrl.p.DeleteRecipe(ctx, uint(id))
	if !existed {
		resp := Response{Message: "No Recipe found"}
		return c.JSON(http.StatusNotFound, resp)
	}
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}
	resp := Response{Message: "Recipe successfully removed!"}
	return c.JSON(http.StatusOK, resp)
}
