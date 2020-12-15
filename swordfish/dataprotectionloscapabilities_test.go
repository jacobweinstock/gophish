//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jacobweinstock/gophish/common"
)

var dataProtectionLoSCapabilitiesBody = `{
		"@odata.context": "/redfish/v1/$metadata#DataProtectionLoSCapabilities.DataProtectionLoSCapabilities",
		"@odata.type": "#DataProtectionLoSCapabilities.v1_1_2.DataProtectionLoSCapabilities",
		"@odata.id": "/redfish/v1/DataProtectionLoSCapabilities",
		"Id": "DataProtectionLoSCapabilities-1",
		"Name": "DataProtectionLoSCapabilitiesOne",
		"Description": "DataProtectionLoSCapabilities One",
		"Links": {
			"SupportedReplicaOptions": [{
				"@odata.id": "/redfish/v1/ClassesOfService/1"
			}],
			"SupportedReplicaOptions@odata.count": 1
		},
		"SupportedLinesOfService": [{
			"@odata.id": "/redfish/v1/LinesOfService/1"
		}],
		"SupportedLinesOfService@odata.count": 1,
		"SupportedMinLifetimes": [
			"P0Y6M0DT0H0M0S"
		],
		"SupportedRecoveryGeographicObjectives": [
			"Datacenter",
			"Region"
		],
		"SupportedRecoveryPointObjectiveTimes": [
			"P0Y0M0DT0H30M0S"
		],
		"SupportedRecoveryTimeObjectives": [
			"OnlinePassive"
		],
		"SupportedReplicaTypes": [
			"Clone"
		],
		"SupportsIsolated": true
	}`

// TestDataProtectionLoSCapabilities tests the parsing of DataProtectionLoSCapabilities objects.
func TestDataProtectionLoSCapabilities(t *testing.T) {
	var result DataProtectionLoSCapabilities
	err := json.NewDecoder(strings.NewReader(dataProtectionLoSCapabilitiesBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataProtectionLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataProtectionLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if !result.SupportsIsolated {
		t.Error("SupportsIsolated should be true")
	}

	if result.SupportedRecoveryTimeObjectives[0] != OnlinePassiveRecoveryAccessScope {
		t.Errorf("Invalid SupportedRecoveryTimeObjective: %s",
			result.SupportedRecoveryTimeObjectives[0])
	}

	if result.SupportedReplicaTypes[0] != CloneReplicaType {
		t.Errorf("Invalid supported replica type: %s", result.SupportedReplicaTypes[0])
	}
}

// TestDataProtectionLoSCapabilitiesUpdate tests the Update call.
func TestDataProtectionLoSCapabilitiesUpdate(t *testing.T) {
	var result DataProtectionLoSCapabilities
	err := json.NewDecoder(strings.NewReader(dataProtectionLoSCapabilitiesBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &common.TestClient{}
	result.SetClient(testClient)

	result.SupportsIsolated = false
	err = result.Update(context.Background())

	if err != nil {
		t.Errorf("Error making Update call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if !strings.Contains(calls[0].Payload, "SupportsIsolated:false") {
		t.Errorf("Unexpected SupportsIsolated update payload: %s", calls[0].Payload)
	}
}
