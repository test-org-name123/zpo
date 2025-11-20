package main

import (
	"fmt"
	"log"
	"os"

	"github.com/destrex271/zpo/internal"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected subcommand: [list, describe]")
	}

	args := os.Args[2:]
	switch os.Args[1] {
	case "list":
		listArgs, err := GetListArgs(args)
		if err != nil {
			log.Fatal("Error parsing list arguments: ", err)
		}

		err = internal.ListPostgresqlClusters(listArgs.Namespace, false)
		if err != nil {
			log.Fatal("Error executing list: ", err.Error())
		}

	case "describe":
		describeArgs, err := GetDescribeArgs(args)
		if err != nil {
			log.Fatal("Error parsing describe args: ", err)
		}
		output, err := internal.DescribePostgresqlCluster(describeArgs.Namespace, describeArgs.ClusterName, describeArgs.OutputFile, false)
		if err != nil {
			log.Fatal("Error executing describe: ", err.Error())
		}

		fmt.Println(output)
	}
}
