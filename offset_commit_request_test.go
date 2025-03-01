package sarama

import (
	"fmt"
	"testing"
)

var (
	offsetCommitRequestNoBlocksV0 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x00,
	}

	offsetCommitRequestNoBlocksV1 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x00,
	}

	offsetCommitRequestNoBlocksV2 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x33,
		0x00, 0x00, 0x00, 0x00,
	}

	offsetCommitRequestOneBlockV0 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x52, 0x21,
		0x00, 0x00, 0x00, 0x00, 0xDE, 0xAD, 0xBE, 0xEF,
		0x00, 0x08, 'm', 'e', 't', 'a', 'd', 'a', 't', 'a',
	}

	offsetCommitRequestOneBlockV1 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x52, 0x21,
		0x00, 0x00, 0x00, 0x00, 0xDE, 0xAD, 0xBE, 0xEF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x08, 'm', 'e', 't', 'a', 'd', 'a', 't', 'a',
	}

	offsetCommitRequestOneBlockV2 = []byte{
		0x00, 0x06, 'f', 'o', 'o', 'b', 'a', 'r',
		0x00, 0x00, 0x11, 0x22,
		0x00, 0x04, 'c', 'o', 'n', 's',
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x33,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x52, 0x21,
		0x00, 0x00, 0x00, 0x00, 0xDE, 0xAD, 0xBE, 0xEF,
		0x00, 0x08, 'm', 'e', 't', 'a', 'd', 'a', 't', 'a',
	}
)

func TestOffsetCommitRequestV0(t *testing.T) {
	t.Parallel()
	request := new(OffsetCommitRequest)
	request.Version = 0
	request.ConsumerGroup = "foobar"
	testRequest(t, "no blocks v0", request, offsetCommitRequestNoBlocksV0)

	request.AddBlock("topic", 0x5221, 0xDEADBEEF, 0, "metadata")
	testRequest(t, "one block v0", request, offsetCommitRequestOneBlockV0)
}

func TestOffsetCommitRequestV1(t *testing.T) {
	t.Parallel()
	request := new(OffsetCommitRequest)
	request.ConsumerGroup = "foobar"
	request.ConsumerID = "cons"
	request.ConsumerGroupGeneration = 0x1122
	request.Version = 1
	testRequest(t, "no blocks v1", request, offsetCommitRequestNoBlocksV1)

	request.AddBlock("topic", 0x5221, 0xDEADBEEF, ReceiveTime, "metadata")
	testRequest(t, "one block v1", request, offsetCommitRequestOneBlockV1)
}

func TestOffsetCommitRequestV2ToV4(t *testing.T) {
	t.Parallel()
	for version := 2; version <= 4; version++ {
		request := new(OffsetCommitRequest)
		request.ConsumerGroup = "foobar"
		request.ConsumerID = "cons"
		request.ConsumerGroupGeneration = 0x1122
		request.RetentionTime = 0x4433
		request.Version = int16(version)
		testRequest(t, fmt.Sprintf("no blocks v%d", version), request, offsetCommitRequestNoBlocksV2)

		request.AddBlock("topic", 0x5221, 0xDEADBEEF, 0, "metadata")
		testRequest(t, fmt.Sprintf("one block v%d", version), request, offsetCommitRequestOneBlockV2)
	}
}
