//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/jacobweinstock/gophish/common"
)

// UpdateService is used to represent the update service offered by the redfish API
type UpdateService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// FirmwareInventory points towards the firmware store endpoint
	FirmwareInventory string
	// HTTPPushURI endpoint is used to push (POST) firmware updates
	HTTPPushURI string `json:"HttpPushUri"`
	// ServiceEnabled indicates whether this service isenabled.
	ServiceEnabled bool
	// Status describes the status and health of a resource and its children.
	Status common.Status
	// TransferProtocol is the list of network protocols used by the UpdateService to retrieve the software image file
	TransferProtocol []string
	// UpdateServiceTarget indicates where theupdate image is to be applied.
	UpdateServiceTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a UpdateService object from the raw JSON.
func (updateService *UpdateService) UnmarshalJSON(b []byte) error {
	type temp UpdateService
	type actions struct {
		SimpleUpdate struct {
			AllowableValues []string `json:"TransferProtocol@Redfish.AllowableValues"`
			Target          string
		} `json:"#UpdateService.SimpleUpdate"`
	}
	var t struct {
		temp
		Actions           actions
		FirmwareInventory common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*updateService = UpdateService(t.temp)
	updateService.FirmwareInventory = string(t.FirmwareInventory)
	updateService.TransferProtocol = t.Actions.SimpleUpdate.AllowableValues
	updateService.UpdateServiceTarget = t.Actions.SimpleUpdate.Target
	updateService.rawData = b
	return nil
}

// GetUpdateService will get a UpdateService instance from the service.
func GetUpdateService(ctx context.Context, c common.Client, uri string) (*UpdateService, error) {
	resp, err := c.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var updateService UpdateService
	err = json.NewDecoder(resp.Body).Decode(&updateService)
	if err != nil {
		return nil, err
	}
	updateService.SetClient(c)
	return &updateService, nil
}
