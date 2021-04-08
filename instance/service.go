package instance

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateHeadlessService() {
	fmt.Println("Creating Service ...")
	var clientset = CreateClientset()
	svcClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "predis-svc",
			Labels: map[string]string{
				"app": "predisdb",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 6379,
					Name: "redis",
				},
			},
			ClusterIP: apiv1.ClusterIPNone,
			PublishNotReadyAddresses: true,
			Selector: map[string]string{
				"app": "predisdb",
			},
		},
	}
	result, err := svcClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Service %q created\n", result.GetObjectMeta().GetName())
}


func CreateSentinelHeadlessService() {
	fmt.Println("Creating Service for sentinel...")
	var clientset = CreateClientset()
	svcClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "sentinel-svc",
			Labels: map[string]string{
				"app": "sentinel",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 26379,
					Name: "sentinel",
				},
			},
			ClusterIP: apiv1.ClusterIPNone,
			PublishNotReadyAddresses: true,
			Selector: map[string]string{
				"app": "sentinel",
			},
		},
	}
	result, err := svcClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Service %q created\n", result.GetObjectMeta().GetName())
}
