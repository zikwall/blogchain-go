package base

type (
	Application struct {
		config ApplicationConfig
	}

	ApplicationConfig struct {
	}

	Service interface {
		Build() error
	}
)

func NewApplication(config ApplicationConfig) *Application {
	a := &Application{}
	a.config = config

	return a
}
