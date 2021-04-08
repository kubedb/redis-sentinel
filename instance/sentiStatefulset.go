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
					Containers: []apiv1.Container{
						{
							Name:            "senti-cont",
							Image:           "redis:6.2.1",
							ImagePullPolicy: "IfNotPresent",
							Command: []string{
								"/scripts/run.sh",
								//"sleep",
								//"360000",

								//"cp",
								//"/config/sentinel.conf",
								//"/data/sentinel.conf",
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
									MountPath: "/config",
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
								ConfigMap: &apiv1.ConfigMapVolumeSource{
									LocalObjectReference: apiv1.LocalObjectReference{
										Name: "senti-conf",
									},
									DefaultMode: int32Ptr(0777),
								},
							},
						},
						{
							Name: "script-vol",
							VolumeSource: apiv1.VolumeSource{
								ConfigMap: &apiv1.ConfigMapVolumeSource{
									LocalObjectReference: apiv1.LocalObjectReference{
										Name: "sentinel-scripts",
									},
									DefaultMode: int32Ptr(0777),
								},
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





