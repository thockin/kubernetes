/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	apiauthenticationv1alpha1 "k8s.io/api/authentication/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

type AuthenticationV1alpha1Interface interface {
	RESTClient() clientgorest.Interface
	SelfSubjectReviewsGetter
}

// AuthenticationV1alpha1Client is used to interact with features provided by the authentication.k8s.io group.
type AuthenticationV1alpha1Client struct {
	restClient clientgorest.Interface
}

func (c *AuthenticationV1alpha1Client) SelfSubjectReviews() SelfSubjectReviewInterface {
	return newSelfSubjectReviews(c)
}

// NewForConfig creates a new AuthenticationV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *clientgorest.Config) (*AuthenticationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := clientgorest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new AuthenticationV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *clientgorest.Config, h *http.Client) (*AuthenticationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := clientgorest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &AuthenticationV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new AuthenticationV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *clientgorest.Config) *AuthenticationV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AuthenticationV1alpha1Client for the given RESTClient.
func New(c clientgorest.Interface) *AuthenticationV1alpha1Client {
	return &AuthenticationV1alpha1Client{c}
}

func setConfigDefaults(config *clientgorest.Config) error {
	gv := apiauthenticationv1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = clientgorest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AuthenticationV1alpha1Client) RESTClient() clientgorest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
