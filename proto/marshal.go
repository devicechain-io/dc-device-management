/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package proto

import (
	"github.com/devicechain-io/dc-device-management/model"
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
	// Encode protobuf event.
	dgid := uint64(*event.DeviceGroupId)
	aid := uint64(*event.AssetId)
	agid := uint64(*event.AssetGroupId)
	cid := uint64(*event.CustomerId)
	cgid := uint64(*event.CustomerGroupId)
	arid := uint64(*event.AreaId)
	argid := uint64(*event.AreaGroupId)
	pbevent := &PResolvedEvent{
		Source:          event.Source,
		AltId:           event.AltId,
		EventType:       int64(event.EventType),
		DeviceId:        uint64(event.DeviceId),
		DeviceGroupId:   &dgid,
		AssetId:         &aid,
		AssetGroupId:    &agid,
		CustomerId:      &cid,
		CustomerGroupId: &cgid,
		AreaId:          &arid,
		AreaGroupId:     &argid,
	}

	// Marshal event to bytes.
	bytes, err := proto.Marshal(pbevent)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
