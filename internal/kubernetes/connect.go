package kubernetes

import (
	"encoding/base64"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/DB-Vincent/inkube/internal/config"
)

func createRestConfig(config *config.Config) (*rest.Config, error) {
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

	return &rest.Config{
		Host: config.Cluster.Server,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   caData,
			CertData: clientCertData,
			KeyData:  clientKeyData,
		},
	}, nil
}

func ConnectToCluster() (*kubernetes.Clientset, error) {
	configPath := "config.toml"
	config, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	restConfig, err := createRestConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create REST config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	return clientset, nil
}
