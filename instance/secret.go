package instance

import (
	"context"
	"fmt"
	_ "gomodules.xyz/pointer"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateSecret() {
	var clientset = CreateClientset()
	fmt.Println("Creating Secret ... ")
	secretClient := clientset.CoreV1().Secrets("demo")

	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: "predis-secret",
		},
		Type: apiv1.SecretType("Opaque"),
		StringData: map[string]string{
			apiv1.BasicAuthUsernameKey: "admin",
			apiv1.BasicAuthPasswordKey: "admin",
		},
	}


	resultSts, errSts := secretClient.Create(context.TODO(), secret, metav1.CreateOptions{})
	if errSts != nil {
		fmt.Println(errSts)
		return
	}
	fmt.Printf("Created Secret: %q\n", resultSts.GetObjectMeta().GetName())

}