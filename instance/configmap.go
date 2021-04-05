package instance

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//CreateClientset-------------------------------------------------------------------- create an clients ------------------------------------------------------
func CreateConfigmap() *apiv1.ConfigMap {
	var clientset = createClientset()
	fmt.Println("Creating Configmap ... ")
	confClient := clientset.CoreV1().ConfigMaps(apiv1.NamespaceDefault)

	configMap := &apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "predis-conf",
		},
		Data: map[string]string{
			"appendonly": "yes",
			"save":       "",
		},
	}

	resultCnf, errCnf := confClient.Create(context.TODO(), configMap, metav1.CreateOptions{})
	if errCnf != nil {
		fmt.Println(errCnf)
		panic(errCnf)
	}
	fmt.Printf("Created ConfigMap : %q\n", resultCnf.GetObjectMeta().GetName())
	return resultCnf
}
