package configs

var (
	ConfigLoaderObject *ConfigLoader
)

type ConfigLoader struct {
	EnvConfig *EnvConfig
	DbConfig  *DbConfig
}

func NewConfigLoader() *ConfigLoader {
	if ConfigLoaderObject == nil {
		ConfigLoaderObject = &ConfigLoader{
			EnvConfig: NewEnvConfig(),
			DbConfig:  NewDbConfig(),
		}
	}

	return ConfigLoaderObject
}

func (c *ConfigLoader) LoadConfig() error {
	err := c.EnvConfig.LoadConfig()
	if err != nil {
		return err
	}

	err = c.DbConfig.ConnectToDatabase(c.EnvConfig)
	if err != nil {
		return err
	}

	return nil
}
