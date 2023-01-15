package socket

import (
	"linktree_server/modules/socket/wsPool"
)

var WsPool = make(wsPool.Pool, 2)
