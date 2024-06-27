package test

type AppConfig struct {
	Sev Server   `mapstructure:"server"`
	DB  Database `mapstructure:"database"`
}
