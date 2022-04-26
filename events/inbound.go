/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package events

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	dmodel "github.com/devicechain-io/dc-device-management/model"
	"github.com/devicechain-io/dc-device-management/proto"
	esmodel "github.com/devicechain-io/dc-event-sources/model"
	esproto "github.com/devicechain-io/dc-event-sources/proto"
	"github.com/devicechain-io/dc-microservice/core"
	kcore "github.com/devicechain-io/dc-microservice/kafka"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const (
	EVENT_RESOLVER_COUNT = 5   // Number of event resolvers running in parallel
	KAFKA_BACKLOG_SIZE   = 100 // Number of Kafka messages that can be read and waiting to be processed
)

type InboundEventsProcessor struct {
	Microservice        *core.Microservice
	InboundEventsReader kcore.KafkaReader
	FailedEventsWriter  kcore.KafkaWriter
	Api                 dmodel.DeviceManagementApi

	messages  chan kafka.Message
	resolvers []*EventResolver

	lifecycle core.LifecycleManager
}

// Create a new inbound events processor.
func NewInboundEventsProcessor(ms *core.Microservice, inbound kcore.KafkaReader, failed kcore.KafkaWriter,
	callbacks core.LifecycleCallbacks, api dmodel.DeviceManagementApi) *InboundEventsProcessor {
	iproc := &InboundEventsProcessor{
		Microservice:        ms,
		InboundEventsReader: inbound,
		FailedEventsWriter:  failed,
		Api:                 api,
	}

	// Create lifecycle manager.
	ipname := fmt.Sprintf("%s-%s", ms.FunctionalArea, "inbound-event-proc")
	iproc.lifecycle = core.NewLifecycleManager(ipname, iproc, callbacks)
	return iproc
}

// Handle case where event failed to process.
func (iproc *InboundEventsProcessor) HandleFailedEvent(ctx context.Context, reason uint, payload []byte, err error) error {
	failed := &dmodel.FailedEvent{
		Reason:  reason,
		Message: err.Error(),
		Payload: payload,
	}

	// Marshal event message to protobuf.
	bytes, err := proto.MarshalFailedEvent(failed)
	if err != nil {
		log.Error().Err(err).Msg("unable to marshal event to protobuf")
	}

	// Create and deliver message.
	msg := kafka.Message{
		Key:   []byte(strconv.FormatInt(int64(reason), 10)),
		Value: bytes,
	}
	err = iproc.FailedEventsWriter.WriteMessages(ctx, msg)
	if err != nil {
		log.Error().Err(err).Msg("unable to send failed event message to kafka")
	}
	return err
}

// Called when a message can not be unmarshaled to an event.
func (iproc *InboundEventsProcessor) OnInvalidEventMessage(msg kafka.Message) {
	iproc.HandleFailedEvent(context.Background(), uint(proto.FailureReason_Invalid),
		msg.Value, errors.New("message could not be parsed"))
}

// Called when an event can not be resolved.
func (iproc *InboundEventsProcessor) onUnresolvedEvent(reason uint, unrez esmodel.UnresolvedEvent, rezerr error) {
	// Marshal event message to protobuf.
	bytes, err := esproto.MarshalUnresolvedEvent(&unrez)
	if err != nil {
		log.Error().Err(err).Msg("unable to marshal unresolved event to protobuf")
	} else {
		iproc.HandleFailedEvent(context.Background(), reason, bytes, rezerr)
	}
}

// Called when an event is successfully resolved.
func (iproc *InboundEventsProcessor) OnResolvedEvent([]EventResolutionResults) {
}

// Process an inbound message.
func (iproc *InboundEventsProcessor) ProcessInboundEvent(msg kafka.Message) {
	iproc.messages <- msg
}

// Initialize pool of workers for resolving events.
func (iproc *InboundEventsProcessor) initializeEventResolvers(ctx context.Context) {
	// Make channels and workers for distributed processing.
	iproc.messages = make(chan kafka.Message, KAFKA_BACKLOG_SIZE)
	iproc.resolvers = make([]*EventResolver, 0)
	for w := 1; w <= EVENT_RESOLVER_COUNT; w++ {
		resolver := NewEventResolver(w, iproc.Api, iproc.messages,
			iproc.OnInvalidEventMessage, iproc.OnResolvedEvent, iproc.onUnresolvedEvent)
		iproc.resolvers = append(iproc.resolvers, resolver)
		go resolver.Process(ctx)
	}
}

// Initialize component.
func (iproc *InboundEventsProcessor) Initialize(ctx context.Context) error {
	return iproc.lifecycle.Initialize(ctx)
}

// Lifecycle callback that runs initialization logic.
func (iproc *InboundEventsProcessor) ExecuteInitialize(ctx context.Context) error {
	// Initialize pool of event resolvers.
	iproc.initializeEventResolvers(ctx)
	return nil
}

// Start component.
func (iproc *InboundEventsProcessor) Start(ctx context.Context) error {
	return iproc.lifecycle.Start(ctx)
}

// Execute primary processing loop. This is done in a goroutine since it runs indefinitely.
func (iproc *InboundEventsProcessor) ProcessMessage(ctx context.Context) bool {
	msg, err := iproc.InboundEventsReader.ReadMessage(ctx)
	if err != nil {
		if errors.Is(err, io.EOF) {
			log.Info().Msg("Detected EOF on inbound events stream")
			return true
		} else {
			log.Error().Err(err).Msg("error reading inbound event message")
		}
	}
	iproc.messages <- msg
	return false
}

// Lifecycle callback that runs startup logic.
func (iproc *InboundEventsProcessor) ExecuteStart(ctx context.Context) error {
	go func() {
		for {
			eof := iproc.ProcessMessage(ctx)
			if eof {
				break
			}
		}
	}()
	return nil
}

// Stop component.
func (iproc *InboundEventsProcessor) Stop(ctx context.Context) error {
	return iproc.lifecycle.Stop(ctx)
}

// Lifecycle callback that runs shutdown logic.
func (iproc *InboundEventsProcessor) ExecuteStop(context.Context) error {
	return nil
}

// Terminate component.
func (iproc *InboundEventsProcessor) Terminate(ctx context.Context) error {
	return iproc.lifecycle.Terminate(ctx)
}

// Lifecycle callback that runs termination logic.
func (iproc *InboundEventsProcessor) ExecuteTerminate(context.Context) error {
	close(iproc.messages)
	return nil
}
