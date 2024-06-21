package dependencies

import "leonardodelira/go-clean-template/internal/gateway"

func initGateways() {
	translatorGateway = gateway.NewTranslatorGatewayDeepl()
}
