package logging

import (
	"fmt"
	"lawyerinyou-backend/pkg/settings"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", settings.AppConfigSetting.App.RuntimeRootPath, settings.AppConfigSetting.App.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format(settings.AppConfigSetting.App.TimeFormat),
		settings.AppConfigSetting.App.LogFileExt,
	)
}
