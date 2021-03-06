package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Focinfi/misa/builders/confparam"

	"github.com/Focinfi/misa/builders"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/pipelines"
	"github.com/jedib0t/go-pretty/table"
	"github.com/mitchellh/go-homedir"
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
			lines, err := pipelines.InitLinesByFile(configPath)
			if err != nil {
				log.Fatalf("init pipelines failed: %v", err)
			}
			h, ok := lines.GetHandlerOK(args[0])
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
	cmdRun.Flags().StringVarP(&requestData, "data", "d", "", "request data")
	cmdRun.Flags().StringVarP(&requestMeta, "meta", "m", "{}", "request meta")
	cmdRun.Flags().BoolVarP(&verbosely, "verbosely", "v", false, "print every step result")

	var cmdLs = &cobra.Command{
		Use:   "ls",
		Short: "list all pipelines",
		Run: func(cmd *cobra.Command, args []string) {
			lines, err := pipelines.InitLinesByFile(configPath)
			if err != nil {
				log.Fatalf("init pipelines failed: %v", err)
			}
			t := table.NewWriter()
			t.SetTitle("Misa Pipelines")

			style := table.StyleBold
			style.Options.SeparateRows = true
			t.SetStyle(style)

			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Line", "Steps"})

			ids := make([]string, 0, len(lines.LineMap))
			for id := range lines.LineMap {
				ids = append(ids, id)
			}
			sort.Strings(ids)
			for _, id := range ids {
				handler := lines.LineMap[id]
				desc := strings.Join(lineDesc(*handler.Handler.(*pipeline.Line)), "\n")
				t.AppendRow(table.Row{id, desc})
			}
			t.Render()
		},
	}

	var cmdBLs = &cobra.Command{
		Use:   "bls",
		Short: "list all handler builders",
		Run: func(cmd *cobra.Command, args []string) {
			t := table.NewWriter()
			t.SetTitle("Misa Handler Builders")

			style := table.StyleBold
			style.Options.SeparateRows = true
			t.SetStyle(style)

			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Builder", "Field Name", "Field Type", "Description", "Validation"})

			names := make([]string, 0, len(builders.Builders))
			for id := range builders.Builders {
				names = append(names, id)
			}
			sort.Strings(names)
			for _, name := range names {
				builder := builders.Builders[name]
				params, err := confparam.GetConfParams(builder)
				if err != nil {
					fmt.Println("get conf params failed:", err)
					os.Exit(1)
				}
				rows := make([]table.Row, 0, len(params))
				for param, definition := range params {
					rows = append(rows, table.Row{
						"", param, definition.Type, definition.Desc, definition.Validation,
					})
				}
				if len(rows) == 0 {
					rows = []table.Row{{name}}
				} else {
					sort.Slice(rows, func(i, j int) bool {
						return rows[i][1].(string) < rows[j][1].(string)
					})
					rows[0][0] = name
				}
				t.AppendRows(rows)
			}
			t.Render()
		},
	}

	var rootCmd = &cobra.Command{
		Use: "misa [-c config path]",
	}
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("get home dir failed: %v", err)
	}
	defaultConfPath := homeDir + "/.misa/conf.json"
	cmdRun.Flags().StringVarP(&configPath, "conf", "c", defaultConfPath, "request data")
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdBLs)
	rootCmd.Execute()
}

func lineDesc(line pipeline.Line) []string {
	descs := make([]string, 0, len(line.Pipes))

	for i, pipe := range line.Pipes {
		var d string
		switch pipe.Type {
		case pipeline.PipeTypeSingle:
			d = pipe.Conf.Desc
		case pipeline.PipeTypeParallel:
			parallel := pipe.Handler.(*pipeline.Parallel)
			ps := make([]string, 0, len(parallel.Pipes))
			for _, p := range parallel.Pipes {
				ps = append(ps, p.Conf.Desc)
			}
			d = "[" + strings.Join(ps, "|") + "]"
		}

		desc := fmt.Sprintf("%d.%v", i+1, d)
		descs = append(descs, desc)
	}
	return descs
}
