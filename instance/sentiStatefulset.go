package instance

import (
	"context"
	"fmt"

	"gomodules.xyz/pointer"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



//CreateStatefulsetForSentinel -------------------------------------------------------------------- create the statefulset ---------------------------------------------------
func CreateStatefulsetForSentinel() {
	var clientset = CreateClientset()
	fmt.Println("Creating Statefulset ... ")
	stsClient := clientset.AppsV1().StatefulSets(apiv1.NamespaceDefault)

	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "sentinel-sts",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "sentinel",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "sentinel",
					},
				},

				Spec: apiv1.PodSpec{
					InitContainers: []apiv1.Container{
						{
							Name:            "senti-init",
							Image:          "pranganmajumder/predis:1.0.0",
							ImagePullPolicy: "IfNotPresent",

							SecurityContext: &apiv1.SecurityContext{

								RunAsUser: pointer.Int64P(0),
							},

							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      "config-vol",
									MountPath: "/conf",
								},
								{
									Name:      "script-vol",
									MountPath: "/scripts",
								},
							},
						},
					},
					Containers: []apiv1.Container{
						{
							Name:            "senti-cont",
							Image:           "redis:6.2.1",
							ImagePullPolicy: "IfNotPresent",
							Command: []string{
								"/scripts/sentinel.sh",
							},

							Env: []apiv1.EnvVar{
								{
									Name: "REPLICA_OF_PREDIS",
									Value: "3",
								},
							},

							SecurityContext: &apiv1.SecurityContext{

								RunAsUser: pointer.Int64P(0),
							},

							Ports: []apiv1.ContainerPort{
								{
									Name:          "sentinel-port",
									ContainerPort: 26379,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name: "config-vol",
									MountPath: "/conf",
								},
								{
									Name: "script-vol",
									MountPath: "/scripts",
								},
								{
									Name:      "senti-vol",
									MountPath: "/data",
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						{
							Name: "config-vol",
							VolumeSource: apiv1.VolumeSource{
								EmptyDir: &apiv1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "script-vol",
							VolumeSource: apiv1.VolumeSource{
								EmptyDir: &apiv1.EmptyDirVolumeSource{},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "senti-vol",
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
			ServiceName: "sentinel-svc",
		},
	}
	resultSts, errSts := stsClient.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if errSts != nil {
		fmt.Println(errSts)
		return
	}
	fmt.Printf("Created StatefulSet: %q\n", resultSts.GetObjectMeta().GetName())
}





