package feedreader

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const (
	defaultTitle = "Feed Reader"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	feeds     []string
	feedLimit int
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		feeds:     wtf.ToStrs(ymlConfig.UList("feeds")),
		feedLimit: ymlConfig.UInt("feedLimit", -1),
	}

	return settings
}
