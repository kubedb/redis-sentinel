package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

//CreateClientset-------------------------------------------------------------------- create an clients ------------------------------------------------------
func CreateClientset() kubernetes.Interface {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func main() {
	var clientset = CreateClientset()
	// emon test
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "predis-sts-0", metav1.GetOptions{})
	if err != nil {
		fmt.Println("........................errrrrr.....", err)
	}
	labels := pod.Labels
	labels["role"] = "master"
	pod.Labels = labels
	pod, err = clientset.CoreV1().Pods("default").Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		fmt.Println("...........fdhjadskjdfaghkad", err)
	}
	fmt.Println(pod.Labels)

}
