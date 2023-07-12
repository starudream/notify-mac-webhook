package main

import (
	"os"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/flag"
	"github.com/starudream/go-lib/log"
)

var rootCmd = &flag.Command{
	Use:     constant.NAME,
	Version: constant.VERSION + " (" + constant.BIDTIME + ")",
	Run: func(cmd *flag.Command, args []string) {
		app.Add(start)
		app.Defer(stop)
		err := app.Go()
		if err != nil {
			log.Fatal().Msgf("%v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().String("addr", "0.0.0.0:5400", "server address")
	_ = config.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
