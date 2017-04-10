package anything_test

//go:generate protoc --go_out=. anything.proto

import (
	"reflect"
	"testing"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pokstad/anything"
)

func TestAnything(t *testing.T) {
	t1 := &timestamp.Timestamp{
		Seconds: 5, // easy to verify
		Nanos:   6, // easy to verify
	}

	serialized, err := proto.Marshal(t1)
	if err != nil {
		t.Fatal("could not serialize timestamp")
	}

	// Blue was a great album by 3EB, before Cadgogan got kicked out
	// and Jenkins went full primadonna
	a := anything.AnythingForYou{
		Anything: &any.Any{
			TypeUrl: "example.com/yaddayaddayadda/" + proto.MessageName(t1),
			Value:   serialized,
		},
	}

	// marshal to simulate going on the wire:
	serializedA, err := proto.Marshal(&a)
	if err != nil {
		t.Fatal("could not serialize anything")
	}

	// unmarshal to simulate coming off the wire
	var a2 anything.AnythingForYou
	if err := proto.Unmarshal(serializedA, &a2); err != nil {
		t.Fatal("could not deserialize anything")
	}

	// unmarshal the timestamp
	var t2 timestamp.Timestamp
	if err := ptypes.UnmarshalAny(a2.Anything, &t2); err != nil {
		t.Fatalf("Could not unmarshal timestamp from anything field: %s", err)
	}

	// Verify the values are as expected
	if !reflect.DeepEqual(t1, &t2) {
		t.Fatalf("Values don't match up:\n %+v \n %+v", t1, t2)
	}
}
