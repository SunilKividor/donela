package config

import "os"

type Config struct {
	ServerConfig *ServerConfig
	AwsS3Config  *AWSS3Config
}

type ServerConfig struct {
	Port string
}

type AWSS3Config struct {
	Bucket       string
	AccessKey    string
	AccessSecret string
	Region       string
}

func Load() *Config {
	return &Config{
		ServerConfig: &ServerConfig{
			Port: os.Getenv("PORT"),
		},
		AwsS3Config: &AWSS3Config{
			Bucket:       os.Getenv("S3Bucket"),
			AccessKey:    os.Getenv("S3AccessKey"),
			AccessSecret: os.Getenv("S3AccessSecret"),
			Region:       os.Getenv("S3Region"),
		},
	}
}
