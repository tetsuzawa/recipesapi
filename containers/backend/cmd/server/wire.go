//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"github.com/tetsuzawa/voyageapi/containers/backend/cmd/server/controller"
	"github.com/tetsuzawa/voyageapi/containers/backend/internal/core"
)

// InitializeControllers - 依存管理. wireでDIする.
func InitializeControllers(db *gorm.DB) *controller.Controllers {
	wire.Build(
		core.NewGateway,
		core.NewProvider,
		controller.NewController,
		controller.NewControllers,
	)
	return &controller.Controllers{}
}

/*
// InitializeControllers - 依存管理. wireでDIする.
// Mock
func InitializeControllers(db *core.MockDB) *controller.Controllers {
	wire.Build(
		core.NewProvider,
		controller.NewController,
		controller.NewControllers,
	)
	return &controller.Controllers{}
}
*/
