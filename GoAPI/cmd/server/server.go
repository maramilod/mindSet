package server

import (
	"fmt"
	database "mind-set/database"
	"mind-set/internal/validator"
	routes "mind-set/routes"
	"time"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "go run main.go server -c config.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			database.InitData()
			validator.InitValidatorTrans("en")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	host string
	port int
)

func init() {
	Cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "监听服务器地址")
	Cmd.Flags().IntVarP(&port, "port", "P", 8005, "监听服务器端口")

}

func run() error {
	r := routes.SetRouters()
	loc, locErr := time.LoadLocation("Africa/Tripoli")
	if locErr != nil {
		return locErr
	}
	time.Local = loc
	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	return nil
}
