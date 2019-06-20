package deps

import "github.com/PayloadPro/svc.pro.payload.bins/configs"

// Config wrapped in a container
type Config struct {
	DB *configs.DBConfig
}

// Setup the config
func (c Config) Setup() {
	c.DB.Setup()
}
