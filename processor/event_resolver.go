/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package processor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	dmproto "github.com/devicechain-io/dc-device-management/proto"
	esmodel "github.com/devicechain-io/dc-event-sources/model"
	esproto "github.com/devicechain-io/dc-event-sources/proto"
	"github.com/devicechain-io/dc-microservice/proto"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Worker used to resolve event entities.
type EventResolver struct {
	WorkerId   int
	Api        model.DeviceManagementApi
	Unresolved <-chan kafka.Message
	Invalid    func(error, kafka.Message)
	Resolved   func([]EventResolutionResults)
	Failed     func(uint, esmodel.UnresolvedEvent, error)
}

// Results of event resolution process.
type EventResolutionResults struct {
	Device     *model.Device
	Assignment *model.DeviceAssignment
	Resolved   *model.ResolvedEvent
}

// Create a new event resolver.
func NewEventResolver(workerId int, api model.DeviceManagementApi,
	unrez <-chan kafka.Message,
	invalid func(error, kafka.Message),
	resolved func([]EventResolutionResults),
	failed func(uint, esmodel.UnresolvedEvent, error)) *EventResolver {
	return &EventResolver{
		WorkerId:   workerId,
		Api:        api,
		Unresolved: unrez,
		Invalid:    invalid,
		Resolved:   resolved,
		Failed:     failed,
	}
}

// Merge device and assignment data with unresolved event in order to create a resolved event.
func (rez *EventResolver) MergeAssignmentToResolveEvent(device *model.Device, assn *model.DeviceAssignment,
	event *esmodel.UnresolvedEvent, rezPayload interface{}) (*EventResolutionResults, error) {

	// Assemble resolved event from initial event and device assignment.
	resolved := &model.ResolvedEvent{
		Source:          event.Source,
		AltId:           event.AltId,
		AssignmentId:    assn.ID,
		DeviceId:        device.ID,
		DeviceGroupId:   assn.DeviceGroupId,
		AssetId:         assn.AssetId,
		AssetGroupId:    assn.AssetGroupId,
		CustomerId:      assn.CustomerId,
		CustomerGroupId: assn.CustomerGroupId,
		AreaId:          assn.AreaId,
		AreaGroupId:     assn.AreaGroupId,
		OccurredTime:    event.OccurredTime,
		ProcessedTime:   event.ProcessedTime,
		EventType:       event.EventType,
		Payload:         rezPayload,
	}

	results := &EventResolutionResults{
		Device:     device,
		Assignment: assn,
		Resolved:   resolved,
	}

	return results, nil
}

// Create a new device assignment based on the event payload.
func (rez *EventResolver) CreateNewAssignment(ctx context.Context,
	device *model.Device, payload esmodel.UnresolvedNewAssignmentPayload) (*model.DeviceAssignment, uint, error) {

	active := true
	req := &model.DeviceAssignmentCreateRequest{
		Token:         uuid.NewString(),
		Device:        device.Token,
		DeviceGroup:   payload.DeviceGroup,
		Asset:         payload.Asset,
		AssetGroup:    payload.AssetGroup,
		Customer:      payload.Customer,
		CustomerGroup: payload.CustomerGroup,
		Area:          payload.Area,
		AreaGroup:     payload.AreaGroup,
		Active:        &active,
	}
	assn, err := rez.Api.CreateDeviceAssignment(ctx, req)
	if err != nil {
		return nil, uint(dmproto.FailureReason_ApiCallFailed), err
	}

	return assn, 0, nil
}

