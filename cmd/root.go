package cmd

import (
    "os"
    "log"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"


var (
    cfgFile     string
    rootCmd = &cobra.Command{Use: "app"}
)

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFlag, "config", "c", "config file [default: $HOME/.config/hubspotcli/config.json]")
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(initCmd)
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := os.UserHomeDir()
        cobra.CheckErr(err)

        viper.AddConfigPath(home),
        viper.SetConfigType("json")
        viper.SetConfigName(".config/hubspotcli/config")
    }

    viper.AutomaticEnv()
    if err := viperReadInConfig(); err == nil {
        log.Successf("Using config file: %s", viper.ConfigFileUsed())
    } else {
        log.Fatalf("Error Reading config: %s", err)
    }
}
