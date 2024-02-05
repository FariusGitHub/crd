# Kubernetes Custom Resource Definition and Intro to GO Language (Golang)

What are Docker, Kubernetes, Terraform, Prometheus, Etcd, CockroachDB having in common? Right, they are all written in Go Language. GO's language features, performance, compatibility, and community support make it an ideal choice for developing operators and custom resources in Kubernetes.

![](go.jpeg)

Let's see the common custom resource available from kubectl api-resources --sort-by=name

| |NAME |SHORTNAMES   |APIVERSION |NAMESPACED   |KIND
|---|---|---|---|---   |---
|01|apiservices| | apiregistration.k8s.io/v1| false| APIService
|02|bindings ||v1 |true| Binding
|03|certificatesigningrequests| csr| certificates.k8s.io/v1 | false |CertificateSigningRequest
|04|clusterrolebindings||rbac.authorization.k8s.io/v1|false|ClusterRoleBinding
|05|clusterroles||rbac.authorization.k8s.io/v1|false|ClusterRole
|06|componentstatuses|cs|v1|false|ComponentStatus
|07|configmaps|cm|v1|true|ConfigMap
|08|controllerrevisions||apps/v1|true|ControllerRevision
|09|cronjobs|cj|batch/v1|true|CronJob
|10|csidrivers||storage.k8s.io/v1|false|CSIDriver
|11|csinodes||storage.k8s.io/v1|false|CSINode
|12|csistoragecapacities||storage.k8s.io/v1|true|CSIStorageCapacity
|13|customresourcedefinitions|crd,crds|apiextensions.k8s.io/v1|false|CustomResourceDefinition
|14|daemonsets|ds|apps/v1|true|DaemonSet|
|15|deployments|deploy|apps/v1|true|Deployment
|16|endpoints|ep|v1|true|Endpoints
|17|endpointslices||discovery.k8s.io/v1|true|EndpointSlice
|18|events|ev|v1|true|Event
|19|events|ev|events.k8s.io/v1|true|Event
|20|flowschemas||flowcontrol.apiserver.k8s.io/v1beta3|false|FlowSchema
|21|horizontalpodautoscalers|hpa|autoscaling/v2|true|HorizontalPodAutoscaler
|22|ingressclasses||networking.k8s.io/v1|false|IngressClass
|23|ingresses|ing|networking.k8s.io/v1|true|Ingress
|24|jobs||batch/v1|true|Job
|25|leases||coordination.k8s.io/v1|true|Lease
|26|limitranges|limits|v1|true|LimitRange
|27|localsubjectaccessreviews||authorization.k8s.io/v1|true|         LocalSubjectAccessReview
|28|mutatingwebhookconfigurations||admissionregistration.k8s.io/v1|false|MutatingWebhookConfiguration
|29|namespaces|ns|v1|false|Namespace
|30|networkpolicies|netpol|networking.k8s.io/v1|true|NetworkPolicy
|31|nodes|no|v1|false|Node
|32|persistentvolumeclaims|pvc|v1|true|PersistentVolumeClaim
|33|persistentvolumes|pv|v1|false|PersistentVolume
|34|poddisruptionbudgets|pdb|policy/v1|true|         PodDisruptionBudget
|35|pods|po|v1|true|Pod
|36|podtemplates||v1|true|PodTemplate
|37|priorityclasses|pc|scheduling.k8s.io/v1|false|PriorityClass
|38|prioritylevelconfigurations||flowcontrol.apiserver.k8s.io/v1beta3|false|PriorityLevelConfiguration
|39|replicasets|rs|apps/v1|true|ReplicaSet
|40|replicationcontrollers|rc|v1|true|ReplicationController
|41|resourcequotas|quota|v1|true|ResourceQuota
|42|rolebindings||rbac.authorization.k8s.io/v1|true|RoleBinding
|43|roles||rbac.authorization.k8s.io/v1|true|Role
|44|runtimeclasses||node.k8s.io/v1|false|RuntimeClass
|46|secrets||v1|true|Secret
|47|selfsubjectaccessreviews||authorization.k8s.io/v1|false|SelfSubjectAccessReview
|48|selfsubjectreviews||authentication.k8s.io/v1|false|SelfSubjectReview
|49|selfsubjectrulesreviews||authorization.k8s.io/v1|false|SelfSubjectRulesReview
|50|serviceaccounts|sa|v1|true|ServiceAccount
|51|services|svc|v1|true|Service
|52|statefulsets|sts|apps/v1|true|StatefulSet
|53|storageclasses|sc|storage.k8s.io/v1|false|StorageClass
|54|subjectaccessreviews||authorization.k8s.io/v1|false|SubjectAccessReview
|55|tokenreviews||authentication.k8s.io/v1|false|TokenReview
|56|validatingwebhookconfigurations||                admissionregistration.k8s.io/v1|false|ValidatingWebhookConfiguration
|57|volumeattachments||storage.k8s.io/v1|false|        VolumeAttachment

