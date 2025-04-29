package authcon

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AuthConnectorBuilder interface { 
	AuthConnector() AuthConnector
}

type authConnectorBuilder struct {
	getter kv.Getter
	comfig.Once
}

func (a *authConnectorBuilder) AuthConnector() AuthConnector {
	return a.Do(func() interface{} {
		var config struct {
			URL string `fig:"url"`
		}

		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(a.getter, "auth_connector")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out auth_connector"))
		}

		connector := authConnector{
			url: config.URL,
		}

		return &connector
	}).(AuthConnector)
}


func NewAuthConnectorBuilder(getter kv.Getter) AuthConnectorBuilder {
	return &authConnectorBuilder{
		getter: getter,
	}
}
