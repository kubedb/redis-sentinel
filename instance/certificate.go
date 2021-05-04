package instance

import (
	"context"
	"flag"
	"fmt"
	cm_api "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	cm "github.com/jetstack/cert-manager/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	cm_util "kmodules.xyz/cert-manager-util/certmanager/v1"
	kutil "kmodules.xyz/client-go"
	"path/filepath"
	"time"
)

func CreateServerCert() (kutil.VerbType, error) {

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

	CertManagerClient, err := cm.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Creating Certificate ... ")

	cert := func(in *cm_api.Certificate) *cm_api.Certificate {

		in.Spec = cm_api.CertificateSpec{
			Subject: &cm_api.X509Subject{
				Organizations:       nil,
				Countries:           nil,
				OrganizationalUnits: nil,
				Localities:          nil,
				Provinces:           nil,
				StreetAddresses:     nil,
				PostalCodes:         nil,
				SerialNumber:        "",
			},
			CommonName: "admin",
			Duration: &metav1.Duration{
				Duration: 2*time.Hour,
			},
			RenewBefore: &metav1.Duration{
				Duration: 1*time.Hour,
			},
			DNSNames: []string{
				"predis-svc.default.svc",
			},
			IPAddresses: []string{
				"127.0.0.1",
			},
			URIs:           nil,
			EmailAddresses: nil,
			SecretName:     "example-com-tls",

			IssuerRef: cmmeta.ObjectReference{
				Name:  "ca-issuer",
				Kind:  "Issuer",
				Group: "cert-manager.io",
			},
			IsCA: false,
			Usages: []cm_api.KeyUsage{
				cm_api.UsageDigitalSignature,
				cm_api.UsageKeyEncipherment,
				cm_api.UsageServerAuth,
				cm_api.UsageClientAuth,
			},
		}
		return in
	}


	_, vt, err := cm_util.CreateOrPatchCertificate(
		context.TODO(),
		CertManagerClient.CertmanagerV1(),
		metav1.ObjectMeta{
			Name:      "example-com",
			Namespace: "demo",
		},
		cert,
		metav1.PatchOptions{})
	return vt, err

}
