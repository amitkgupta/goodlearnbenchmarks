package parameterestimator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParameterEstimator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parameter Estimator Suite")
}
