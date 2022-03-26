package github

import (
	"context"
	"fmt"
)

type CustomRole struct {
	ID		*int64	`json:"id",omitempty"`
	Name	*string	`json:"name",omitempty"`
}

type CustomRoleList struct {
	TotalCount	*int64			`json:"total_count",omitempty"`
	CustomRoles	[]CustomRole	`json:"custom_roles",omitempty"`
}

func (s *OrganizationsService) ListCustomRoles(ctx context.Context, org string) (*CustomRoleList, *Response, error) {
	u := fmt.Sprintf("orgs/%v/custom_roles", org)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	customRoleList := new(CustomRoleList)
	resp, err := s.client.Do(ctx, req, customRoleList)
	if err != nil {
		return nil, resp, err
	}

	return customRoleList, resp, nil
}

