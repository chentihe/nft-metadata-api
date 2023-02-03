package config

import (
	"os"
	"bytes"
	_ "embed"

	"github.com/spf13/viper"
)

//go:embed .env.yaml
var defaultConfiguration []byte

const gcloudFuncSourceDir = "serverless_function_source_code"

type Config struct {
	Stage string `mapstructure:"STAGE"`
	BaseUrl string `mapstructure:"BASE_URL"`
	ContractAddress string `mapstructure:"CONTRACT_ADDRESS"`
	MainnetRpc string `mapstructure:"MAINNET_RPC"`
	GoerliRpc string `mapstructure:"GOERLI_RPC"`
}

func LoadConfig() (*Config, error) {
	// Get the root folder of cloud functions
	fileInfo, err := os.Stat(gcloudFuncSourceDir)
    if err == nil && fileInfo.IsDir() {
        _ = os.Chdir(gcloudFuncSourceDir)
    }

	// Watching changes in .env.yaml
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	// Reading the config file
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

