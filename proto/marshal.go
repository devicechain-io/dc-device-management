/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package proto

import (
	"github.com/devicechain-io/dc-device-management/model"
	esproto "github.com/devicechain-io/dc-event-sources/proto"
	util "github.com/devicechain-io/dc-microservice/proto"
	"google.golang.org/protobuf/proto"
)

// Marshal a failed event to protobuf bytes.
func MarshalFailedEvent(event *model.FailedEvent) ([]byte, error) {
	// Encode protobuf event.
	pbevent := &PFailedEvent{
		Reason:  FailureReason(event.Reason),
		Message: event.Message,
		Payload: event.Payload,
	}

	// Marshal event to bytes.
	bytes, err := proto.Marshal(pbevent)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Unmarshal encoded unresolved event.
func UnmarshalFailedEvent(encoded []byte) (*model.FailedEvent, error) {
	// Unmarshal protobuf event.
	pbevent := &PFailedEvent{}
	err := proto.Unmarshal(encoded, pbevent)
	if err != nil {
		return nil, err
	}

	event := &model.FailedEvent{
		Reason:  uint(pbevent.Reason),
		Message: pbevent.Message,
		Payload: pbevent.Payload,
	}

	return event, nil
}

// Marshal a resolved event to protobuf bytes.
func MarshalResolvedEvent(event *model.ResolvedEvent) ([]byte, error) {
	// Encode payload.
	pybytes, err := esproto.MarshalPayloadForEventType(event.EventType, event.Payload)
	if err != nil {
		return nil, err
	}

	// Encode protobuf event.
	pbevent := &PResolvedEvent{
		Source:          event.Source,
		AltId:           event.AltId,
		AssignmentId:    uint64(event.AssignmentId),
		DeviceId:        uint64(event.DeviceId),
		DeviceGroupId:   util.NullUint64Of(event.DeviceGroupId),
		AssetId:         util.NullUint64Of(event.AssetId),
		AssetGroupId:    util.NullUint64Of(event.AssetGroupId),
		CustomerId:      util.NullUint64Of(event.CustomerId),
		CustomerGroupId: util.NullUint64Of(event.CustomerGroupId),
		AreaId:          util.NullUint64Of(event.AreaId),
		AreaGroupId:     util.NullUint64Of(event.AreaGroupId),
		EventType:       int64(event.EventType),
		Payload:         pybytes,
	}

	// Marshal event to bytes.
	bytes, err := proto.Marshal(pbevent)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
