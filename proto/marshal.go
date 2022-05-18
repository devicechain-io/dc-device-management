/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package proto

import (
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	esmodel "github.com/devicechain-io/dc-event-sources/model"
	util "github.com/devicechain-io/dc-microservice/proto"
	"google.golang.org/protobuf/proto"
)

// Marshal a failed event to protobuf bytes.
func MarshalFailedEvent(event *model.FailedEvent) ([]byte, error) {
	// Encode protobuf event.
	pbevent := &PFailedEvent{
		Reason:  FailureReason(event.Reason),
		Service: event.Service,
		Message: event.Message,
		Error:   event.Error,
		Payload: event.Payload,
	}

	// Marshal event to bytes.
	bytes, err := proto.Marshal(pbevent)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Unmarshal encoded failed event.
func UnmarshalFailedEvent(encoded []byte) (*model.FailedEvent, error) {
	// Unmarshal protobuf event.
	pbevent := &PFailedEvent{}
	err := proto.Unmarshal(encoded, pbevent)
	if err != nil {
		return nil, err
	}

	event := &model.FailedEvent{
		Reason:  uint(pbevent.Reason),
		Service: pbevent.Service,
		Message: pbevent.Message,
		Error:   pbevent.Error,
		Payload: pbevent.Payload,
	}

	return event, nil
}

// Marshal payload for a locations event.
func MarshalPayloadForNewAssignmentEvent(payload *model.ResolvedNewAssignmentPayload) ([]byte, error) {
	pbpayload := &PResolvedNewAssignmentPayload{
		AssignmentId:  payload.AssignmentId,
		DeviceGroup:   payload.DeviceGroup,
		Asset:         payload.Asset,
		AssetGroup:    payload.AssetGroup,
		Customer:      payload.Customer,
		CustomerGroup: payload.CustomerGroup,
		Area:          payload.Area,
		AreaGroup:     payload.AreaGroup,
	}

	bytes, err := proto.Marshal(pbpayload)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Marshal payload for a locations event.
func MarshalPayloadForLocationsEvent(payload *model.ResolvedLocationsPayload) ([]byte, error) {
	pbpayload := &PResolvedLocationsPayload{}
	for _, entry := range payload.Entries {
		pbentry := &PResolvedLocationEntry{
			Latitude:     entry.Latitude,
			Longitude:    entry.Longitude,
			Elevation:    entry.Elevation,
			OccurredTime: entry.OccurredTime,
		}
		pbpayload.Entries = append(pbpayload.Entries, pbentry)
	}
	bytes, err := proto.Marshal(pbpayload)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Marshal payload for a measurements event.
func MarshalPayloadForMeasurementsEvent(payload *model.ResolvedMeasurementsPayload) ([]byte, error) {
	pbpayload := &PResolvedMeasurementsPayload{}
	for _, mxsentry := range payload.Entries {
		pmxentries := make([]*PResolvedMeasurementEntry, 0)
		for _, mxentry := range mxsentry.Entries {
			pmxentry := &PResolvedMeasurementEntry{
				Name:       mxentry.Name,
				Value:      mxentry.Value,
				Classifier: mxentry.Classifier,
			}
			pmxentries = append(pmxentries, pmxentry)
		}
		pbentry := &PResolvedMeasurementsEntry{
			Measurements: pmxentries,
			OccurredTime: mxsentry.OccurredTime,
		}
		pbpayload.Entries = append(pbpayload.Entries, pbentry)
	}
	bytes, err := proto.Marshal(pbpayload)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Marshal payload for an alerts event.
func MarshalPayloadForAlertsEvent(payload *model.ResolvedAlertsPayload) ([]byte, error) {
	pbpayload := &PResolvedAlertsPayload{}
	for _, entry := range payload.Entries {
		pbentry := &PResolvedAlertEntry{
			Type:         entry.Type,
			Level:        entry.Level,
			Message:      entry.Message,
			Source:       entry.Source,
			OccurredTime: entry.OccurredTime,
		}
		pbpayload.Entries = append(pbpayload.Entries, pbentry)
	}
	bytes, err := proto.Marshal(pbpayload)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal a payload into a new assignment event.
func UnmarshalPayloadForNewAssignmentEvent(encoded []byte) (*model.ResolvedNewAssignmentPayload, error) {
	pbpayload := &PResolvedNewAssignmentPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.ResolvedNewAssignmentPayload{
		AssignmentId:  pbpayload.AssignmentId,
		DeviceGroup:   pbpayload.DeviceGroup,
		Asset:         pbpayload.Asset,
		AssetGroup:    pbpayload.AssetGroup,
		Customer:      pbpayload.Customer,
		CustomerGroup: pbpayload.CustomerGroup,
		Area:          pbpayload.Area,
		AreaGroup:     pbpayload.AreaGroup,
	}

	return payload, nil
}

// Unmarshal a payload into a locations event.
func UnmarshalPayloadForLocationsEvent(encoded []byte) (*model.ResolvedLocationsPayload, error) {
	pbpayload := &PResolvedLocationsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.ResolvedLocationsPayload{}
	entries := make([]model.ResolvedLocationEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		entry := model.ResolvedLocationEntry{
			Latitude:     pbentry.Latitude,
			Longitude:    pbentry.Longitude,
			Elevation:    pbentry.Elevation,
			OccurredTime: pbentry.OccurredTime,
		}
		entries = append(entries, entry)
	}
	payload.Entries = entries
	return payload, nil
}

// Unmarshal a payload into a measurements event.
func UnmarshalPayloadForMeasurementsEvent(encoded []byte) (*model.ResolvedMeasurementsPayload, error) {
	pbpayload := &PResolvedMeasurementsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.ResolvedMeasurementsPayload{}
	entries := make([]model.ResolvedMeasurementsEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		mxs := make([]model.ResolvedMeasurementEntry, 0)
		for _, pmx := range pbentry.Measurements {
			mx := model.ResolvedMeasurementEntry{
				Name:       pmx.Name,
				Value:      pmx.Value,
				Classifier: pmx.Classifier,
			}
			mxs = append(mxs, mx)
		}
		entry := model.ResolvedMeasurementsEntry{
			Entries:      mxs,
			OccurredTime: pbentry.OccurredTime,
		}
		entries = append(entries, entry)
	}
	payload.Entries = entries
	return payload, nil
}

// Unmarshal a payload into an alerts event.
func UnmarshalPayloadForAlertsEvent(encoded []byte) (*model.ResolvedAlertsPayload, error) {
	pbpayload := &PResolvedAlertsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.ResolvedAlertsPayload{}
	entries := make([]model.ResolvedAlertEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		entry := model.ResolvedAlertEntry{
			Type:         pbentry.Type,
			Level:        pbentry.Level,
			Message:      pbentry.Message,
			Source:       pbentry.Source,
			OccurredTime: pbentry.OccurredTime,
		}
		entries = append(entries, entry)
	}
	payload.Entries = entries
	return payload, nil
}

// Marshal unresolved payload based on event type.
func MarshalResolvedPayload(etype esmodel.EventType, payload interface{}) ([]byte, error) {
	switch etype {
	case esmodel.NewAssignment:
		if rnapayload, ok := payload.(*model.ResolvedNewAssignmentPayload); ok {
			return MarshalPayloadForNewAssignmentEvent(rnapayload)
		}
		return nil, fmt.Errorf("invalid new assignment payload: %+v", payload)
	case esmodel.Location:
		if locpayload, ok := payload.(*model.ResolvedLocationsPayload); ok {
			return MarshalPayloadForLocationsEvent(locpayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	case esmodel.Measurement:
		if mxpayload, ok := payload.(*model.ResolvedMeasurementsPayload); ok {
			return MarshalPayloadForMeasurementsEvent(mxpayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	case esmodel.Alert:
		if apayload, ok := payload.(*model.ResolvedAlertsPayload); ok {
			return MarshalPayloadForAlertsEvent(apayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	default:
		return nil, fmt.Errorf("unable to marshal unresolved payload for event type: %s", etype.String())
	}
}

// Unmarshal unresolved payload based on event type.
func UnmarshalResolvedPayload(etype esmodel.EventType, payload []byte) (interface{}, error) {
	switch etype {
	case esmodel.NewAssignment:
		return UnmarshalPayloadForNewAssignmentEvent(payload)
	case esmodel.Location:
		return UnmarshalPayloadForLocationsEvent(payload)
	case esmodel.Measurement:
		return UnmarshalPayloadForMeasurementsEvent(payload)
	case esmodel.Alert:
		return UnmarshalPayloadForAlertsEvent(payload)
	default:
		return nil, fmt.Errorf("unable to unmarshal resolved payload for event type: %s", etype.String())
	}
}

// Marshal a resolved event to protobuf bytes.
func MarshalResolvedEvent(event *model.ResolvedEvent) ([]byte, error) {
	// Encode payload.
	pybytes, err := MarshalResolvedPayload(event.EventType, event.Payload)
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

// Unmarshal encoded resolved event.
func UnmarshalResolvedEvent(encoded []byte) (*model.ResolvedEvent, error) {
	// Unmarshal protobuf event.
	pbevent := &PResolvedEvent{}
	err := proto.Unmarshal(encoded, pbevent)
	if err != nil {
		return nil, err
	}

	// Unmarshal payload based on event type.
	payload, err := UnmarshalResolvedPayload(esmodel.EventType(pbevent.EventType), pbevent.Payload)
	if err != nil {
		return nil, err
	}

	event := &model.ResolvedEvent{
		Source:          pbevent.Source,
		AltId:           pbevent.AltId,
		AssignmentId:    uint(pbevent.AssignmentId),
		DeviceId:        uint(pbevent.DeviceId),
		DeviceGroupId:   util.NullUintOf(pbevent.DeviceGroupId),
		AssetId:         util.NullUintOf(pbevent.AssetId),
		AssetGroupId:    util.NullUintOf(pbevent.AssetGroupId),
		CustomerId:      util.NullUintOf(pbevent.CustomerId),
		CustomerGroupId: util.NullUintOf(pbevent.CustomerGroupId),
		AreaId:          util.NullUintOf(pbevent.AreaId),
		AreaGroupId:     util.NullUintOf(pbevent.AreaGroupId),
		EventType:       esmodel.EventType(pbevent.EventType),
		Payload:         payload,
	}

	return event, nil
}
