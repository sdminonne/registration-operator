// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/open-cluster-management/api/client/nucleus/clientset/versioned/typed/nucleus/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNucleusV1 struct {
	*testing.Fake
}

func (c *FakeNucleusV1) HubCores() v1.HubCoreInterface {
	return &FakeHubCores{c}
}

func (c *FakeNucleusV1) SpokeCores() v1.SpokeCoreInterface {
	return &FakeSpokeCores{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNucleusV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
