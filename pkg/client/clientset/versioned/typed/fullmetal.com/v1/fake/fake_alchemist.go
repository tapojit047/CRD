/*
Copyright Mehedi Hasan.

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

package fake

import (
	"context"

	fullmetalcomv1 "github.com/tapojit047/CRD/pkg/apis/fullmetal.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlchemists implements AlchemistInterface
type FakeAlchemists struct {
	Fake *FakeFullmetalV1
	ns   string
}

var alchemistsResource = schema.GroupVersionResource{Group: "fullmetal.com", Version: "v1", Resource: "alchemists"}

var alchemistsKind = schema.GroupVersionKind{Group: "fullmetal.com", Version: "v1", Kind: "Alchemist"}

// Get takes name of the alchemist, and returns the corresponding alchemist object, and an error if there is any.
func (c *FakeAlchemists) Get(ctx context.Context, name string, options v1.GetOptions) (result *fullmetalcomv1.Alchemist, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alchemistsResource, c.ns, name), &fullmetalcomv1.Alchemist{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fullmetalcomv1.Alchemist), err
}

// List takes label and field selectors, and returns the list of Alchemists that match those selectors.
func (c *FakeAlchemists) List(ctx context.Context, opts v1.ListOptions) (result *fullmetalcomv1.AlchemistList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alchemistsResource, alchemistsKind, c.ns, opts), &fullmetalcomv1.AlchemistList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fullmetalcomv1.AlchemistList{ListMeta: obj.(*fullmetalcomv1.AlchemistList).ListMeta}
	for _, item := range obj.(*fullmetalcomv1.AlchemistList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alchemists.
func (c *FakeAlchemists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alchemistsResource, c.ns, opts))

}

// Create takes the representation of a alchemist and creates it.  Returns the server's representation of the alchemist, and an error, if there is any.
func (c *FakeAlchemists) Create(ctx context.Context, alchemist *fullmetalcomv1.Alchemist, opts v1.CreateOptions) (result *fullmetalcomv1.Alchemist, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alchemistsResource, c.ns, alchemist), &fullmetalcomv1.Alchemist{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fullmetalcomv1.Alchemist), err
}

// Update takes the representation of a alchemist and updates it. Returns the server's representation of the alchemist, and an error, if there is any.
func (c *FakeAlchemists) Update(ctx context.Context, alchemist *fullmetalcomv1.Alchemist, opts v1.UpdateOptions) (result *fullmetalcomv1.Alchemist, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alchemistsResource, c.ns, alchemist), &fullmetalcomv1.Alchemist{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fullmetalcomv1.Alchemist), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlchemists) UpdateStatus(ctx context.Context, alchemist *fullmetalcomv1.Alchemist, opts v1.UpdateOptions) (*fullmetalcomv1.Alchemist, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alchemistsResource, "status", c.ns, alchemist), &fullmetalcomv1.Alchemist{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fullmetalcomv1.Alchemist), err
}

// Delete takes name of the alchemist and deletes it. Returns an error if one occurs.
func (c *FakeAlchemists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(alchemistsResource, c.ns, name, opts), &fullmetalcomv1.Alchemist{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlchemists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alchemistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &fullmetalcomv1.AlchemistList{})
	return err
}

// Patch applies the patch and returns the patched alchemist.
func (c *FakeAlchemists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fullmetalcomv1.Alchemist, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alchemistsResource, c.ns, name, pt, data, subresources...), &fullmetalcomv1.Alchemist{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fullmetalcomv1.Alchemist), err
}
