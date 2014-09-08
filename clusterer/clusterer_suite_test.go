package clusterer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestClusterer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clusterer Suite")
}
