package cmd

import (
	"fmt"

	helper "github.com/home-assistant/hassio-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var homeassistantRestartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"re"},
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("homeassistant restart")

		section := "homeassistant"
		command := "restart"
		base := viper.GetString("endpoint")

		resp, err := helper.GenericJSONPost(base, section, command, nil)
		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}

		return
	},
}

func init() {
	homeassistantCmd.AddCommand(homeassistantRestartCmd)
}
