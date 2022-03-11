package models

type Config struct {
	Port string `json:"port"`
	DbHost string `json:"db_host"`
	DbName string `json:"db_name"`
	DbUser string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbPort string `json:"db_port"`
}
