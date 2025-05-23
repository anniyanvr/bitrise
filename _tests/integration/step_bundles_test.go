//go:build linux_and_mac
// +build linux_and_mac

package integration

import (
	"testing"

	"github.com/bitrise-io/go-utils/command"
	"github.com/stretchr/testify/require"
)

func TestStepBundleInputs(t *testing.T) {
	configPth := "step_bundles_test_bitrise.yml"

	cmd := command.New(binPath(), "run", "test_step_bundle_inputs", "-c", configPth)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	require.NoError(t, err, out)
	require.Contains(t, out, "Hello Bitrise!")
}

func TestNestedStepBundle(t *testing.T) {
	configPth := "step_bundles_test_bitrise.yml"

	cmd := command.New(binPath(), "run", "--output-format", "json", "test_nested_step_bundle", "-c", configPth)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	require.NoError(t, err, out)
	stepOutputs := collectStepOutputs(out, t)
	require.Equal(t, stepOutputs, []string{
		`bundle1
bundle1_input1: bundle1_input1
bundle2_input1: bundle2_input1
`,
		`bundle1
bundle1_input1: bundle1_input1_override
bundle2_input1: bundle2_input1
`,
		`bundle2
bundle1_input1: 
bundle2_input1: bundle2_input1
`,
		`workflow step
bundle1_input1: 
bundle2_input1: 
`,
	})
}
