package conf

import (
	"github.com/go-ini/ini"
	"log"
	"sync"
)

type GlobalSetting struct {
	App *App
	Database *Database
}

type App struct {
	Port int
	RunMode string
}

type Database struct {
	Type  string
	User  string
	Password  string
	Host  string
	Name string
	TablePrefix string
}

var app = &App{}
var database = &Database{}

var globalSetting = &GlobalSetting{
	App: app,
	Database: database,
}

var configLock = new(sync.RWMutex)

var cfg *ini.File

func Config() *GlobalSetting {
	configLock.RLock()
	defer configLock.RUnlock()
	return globalSetting
}

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app",app)
	mapTo("database",database)
}

//刷新配置,刷新端口不能实时生效，需要重启应用
func Reload()  {
	configLock.Lock()
	defer configLock.Unlock()
	Setup()
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
