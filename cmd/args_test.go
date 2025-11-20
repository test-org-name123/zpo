package main

import (
	"testing"
)

func TestGetListArgs_NoArgs(t *testing.T) {
	args := []string{}
	got, err := GetListArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Namespace != "" {
		t.Errorf("expected Namespace '', got '%s'", got.Namespace)
	}
}

func TestGetListArgs_NamespaceFullFlag(t *testing.T) {
	args := []string{"--namespace", "dev"}
	got, err := GetListArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Namespace != "dev" {
		t.Errorf("expected Namespace 'dev', got '%s'", got.Namespace)
	}
}

func TestGetListArgs_NamespaceShortFlag(t *testing.T) {
	args := []string{"-n", "staging"}
	got, err := GetListArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Namespace != "staging" {
		t.Errorf("expected Namespace 'staging', got '%s'", got.Namespace)
	}
}

func TestGetListArgs_InvalidFlag(t *testing.T) {
	args := []string{"--invalid"}
	_, err := GetListArgs(args)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestGetDescribeArgs_MissingClusterName(t *testing.T) {
	args := []string{"--namespace", "dev"}
	_, err := GetDescribeArgs(args)
	if err == nil {
		t.Fatal("expected error for missing cluster name, got nil")
	}
}

func TestGetDescribeArgs_MissingNamespace(t *testing.T) {
	args := []string{"my-cluster"}
	_, err := GetDescribeArgs(args)
	if err == nil {
		t.Fatal("expected error for missing namespace, got nil")
	}
}

func TestGetDescribeArgs_ValidYamlOutput(t *testing.T) {
	args := []string{"--namespace", "prod", "my-cluster"}
	got, err := GetDescribeArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Namespace != "prod" {
		t.Errorf("expected Namespace 'prod', got '%s'", got.Namespace)
	}
	if got.ClusterName != "my-cluster" {
		t.Errorf("expected ClusterName 'my-cluster', got '%s'", got.ClusterName)
	}
	if got.OutputFile != "yaml" {
		t.Errorf("expected OutputFile 'yaml', got '%s'", got.OutputFile)
	}
}

func TestGetDescribeArgs_ValidJsonOutput(t *testing.T) {
	args := []string{"--namespace", "prod", "--output", "json", "my-cluster"}
	got, err := GetDescribeArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Namespace != "prod" {
		t.Errorf("expected Namespace 'prod', got '%s'", got.Namespace)
	}
	if got.ClusterName != "my-cluster" {
		t.Errorf("expected ClusterName 'my-cluster', got '%s'", got.ClusterName)
	}
	if got.OutputFile != "json" {
		t.Errorf("expected OutputFile 'json', got '%s'", got.OutputFile)
	}
}

func TestGetDescribeArgs_InvalidOutputFormat(t *testing.T) {
	args := []string{"--namespace", "prod", "--output", "xml", "my-cluster"}
	_, err := GetDescribeArgs(args)
	if err == nil {
		t.Fatal("expected error for invalid output format, got nil")
	}
}
