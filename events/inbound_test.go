/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package events

import (
	"context"
	"io"
	"testing"
	"time"

	dmodel "github.com/devicechain-io/dc-device-management/model"
	dmtest "github.com/devicechain-io/dc-device-management/test"
	"github.com/devicechain-io/dc-event-sources/model"
	esproto "github.com/devicechain-io/dc-event-sources/proto"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/devicechain-io/dc-microservice/rdb"
	test "github.com/devicechain-io/dc-microservice/test"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type InboundEventsProcessorTestSuite struct {
	suite.Suite
	IP       *InboundEventsProcessor
	Inbound  *test.MockKafkaReader
	Resolved *test.MockKafkaWriter
	Failed   *test.MockKafkaWriter
	API      *dmtest.MockApi
}

// Perform common setup tasks.
func (suite *InboundEventsProcessorTestSuite) SetupTest() {
	suite.Inbound = new(test.MockKafkaReader)
	suite.Resolved = new(test.MockKafkaWriter)
	suite.Failed = new(test.MockKafkaWriter)
	suite.API = new(dmtest.MockApi)
	suite.IP = NewInboundEventsProcessor(
		dmtest.DeviceManagementMicroservice,
		suite.Inbound,
		suite.Resolved,
		suite.Failed,
		core.NewNoOpLifecycleCallbacks(),
		suite.API)
	ctx := context.Background()
	suite.IP.Initialize(ctx)
}

// Test processing loop termination on EOF.
func (suite *InboundEventsProcessorTestSuite) TestProcessingLoopEof() {
	suite.Inbound.Mock.On("ReadMessage", mock.Anything).Return(kafka.Message{}, io.EOF)

	eof := suite.IP.ProcessMessage(context.Background())

	assert.Equal(suite.T(), eof, true)
}

// Test processing loop without EOF.
func (suite *InboundEventsProcessorTestSuite) TestProcessingLoopNonEof() {
	suite.Inbound.Mock.On("ReadMessage", mock.Anything).Return(kafka.Message{}, nil)

	eof := suite.IP.ProcessMessage(context.Background())

	assert.Equal(suite.T(), eof, false)
}

// Test event with invalid protobuf content.
func (suite *InboundEventsProcessorTestSuite) TestInvalidEvent() {
	// Assuming invalid binary message format..
	key := []byte("test")
	value := []byte("badvalue")
	badmsg := kafka.Message{Key: key, Value: value}

	// Emulate kafka read/write.
	suite.Inbound.Mock.On("ReadMessage", mock.Anything).Return(badmsg, nil)
	suite.Failed.Mock.On("WriteMessages", mock.Anything, mock.Anything).Return(nil)

	// Send message and wait for event to be processed by resolver.
	ctx := context.Background()
	suite.IP.ProcessMessage(ctx)
	suite.IP.ProcessFailedEvent(ctx)

	// Verify a message was written to failed messages writer.
	suite.Failed.AssertCalled(suite.T(), "WriteMessages", mock.Anything, mock.Anything)
}

// Build a location event.
func buildLocationEvent() *model.UnresolvedEvent {
	lat := "33.7490"
	lon := "-84.3880"
	ele := "738"
	loc := &model.LocationPayload{
		Latitude:  &lat,
		Longitude: &lon,
		Elevation: &ele,
	}
	altid := "alternateId"
	event := &model.UnresolvedEvent{
		Source:    "mysource",
		AltId:     &altid,
		Device:    "TEST-123",
		EventType: model.Location,
		Payload:   loc,
	}
	return event
}

// Build a location event.
func buildDevice() *dmodel.Device {
	name := "Test 123"
	desc := "Device test 123"
	dtype := uint(123)
	device := &dmodel.Device{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		TokenReference: rdb.TokenReference{
			Token: "TEST-123",
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(&name),
			Description: rdb.NullStrOf(&desc),
		},
		DeviceTypeId: &dtype,
	}
	return device
}

// Build a device assignment.
func buildAssignments() []dmodel.DeviceAssignment {
	dgrp := uint(2)
	asset := uint(3)
	agrp := uint(4)
	cust := uint(5)
	cgrp := uint(6)
	area := uint(7)
	argrp := uint(8)
	assn := &dmodel.DeviceAssignment{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		TokenReference: rdb.TokenReference{
			Token: "assn1",
		},
		DeviceId:        1,
		DeviceGroupId:   &dgrp,
		AssetId:         &asset,
		AssetGroupId:    &agrp,
		CustomerId:      &cust,
		CustomerGroupId: &cgrp,
		AreaId:          &area,
		AreaGroupId:     &argrp,
	}
	result := make([]dmodel.DeviceAssignment, 0)
	result = append(result, *assn)
	return result
}

// Test valid location event.
func (suite *InboundEventsProcessorTestSuite) TestValidLocationEvent() {
	loc := buildLocationEvent()
	bytes, err := esproto.MarshalUnresolvedEvent(loc)
	assert.Nil(suite.T(), err)

	// Assuming invalid binary message format..
	key := []byte(loc.Device)
	msg := kafka.Message{Key: key, Value: bytes}

	// Emulate kafka read/write.
	suite.Inbound.Mock.On("ReadMessage", mock.Anything).Return(msg, nil)
	suite.Resolved.Mock.On("WriteMessages", mock.Anything, mock.Anything).Return(nil)
	suite.API.Mock.On("DeviceByToken", mock.Anything, mock.Anything).Return(buildDevice(), nil)
	suite.API.Mock.On("ActiveDeviceAssignmentsForDevice", mock.Anything, mock.Anything).Return(buildAssignments(), nil)

	// Send message and wait for event to be processed by resolver.
	ctx := context.Background()
	suite.IP.ProcessMessage(ctx)
	suite.IP.ProcessResolvedEvent(ctx)

	// Verify a message was written to failed messages writer.
	suite.Resolved.AssertCalled(suite.T(), "WriteMessages", mock.Anything, mock.Anything)
}

// Run all tests.
func TestInboundEventsProcessorTestSuite(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	suite.Run(t, new(InboundEventsProcessorTestSuite))
}
