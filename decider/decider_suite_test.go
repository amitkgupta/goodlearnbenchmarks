package decider_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDecider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decider Suite")
}
