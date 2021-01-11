package app

import (
	"encoding/json"
	"fmt"
	"github.com/turtacn/alameda/cmd/app"
	"github.com/turtacn/alameda/datahub"
	"github.com/spf13/cobra"
)

const (
	envVarPrefix                    = "ALAMEDA_DATAHUB"
	defaultRotationMaxSizeMegabytes = 100
	defaultRotationMaxBackups       = 7
	defaultLogRotateOutputFile      = "/var/log/alameda/alameda-datahub.log"
)

var (
	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "start alameda datahub server",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			var (
				err error

				server *datahub.Server
			)
			app.PrintSoftwareVer()
			initConfig()
			initLogger()
			initEventMgt()
			initKeycode()
			initNotifier()
			setLoggerScopesWithConfig(*config.Log)
			displayConfig()
			server, err = datahub.NewServer(config)
			if err != nil {
				panic(err)
			}

			server.InitInfluxdbDatabase()

			if err = server.Run(); err != nil {
				server.Stop()
				panic(err)
			}
		},
	}
)

func displayConfig() {
	if configBin, err := json.MarshalIndent(config, "", "  "); err != nil {
		scope.Error(err.Error())
	} else {
		scope.Infof(fmt.Sprintf("Alameda datahub server configuration: %s", string(configBin)))
	}
}
