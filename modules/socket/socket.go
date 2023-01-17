package socket

import (
	"linktree_core/modules/socket/wsPool"
)

var WsPool = make(wsPool.Pool, 2)
