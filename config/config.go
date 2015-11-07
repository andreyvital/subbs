package config

type Config struct {
	Languages []string
	OS        struct {
		User     string
		Password string
	}
}