// Handle a new assignment event.
func (rez *EventResolver) HandleNewAssignmentEvent(ctx context.Context,
	device *model.Device, event *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	assncreate, ok := event.Payload.(*esmodel.UnresolvedNewAssignmentPayload)
	if !ok {
		return nil, uint(dmproto.FailureReason_Invalid), errors.New("new assignment payload was not of expected type")
	}

	// Create new device assignment from the event payload.
	created, reason, err := rez.CreateNewAssignment(ctx, device, *assncreate)
	if err != nil {
		return nil, reason, errors.New("could not create device assignment")
	}

	// Convert to resolved payload.
	payload := &model.ResolvedNewAssignmentPayload{
		AssignmentId:  uint64(created.ID),
		DeviceGroup:   proto.NullUint64Of(created.DeviceGroupId),
		Asset:         proto.NullUint64Of(created.AssetId),
		AssetGroup:    proto.NullUint64Of(created.AssetGroupId),
		Customer:      proto.NullUint64Of(created.CustomerId),
		CustomerGroup: proto.NullUint64Of(created.CustomerGroupId),
		Area:          proto.NullUint64Of(created.AreaId),
		AreaGroup:     proto.NullUint64Of(created.AreaGroupId),
	}

	// Merge info from device and created assignment into event.
	resolved, err := rez.MergeAssignmentToResolveEvent(device, created, event, payload)
	if err != nil {
		return nil, uint(dmproto.FailureReason_Unknown), errors.New("unable to merge info to resolve event")
	}

	return []EventResolutionResults{*resolved}, 0, nil
}

// Resolve a locations event payload.
func (rez *EventResolver) ResolveLocationsEventPayload(ctx context.Context, device *model.Device,
	assignment *model.DeviceAssignment, event *esmodel.UnresolvedEvent) (interface{}, error) {
	if lpayload, ok := event.Payload.(*esmodel.UnresolvedLocationsPayload); ok {
		rlpayload := &model.ResolvedLocationsPayload{}
		rlentries := make([]model.ResolvedLocationEntry, 0)
		for _, ulentry := range lpayload.Entries {
			rlentry := model.ResolvedLocationEntry{
				Latitude:     ulentry.Latitude,
				Longitude:    ulentry.Longitude,
				Elevation:    ulentry.Elevation,
				OccurredTime: ulentry.OccurredTime,
			}
			rlentries = append(rlentries, rlentry)
		}
		rlpayload.Entries = rlentries
		return rlpayload, nil
	}
	return nil, fmt.Errorf("can not resolve locations payload. invalid unresolved payload type")
}

// Resolve a measurements event payload.
func (rez *EventResolver) ResolveMeasurementsEventPayload(ctx context.Context, device *model.Device,
	assignment *model.DeviceAssignment, event *esmodel.UnresolvedEvent) (interface{}, error) {
	if mpayload, ok := event.Payload.(*esmodel.UnresolvedMeasurementsPayload); ok {
		rmpayload := &model.ResolvedMeasurementsPayload{}
		rmsentries := make([]model.ResolvedMeasurementsEntry, 0)
		for _, umsentry := range mpayload.Entries {
			rmentries := make([]model.ResolvedMeasurementEntry, 0)
			for mxkey, mxvalue := range umsentry.Measurements {
				rmentry := model.ResolvedMeasurementEntry{
					Name:       mxkey,
					Value:      mxvalue,
					Classifier: nil,
				}
				rmentries = append(rmentries, rmentry)
			}
			rmsentry := model.ResolvedMeasurementsEntry{
				Entries:      rmentries,
				OccurredTime: umsentry.OccurredTime,
			}
			rmsentries = append(rmsentries, rmsentry)
		}
		rmpayload.Entries = rmsentries
		return rmpayload, nil
	}
	return nil, fmt.Errorf("can not resolve measurements payload. invalid unresolved payload type")
}

// Resolve a alerts event payload.
func (rez *EventResolver) ResolveAlertsEventPayload(ctx context.Context, device *model.Device,
	assignment *model.DeviceAssignment, event *esmodel.UnresolvedEvent) (interface{}, error) {
	if apayload, ok := event.Payload.(*esmodel.UnresolvedAlertsPayload); ok {
		rapayload := &model.ResolvedAlertsPayload{}
		raentries := make([]model.ResolvedAlertEntry, 0)
		for _, uaentry := range apayload.Entries {
			raentry := model.ResolvedAlertEntry{
				Type:         uaentry.Type,
				Level:        uaentry.Level,
				Message:      uaentry.Message,
				Source:       uaentry.Source,
				OccurredTime: uaentry.OccurredTime,
			}
			raentries = append(raentries, raentry)
		}
		rapayload.Entries = raentries
		return rapayload, nil
	}
	return nil, fmt.Errorf("can not resolve alerts payload. invalid unresolved payload type")
}

