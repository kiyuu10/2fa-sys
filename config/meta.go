package config

type Config struct {
	DB DBConfig `json:"db"`
}

type DBConfig struct {
	URL string `json:"url"`
}
