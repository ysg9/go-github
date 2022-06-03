// Copyright 2022 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// TagProtection represents a GitHub tag protection state in a repository.
type RepoTagProtectionState struct {
	ID          *int64          `json:"id,omitempty"`
	CreatedAt   *Timestamp      `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp      `json:"updated_at,omitempty"`
	Enabled     *bool           `json:"enabled,omitempty"`
	Pattern     *string         `json:"pattern,omitempty"`
}

// ListRepoTagProtectionState lists the tag protection states of a repository.
// GitHub API docs: https://docs.github.com/en/rest/repos/tags
func (s *RepositoriesService) ListRepoTagProtectionState(ctx context.Context, owner, repo string) ([]*RepoTagProtectionState, *Response, error) {
	u := fmt.Sprintf("repos/%s/%s/tags/protection", owner, repo)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tags []*RepoTagProtectionState
	resp, err := s.client.Do(ctx, req, &tags)
	if err != nil {
		return nil, resp, err
	}
	return tags, resp, nil
}

// repoTagProtectionRequest is a subset of RepoTagProtectionState and
// is used internally by CreateRepoTagProtection to pass
// only the known fields for these endpoints.
//
type repoTagProtectionRequest struct {
	Pattern *string `json:"pattern,omitempty"`
}

// CreateRepoTagProtection adds a new tag protection state for a repository.
// GitHub API docs: https://docs.github.com/en/rest/repos/tags#create-a-tag-protection-state-for-a-repository
func (s *RepositoriesService) CreateRepoTagProtection(ctx context.Context, owner, repo string, tagState *RepoTagProtectionState) (*RepoTagProtectionState, *Response, error) {
	u := fmt.Sprintf("repos/%s/%s/tags/protection", owner, repo)

	tagProtectionReq := &repoTagProtectionRequest{
		Pattern: tagState.Pattern,
	}

	req, err := s.client.NewRequest("POST", u, tagProtectionReq)
	if err != nil {
		return nil, nil, err
	}

	r := new(RepoTagProtectionState)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	return r, resp, nil
}

// DeleteRepoTagProtection deletes a tag protection state for a repository.
// GitHub API docs: https://docs.github.com/en/rest/repos/tags#delete-a-tag-protection-state-for-a-repository
func (s *RepositoriesService) DeleteRepoTagProtection(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	u := fmt.Sprintf("repos/%s/%s/tags/protection/%d", owner, repo, id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

