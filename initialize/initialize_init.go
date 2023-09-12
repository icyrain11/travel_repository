package initialize

// InitializeGlobalVaribale 全局配置初始化化
func InitializeGlobalVaribale() {
	//初始化gorm
	Gorm()
	//初始化session
	Session()
}
