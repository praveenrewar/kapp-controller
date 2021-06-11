// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/internalpackaging/v1alpha1"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type InternalV1alpha1Interface interface {
	RESTClient() rest.Interface
	InternalPackagesGetter
	InternalPackageVersionsGetter
}

// InternalV1alpha1Client is used to interact with features provided by the internal.packaging.carvel.dev group.
type InternalV1alpha1Client struct {
	restClient rest.Interface
}

func (c *InternalV1alpha1Client) InternalPackages(namespace string) InternalPackageInterface {
	return newInternalPackages(c, namespace)
}

func (c *InternalV1alpha1Client) InternalPackageVersions(namespace string) InternalPackageVersionInterface {
	return newInternalPackageVersions(c, namespace)
}

// NewForConfig creates a new InternalV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*InternalV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &InternalV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new InternalV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *InternalV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new InternalV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *InternalV1alpha1Client {
	return &InternalV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *InternalV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
