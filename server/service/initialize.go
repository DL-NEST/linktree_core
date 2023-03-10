package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"linktree_core/global"
	"linktree_core/server/model/dto"
	"linktree_core/server/model/dto/request"
	"linktree_core/server/model/entity"
	"linktree_core/server/modules/jwt"
	"linktree_core/utils/gos"
	"time"
)

type initializeService struct{}

func (initializeService) VerifyDBLink(ds dto.Dsn) error {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		ds.Username,
		ds.Password,
		ds.Addr,
		ds.Dataname,
		"utf8")
	// 打开数据库
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		return err
	}
	return nil
}

func (initializeService) VerifyRedisLink(rOpt dto.RedisOpt) error {
	ctx := context.Background()
	c := redis.NewClient(&redis.Options{
		Addr:        rOpt.Addr,
		Password:    rOpt.Password,
		DB:          0,
		DialTimeout: 500 * time.Millisecond,
	})
	defer c.Close()
	_, err := c.Ping(ctx).Result()
	if err != nil {
		global.GLOG.Debugf("%v", rOpt)
		return err
	}
	return nil
}

// InitLogin 初始化登录
func (initializeService) InitLogin(request request.LoginRequest) (error, entity.User, string) {
	// 系统未初始化
	if global.SysInit {
		return errors.New("系统已初始化"), entity.User{}, ""
	}

	err, use, pwd := gos.ReadPassFile()
	if err != nil {
		fmt.Printf("There is no pass file")
		return errors.New("there is no pass file"), entity.User{}, ""
	}
	if request.UserName == use && request.Password == pwd {
		// 签发token
		j := jwt.NewJWT()
		token, err2 := j.CreateToken(j.CreateTemporaryClaims(jwt.BaseClaims{
			Username:   "root",
			LoginPlace: request.LoginPlace,
			LoginIp:    request.LoginIp,
		}, time.Now().Add(5*time.Minute))) // 有效期5分钟
		if err2 != nil {
			return err2, entity.User{}, token
		}
		return nil, entity.User{}, token
	} else {
		return errors.New("初始密码错误"), entity.User{}, ""
	}
}

func (initializeService) CreateConfig(setupOpt dto.SetupOpt) error {
	// 设置数据库-判断数据库连接类型
	if setupOpt.Dsn.Addr != "" {
		viper.SetDefault("persistence.mysql.addr", setupOpt.Dsn.Addr)
		viper.SetDefault("persistence.mysql.username", setupOpt.Dsn.Username)
		viper.SetDefault("persistence.mysql.password", setupOpt.Dsn.Password)
		viper.SetDefault("persistence.mysql.dataname", setupOpt.Dsn.Dataname)
		viper.SetDefault("persistence.mysql.charset", "utf8")
	} else {
		viper.SetDefault("persistence.mysql.addr", "")
		viper.SetDefault("persistence.mysql.username", "")
		viper.SetDefault("persistence.mysql.password", "")
		viper.SetDefault("persistence.mysql.dataname", "")
		viper.SetDefault("persistence.mysql.charset", "")
	}
	// 设置缓存方式-判断缓存类型
	if setupOpt.RedisOpt.Addr != "" {
		viper.SetDefault("persistence.redis.addr", setupOpt.RedisOpt.Addr)
		viper.SetDefault("persistence.redis.password", setupOpt.RedisOpt.Password)
	} else {
		viper.SetDefault("persistence.redis.addr", "")
	}
	// mqtt-broker
	viper.SetDefault("broker.addr", setupOpt.Emq.Addr)
	viper.SetDefault("broker.client", setupOpt.Emq.Client)
	viper.SetDefault("broker.username", setupOpt.Emq.Username)
	viper.SetDefault("broker.password", setupOpt.Emq.Password)

	err := viper.SafeWriteConfig()
	if err != nil {
		global.GLOG.Errorf("创建配置文件失败:%v", err)
		return err
	}
	return nil
}