// Convert an unresolved event payload into a resolved payload.
func (rez *EventResolver) ResolveEventPayload(ctx context.Context, device *model.Device,
	assignment *model.DeviceAssignment, event *esmodel.UnresolvedEvent) (interface{}, error) {
	switch event.EventType {
	case esmodel.Location:
		return rez.ResolveLocationsEventPayload(ctx, device, assignment, event)
	case esmodel.Measurement:
		return rez.ResolveMeasurementsEventPayload(ctx, device, assignment, event)
	case esmodel.Alert:
		return rez.ResolveAlertsEventPayload(ctx, device, assignment, event)
	default:
		return nil, fmt.Errorf("unable to handle resolution for payload type: %s", event.EventType.String())
	}
}

// Create resolved events by looking up device assignment info and merging it into other event data.
func (rez *EventResolver) HandleStandardEvent(ctx context.Context,
	device *model.Device, event *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	// Look up active assignments for device.
	assns, err := rez.Api.ActiveDeviceAssignmentsForDevice(ctx, device.ID)
	if err != nil {
		return nil, uint(dmproto.FailureReason_ApiCallFailed), err
	}
	// Device should have an active assignment.
	if len(assns) == 0 {
		return nil, uint(dmproto.FailureReason_NoActiveDeviceAssignments), fmt.Errorf("no active assignments found for device: %s", device.Token)
	}

	// Create separate merged event for each assignment.
	results := make([]EventResolutionResults, 0)
	for _, assn := range assns {
		resolved, err := rez.ResolveEventPayload(ctx, device, &assn, event)
		if err != nil {
			return nil, uint(dmproto.FailureReason_ApiCallFailed), err
		}

		result, err := rez.MergeAssignmentToResolveEvent(device, &assn, event, resolved)
		if err != nil {
			return nil, uint(dmproto.FailureReason_ApiCallFailed), err
		}
		results = append(results, *result)
	}

	return results, 0, nil
}

// Route event to handlers based on event type.
func (rez *EventResolver) HandleEvent(ctx context.Context,
	device *model.Device, unresolved *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	switch unresolved.EventType {
	case esmodel.NewAssignment:
		return rez.HandleNewAssignmentEvent(ctx, device, unresolved)
	case esmodel.Location, esmodel.Measurement, esmodel.Alert:
		return rez.HandleStandardEvent(ctx, device, unresolved)
	default:
		return nil, uint(dmproto.FailureReason_Invalid), fmt.Errorf("unhandled event type: %s", unresolved.EventType.String())
	}
}

// Execute logic to resolve event.
func (rez *EventResolver) ResolveEvent(ctx context.Context, unrez *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	device, err := rez.Api.DeviceByToken(context.Background(), unrez.Device)
	if err != nil {
		return nil, uint(dmproto.FailureReason_DeviceNotFound), err
	}
	return rez.HandleEvent(ctx, device, unrez)
}

// Converts unresolved events into resolved events.
func (rez *EventResolver) Process(ctx context.Context) {
	for {
		unresolved, more := <-rez.Unresolved
		if more {
			log.Debug().Msg(fmt.Sprintf("Event resolution handled by resolver id %d", rez.WorkerId))

			// Attempt to unmarshal event.
			event, err := esproto.UnmarshalUnresolvedEvent(unresolved.Value)
			if err != nil {
				rez.Invalid(err, unresolved)
				continue
			}

			if log.Debug().Enabled() {
				jevent, err := json.MarshalIndent(event, "", "  ")
				if err == nil {
					log.Debug().Msg(fmt.Sprintf("Received %s event:\n%s", event.EventType.String(), jevent))
				}
			}

			// Attempt to resolve event.
			resolved, reason, err := rez.ResolveEvent(ctx, event)
			if err != nil {
				rez.Failed(reason, *event, err)
			} else {
				rez.Resolved(resolved)
			}
		} else {
			log.Debug().Msg("Event resolver received shutdown signal.")
			return
		}
	}
}
