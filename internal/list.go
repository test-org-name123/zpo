package internal

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ListPostgresqlClusters(namespace string, istest bool) error {
	client, err := GetDynamicClient(istest)
	if err != nil {
		return err
	}

	var list *unstructured.UnstructuredList
	if namespace == "" {
		list, err = client.Resource(postgresGVR).List(context.TODO(), metav1.ListOptions{})
	} else {
		list, err = client.Resource(postgresGVR).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	}
	if err != nil {
		return err
	}

	if len(list.Items) == 0 {
		if namespace == "" {
			return fmt.Errorf("No postgres resources available across namespaces!")
		}
		return fmt.Errorf("No postgres resources available in the namespace '%s'", namespace)
	}

	for _, item := range list.Items {
		name := item.GetName()
		ns := item.GetNamespace()
		status, _, _ := unstructured.NestedString(item.Object, "status", "PostgresClusterStatus")
		fmt.Printf("Name: %-30s Namespace: %-20s Status: %s\n", name, ns, status)
	}
	return nil
}
