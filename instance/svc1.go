package instance

import (
	"context"
	"fmt"


	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

func CreateSvc1() {
	fmt.Println("Creating svc1 ...")
	var clientset = CreateClientset()
	svcClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "predis-svc1-master",
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
			Selector: map[string]string{
				"app": "predisdb",
				"role": "master",
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
