package test_dir

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestTester(t *testing.T) {
	suite.Run(t, &Model1{Snake: &Snake{}})
	suite.Run(t, &Model2{Snake: &Snake{}})
	suite.Run(t, &Model3{Snake: &Snake{}})
}
