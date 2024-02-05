package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// Path to the kubeconfig file
	kubeconfig := flag.String("kubeconfig", filepath.Join(
		homeDir(), ".kube", "config"), "path to your kubeconfig file")

	// Create the config object from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Example: List all pods in the cluster
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}

	// Example: List all nodes in the cluster
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
	for _, node := range nodes.Items {
		fmt.Printf("- %s\n", node.Name)
	}
	fmt.Printf("- %s\n", homeDir())
}

// Function to get the home directory
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}
