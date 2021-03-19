package clients

// Config interface
type Config interface {
	GetConfigCustom() *ConfigCustomServer
}

// ConfigOgmioServer holds the configuration of your Ogmio
type ConfigCustomServer struct {
	Name              			string
	LocationFiles               string
}

// NewConfig init point to play with Config
func NewConfig(configCustomServer *ConfigCustomServer) Config {
	config := &configClient{
		configCustomServer: configCustomServer,
	}
	return config
}

type configClient struct {
	configCustomServer *ConfigCustomServer
}

// AuthHandle takes care of the auth
func (j *configClient) GetConfigCustom() *ConfigCustomServer {
	return j.configCustomServer
}
