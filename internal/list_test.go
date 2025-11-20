package internal

import (
	"testing"
)

func TestListPostgresqlClusters(t *testing.T) {
	err := ListPostgresqlClusters("", true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
