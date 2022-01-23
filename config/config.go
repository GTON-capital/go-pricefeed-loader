package config

type Config struct {
	ReportEndpoint string `mapstructure:PRICELOADER_REPORT_ENDPOINT`
}

var MainConfig Config
