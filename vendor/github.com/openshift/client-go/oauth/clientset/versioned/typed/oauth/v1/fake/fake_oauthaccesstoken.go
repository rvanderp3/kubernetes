// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	oauthv1 "github.com/openshift/api/oauth/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOAuthAccessTokens implements OAuthAccessTokenInterface
type FakeOAuthAccessTokens struct {
	Fake *FakeOauthV1
}

var oauthaccesstokensResource = schema.GroupVersionResource{Group: "oauth.openshift.io", Version: "v1", Resource: "oauthaccesstokens"}

var oauthaccesstokensKind = schema.GroupVersionKind{Group: "oauth.openshift.io", Version: "v1", Kind: "OAuthAccessToken"}

// Get takes name of the oAuthAccessToken, and returns the corresponding oAuthAccessToken object, and an error if there is any.
func (c *FakeOAuthAccessTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *oauthv1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(oauthaccesstokensResource, name), &oauthv1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAccessToken), err
}

// List takes label and field selectors, and returns the list of OAuthAccessTokens that match those selectors.
func (c *FakeOAuthAccessTokens) List(ctx context.Context, opts v1.ListOptions) (result *oauthv1.OAuthAccessTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(oauthaccesstokensResource, oauthaccesstokensKind, opts), &oauthv1.OAuthAccessTokenList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &oauthv1.OAuthAccessTokenList{ListMeta: obj.(*oauthv1.OAuthAccessTokenList).ListMeta}
	for _, item := range obj.(*oauthv1.OAuthAccessTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oAuthAccessTokens.
func (c *FakeOAuthAccessTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(oauthaccesstokensResource, opts))
}

// Create takes the representation of a oAuthAccessToken and creates it.  Returns the server's representation of the oAuthAccessToken, and an error, if there is any.
func (c *FakeOAuthAccessTokens) Create(ctx context.Context, oAuthAccessToken *oauthv1.OAuthAccessToken, opts v1.CreateOptions) (result *oauthv1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(oauthaccesstokensResource, oAuthAccessToken), &oauthv1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAccessToken), err
}

// Update takes the representation of a oAuthAccessToken and updates it. Returns the server's representation of the oAuthAccessToken, and an error, if there is any.
func (c *FakeOAuthAccessTokens) Update(ctx context.Context, oAuthAccessToken *oauthv1.OAuthAccessToken, opts v1.UpdateOptions) (result *oauthv1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(oauthaccesstokensResource, oAuthAccessToken), &oauthv1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAccessToken), err
}

// Delete takes name of the oAuthAccessToken and deletes it. Returns an error if one occurs.
func (c *FakeOAuthAccessTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(oauthaccesstokensResource, name), &oauthv1.OAuthAccessToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOAuthAccessTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(oauthaccesstokensResource, listOpts)

	_, err := c.Fake.Invokes(action, &oauthv1.OAuthAccessTokenList{})
	return err
}

// Patch applies the patch and returns the patched oAuthAccessToken.
func (c *FakeOAuthAccessTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *oauthv1.OAuthAccessToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(oauthaccesstokensResource, name, pt, data, subresources...), &oauthv1.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAccessToken), err
}