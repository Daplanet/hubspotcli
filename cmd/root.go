package cmd

import (
    "os"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    cfgFile     string
    rootCmd = &cobra.Command{Use: "hscli"}
)

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := os.UserHomeDir()
        cobra.CheckErr(err)

        viper.AddConfigPath(home)
        viper.SetConfigType("json")
        viper.SetConfigName(".config/hubspotcli/config.json")
    }

    viper.AutomaticEnv()

    log.SetFormatter(&log.TextFormatter{
        DisableColors: true,
        FullTimestamp: true,
    })

    if err := viper.ReadInConfig(); err == nil {
        log.Printf("Using config file: %s", viper.ConfigFileUsed())
    } else {
        log.Fatalf("Error Reading config: %s", err)
    }
}
