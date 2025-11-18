package config

import "os"

type Config struct {
	ServerConfig *ServerConfig
	AwsS3Config  *AWSS3Config
	AwsSQSConfig *AWSSQSConfig
	AwsIAMConfig *AWSIAMConfig
}

type ServerConfig struct {
	Port string
}

type AWSIAMConfig struct {
	AccessKey    string
	AccessSecret string
	Region       string
}

type AWSS3Config struct {
	Bucket string
}

type AWSSQSConfig struct {
	QueueURL string
}

func Load() *Config {
	return &Config{
		ServerConfig: &ServerConfig{
			Port: os.Getenv("PORT"),
		},
		AwsS3Config: &AWSS3Config{
			Bucket: os.Getenv("S3Bucket"),
		},
		AwsSQSConfig: &AWSSQSConfig{
			QueueURL: os.Getenv("SQSQueueURL"),
		},
		AwsIAMConfig: &AWSIAMConfig{
			AccessKey:    os.Getenv("S3AccessKey"),
			AccessSecret: os.Getenv("S3AccessSecret"),
			Region:       os.Getenv("S3Region"),
		},
	}
}
