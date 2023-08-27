package config

// import (
// 	"fmt"
// 	"log"

// 	"github.com/spf13/viper"
// )


// type Config struct {
// 	HOST string `mapstructure:"HOST"`
// 	PORT string `mapstructure:"PORT"`
// 	PASSWORD string `mapstructure:"PASSWORD"`
// 	USER string `mapstructure:"USER"`
// 	DBNAME string `mapstructure:"DBNAME"`
// 	SSLMODE string `mapstructure:"SSLMODE"`
// 	TIMEZONE string `mapstructure:"TIMEZONE"`
// }


// func LoadConfig(path string) (config Config,err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err=viper.ReadInConfig()
// 	if err !=nil {
// 		return 
// 	}
// 	err= viper.Unmarshal(&config)
// 	return
// }


// func EnvironmentVar()Config{
// 	// LOAD ENV 
// 	fmt.Println("🥳🇧🇩🎉👨‍💻 loading Environment variables ")
// 	config,err:=LoadConfig(".")
// 	if err !=nil {
// 		log.Fatalln("🔥❌ viper error : ",err)
		
// 	}else {
// 		fmt.Println("🎉 Config : ",config)
// 	}
// 	return config

// }