Now we are going to make a new custom resource DEFINITION, one way would let NAMESPACED value True like below
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: namespacedtrue.farius.com
spec:
  group: farius.com
  scope: Namespaced
  names:
    plural: namespacedtrue
    singular: namespacedtrues
    kind: namespacedtrue
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                # Define your spec properties here
                foo:
                  type: string
                bar:
                  type: integer
          required:
            - spec
```
The other way would let the NAMESPACED value to False like below
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: namespacedfalse.farius.com
spec:
  group: farius.com
  scope: Cluster
  names:
    plural: namespacedfalse
    singular: namespacedfalses
    kind: namespacedfalse
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                # Define your spec properties here
                foo:
                  type: string
                bar:
                  type: integer
          required:
            - spec
```
We will see new two custom resources as follow
| |NAME |SHORTNAMES   |APIVERSION |NAMESPACED   |KIND
|---|---|---|---|---   |---
|29|namespacedfalse||farius.com/v1|false|namespacedfalse
|30|namespacetrue||farius.com/v1|true|namespacetrue

And we can make new resources for the two, like making below yaml 

```yaml
apiVersion: farius.com/v1
kind: namespacedtrue
metadata:
  name: namespacedtrue-sample
spec:
  foo: "example"
  bar: 123

---

apiVersion: farius.com/v1
kind: namespacedfalse
metadata:
  name: namespacedfalse-sample
spec:
  foo: "example"
  bar: 123
```

Finally, we could create the custom resources and review them like below

```sh
kubectl get namespacedtrue,namespacedfalse
    NAME                                              AGE
    namespacedtrue.farius.com/namespacedtrue-sample   54s

    NAME                                                AGE
    namespacedfalse.farius.com/namespacedfalse-sample   54s
```

## INTRO TO GO

GO is a programming language developed by Google. It is known for its simplicity, efficiency, and strong support for concurrent programming. GO is often used for building scalable and reliable software systems.

In the context of Kubernetes, GO can be used to build custom resources. Kubernetes allows users to define their own custom resources, which are extensions of the Kubernetes API. These custom resources can represent any kind of object or service that is not natively supported by Kubernetes.

GO provides a Kubernetes client library called "client-go" that makes it easier to interact with the Kubernetes API and build custom resources. The client-go library provides a set of functions and utilities that handle the communication with the Kubernetes API server, allowing developers to create, update, and delete custom resources.

By using GO and the client-go library, developers can write code to define the structure and behavior of custom resources, as well as implement the logic for handling various operations on these resources. GO's simplicity and concurrency features make it a suitable language for building custom resources in Kubernetes, as it allows developers to efficiently handle multiple requests and scale their applications as needed.

I am using GO go1.21.1 linux/amd64
```sh
go version
    go version go1.21.1 linux/amd64
```
Let's pull Gin and client-go like below

```txt
go mod init test # this will create empty go.mod file
go get -u github.com/gin-gonic/gin
go get -u k8s.io/client-go@v0.24.0
```
By the end of both pulls we should see go.mod like below

```sh
cat go.mod
module test

go 1.21.1

require (
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.17.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/client-go v0.24.0 // indirect
)
```

