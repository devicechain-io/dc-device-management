/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventResolverTestSuite struct {
	suite.Suite
}

// Perform common setup tasks.
func (suite *EventResolverTestSuite) SetupTest() {
}

// Test 1
func (suite *EventResolverTestSuite) Test1() {
	assert.Equal(suite.T(), 1, 1)
}

// Test 2
func (suite *EventResolverTestSuite) Test2() {
	assert.Equal(suite.T(), 2, 2)
}

// Run all tests.
func TestWorkerTestSuite(t *testing.T) {
	suite.Run(t, new(EventResolverTestSuite))
}
