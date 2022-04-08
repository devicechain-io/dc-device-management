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

// Base type for events.
type ResolvedEvent struct {
	Source        string
	AltId         *string
	DeviceId      uint
	AssignmentId  *uint
	CustomerId    *uint
	AreaId        *uint
	AssetId       *uint
	OccurredTime  time.Time
	ProcessedTime time.Time
	EventType     esmodel.EventType
	Payload       interface{}
}
