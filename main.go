/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package main

import (
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/rs/zerolog/log"
)

func main() {
	callbacks := core.LifecycleCallbacks{
		Initializer: core.NewNoOpLifecycleCallback(),
		Starter:     core.NewNoOpLifecycleCallback(),
		Stopper:     core.NewNoOpLifecycleCallback(),
		Terminator:  core.NewNoOpLifecycleCallback(),
	}
	ms := core.NewMicroservice("device-management", callbacks)
	err := ms.Run()
	if err != nil {
		log.Error().Err(err).Msg("Unable to start microservice")
	}
}
