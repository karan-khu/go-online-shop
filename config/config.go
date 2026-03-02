package config

type Config struct {
	env *Env
}

func NewConfig() *Config {
	env := NewEnv()

	return &Config{
		env: env,
	}
}
