package cmd

import (
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/bold-commerce/go-hubspot/hubspot"
)

var (
    apiKey          string
    emailId         int
    hubspotClient   hubspot.Client

    emailSingle = &cobra.Command{
        Use: "email client",
        Short: "Send a single email to a client",
        Long: "Send a single email by id to a client",
        Args: cobra.MinimumNArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {

            hubspotClient := hubspot.NewClient("https://api.hubspot.com", apiKey)

            if err := hubspotClient.SingleEmail(emailId, "tyler.durden@gmail.com"); err != nil {
                log.Fatalf("hubspot error: %s", err.Error())
                return err
            }
            return nil
        },
    }
)

func init() {

    emailSingle.Flags().StringVar(&apiKey, "apikey", "k", "Hubspot API Key")
    emailSingle.Flags().IntVar(&emailId, "emailid", 1, "Email id in hubspot to send")

    rootCmd.AddCommand(emailSingle)
}
