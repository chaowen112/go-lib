package json

import (
	jsoniter "github.com/json-iterator/go" // nolint:goimports
	"github.com/json-iterator/go/extra"
	"github.com/pkg/errors"
)

// nolint:gochecknoinits
func init() {
	jsoniter.RegisterExtension(&extra.BinaryAsStringExtension{})
}

var (
	json                = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal             = json.Marshal
	MarshalIndent       = json.MarshalIndent
	Unmarshal           = json.Unmarshal
	MarshalToString     = json.MarshalToString
	UnmarshalFromString = json.UnmarshalFromString
	NewDecoder          = json.NewDecoder
	NewEncoder          = json.NewEncoder
)

func DeepCopy(dst interface{}, src interface{}) error {
	srcJSON, err := json.Marshal(src)
	if err != nil {
		return errors.Wrapf(err, "marshal failed when deep copy")
	}

	err = json.Unmarshal(srcJSON, dst)
	if err != nil {
		return errors.Wrapf(err, "unmarshal failed when deep copy")
	}

	return nil
}

func MarshalPretty(v interface{}) ([]byte, error) {
	return MarshalIndent(v, "", "    ")
}
