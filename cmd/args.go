package main

import (
	"errors"
	"flag"
	"log"
)

type ListArgs struct {
	Namespace string
}

type DescribeArgs struct {
	Namespace   string
	ClusterName string
	OutputFile  string
}

func GetListArgs(args []string)            (ListArgs, error) {
	listCmd := flag.NewFlagSet("list", flag.ContinueOnError)
	namespace := listCmd.String("namespace", "", "Target namespace (optional)")
	n := listCmd.String("n", "", "Shorthand for --namespace")
	if err := listCmd.Parse(args); err != nil {
		return ListArgs{}, err
	}

	ns := *namespace
	if ns == "" {
		ns = *n
	}

	return ListArgs{
		Namespace: ns,
	}, nil
}

func GetDescribeArgs(args []string) (DescribeArgs, error) {
	describeCmd := flag.NewFlagSet("describe", flag.ContinueOnError)

	namespace := describeCmd.String("namespace", "", "Target namespace (required)")

	outputFile := describeCmd.String("output", "yaml", "Output format: yaml or json (default yaml)")

	if err := describeCmd.Parse(args); err != nil {
		log.Fatal("Error parsing flags:", err)
	}

	if f := describeCmd.Lookup("n"); f != nil && f.Value.String() != "" && f.DefValue != f.Value.String() {
		*namespace = f.Value.String()
	}
	if f := describeCmd.Lookup("o"); f != nil && f.Value.String() != "" && f.DefValue != f.Value.String() {
		*outputFile = f.Value.String()
	}

	allArgs := describeCmd.Args()
	if len(allArgs) < 1 {
		return DescribeArgs{}, errors.New("missing required positional argument: <clustername>")
	}

	if *namespace == "" {
		return DescribeArgs{}, errors.New("--namespace (-n) is required")
	}

	if *outputFile != "json" && *outputFile != "yaml" {
		return DescribeArgs{}, errors.New("--output (-o) must be 'json' or 'yaml'")
	}

	return DescribeArgs{
		ClusterName: allArgs[0],
		Namespace:   *namespace,
		OutputFile:  *outputFile,
	}, nil
}