### Hello World with GO
Herewith a little example to port forward Hello World to localhost:8080<br>

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin")

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.Run(":8080")
}
```
Assuming above file was located in main.go, the Hello World should be accessible from the browser

```sh
go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2024/02/04 - 18:19:27 | 200 |      13.886µs |             ::1 | GET      "/"
[GIN] 2024/02/04 - 18:23:45 | 200 |      14.097µs |             ::1 | GET      "/"
[GIN] 2024/02/04 - 18:23:45 | 200 |      28.447µs |             ::1 | GET      "/"
[GIN] 2024/02/04 - 18:23:45 | 200 |       5.313µs |             ::1 | GET      "/"
```

### List all Pods with GO

Let's try another GO script to mimic below command and output
```sh
kubectl get pods-A
    NAMESPACE     NAME                                     READY   STATUS    RESTARTS      AGE
    kube-system   coredns-5dd5756b68-gl66r                 1/1     Running   0             7d4h
    kube-system   coredns-5dd5756b68-p75fr                 1/1     Running   0             7d4h
    kube-system   etcd-docker-desktop                      1/1     Running   0             7d4h
    kube-system   kube-apiserver-docker-desktop            1/1     Running   1 (30h ago)   7d4h
    kube-system   kube-controller-manager-docker-desktop   1/1     Running   1 (30h ago)   7d4h
    kube-system   kube-proxy-4cr8h                         1/1     Running   0             7d4h
    kube-system   kube-scheduler-docker-desktop            1/1     Running   3 (11h ago)   7d4h
    kube-system   storage-provisioner                      1/1     Running   5 (39m ago)   7d4h
    kube-system   vpnkit-controller                        1/1     Running   0             7d4h
```

#### DEPENDING ON YOUR PREVIOUS INSTALLATION, YOU MAY NEED TO DO ADDITIONAL INSTALL BELOW

```sh
go get k8s.io/client-go/plugin/pkg/client/auth/exec@v0.24.0
    go: downloading github.com/davecgh/go-spew v1.1.1
    go: downloading golang.org/x/term v0.16.0
    go: downloading k8s.io/apimachinery v0.24.0
    go: downloading k8s.io/klog/v2 v2.60.1
    go: downloading k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
    go: downloading golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
    go: downloading github.com/go-logr/logr v1.2.0
    go: downloading github.com/gogo/protobuf v1.3.2
    go: downloading sigs.k8s.io/structured-merge-diff/v4 v4.2.1
    go: downloading github.com/google/gofuzz v1.1.0
    go: downloading golang.org/x/time v0.0.0-20220210224613-90d013bbcef8
    go: downloading sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2
    go: downloading sigs.k8s.io/yaml v1.2.0
    go: downloading gopkg.in/inf.v0 v0.9.1
    go: downloading gopkg.in/yaml.v2 v2.4.0
    go: downloading google.golang.org/appengine v1.6.7
    go: downloading github.com/golang/protobuf v1.5.2
go get k8s.io/client-go/discovery@v0.24.0compiler
    go: k8s.io/client-go/discovery@v0.24.0compiler: invalid version: unknown revision v0.24.0compiler
go get k8s.io/client-go/discovery@v0.24.0compiler
    go: k8s.io/client-go/discovery@v0.24.0compiler: invalid version: unknown revision v0.24.0compiler
go get k8s.io/client-go/openapi@v0.24.0compiler
    go: k8s.io/client-go/openapi@v0.24.0compiler: invalid version: unknown revision v0.24.0compiler
go get k8s.io/client-go/discovery@v0.24.0
    go: downloading github.com/google/gnostic v0.5.7-v3refs
    go: downloading k8s.io/api v0.24.0
    go: downloading k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42
    go: downloading github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
    go: downloading github.com/go-openapi/jsonreference v0.19.5
    go: downloading github.com/emicklei/go-restful v2.9.5+incompatible
    go: downloading github.com/go-openapi/swag v0.19.14
    go: downloading github.com/go-openapi/jsonpointer v0.19.5
    go: downloading github.com/PuerkitoBio/purell v1.1.1
    go: downloading github.com/mailru/easyjson v0.7.6
    go: downloading github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
    go: downloading github.com/josharian/intern v1.0.0
