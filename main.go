/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package main

import (
	"context"

	"github.com/devicechain-io/dc-devicemanagement/model"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/devicechain-io/dc-microservice/rdb"
)

var (
	Microservice *core.Microservice
	RdbManager   *rdb.RdbManager
)

func main() {
	callbacks := core.LifecycleCallbacks{
		Initializer: core.LifecycleCallback{
			Preprocess:  func(context.Context) error { return nil },
			Postprocess: afterMicroserviceInitialized,
		},
		Starter: core.LifecycleCallback{
			Preprocess:  func(context.Context) error { return nil },
			Postprocess: afterMicroserviceStarted,
		},
		Stopper: core.LifecycleCallback{
			Preprocess:  beforeMicroserviceStopped,
			Postprocess: func(context.Context) error { return nil },
		},
		Terminator: core.LifecycleCallback{
			Preprocess:  beforeMicroserviceTerminated,
			Postprocess: func(context.Context) error { return nil },
		},
	}
	Microservice = core.NewMicroservice(callbacks)
	Microservice.Run()
}

// Called after microservice has been initialized.
func afterMicroserviceInitialized(ctx context.Context) error {
	// Create and initialize rdb manager.
	callbacks := core.NewNoOpLifecycleCallbacks()
	RdbManager = rdb.NewRdbManager(Microservice, callbacks, model.Migrations)
	return RdbManager.Initialize(context.Background())
}

// Called after microservice has been started.
func afterMicroserviceStarted(ctx context.Context) error {
	return RdbManager.Start(context.Background())
}

// Called before microservice has been stopped.
func beforeMicroserviceStopped(ctx context.Context) error {
	return RdbManager.Stop(context.Background())
}

// Called before microservice has been terminated.
func beforeMicroserviceTerminated(ctx context.Context) error {
	return RdbManager.Terminate(context.Background())
}
