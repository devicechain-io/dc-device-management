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

	dmtest "github.com/devicechain-io/dc-device-management/test"
	"github.com/devicechain-io/dc-microservice/core"
	test "github.com/devicechain-io/dc-microservice/test"
	"github.com/segmentio/kafka-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type InboundEventsProcessorTestSuite struct {
	suite.Suite
	IP      *InboundEventsProcessor
	Inbound *test.MockKafkaReader
	Failed  *test.MockKafkaWriter
	API     *dmtest.MockApi
}

// Perform common setup tasks.
func (suite *InboundEventsProcessorTestSuite) SetupTest() {
	suite.Inbound = new(test.MockKafkaReader)
	suite.Failed = new(test.MockKafkaWriter)
	suite.API = new(dmtest.MockApi)
	suite.IP = NewInboundEventsProcessor(
		dmtest.DeviceManagementMicroservice,
		suite.Inbound,
		suite.Failed,
		core.NewNoOpLifecycleCallbacks(),
		suite.API)
	ctx := context.Background()
	suite.IP.Initialize(ctx)
}

// Test event with invalid protobuf content.
func (suite *InboundEventsProcessorTestSuite) TestInvalidEvent() {
	key := []byte("test")
	value := []byte("badvalue")
	msg := kafka.Message{Key: key, Value: value}
	suite.IP.ProcessInboundEvent(msg)

	assert.NotNil(suite.T(), msg)
}

// Test processing loop termination on EOF.
func (suite *InboundEventsProcessorTestSuite) TestProcessingLoopEof() {
	suite.Inbound.Mock.On("ReadMessage", mock.Anything).Return(kafka.Message{}, io.EOF)

	eof := suite.IP.ProcessMessage(context.Background())

	assert.Equal(suite.T(), eof, true)
}

// Run all tests.
func TestInboundEventsProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(InboundEventsProcessorTestSuite))
}
