package claims

import (
	"net/http"

	"github.com/gophercloud/gophercloud"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToClaimCreateRequest() (map[string]interface{}, string, error)
}

// CreateOpts params to be used with Create.
type CreateOpts struct {
	// Sets the TTL for the claim. When the claim expires un-deleted messages will be able to be claimed again.
	TTL int `json:"ttl,omitempty"`

	// Sets the Grace period for the claimed messages. The server extends the lifetime of claimed messages
	// to be at least as long as the lifetime of the claim itself, plus the specified grace period.
	Grace int `json:"grace,omitempty"`

	// Set the limit of messages returned by create.
	Limit int `q:"limit" json:"-"`
}

// ToClaimCreateRequest assembles a body and URL for a Create request based on
// the contents of a CreateOpts.
func (opts CreateOpts) ToClaimCreateRequest() (map[string]interface{}, string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return nil, q.String(), err
	}

	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return b, "", err
	}
	return b, q.String(), err
}

// Create creates a Claim that claims messages on a specified queue.
func Create(client *gophercloud.ServiceClient, queueName string, opts CreateOptsBuilder) (r CreateResult) {
	b, q, err := opts.ToClaimCreateRequest()
	if err != nil {
		r.Err = err
		return
	}

	url := createURL(client, queueName)
	if q != "" {
		url += q
	}

	var resp *http.Response
	resp, r.Err = client.Post(url, b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{201, 204},
	})
	// If the Claim has no content return an empty CreateResult
	if resp.StatusCode == 204 {
		r.Body = CreateResult{}
	} else {
		r.Body = resp.Body
	}
	return
}
