package instance

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

//createClientset-------------------------------------------------------------------- create an clients ------------------------------------------------------
func createClientset() kubernetes.Interface {
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

//CreateStatefulset -------------------------------------------------------------------- create the statefulset ---------------------------------------------------
func CreateStatefulset(image string, replica int32) {

	var clientset = createClientset()
	fmt.Println("Creating Statefulset ... ")
	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)

	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "predis-sts",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: int32Ptr(replica),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "predisdb",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "predisdb",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "predis",
							Image:           "docker.io/bitnami/redis:6.0.12-debian-10-r3",
							ImagePullPolicy: "IfNotPresent",

							Ports: []apiv1.ContainerPort{
								{
									Name:          "predis-port",
									ContainerPort: 6379,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "predis-vol",
									MountPath: "/data/db",
								},
								{
									Name:      "config-vol",
									MountPath: "/data/predis-data",
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name: "ALLOW_EMPTY_PASSWORD",
									Value: "yes",
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						{
							Name: "config-vol",
							VolumeSource: apiv1.VolumeSource{
								ConfigMap: &apiv1.ConfigMapVolumeSource{
									LocalObjectReference: apiv1.LocalObjectReference{
										Name: "predis-conf",
									},
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "predis-vol",
					},
					Spec: apiv1.PersistentVolumeClaimSpec{
						AccessModes:      []apiv1.PersistentVolumeAccessMode{apiv1.ReadWriteOnce},
						StorageClassName: strPtr("standard"),
						Resources: apiv1.ResourceRequirements{
							Requests: apiv1.ResourceList{
								apiv1.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
					},
				},
			},
			ServiceName: "predis-service",
		},
	}

	resultSts, errSts := stsClient.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if errSts != nil {
		fmt.Println(errSts)
		return
	}
	fmt.Printf("Created StatefulSet: %q\n", resultSts.GetObjectMeta().GetName())
}

func ListStatefulSet() {
	fmt.Println("*****   Listing all StatefulSets   ******  ")
	var clientset = createClientset()

	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)
	list, err := stsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range list.Items {
		fmt.Printf("---> %s (%d replicas)\n", item.Name, *item.Spec.Replicas)
	}
}

func DeleteStatefulSet() {
	var clientset = createClientset()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	fmt.Println("Deleting deployment...")
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), "demo-deployment", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment.")
}

//func UpdateStatefulSet() {
//	fmt.Printf("Updating StatefulSet %q replicas to %d\n", stsName, replicas)
//	var clientset := createClientset()
//	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)
//	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
//		result, getErr := stsClient.Get(context.TODO(), stsName, metav1.GetOptions{})
//		if getErr != nil {
//			panic(fmt.Errorf("Failed to get latest version of StatefulSet: %v", getErr))
//		}
//		result.Spec.Replicas = int32Ptr(replicas)
//		result.Spec.Template.Spec.Containers[0].Image = image
//		_, updateErr := stsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
//		return updateErr
//	})
//	if retryErr != nil {
//		panic(fmt.Errorf("Update failed: %v", retryErr))
//	}
//	fmt.Printf("Statefulset %q Successfully updated\n", stsName)
//}

func int32Ptr(i int32) *int32 { return &i }
func strPtr(s string) *string {
	return &s
}
