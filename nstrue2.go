package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

type NamespacedTrue struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Foo string `yaml:"foo"`
	Bar int    `yaml:"bar"`
}

func main() {
	// Path to the kubeconfig file
	kubeconfig := flag.String("kubeconfig", filepath.Join(
		"/home/devops", ".kube", "config"), "path to your kubeconfig file")

	// Create the config object from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Kubernetes clientset
	// clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Create a dynamic client
	dynamicClient := dynamic.NewForConfigOrDie(config)

	// Define the CustomResourceDefinition
	crd := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": "namespacedtrue.farius.com",
			},
			"spec": map[string]interface{}{
				"group": "farius.com",
				"scope": "Namespaced",
				"names": map[string]interface{}{
					"plural":   "namespacedtrue",
					"singular": "namespacedtrues",
					"kind":     "namespacedtrue",
				},
				"versions": []interface{}{
					map[string]interface{}{
						"name":    "v1",
						"served":  true,
						"storage": true,
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"type": "object",
								"properties": map[string]interface{}{
									"spec": map[string]interface{}{
										"type": "object",
										"properties": map[string]interface{}{
											"foo": map[string]interface{}{
												"type": "string",
											},
											"bar": map[string]interface{}{
												"type": "integer",
											},
										},
									},
								},
								"required": []interface{}{
									"spec",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the CustomResourceDefinition
	_, err = dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	}).Create(context.TODO(), crd, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	// Create the NamespacedTrue object
	namespacedTrue := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "farius.com/v1",
			"kind":       "namespacedtrue",
			"metadata": map[string]interface{}{
				"name": "namespacedtrue-sample",
			},
			"spec": map[string]interface{}{
				"foo": "example",
				"bar": 123,
			},
		},
	}

	// Delay for 10 seconds
	time.Sleep(10 * time.Second)

	// Create the NamespacedTrue resource
	_, err = dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "farius.com",
		Version:  "v1",
		Resource: "namespacedtrue",
	}).Namespace("default").Create(context.TODO(), namespacedTrue, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

}
