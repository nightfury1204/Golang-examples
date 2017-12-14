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
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
)

// cleanCmd represents the clean command
var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "delete deployment and service",
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
		log.Println("Deleting deployment..")
		deletePolicy := metav1.DeletePropagationForeground
		err = deploymentsClient.Delete("hello-deployment", &metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deployment deleted")

		serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
		log.Println("Deleting service..")
		err = serviceClient.Delete("hello-service", &metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("service deleted")

	},
}