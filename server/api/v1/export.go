package v1

import (
	"linktree_core/server/api/v1/model"
	"linktree_core/server/api/v1/system"
)

type apiGroup struct {
	SystemApiGroup system.ApiGroup
	ModelApiGroup  model.ApiGroup
}

var ApiGroup apiGroup
