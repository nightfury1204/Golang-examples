package main

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
	"k8s.io/client-go/util/homedir"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var patchTypes = []string{"json", "merge", "strategic"}

func main() {
	var kubeconfig string //contain kubeconfig file path
	home := homedir.HomeDir()
	//log.Println(home)
	cmd := &cobra.Command{
		Use:               "deploy",
		Short:             "Create deployment and service",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err)
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err)
			}

			deploymentsClient := clientset.AppsV1beta2().Deployments(apiv1.NamespaceDefault)
			deployment, err := createDeployment(deploymentsClient)
			if err != nil {
				panic(err)
			}

			log.Printf("Created deployment %q\n", deployment.GetObjectMeta().GetName())

			serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
			service, err := createService(serviceClient)
			if err!= nil {
				panic(err)
			}
			log.Printf("Service created : %q\n",service.GetObjectMeta().GetName())

		},
	}

	cleanCmd := &cobra.Command{
		Use:               "clean",
		Short:             "delete deployment",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err)
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err)
			}

			deploymentsClient := clientset.AppsV1beta2().Deployments(apiv1.NamespaceDefault)
			log.Println("Deleting deployment..")
			deletePolicy := metav1.DeletePropagationForeground
			err = deploymentsClient.Delete("hello-deployment", &metav1.DeleteOptions{
				PropagationPolicy: &deletePolicy,
			})
			if err != nil {
				panic(err)
			}
			log.Println("Deployment deleted")
		},
	}
	cmd.Flags().StringVarP(&kubeconfig, "configPath", "c", home+"/.kube/config", "kube config path")

	cmd.AddCommand(cleanCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
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
							Image: "nightfury1204/hello_serve",
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

	result, err := serviceClient.Create(service)

	return result,err
}

func int32Ptr(i int32) *int32 {
	return &i
}
