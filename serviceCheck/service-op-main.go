//Author: Nic Hima
//goal of this is to create or modify services of type cluster IP
//as of right now, only flag that works will be --deployment [deploymentName]
//will check if service matching deployment selectors is already made
//if not, creates the service
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	//"time"

	//"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//check to make sure commands present
	if len(os.Args[1:]) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [--deployment deploymentName]\n", os.Args[0])
		os.Exit(1)
	}

	cmdArg := os.Args[1]
	cmdParams := os.Args[2]

	var conf configAction

	//check what method by which to create service from
	switch cmdArg {
	//creates interface to handle different flags & thus different logic
	case "--deployment":
		{

			conf = &deploymentMethod{
				deploymentName: cmdParams,
			}
		}
	case "--selector":
		{
			//TODO: grab next arg as comma seperated key:val pairs, and use to create service
		}
	default:
		{
			fmt.Fprintf(os.Stderr, "unknown operation: %s\n", cmdArg)
			os.Exit(2)
		}
	}

	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	//create clientset
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	//pass in c as param to whatever method is being executed based on cmdArgs
	conf.Execute(c)
}
