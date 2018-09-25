//list all pod names, showing what namespaces they are in
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{LabelSelector: "k8s-app=heapster"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d pods in the label 'k8s-app=heapster' \n", len(pods.Items))

	for _, s := range pods.Items {
		fmt.Println("Pod name:", s.Name, " in namespace:", s.Namespace)
	}

}
