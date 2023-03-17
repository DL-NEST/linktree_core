package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"linktree_core/global"
	"linktree_core/modules/emqx"
	"strings"
	"time"
)

const KeyName = "logCache"

func MsgWork() {
	for {
		time.Sleep(1 * time.Second)
		if global.RdGroup.MqMsg == nil {
			continue
		}
		lens := global.RdGroup.MqMsg.LLen(context.Background(), KeyName).Val()
		if lens > 0 {
			dump(lens)
		}
	}
}

func dump(list int64) {
	// mysql的max_allowed_packet
	maxAllowed := 1024
	outStrTemplate := "INSERT INTO `device_msgs` (`created_at`,`updated_at`,`deleted_at`,`device_subject`,`device_msg`) VALUES "

	dumpList := global.RdGroup.MqMsg.LRange(context.Background(), "logCache", 0, list).Val()
	var sqlList = make([]string, 1)
	var i uint = 0
	// 可以提取出来统一转换
	for ii, s := range dumpList {
		// 反射取到的数据
		var msg emqx.MsgType
		err := json.Unmarshal([]byte(s), &msg)
		if err != nil {
			fmt.Printf("erro\n")
			return
		}
		// 添加的插入
		times := time.UnixMilli(int64(msg.Time)).Format("2006-01-02 15:04:05.000")
		newStr := "(" + "'" + times + "','" + times + "',NULL,'" + msg.Topic + "','" + msg.Msg + "')"
		newLen := len(sqlList[i]) + len(outStrTemplate) + len(newStr)
		// 拼接sql语句
		if newLen < maxAllowed {
			sqlList[i] += newStr + ","
		} else {
			sqlList[i] = strings.TrimRight(sqlList[i], ",")
			sqlList = append(sqlList, newStr+",")
			i++
		}
		if ii == len(dumpList)-1 {
			sqlList[i] = strings.TrimRight(sqlList[i], ",")
		}
	}
	// 提交事务
	transaction := global.DB.Begin()
	for _, ss := range sqlList {
		transaction.Exec(outStrTemplate + ss + ";")
	}
	err := transaction.Commit().Error
	if err != nil {
		// 回滚事件
		fmt.Println(transaction.Rollback().Error)
		return
	}
	for i2 := 0; i2 < len(dumpList); i2++ {
		global.RdGroup.MqMsg.RPop(context.Background(), "logCache")
	}
	// 清除redis数据
	//Rdb.LTrim(context.Background(), "logCache", int64(len(dumpList)), 0)
}
