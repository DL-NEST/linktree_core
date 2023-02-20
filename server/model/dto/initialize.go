package dto

type Dsn struct {
	Addr     string `json:"addr,omitempty"`
	Dataname string `json:"dataname,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type RedisOpt struct {
	Addr     string `json:"addr,omitempty"`
	Password string `json:"password,omitempty"`
}

type Emq struct {
	Addr     string `json:"addr,omitempty"`
	Client   string `json:"client,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type SetupOpt struct {
	Dsn      Dsn      `json:"dsn,omitempty"`
	RedisOpt RedisOpt `json:"redis_opt,omitempty"`
	Emq      Emq      `json:"emq,omitempty"`
}
