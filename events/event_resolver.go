/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	"github.com/devicechain-io/dc-device-management/proto"
	esmodel "github.com/devicechain-io/dc-event-sources/model"
	esproto "github.com/devicechain-io/dc-event-sources/proto"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Worker used to resolve event entities.
type EventResolver struct {
	WorkerId   int
	Api        model.DeviceManagementApi
	Unresolved <-chan kafka.Message
	Invalid    func(kafka.Message)
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
	invalid func(kafka.Message),
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
	event *esmodel.UnresolvedEvent) (*EventResolutionResults, error) {

	// Assemble resolved event from initial event and device assignment.
	resolved := &model.ResolvedEvent{
		Source:          event.Source,
		AltId:           event.AltId,
		AssignmentId:    &assn.ID,
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
		Payload:         event.Payload,
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
	device *model.Device, payload esmodel.NewAssignmentPayload) (*model.DeviceAssignment, uint, error) {

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
		return nil, uint(proto.FailureReason_ApiCallFailed), err
	}

	return assn, 0, nil
}

// Handle a new assignment event.
func (rez *EventResolver) HandleNewAssignmentEvent(ctx context.Context,
	device *model.Device, event *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	assncreate, ok := event.Payload.(esmodel.NewAssignmentPayload)
	if !ok {
		return nil, uint(proto.FailureReason_Invalid), errors.New("new assignment payload was not of expected type")
	}

	// Create new device assignment from the event payload.
	created, reason, err := rez.CreateNewAssignment(ctx, device, assncreate)
	if err != nil {
		return nil, reason, errors.New("could not create device assignment")
	}

	// Merge info from device and created assignment into event.
	resolved, err := rez.MergeAssignmentToResolveEvent(device, created, event)
	if err != nil {
		return nil, uint(proto.FailureReason_Unknown), errors.New("unable to merge info to resolve event")
	}

	return []EventResolutionResults{*resolved}, 0, nil
}

// Create resolved events by looking up device assignment info and merging it into other event data.
func (rez *EventResolver) HandleStandardEvents(ctx context.Context,
	device *model.Device, event *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	// Look up active assignments for device.
	assns, err := rez.Api.ActiveDeviceAssignmentsForDevice(ctx, device.ID)
	if err != nil {
		return nil, uint(proto.FailureReason_ApiCallFailed), err
	}
	// Device should have an active assignment.
	if len(assns) == 0 {
		return nil, uint(proto.FailureReason_NoActiveDeviceAssignments), fmt.Errorf("no active assignments found for device: %s", device.Token)
	}
	return nil, 0, nil
}

// Route event to handlers based on event type.
func (rez *EventResolver) HandleEvent(ctx context.Context,
	device *model.Device, unresolved *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	switch unresolved.EventType {
	case esmodel.NewAssignment:
		return rez.HandleNewAssignmentEvent(ctx, device, unresolved)
	case esmodel.Location, esmodel.Measurement, esmodel.Alert:
		return rez.HandleStandardEvents(ctx, device, unresolved)
	default:
		return nil, uint(proto.FailureReason_Invalid), fmt.Errorf("unhandled event type: %s", unresolved.EventType.String())
	}
}

// Execute logic to resolve event.
func (rez *EventResolver) ResolveEvent(ctx context.Context, unrez *esmodel.UnresolvedEvent) ([]EventResolutionResults, uint, error) {
	device, err := rez.Api.DeviceByToken(context.Background(), unrez.Device)
	if err != nil {
		return nil, uint(proto.FailureReason_DeviceNotFound), err
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
				rez.Invalid(unresolved)
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
