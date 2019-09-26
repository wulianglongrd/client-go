// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "istio.io/client-go/pkg/clientset/versioned"
	authenticationv1alpha1 "istio.io/client-go/pkg/clientset/versioned/typed/authentication/v1alpha1"
	fakeauthenticationv1alpha1 "istio.io/client-go/pkg/clientset/versioned/typed/authentication/v1alpha1/fake"
	configv1alpha2 "istio.io/client-go/pkg/clientset/versioned/typed/config/v1alpha2"
	fakeconfigv1alpha2 "istio.io/client-go/pkg/clientset/versioned/typed/config/v1alpha2/fake"
	networkingv1alpha3 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1alpha3"
	fakenetworkingv1alpha3 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1alpha3/fake"
	rbacv1alpha1 "istio.io/client-go/pkg/clientset/versioned/typed/rbac/v1alpha1"
	fakerbacv1alpha1 "istio.io/client-go/pkg/clientset/versioned/typed/rbac/v1alpha1/fake"
	securityv1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/security/v1beta1"
	fakesecurityv1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/security/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// AuthenticationV1alpha1 retrieves the AuthenticationV1alpha1Client
func (c *Clientset) AuthenticationV1alpha1() authenticationv1alpha1.AuthenticationV1alpha1Interface {
	return &fakeauthenticationv1alpha1.FakeAuthenticationV1alpha1{Fake: &c.Fake}
}

// Authentication retrieves the AuthenticationV1alpha1Client
func (c *Clientset) Authentication() authenticationv1alpha1.AuthenticationV1alpha1Interface {
	return &fakeauthenticationv1alpha1.FakeAuthenticationV1alpha1{Fake: &c.Fake}
}

// ConfigV1alpha2 retrieves the ConfigV1alpha2Client
func (c *Clientset) ConfigV1alpha2() configv1alpha2.ConfigV1alpha2Interface {
	return &fakeconfigv1alpha2.FakeConfigV1alpha2{Fake: &c.Fake}
}

// Config retrieves the ConfigV1alpha2Client
func (c *Clientset) Config() configv1alpha2.ConfigV1alpha2Interface {
	return &fakeconfigv1alpha2.FakeConfigV1alpha2{Fake: &c.Fake}
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return &fakenetworkingv1alpha3.FakeNetworkingV1alpha3{Fake: &c.Fake}
}

// Networking retrieves the NetworkingV1alpha3Client
func (c *Clientset) Networking() networkingv1alpha3.NetworkingV1alpha3Interface {
	return &fakenetworkingv1alpha3.FakeNetworkingV1alpha3{Fake: &c.Fake}
}

// RbacV1alpha1 retrieves the RbacV1alpha1Client
func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return &fakerbacv1alpha1.FakeRbacV1alpha1{Fake: &c.Fake}
}

// Rbac retrieves the RbacV1alpha1Client
func (c *Clientset) Rbac() rbacv1alpha1.RbacV1alpha1Interface {
	return &fakerbacv1alpha1.FakeRbacV1alpha1{Fake: &c.Fake}
}

// SecurityV1beta1 retrieves the SecurityV1beta1Client
func (c *Clientset) SecurityV1beta1() securityv1beta1.SecurityV1beta1Interface {
	return &fakesecurityv1beta1.FakeSecurityV1beta1{Fake: &c.Fake}
}

// Security retrieves the SecurityV1beta1Client
func (c *Clientset) Security() securityv1beta1.SecurityV1beta1Interface {
	return &fakesecurityv1beta1.FakeSecurityV1beta1{Fake: &c.Fake}
}
