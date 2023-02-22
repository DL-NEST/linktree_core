package request

type LoginRequest struct {
	UserName   string `json:"userName,omitempty" binding:"required"`
	Password   string `json:"password,omitempty" binding:"required"`
	LoginPlace string `json:"loginPlace,omitempty" binding:"required"`
	LoginIp    string `json:"loginIp,omitempty" binding:"required"`
}
