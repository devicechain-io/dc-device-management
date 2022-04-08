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

	esproto "github.com/devicechain-io/dc-event-sources/proto"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type InboundEventsProcessor struct {
	Microservice        *core.Microservice
	InboundEventsReader *kafka.Reader

	lifecycle core.LifecycleManager
}

// Create a new inbound events processor.
func NewInboundEventsProcessor(ms *core.Microservice, reader *kafka.Reader,
	callbacks core.LifecycleCallbacks) *InboundEventsProcessor {
	iproc := &InboundEventsProcessor{
		Microservice:        ms,
		InboundEventsReader: reader,
	}

	// Create lifecycle manager.
	ipname := fmt.Sprintf("%s-%s", ms.FunctionalArea, "inbound-event-proc")
	iproc.lifecycle = core.NewLifecycleManager(ipname, iproc, callbacks)
	return iproc
}

// Process an inbound message.
func (iproc *InboundEventsProcessor) ProcessInboundEvent(msg kafka.Message) error {
	event, payload, err := esproto.UnmarshalEvent(msg.Value)
	if err != nil {
		return err
	}
	if log.Debug().Enabled() {
		log.Debug().Msg(fmt.Sprintf("Received event %+v with payload %+v", event, payload))
	}
	return nil
}

// Initialize component.
func (iproc *InboundEventsProcessor) Initialize(ctx context.Context) error {
	return iproc.lifecycle.Initialize(ctx)
}

// Lifecycle callback that runs initialization logic.
func (iproc *InboundEventsProcessor) ExecuteInitialize(context.Context) error {
	return nil
}

// Start component.
func (iproc *InboundEventsProcessor) Start(ctx context.Context) error {
	return iproc.lifecycle.Start(ctx)
}

// Lifecycle callback that runs startup logic.
func (iproc *InboundEventsProcessor) ExecuteStart(ctx context.Context) error {
	go func() {
		for {
			msg, err := iproc.InboundEventsReader.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Info().Msg("Detected EOF on inbound events stream")
					return
				} else {
					log.Error().Err(err).Msg("error reading inbound event message")
				}
			}
			iproc.ProcessInboundEvent(msg)
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
	return nil
}
