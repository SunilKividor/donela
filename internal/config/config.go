package config

import "os"

type Config struct {
	ServerConfig   *ServerConfig
	AwsS3Config    *AWSS3Config
	AwsSQSConfig   *AWSSQSConfig
	AwsIAMConfig   *AWSIAMConfig
	R2Config       *R2Config
	PostgresConfig *PostgresConfig
	RedisConfig    *RedisConfig
	JWTConfig      *JWTConfig
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

type R2Config struct {
	AccountId    string
	AccessKey    string
	AccessSecret string
	Bucket       string
	Region       string
}

type AWSSQSConfig struct {
	QueueURL string
}

type PostgresConfig struct {
	ConnectionString string
}

type RedisConfig struct {
	ConnectionString string
}

type JWTConfig struct {
	Secret string
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
		R2Config: &R2Config{
			AccountId:    os.Getenv("R2AccountId"),
			AccessKey:    os.Getenv("R2AccessKey"),
			AccessSecret: os.Getenv("R2AccessSecret"),
			Bucket:       os.Getenv("R2Bucket"),
			Region:       os.Getenv("R2Region"),
		},
		PostgresConfig: &PostgresConfig{
			ConnectionString: os.Getenv("Postgres_URI"),
		},
		RedisConfig: &RedisConfig{
			ConnectionString: os.Getenv("Redis_URI"),
		},
		JWTConfig: &JWTConfig{
			Secret: os.Getenv("JWTAPISECRET"),
		},
	}
}
