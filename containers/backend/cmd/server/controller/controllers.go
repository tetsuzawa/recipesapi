package controller

// Controllers ...
type Controllers struct {
	Ctrl *Controller
}

// NewControllers ...
func NewControllers(ctrl *Controller) *Controllers {
	return &Controllers{ctrl}
}
