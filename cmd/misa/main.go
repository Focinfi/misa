package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/mitchellh/go-homedir"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers"
	"github.com/spf13/cobra"
)

func main() {
	var (
		requestData string
		requestMeta string
		configPath  string
		verbosely   bool
	)

	var cmdRun = &cobra.Command{
		Use:   "run [pipeline name] [-d request data] [-m request meta] [-v verbosely]",
		Short: "run a pipeline",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := handlers.InitHandlers(configPath); err != nil {
				log.Fatalf("init pipelines failed: %v", err)
			}
			h, ok := handlers.Handlers.GetOK(args[0])
			if !ok {
				log.Fatalf("pipeline[%v] not found", args[0])
			}

			meta := make(map[string]interface{})
			if err := json.Unmarshal([]byte(requestMeta), &meta); err != nil {
				log.Fatalf("meta is not a json object string, err: %v", err)
			}

			req := &pipeline.HandleRes{
				Data: requestData,
				Meta: meta,
			}

			var (
				resp interface{}
				err  error
			)
			if verbosely {
				resp, err = h.(*pipeline.Line).HandleVerbosely(context.Background(), req)
			} else {
				resp, err = h.Handle(context.Background(), req)
			}
			if err != nil {
				log.Fatalf("failed: %v", err)
			}
			p, _ := json.MarshalIndent(resp, "", "  ")
			log.Println(string(p))
		},
	}
	cmdRun.Flags().StringVarP(&requestData, "data", "r", "", "request data")
	cmdRun.Flags().StringVarP(&requestMeta, "meta", "m", "{}", "request meta")
	cmdRun.Flags().BoolVarP(&verbosely, "verbosely", "v", false, "print every step result")

	var rootCmd = &cobra.Command{Use: "misa [-c config path]"}
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("get home dir failed: %v", err)
	}
	defaultConfPath := homeDir + "/.misa/conf.json"
	cmdRun.Flags().StringVarP(&configPath, "conf", "c", defaultConfPath, "request data")
	rootCmd.AddCommand(cmdRun)
	rootCmd.Execute()
}
