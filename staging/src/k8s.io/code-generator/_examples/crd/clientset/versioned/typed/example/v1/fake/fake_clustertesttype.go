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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	example_v1 "k8s.io/code-generator/_examples/crd/apis/example/v1"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

// FakeClusterTestTypes implements ClusterTestTypeInterface
type FakeClusterTestTypes struct {
	Fake *FakeExampleV1
}

var clustertesttypesResource = schema.GroupVersionResource{Group: "example.crd.code-generator.k8s.io", Version: "v1", Resource: "clustertesttypes"}

var clustertesttypesKind = schema.GroupVersionKind{Group: "example.crd.code-generator.k8s.io", Version: "v1", Kind: "ClusterTestType"}

// Get takes name of the clusterTestType, and returns the corresponding clusterTestType object, and an error if there is any.
func (c *FakeClusterTestTypes) Get(name string, options v1.GetOptions) (result *example_v1.ClusterTestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustertesttypesResource, name), &example_v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.ClusterTestType), err
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *FakeClusterTestTypes) List(opts v1.ListOptions) (result *example_v1.ClusterTestTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustertesttypesResource, clustertesttypesKind, opts), &example_v1.ClusterTestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &example_v1.ClusterTestTypeList{}
	for _, item := range obj.(*example_v1.ClusterTestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTestTypes.
func (c *FakeClusterTestTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustertesttypesResource, opts))
}

// Create takes the representation of a clusterTestType and creates it.  Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *FakeClusterTestTypes) Create(clusterTestType *example_v1.ClusterTestType) (result *example_v1.ClusterTestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustertesttypesResource, clusterTestType), &example_v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.ClusterTestType), err
}

// Update takes the representation of a clusterTestType and updates it. Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *FakeClusterTestTypes) Update(clusterTestType *example_v1.ClusterTestType) (result *example_v1.ClusterTestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustertesttypesResource, clusterTestType), &example_v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.ClusterTestType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterTestTypes) UpdateStatus(clusterTestType *example_v1.ClusterTestType) (*example_v1.ClusterTestType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustertesttypesResource, "status", clusterTestType), &example_v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.ClusterTestType), err
}

// Delete takes name of the clusterTestType and deletes it. Returns an error if one occurs.
func (c *FakeClusterTestTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clustertesttypesResource, name), &example_v1.ClusterTestType{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTestTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustertesttypesResource, listOptions)

	_, err := c.Fake.Invokes(action, &example_v1.ClusterTestTypeList{})
	return err
}

// Patch applies the patch and returns the patched clusterTestType.
func (c *FakeClusterTestTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *example_v1.ClusterTestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustertesttypesResource, name, data, subresources...), &example_v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example_v1.ClusterTestType), err
}

// GetScale takes name of the clusterTestType, and returns the corresponding scale object, and an error if there is any.
func (c *FakeClusterTestTypes) GetScale(clusterTestTypeName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(clustertesttypesResource, clusterTestTypeName), &autoscaling.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeClusterTestTypes) UpdateScale(clusterTestTypeName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustertesttypesResource, "scale", scale), &autoscaling.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}