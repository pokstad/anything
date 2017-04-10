package anything

//go:generate protoc --go_out=. anything.proto

import (
	"log"

	"reflect"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	t1 := &timestamp.Timestamp{
		Seconds: 5, // easy to verify
		Nanos:   6, // easy to verify
	}

	serialized, err := proto.Marshal(t1)
	if err != nil {
		log.Fatal("could not serialize timestamp")
	}

	// Blue was a great album by 3EB, before Cadgogan got kicked out
	// and Jenkins went full primadonna
	a := AnythingForYou{
		Anything: &any.Any{
			TypeUrl: "google.protobuf.Timestamp",
			Value:   serialized,
		},
	}

	// marshal to simulate going on the wire:
	serializedA, err := proto.Marshal(&a)
	if err != nil {
		log.Fatal("could not serialize anything")
	}

	// unmarshal to simulate coming off the wire
	var a2 AnythingForYou
	if err := proto.Unmarshal(serializedA, &a2); err != nil {
		log.Fatal("could not deserialize anything")
	}

	// unmarshal the timestamp
	var t2 timestamp.Timestamp
	if err := ptypes.UnmarshalAny(a2.Anything, &t2); err != nil {
		log.Fatal("Could not unmarshal timestamp from anything field")
	}

	// Verify the values are as expected
	if !reflect.DeepEqual(t1, t2) {
		log.Fatal("Values don't match up")
	}
}
