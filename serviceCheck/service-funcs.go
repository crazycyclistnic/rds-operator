//Author: Nic Hima
package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/util/intstr"

	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const NAMESPACE string = "default"

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
	//Grab all deployments in default namespace (could be changed to select different possible namespaces)
	deploymentInterface := c.ExtensionsV1beta1().Deployments(NAMESPACE)
	//Check to see if deployment by name passed in as arg is present
	deployment, err := deploymentInterface.Get(conf.deploymentName, metav1.GetOptions{})
	if err != nil || deployment == nil { //if deployment not present:
		fmt.Fprintf(os.Stderr, "Could not find deployment %s\n", conf.deploymentName)
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "Deployment %s is present in cluster\n", deployment.GetName())

	//Define service spec based on deployment
	serviceSpec := &api.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.Name + "-service",
			Namespace: deployment.Namespace,
		},
		Spec: api.ServiceSpec{
			Type:     api.ServiceTypeClusterIP,
			Selector: deployment.Spec.Selector.MatchLabels,
			Ports: []api.ServicePort{
				api.ServicePort{
					Port:       80,
					TargetPort: intstr.FromInt(80),
				},
			},
		},
	}

	fmt.Fprintf(os.Stdout, "Creating service named: %s...\n", serviceSpec.GetName())
	//Create service
	_, err = c.CoreV1().Services(deployment.Namespace).Create(serviceSpec)
	if err != nil {
		log.Fatal(err)
	}

}
