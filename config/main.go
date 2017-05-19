package config

import (
    "log"
    "sync"
    "io/ioutil"
    "encoding/json"
)

var environments = map[string]string {
    "prod": "config/json/prod.json", // production environment config
    "dev": "config/json/dev.json", // development environment config
    "dev-local": "config/json/dev-local.json",
}  

type config struct {
    MongoAddress string
    MongoDBName string
}

var (
    currentConfig *config
    currentEnv string
    once sync.Once
)

func Load(env string) {
    once.Do(func() {
        currentEnv = env
        log.Printf("Configuration environment set to: %s", currentEnv)
        // load the config file
        loadConfigurationByEnv(currentEnv)
    })
}

func loadConfigurationByEnv(env string) {
    configurationFile, err := ioutil.ReadFile(environments[env])
    if err != nil {
        log.Println("No config file has been found in the provided path", err)
    }

    if err := json.Unmarshal(configurationFile, &currentConfig); err != nil {
        log.Println("Error while parsing the config file", err)
    }
}

func Env() string {
    return currentEnv
}

// Current returns current configuration
func Current() *config {
    if currentConfig == nil {
        Load(currentEnv)
    }
    
    return currentConfig
}