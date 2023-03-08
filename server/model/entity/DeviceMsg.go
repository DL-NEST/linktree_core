package entity

// DeviceMsg 设备消息表
type DeviceMsg struct {
	BaseModel
	DeviceSubject string
	DeviceMsg     string
}
