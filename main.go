package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// v1 "k8s.io/api/apps/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/tools/clientcmd"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Happy Birthday, Dad!")
	})
	r.Run(":8080")
}

// func GetKubernetesClient() (*kubernetes.Clientset, error) {
// 	var kubeconfig = "your kube config file path"
// 	/***
// 	BuildConfigFromFlags used as helper function that
// 	builds configs from a master url or a kubeconfig filepath
// 	***/
// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	/**
// 	setup the client with configuration config
// 	**/

// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return clientset, nil
// }
