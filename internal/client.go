package internal

import (
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/tools/clientcmd"
)

var postgresGVR = schema.GroupVersionResource{
	Group:    "acid.zalan.do",
	Version:  "v1",
	Resource: "postgresqls",
}

func GetDynamicClient(isTest bool) (dynamic.Interface, error) {
	if isTest {
		return getDynamicClient_Test()
	}
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.ExpandEnv("$HOME/.kube/config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(config)
}

func getDynamicClient_Test() (dynamic.Interface, error) {
	postgresGVR = schema.GroupVersionResource{
		Group:    "acid.zalan.do",
		Version:  "v1",
		Resource: "postgresqls",
	}

	// Define GVK for the fake objects
	gvk := schema.GroupVersionKind{
		Group:   "acid.zalan.do",
		Version: "v1",
		Kind:    "postgresql",
	}

	cluster1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "acid.zalan.do/v1",
			"kind":       "postgresql",
			"metadata": map[string]interface{}{
				"name":      "pg-cluster-1",
				"namespace": "default",
			},
			"status": map[string]interface{}{
				"PostgresClusterStatus": "Running",
			},
		},
	}
	cluster1.SetGroupVersionKind(gvk)

	cluster2 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "acid.zalan.do/v1",
			"kind":       "postgresql",
			"metadata": map[string]interface{}{
				"name":      "pg-cluster-2",
				"namespace": "dev",
			},
			"status": map[string]interface{}{
				"PostgresClusterStatus": "Pending",
			},
		},
	}
	cluster2.SetGroupVersionKind(gvk)

	scheme := runtime.NewScheme()
	client := fake.NewSimpleDynamicClient(scheme, cluster1, cluster2)
	return client, nil
}
