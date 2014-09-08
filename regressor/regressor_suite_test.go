package regressor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRegressor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Regressor Suite")
}
