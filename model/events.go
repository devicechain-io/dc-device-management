/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"time"

	esmodel "github.com/devicechain-io/dc-event-sources/model"
)

// Payload with resolved device relationship info.
type ResolvedNewRelationshipPayload struct {
	DeviceRelationshipTypeId uint64
	TargetDeviceId           *uint64
	TargetDeviceGroupId      *uint64
	TargetAssetId            *uint64
	TargetAssetGroupId       *uint64
	TargetCustomerId         *uint64
	TargetCustomerGroupId    *uint64
	TargetAreaId             *uint64
	TargetAreaGroupId        *uint64
}

// Entry with resolved location information.
type ResolvedLocationEntry struct {
	Latitude     *string
	Longitude    *string
	Elevation    *string
	OccurredTime *string
}

// Payload with resolved location entries.
type ResolvedLocationsPayload struct {
	Entries []ResolvedLocationEntry
}

// Entry with resolved info for a single measurement.
type ResolvedMeasurementEntry struct {
	Name       string
	Value      string
	Classifier *uint64
}

// Information for a measurements entry.
type ResolvedMeasurementsEntry struct {
	Entries      []ResolvedMeasurementEntry
	OccurredTime *string
}

// Payload with resolved measurement entries.
type ResolvedMeasurementsPayload struct {
	Entries []ResolvedMeasurementsEntry
}

// Information for an alert entry.
type ResolvedAlertEntry struct {
	Type         string
	Level        uint32
	Message      string
	Source       string
	OccurredTime *string
}

// Payload with resolved alert entries.
type ResolvedAlertsPayload struct {
	Entries []ResolvedAlertEntry
}

// Event with token references resolved and info from device relationship merged.
type ResolvedEvent struct {
	Source                string
	AltId                 *string
	SourceDeviceId        uint
	DeviceRelationshipId  uint
	TargetDeviceId        *uint
	TargetDeviceGroupId   *uint
	TargetCustomerId      *uint
	TargetCustomerGroupId *uint
	TargetAreaId          *uint
	TargetAreaGroupId     *uint
	TargetAssetId         *uint
	TargetAssetGroupId    *uint
	OccurredTime          time.Time
	ProcessedTime         time.Time
	EventType             esmodel.EventType
	Payload               interface{}
}

// Captures failure information for events that could not be processed.
type FailedEvent struct {
	Reason  uint
	Service string
	Message string
	Error   string
	Payload []byte
}

// Create a new FailedEvent.
func NewFailedEvent(reason uint, service string, message string, err error, payload []byte) *FailedEvent {
	return &FailedEvent{
		Reason:  reason,
		Service: service,
		Message: message,
		Error:   err.Error(),
		Payload: payload,
	}
}
