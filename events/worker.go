/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package events

import (
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	esmodel "github.com/devicechain-io/dc-event-sources/model"
	"github.com/rs/zerolog/log"
)

// Worker used to resolve event entities.
type EventResolver struct {
	WorkerId         int
	Api              model.Api
	UnresolvedEvents <-chan esmodel.UnresolvedEvent
	Resolved         func(*model.ResolvedEvent)
	Failed           func(esmodel.UnresolvedEvent, error)
}

// Create a new event resolver.
func NewEventResolver(workerId int, api model.Api,
	unrez <-chan esmodel.UnresolvedEvent,
	resolved func(*model.ResolvedEvent),
	failed func(esmodel.UnresolvedEvent, error)) *EventResolver {
	return &EventResolver{
		WorkerId:         workerId,
		Api:              api,
		UnresolvedEvents: unrez,
		Resolved:         resolved,
		Failed:           failed,
	}
}

// Execute logic to resolve event.
func (rez *EventResolver) ResolveEvent(unrez esmodel.UnresolvedEvent) (*model.ResolvedEvent, error) {
	rez.Api.DeviceTypeByToken(unrez.Device)
	return nil, nil
}

// Converts unresolved events into resolved events.
func (rez *EventResolver) Process() {
	for {
		unresolved, more := <-rez.UnresolvedEvents
		if more {
			log.Debug().Msg(fmt.Sprintf("Event resolution handled by resolver id %d", rez.WorkerId))
			resolved, err := rez.ResolveEvent(unresolved)
			if err != nil {
				rez.Failed(unresolved, err)
			} else {
				rez.Resolved(resolved)
			}
		} else {
			log.Debug().Msg("Event resolver received shutdown signal.")
			return
		}
	}
}
