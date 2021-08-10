package cmd

import (
    "log"
    "github.com/spf13/cobra"
    "github.com/bold-commerce/go-hubspot/hubspot"
)

var (
    apiKey string
    emailId int

    emailSingle = &cobra.Command{
        Use: "email client",
        Short: "Send a single email to a client",
        Long: "Send a single email by id to a client",
        Args: cobra.MiniumNArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {

            var hubspotClient *hubspot.NewClient("https://api.hubspot.com", apiKey)
            err, _ := hubspotClient.SingleEmail(emailId, "tyler.durden@gmail.com")
            if err != nil {
                log.Fatalf("hubspot error: %s", err.Error())
                return err
            }
            return nil
            },
    }
)

func init() {

    emailSingle.Flags().String(&apiKey, "apikey", "k", 1, "Hubspot API Key")
    emailSingle.Flags().IntVarP(&emailId, "emailid", "e", 1, "Email id in hubspot to send")

    rootCmd.AddCommand(emailSingle)
}
