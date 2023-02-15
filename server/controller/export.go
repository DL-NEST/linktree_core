package controller

import (
	"linktree_core/server/controller/initialize"
	"linktree_core/server/controller/model"
	"linktree_core/server/controller/sysInfo"
	"linktree_core/server/controller/user"
)

var (
	InitializeController initialize.InitializeController
	SysInfoController    sysInfo.SysInfoController
	ModelController      model.ModeController
	UserController       user.UserController
)
