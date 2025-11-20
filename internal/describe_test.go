package internal

import (
	"strings"
	"testing"
)

func TestDescribePostgresqlCluster_JSON(t *testing.T) {
	_, err := DescribePostgresqlCluster("default", "pg-cluster-1", "json", true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDescribePostgresqlCluster_YAML(t *testing.T) {
	_, err := DescribePostgresqlCluster("dev", "pg-cluster-2", "yaml", true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDescribePostgresqlCluster_InvalidFormat(t *testing.T) {
	_, err := DescribePostgresqlCluster("default", "pg-cluster-1", "xml", true)
	if err == nil || !strings.Contains(err.Error(), "unsupported output format") {
		t.Errorf("Expected unsupported format error, got: %v", err)
	}
}

func TestDescribePostgresqlCluster_NotFound(t *testing.T) {
	_, err := DescribePostgresqlCluster("nonexistent", "pg-missing", "json", true)
	if err == nil {
		t.Errorf("Expected error for missing cluster, got nil")
	}
}