go get k8s.io/client-go/kubernetes/typed/admissionregistration/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/apps/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/apps/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/apps/v1beta2@v0.24.0
go get k8s.io/client-go/kubernetes/typed/authentication/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/authentication/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/authorization/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/authorization/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/apps/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/autoscaling/v2@v0.24.0
go get k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2@v0.24.0
go get k8s.io/client-go/kubernetes/typed/batch/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/batch/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/certificates/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/certificates/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/coordination/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/coordination/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/core/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/discovery/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/discovery/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/events/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/events/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/extensions/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2@v0.24.0
go get k8s.io/client-go/kubernetes/typed/networking/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/networking/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/node/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/node/v1alpha1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/node/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/core/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/core/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/rbac/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/rbac/v1alpha1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/rbac/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/scheduling/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/scheduling/v1beta1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/storage/v1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/storage/v1alpha1@v0.24.0
go get k8s.io/client-go/kubernetes/typed/storage/v1beta1@v0.24.0
go get k8s.io/client-go/openapi@v0.24.0
go get k8s.io/client-go/openapi@v0.24.0
go get k8s.io/apimachinery/pkg/util/managedfields@v0.24.0
go get k8s.io/apimachinery/pkg/util/managedfields@v0.24.0
go get k8s.io/client-go/tools/clientcmd@v0.24.0
    go: downloading github.com/spf13/pflag v1.0.5
    go: downloading github.com/imdario/mergo v0.3.5
go get k8s.io/client-go/tools/clientcmd@v0.24.0
```

And herewith the GO script to connect to K8S and list all existing pods

```go
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
}

// Function to get the home directory
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}
```
The command to run above command and the output would be
```txt
go run k8s.go
There are 9 pods in the cluster
- coredns-5dd5756b68-gl66r
- coredns-5dd5756b68-p75fr
- etcd-docker-desktop
- kube-apiserver-docker-desktop
- kube-controller-manager-docker-desktop
- kube-proxy-4cr8h
- kube-scheduler-docker-desktop
- storage-provisioner
- vpnkit-controller
```

Back to earlier example that we did multiple yaml file definition run them through kubectl apply, the same process can be applied with GO like below

```go
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
```

When we run above program through
```txt
go run nstrue2.go
```
The same new custom resource definition namespacedtrue and the same newly crd namespacedtrue named namespacedtrue-sample were created through GO.

```txt
kubectl api-resources | grep namespacedtrue
    namespacedtrue                                 farius.com/v1                          true         namespacedtrue

kubectl get namespacedtrue
    NAME                    AGE
    namespacedtrue-sample   19s
```

## SUMMARY

There are several reasons why we might choose to use Golang code instead of a simple Linux command (like a lengthy Go codes to list all pods compared to Linux command `kubectl get pods -A`)

1. Automation: Golang code allows us to automate tasks and build more complex workflows. We can write logic, perform calculations, and make decisions based on the data retrieved from the Kubernetes cluster.

2. Customization: With Golang, we have the flexibility to customize and extend the functionality according to our specific requirements. We can add additional logic, error handling, or data processing that may not be available with a simple command.

3. Integration: Golang code can be easily integrated into larger systems or applications. We can incorporate the code into existing software, build APIs, or create microservices that interact with the Kubernetes cluster.

4. Portability: Golang code can be compiled into a binary that can run on different platforms and architectures. This allows us to use the same codebase across different environments, making it more portable and scalable.

5. Error handling: Golang provides robust error handling mechanisms, allowing us to handle errors gracefully and provide meaningful error messages. This can be especially useful when dealing with complex operations or when writing production-grade code.

While a simple Linux command like `kubectl get pods -A` is convenient for quick one-off tasks or manual operations, using Golang code provides more flexibility, control, and scalability for building complex applications and automation workflows.