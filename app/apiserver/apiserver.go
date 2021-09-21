package apiserver

type ApiServer struct {
	config *Config
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
	}
}

func (s *ApiServer) Start() error {
	return nil
}
