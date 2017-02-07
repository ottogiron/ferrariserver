// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "ferrariserver": worker Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateWorkerPath computes a request path to the create action of worker.
func CreateWorkerPath() string {

	return fmt.Sprintf("/v1/workers")
}

// Create a new worker
func (c *Client) CreateWorker(ctx context.Context, path string, payload *WorkerPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateWorkerRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateWorkerRequest create the request corresponding to the create action endpoint of the worker resource.
func (c *Client) NewCreateWorkerRequest(ctx context.Context, path string, payload *WorkerPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteWorkerPath computes a request path to the delete action of worker.
func DeleteWorkerPath(workerID string) string {
	param0 := workerID

	return fmt.Sprintf("/v1/workers/%s", param0)
}

// DeleteWorker makes a request to the delete action endpoint of the worker resource
func (c *Client) DeleteWorker(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteWorkerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteWorkerRequest create the request corresponding to the delete action endpoint of the worker resource.
func (c *Client) NewDeleteWorkerRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListWorkerPath computes a request path to the list action of worker.
func ListWorkerPath() string {

	return fmt.Sprintf("/v1/workers")
}

// ListWorker makes a request to the list action endpoint of the worker resource
func (c *Client) ListWorker(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListWorkerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListWorkerRequest create the request corresponding to the list action endpoint of the worker resource.
func (c *Client) NewListWorkerRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowWorkerPath computes a request path to the show action of worker.
func ShowWorkerPath(workerID string) string {
	param0 := workerID

	return fmt.Sprintf("/v1/workers/%s", param0)
}

// Retrieve a worker given an id
func (c *Client) ShowWorker(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowWorkerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowWorkerRequest create the request corresponding to the show action endpoint of the worker resource.
func (c *Client) NewShowWorkerRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateWorkerPath computes a request path to the update action of worker.
func UpdateWorkerPath(workerID string) string {
	param0 := workerID

	return fmt.Sprintf("/v1/workers/%s", param0)
}

// UpdateWorker makes a request to the update action endpoint of the worker resource
func (c *Client) UpdateWorker(ctx context.Context, path string, payload *WorkerPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateWorkerRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateWorkerRequest create the request corresponding to the update action endpoint of the worker resource.
func (c *Client) NewUpdateWorkerRequest(ctx context.Context, path string, payload *WorkerPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
