package settings

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int    `mapstructure:"machine_id"`
	//*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*JwtConfig   `mapstructure:"jwt"`
	*OssConfig   `mapstructure:"oss"`
	*CodeConfig  `mapstructure:"code"`
}

type OssConfig struct {
	AccessKeyID     string
	AccessKeySecret string
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	Debug        bool   `mapstructure:"debug"`
}

type JwtConfig struct {
	Expire int    `mapstructure:"expire"`
	Header string `mapstructure:"header"`
	Key    string `mapstructure:"key"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type CodeConfig struct {
	StudentCode string `yaml:"studentCode"`
	TeacherCode string `yaml:"teacherCode"`
	ServiceCode string `yaml:"serviceCode"`
	StudentRole int    `yaml:"studentRole"`
	TeacherRole int    `yaml:"studentRole"`
	ServiceRole int    `yaml:"studentRole"`
}

/*
   studentCode: stu
   teacherCode: teac
   serviceCode: serv
   studentRole: 3
   teacherRole: 8
   serviceRole: 10
*/
//
//type LogConfig struct {
//	Level      string `mapstructure:"level"`
//	Filename   string `mapstructure:"filename"`
//	MaxSize    int    `mapstructure:"max_size"`
//	MaxAge     int    `mapstructure:"max_age"`
//	MaxBackups int    `mapstructure:"max_backups"`
//}

func Init() error {
	mode := flag.String("mode", "dev", "设置运行模式:dev pro local")
	flag.Parse()
	filePath := "./config/config.yaml"
	switch *mode {
	case "dev":
		filePath = "./config/config.yaml"
	case "pro":
		filePath = "./config/config_pro.yaml"
	case "local":
		filePath = "./config/config_local.yaml"
	default:
		filePath = "./config/config.yaml"
	}
	viper.SetConfigFile(filePath)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
