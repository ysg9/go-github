package github

import (
	"context"
	"fmt"
	"time"
)

// Credential represents an individual credential authorization
type Credential struct {
	Login							*string			`json:"login,omitempty"`
	CredentialID					*int64			`json:"credential_id,omitempty"`
	CredentialType					*string			`json:"credential_type,omitempty"`
	TokenLastEight					*string			`json:"token_last_eight,omitempty"`
	CredentialAuthorizedAt			*time.Time		`json:"credential_authorized_at,omitempty"`
	AuthorizedCredentialExpiresAt	*time.Time		`json:"authorized_credential_expires_at,omitempty"`
	Scopes							[]Scope	`json:"scopes,omitempty"`
}

// CredentialAuthorizationsOptions specifies the optional parameters to the
// OrganizationsService.CredentialAuthorizations method.
type CredentialAuthorizationsOptions struct {
	// Since filters Organizations by ID.
	Login *string `url:"login,omitempty"`

	// Note: Pagination is powered exclusively by the Since parameter,
	// ListOptions.Page has no effect.
	// ListOptions.PerPage controls an undocumented GitHub API parameter.
	ListOptions
}

// GitHub API docs: https://docs.github.com/en/rest/orgs/orgs#list-saml-sso-authorizations-for-an-organization
func (s *OrganizationsService) CredentialAuthorizations(ctx context.Context, org string, opts *CredentialAuthorizationsOptions) ([]*Credential, *Response, error) {
	u := fmt.Sprintf("orgs/%s/credential-authorizations", org)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	creds := []*Credential{}
	resp, err := s.client.Do(ctx, req, &creds)
	if err != nil {
		return nil, resp, err
	}
	return creds, resp, nil
}

