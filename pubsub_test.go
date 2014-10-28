package pubsub_test

import (
	"encoding/json"

	"github.com/nubo/pubsub"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestPayload struct {
	A int
	B int
	C int
}

var _ = Describe("connection", func() {
	It("should close w/o panic if it is a nil receiver", func() {
		con := pubsub.Conn{}
		Ω(con.Close).ShouldNot(Panic())
	})
})

var _ = Describe("Publish and subscribe handling", func() {
	var chn chan []byte
	var payload TestPayload
	var data []byte

	BeforeEach(func() {
		chn = make(chan []byte)
		payload = TestPayload{0, 1, 2}

		var err error
		data, err = json.Marshal(payload)
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should subscribe once to a single topic and receive a single message", func() {
		ps.Subscribe("testchannel", chn)
		ps.Publish("testchannel", data)

		Eventually(chn).Should(Receive(Equal(data)))
	})
	It("should subscribe twice to a single topic and receive every message twice", func() {
		ps.Subscribe("testchannel", chn)
		ps.Subscribe("testchannel", chn)
		ps.Publish("testchannel", data)

		Eventually(chn).Should(Receive(Equal(data)))
		Eventually(chn).Should(Receive(Equal(data)))
	})
	It("should subscribe to different topics and receive all messages for this topics", func() {
		ps.Subscribe("testchannel.0", chn)
		ps.Subscribe("testchannel.1", chn)
		ps.Publish("testchannel.0", data)
		ps.Publish("testchannel.1", data)

		Eventually(chn).Should(Receive(Equal(data)))
		Eventually(chn).Should(Receive(Equal(data)))
	})
})
