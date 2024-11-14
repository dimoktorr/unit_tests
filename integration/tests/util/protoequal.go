package util

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/encoding/protojson"
)

type protoEqualFn interface {
	True(bool, ...any) bool
}

func NewProtoEqual(fn protoEqualFn, expected, actual proto.Message) bool {
	if is := proto.Equal(expected, actual); is {
		return true
	}

	if is := reflect.TypeOf(expected) == reflect.TypeOf(actual); !is {
		return fn.True(false, fmt.Sprintf("These two protobuf messages are not equal:\nexpected: %T\nactual: %T\n", expected, actual))
	}

	protoEqualSpewConfig := spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
		DisableMethods:          true,
		MaxDepth:                10,
	}

	expectedByte, _ := protojson.Marshal(expected)
	actualByte, _ := protojson.Marshal(actual)

	expectedJson := map[string]any{}
	actualJson := map[string]any{}

	_ = json.Unmarshal(expectedByte, &expectedJson)
	_ = json.Unmarshal(actualByte, &actualJson)

	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(protoEqualSpewConfig.Sdump(expectedJson)),
		B:        difflib.SplitLines(protoEqualSpewConfig.Sdump(actualJson)),
		FromFile: "Expected",
		ToFile:   "Actual",
		Context:  1,
	})

	return fn.True(false, fmt.Sprintf("These two protobuf messages are not equal:\nexpected: %v\nactual: %v\n\nDiff:\n%s", expected, actual, diff))
}
