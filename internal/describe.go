package internal

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func DescribePostgresqlCluster(namespace, name, output string, istest bool) (string, error) {
	client, err := GetDynamicClient(istest)
	if err != nil {
		return "", err
	}

	resource, err := client.Resource(postgresGVR).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	jsonBytes, err := resource.MarshalJSON()
	if err != nil {
		return "", err
	}

	switch output {
	case "json":
		return string(jsonBytes), nil
	case "yaml":
		yamlBytes, err := yaml.JSONToYAML(jsonBytes)
		if err != nil {
			return "", err
		}
		return string(yamlBytes), nil
	default:
		return "", fmt.Errorf("unsupported output format: %s", output)
	}
}
