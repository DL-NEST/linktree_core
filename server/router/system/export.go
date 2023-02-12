package system

type RouterGroup struct {
	ModelRouter
	SysInfoRouter
	InitializeRouter
}

var SysRouterGroup RouterGroup
