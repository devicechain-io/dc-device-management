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

// Event with token references resolved and info from assignment merged.
type ResolvedEvent struct {
	Source          string
	AltId           *string
	DeviceId        uint
	DeviceGroupId   *uint
	AssignmentId    *uint
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
	Message string
	Payload []byte
}
