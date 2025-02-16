package kubernetes

import (
	"encoding/base64"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"

	"github.com/DB-Vincent/inkube/internal/config"
)

type KubernetesConnection struct {
	Clientset     *kubernetes.Clientset
	RestConfig    *rest.Config
	MetricsClient *metricsclientset.Clientset
}

func (kubeconn *KubernetesConnection) createRestConfig(config *config.Config) (*rest.Config, error) {
	caData, err := base64.StdEncoding.DecodeString(config.Cluster.CertificateAuthorityData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode CA data: %v", err)
	}

	clientCertData, err := base64.StdEncoding.DecodeString(config.Auth.ClientCertificateData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode client certificate data: %v", err)
	}

	clientKeyData, err := base64.StdEncoding.DecodeString(config.Auth.ClientKeyData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode client key data: %v", err)
	}

	kubeconn.RestConfig = &rest.Config{
		Host: config.Cluster.Server,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   caData,
			CertData: clientCertData,
			KeyData:  clientKeyData,
		},
	}

	return kubeconn.RestConfig, nil
}

func (kubeconn *KubernetesConnection) ConnectToCluster(config *config.Config) (*kubernetes.Clientset, error) {
  err := fmt.Errorf("") 

  kubeconn.RestConfig, err = kubeconn.createRestConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create REST config: %v", err)
	}

	kubeconn.Clientset, err = kubernetes.NewForConfig(kubeconn.RestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	kubeconn.MetricsClient, err = metricsclientset.NewForConfig(kubeconn.RestConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics client: %v", err)
	}

	return kubeconn.Clientset, nil
}
