//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// TestAPICall captures the arguments to one of the API calls.
type TestAPICall struct {
	// Action is the REST action (GET, PUT, etc) of the call
	Action string
	// URL is the URL to send to
	URL string
	// Payload is the string representation of the payload
	Payload string
}

// TestClient is a mock client to use for unit testing some of the
// function calls and actions that would normally need to connect
// with a host.
type TestClient struct {
	// calls collects any API calls made through the client
	calls []TestAPICall
	// CustomReturnForActions can be used to define custom
	// return for actions, valid keys are:
	// http.MethodGet, http.MethodPost, http.MethodPut,
	// http.MethodPatch and http.MethodDelete.
	// For each key it is possible to define a list of
	// returns (in the order they should be returned).
	CustomReturnForActions map[string][]interface{}
}

// CapturedCalls gets all calls that were made through this instance
func (c *TestClient) CapturedCalls() []TestAPICall {
	return c.calls
}

// actionCount returns how many actions
// of a specific type were already recorded.
func (c *TestClient) actionCount(action string) int {
	var actionCount int
	for _, call := range c.calls {
		if call.Action == action {
			actionCount = actionCount + 1
		}
	}
	return actionCount
}

// getCustomReturnForAction gets the custom return for the action
func (c *TestClient) getCustomReturnForAction(action string) interface{} {
	switch action {
	case http.MethodGet, http.MethodPost,
		http.MethodPut, http.MethodPatch,
		http.MethodDelete:
		customReturnForAction, ok := c.CustomReturnForActions[action]
		if !ok ||
			customReturnForAction == nil ||
			customReturnForAction[c.actionCountIndex(action)] == nil {
			return nil
		}
		return customReturnForAction[c.actionCountIndex(action)]
	}
	return nil
}

// actionCountIndex returns the index that should be used
// to get the custom return from CustomReturnForActions.
func (c *TestClient) actionCountIndex(action string) int {
	return c.actionCount(action) - 1
}

// getPayloadToBeRecorded returns the payload that will
// be recorded for the call.
func (c *TestClient) getPayloadToBeRecorded(payload interface{}) string {
	// when possible do Marshal/Unmarshal of the payload
	// in order to have the json keys when using interfaces
	// in the payload.
	if payload != nil {
		payloadMarshaled, err := json.Marshal(payload)
		if err != nil {
			return fmt.Sprintf("%v", payload)
		}
		var payloadInterface interface{}
		err = json.Unmarshal(payloadMarshaled, &payloadInterface)
		if err != nil {
			return fmt.Sprintf("%v", payload)
		}
		return fmt.Sprintf("%v", payloadInterface)
	}

	return ""
}

// Reset resets the captured information for this mock client.
func (c *TestClient) Reset() {
	c.calls = []TestAPICall{}
	c.CustomReturnForActions = map[string][]interface{}{}
}

// recordCall is a helper to record any API calls made through this client
func (c *TestClient) recordCall(action string, url string, payload interface{}) {
	call := TestAPICall{
		Action:  action,
		URL:     url,
		Payload: c.getPayloadToBeRecorded(payload),
	}

	c.calls = append(c.calls, call)
}

// Get performs a GET request against the Redfish service.
func (c *TestClient) Get(ctx context.Context, url string) (*http.Response, error) {
	c.recordCall(http.MethodGet, url, nil)
	customReturnForAction := c.getCustomReturnForAction(http.MethodGet)
	if customReturnForAction == nil {
		return nil, nil
	}
	return customReturnForAction.(*http.Response), nil
}

// Post performs a Post request against the Redfish service.
func (c *TestClient) Post(ctx context.Context, url string, payload interface{}) (*http.Response, error) {
	c.recordCall(http.MethodPost, url, payload)
	customReturnForAction := c.getCustomReturnForAction(http.MethodPost)
	if customReturnForAction == nil {
		return nil, nil
	}
	return customReturnForAction.(*http.Response), nil
}

// Put performs a Put request against the Redfish service.
func (c *TestClient) Put(ctx context.Context, url string, payload interface{}) (*http.Response, error) {
	c.recordCall(http.MethodPut, url, payload)
	customReturnForAction := c.getCustomReturnForAction(http.MethodPut)
	if customReturnForAction == nil {
		return nil, nil
	}
	return customReturnForAction.(*http.Response), nil
}

// Patch performs a Patch request against the Redfish service.
func (c *TestClient) Patch(ctx context.Context, url string, payload interface{}) (*http.Response, error) {
	c.recordCall(http.MethodPatch, url, payload)
	customReturnForAction := c.getCustomReturnForAction(http.MethodPatch)
	if customReturnForAction == nil {
		return nil, nil
	}
	return customReturnForAction.(*http.Response), nil
}

// Delete performs a Delete request against the Redfish service.
func (c *TestClient) Delete(ctx context.Context, url string) (*http.Response, error) {
	c.recordCall(http.MethodDelete, url, nil)
	customReturnForAction := c.getCustomReturnForAction(http.MethodDelete)
	if customReturnForAction == nil {
		return nil, nil
	}
	return customReturnForAction.(*http.Response), nil
}
