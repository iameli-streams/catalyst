package config

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

//go:embed full-stack.spec.json
var spec []byte

func TestSpecIsValid(t *testing.T) {
	err := json.Unmarshal(fullstack, &MistConfig{})
	require.NoError(t, err)
	err = json.Unmarshal(spec, &MistConfig{})
	require.NoError(t, err)
}

func TestItCanPassthroughEmptyConfig(t *testing.T) {
	generated, _, err := GenerateConfig(&Cli{
		Secret:    "44444444-4444-4444-4444-444444444444",
		PublicURL: "https://example.com",
	})
	require.NoError(t, err)
	require.Empty(t, jsonEQ(spec, generated))
}

// returns empty if equal
func jsonEQ(x, y []byte) string {
	return cmp.Diff(x, y, cmp.Transformer("ParseJSON", func(in []byte) (out any) {
		if err := json.Unmarshal(in, &out); err != nil {
			return err
		}
		return out
	}))

}