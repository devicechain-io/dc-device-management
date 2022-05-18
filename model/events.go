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

// Payload with resolved assignment info.
type ResolvedNewAssignmentPayload struct {
	AssignmentId  uint64
	DeviceGroup   *uint64
	Asset         *uint64
	AssetGroup    *uint64
	Customer      *uint64
	CustomerGroup *uint64
	Area          *uint64
	AreaGroup     *uint64
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

// Event with token references resolved and info from assignment merged.
type ResolvedEvent struct {
	Source          string
	AltId           *string
	DeviceId        uint
	AssignmentId    uint
	DeviceGroupId   *uint
	CustomerId      *uint
	CustomerGroupId *uint
	AreaId          *uint
	AreaGroupId     *uint
	AssetId         *uint
	AssetGroupId    *uint
	OccurredTime    time.Time
	ProcessedTime   time.Time
	EventType       esmodel.EventType
	Payload         interface{}
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
