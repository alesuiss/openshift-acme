// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	config_v1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngresses implements IngressInterface
type FakeIngresses struct {
	Fake *FakeConfigV1
}

var ingressesResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "ingresses"}

var ingressesKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "Ingress"}

// Get takes name of the ingress, and returns the corresponding ingress object, and an error if there is any.
func (c *FakeIngresses) Get(name string, options v1.GetOptions) (result *config_v1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ingressesResource, name), &config_v1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.Ingress), err
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors.
func (c *FakeIngresses) List(opts v1.ListOptions) (result *config_v1.IngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ingressesResource, ingressesKind, opts), &config_v1.IngressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &config_v1.IngressList{ListMeta: obj.(*config_v1.IngressList).ListMeta}
	for _, item := range obj.(*config_v1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingresses.
func (c *FakeIngresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ingressesResource, opts))
}

// Create takes the representation of a ingress and creates it.  Returns the server's representation of the ingress, and an error, if there is any.
func (c *FakeIngresses) Create(ingress *config_v1.Ingress) (result *config_v1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ingressesResource, ingress), &config_v1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.Ingress), err
}

// Update takes the representation of a ingress and updates it. Returns the server's representation of the ingress, and an error, if there is any.
func (c *FakeIngresses) Update(ingress *config_v1.Ingress) (result *config_v1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ingressesResource, ingress), &config_v1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.Ingress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIngresses) UpdateStatus(ingress *config_v1.Ingress) (*config_v1.Ingress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ingressesResource, "status", ingress), &config_v1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.Ingress), err
}

// Delete takes name of the ingress and deletes it. Returns an error if one occurs.
func (c *FakeIngresses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ingressesResource, name), &config_v1.Ingress{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ingressesResource, listOptions)

	_, err := c.Fake.Invokes(action, &config_v1.IngressList{})
	return err
}

// Patch applies the patch and returns the patched ingress.
func (c *FakeIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *config_v1.Ingress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ingressesResource, name, data, subresources...), &config_v1.Ingress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.Ingress), err
}
