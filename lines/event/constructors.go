package event

import (
	"JakesBettingAlgorithmV2/lines/logging"
	"JakesBettingAlgorithmV2/lines/utils"
	"fmt"
)

func NewRabbitMQEventHandlerConfig(appName string) RabbitMQEventHandlerConfig {
	logLevel := utils.GetEnvOrDefault("LOG_LEVEL", "info", "string").(string)
	testRunner := utils.GetEnvOrDefault("TEST_RUNNER", "false", "bool").(bool)
	logger := logging.NewLogrusHandler(logLevel)
	urlEnv := ""
	exchangeEnv := ""
	if testRunner {
		urlEnv = fmt.Sprintf("%v_RABBITMQ_URL_TEST", appName)
		exchangeEnv = fmt.Sprintf("%v_RABBITMQ_EXCHANGE_NAME_TEST", appName)
	} else {
		urlEnv = fmt.Sprintf("%v_RABBITMQ_URL", appName)
		exchangeEnv = fmt.Sprintf("%v_RABBITMQ_EXCHANGE_NAME", appName)
	}
	connString := utils.GetEnvOrDefault(
		urlEnv,
		"NODEFAULT",
		"string",
	).(string)
	exchangeName := utils.GetEnvOrDefault(
		exchangeEnv,
		"NODEFAULT",
		"string",
	).(string)
	return RabbitMQEventHandlerConfig{
		Logger:       logger,
		URL:          connString,
		AppName:      appName,
		ExchangeName: exchangeName,
		Concurrency:  utils.GetEnvOrDefault(fmt.Sprintf("%v_RABBITMQ_CONCURRENCY", appName), "50", "int").(int),
	}
}

func NewRabbitEventHandler(config RabbitMQEventHandlerConfig) *RabbitEventHandler {
	handler := RabbitEventHandler{
		Config: config,
	}
	err := handler.Connect()
	if err != nil {
		config.Logger.Fatal("event", "NewRabbitEventHandler", err.Error())
	}
	return &handler
}
