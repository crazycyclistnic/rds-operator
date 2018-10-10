//Author: Nic Hima
package main

import (
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const namespace string = "default"

//interface used to dictate which flags func to run to create service
type configAction interface {
	Execute(c *kubernetes.Clientset)
}

type deploymentMethod struct {
	deploymentName string
}

type selectorMethod struct {
	//TODO: create list of label:value pairs to be used to create service
}

//creation of service by taking selectors from a deployment
func (conf *deploymentMethod) Execute(c *kubernetes.Clientset) {
	fmt.Fprintf(os.Stdout, "Deployment that service is being created for is: %s\n", conf.deploymentName)
	deploymentInterface := c.ExtensionsV1beta1().Deployments(namespace)
	deployment, err := deploymentInterface.Get(conf.deploymentName, metav1.GetOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find deployment %s\n", conf.deploymentName)
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "Deployment %s is present in cluster\n", deployment.GetName())

	//TODO: using the found deployment, create a service of type clusterIP that matches the selectors in
	//the deployment
}
