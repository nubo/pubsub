package pubsub_test

import (
	"os"
	"testing"

	"github.com/nubo/pubsub"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ps pubsub.Conn

func TestGoPubSub(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoPubsub Suite")
}

var _ = BeforeSuite(func() {
	addr := os.Getenv("REDIS_ADDR")
	Ω(addr).ShouldNot(BeEmpty())
	ps = pubsub.Dial("tcp", addr, 10, 1000)
})

var _ = AfterSuite(func() {
	ps.Close()
})
