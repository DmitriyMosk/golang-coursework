package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var GConfig Config

type Config struct {
	Port             int `mapstructure:"port"`
	ResourseTimeout  int `mapstructure:"resourseTimeout"`
	AnalyticsTimeout int `mapstructure:"analyticsTimeout"`
	DBSettings       struct {
		DBUser     string `mapstructure:"dbUser"`
		DBPassword string `mapstructure:"dbPassword"`
		DBHost     string `mapstructure:"dbHost"`
		DBPort     int    `mapstructure:"dbPort"`
		DBName     string `mapstructure:"dbName"`
	} `mapstructure:"DBSettings"`
	ProgramSettings struct {
		JiraURL           string `mapstructure:"jiraUrl"`
		JiraUser          string `mapstructure:"jiraUrl"`
		JiraToken         string `mapstructure:"jiraUrl"`
		ThreadCount       int    `mapstructure:"threadCount"`
		IssueInOneRequest int    `mapstructure:"issueInOneRequest"`
		MaxTimeSleep      int    `mapstructure:"maxTimeSleep"`
		MinTimeSleep      int    `mapstructure:"minTimeSleep"`
	} `mapstructure:"ProgramSettings"`
}

func InitConfig(path string) (GConfig Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can not read CFG")
		return
	}

	err = viper.Unmarshal(&GConfig)
	fmt.Println("GConfig")
	fmt.Println(GConfig)
	return
}
