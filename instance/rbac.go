package instance

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	v1alpha1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func CreateRoleBinding(ns string, roleName string, saName string)  {

	rolebinding := &v1alpha1.ClusterRoleBinding{
		ObjectMeta :	metav1.ObjectMeta{
				Name:      roleName,
				Namespace: ns,
		},
		Subjects: []v1alpha1.Subject{
			{
				Kind:      v1alpha1.ServiceAccountKind,
				Name:      saName,
				Namespace: ns,
			},
		},
		RoleRef: v1alpha1.RoleRef{
			APIGroup: v1alpha1.GroupName,
			Kind:     "ClusterRole",
			Name:     roleName,
		},

	}

	var clientset = CreateClientset()
	fmt.Println("Creating Rbac Rolebinding... ")
	results, err := clientset.RbacV1().ClusterRoleBindings().Create(context.TODO(), rolebinding ,metav1.CreateOptions{})


	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created StatefulSet: %q\n", results.GetObjectMeta().GetName())

}




func CreateRole(ns string, name string)  {



	role := &v1alpha1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       ns,
		},
		Rules: []v1alpha1.PolicyRule{
			{
				APIGroups: []string{"*"},
				Resources: []string{"*"},
				Verbs:     []string{"*"},
			},
		},
	}

	var clientset = CreateClientset()
	fmt.Println("Creating Rbac ... ")
	results, err := clientset.RbacV1().ClusterRoles().Create(context.TODO(), role , metav1.CreateOptions{})


	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created StatefulSet: %q\n", results.GetObjectMeta().GetName())

}



func  CreateServiceAccount(ns string , saName string)  {

	sa := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      saName,
			Namespace: ns,
		},
	}

	var clientset = CreateClientset()
	fmt.Println("Creating Service account... ")
	results, err := clientset.CoreV1().ServiceAccounts(sa.Namespace).Create(context.TODO(), sa, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created StatefulSet: %q\n", results.GetObjectMeta().GetName())

}