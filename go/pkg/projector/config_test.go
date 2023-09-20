package projector_test

import (
	"reflect"
	"testing"

	"github.com/Asteromorph/aoc/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("expected no error %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("expected args to be an empty array but got %v", config.Args)
	}

	if config.Operation != operation {
		t.Errorf("expected operation %v but got %v", config.Operation, operation)
	}

}

func testConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func testConfigPrintWithArgs(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func testConfigAdd(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"},[]string{"foo", "bar"}, projector.Add)
}

func testConfigRemove(t *testing.T) {
	testConfig(t, []string{"rm", "foo"}, []string{"foo"}, projector.Remove)
}
