package global

import (
	"Go-000/Week04/pkg/logger"
	"Go-000/Week04/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	AppSetting      *setting.AppSettingS
	Logger          *logger.Logger
)
