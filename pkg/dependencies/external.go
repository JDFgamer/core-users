package dependencies

import "core-users/pkg/configs"

type External interface{}

type external struct {
	config configs.Config
}

func InitExternal(config configs.Config) External {
	return &external{
		config: config,
	}
}
