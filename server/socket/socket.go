package socket

import (
	"linktree_server/server/socket/wsPool"
)

var WsPool = make(wsPool.Pool, 2)
