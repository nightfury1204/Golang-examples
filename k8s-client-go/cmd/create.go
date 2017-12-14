// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typedappsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/util/intstr"
	"fmt"
)

var Kubeconfig string //contain Kubeconfig file path
// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create deployment and service",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := clientcmd.BuildConfigFromFlags("", Kubeconfig)
		if err != nil {
			log.Fatal(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}

		deploymentsClient := clientset.AppsV1beta2().Deployments(apiv1.NamespaceDefault)
		deployment, err := createDeployment(deploymentsClient)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Created deployment %q\n", deployment.GetObjectMeta().GetName())

		serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
		service, err := createService(serviceClient)
		if err!= nil {
			log.Fatal(err)
		}
		log.Printf("Service created : %q\n",service.GetObjectMeta().GetName())

		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		if err!=nil {
			log.Fatal(err)
		}
		if nodes!=nil {
			for _, node := range nodes.Items {
				if node.Name == "minikube" {
					fmt.Println("Access Url:")
					fmt.Printf("%v:%v\n",node.Status.Addresses[0].Address,service.Spec.Ports[0].NodePort)
				}
			}
		}else {
			log.Fatal("No nodes found!!")
		}
	},
}

func createDeployment(deploymentsClient typedappsv1beta2.DeploymentInterface) (*appsv1beta2.Deployment, error) {
	deployment := &appsv1beta2.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "hello-deployment",
		},
		Spec: appsv1beta2.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "hello",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "hello",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "hello-server",
							Image: "nightfury1204/hello_server",
							Args: []string{
								"serve",
							},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	//create deployment
	log.Println("Creating deployment....")

	result, err := deploymentsClient.Create(deployment)
	return result, err
}

func createService(serviceClient typedcorev1.ServiceInterface) ( *apiv1.Service, error) {
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "hello-service",
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": "hello",
			},
			Ports: []apiv1.ServicePort{
				{
					Name:       "http",
					Protocol:   apiv1.ProtocolTCP,
					Port:       8888,
					TargetPort: intstr.IntOrString{
						Type: intstr.Int,
						IntVal:8080,
					},
				},
			},
		},
	}

	log.Println("Creating service...")
	result, err := serviceClient.Create(service)

	return result,err
}

func int32Ptr(i int32) *int32 {
	return &i
}