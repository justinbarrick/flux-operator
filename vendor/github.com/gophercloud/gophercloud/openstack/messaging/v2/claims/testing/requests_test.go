package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	createOpts := claims.CreateOpts{
		TTL:   3600,
		Grace: 3600,
		Limit: 10,
	}

	actual, err := claims.Create(fake.ServiceClient(), QueueName, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CreatedClaim, actual)
}
