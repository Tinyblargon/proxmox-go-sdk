package proxmox

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_VmRef_nilCheck(t *testing.T) {
	tests := []struct {
		name   string
		input  *VmRef
		output error
	}{
		{name: "Populated pointer",
			input:  &VmRef{},
			output: nil,
		},
		{name: "Empty pointer",
			input:  nil,
			output: errors.New(VmRef_Error_Nil),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.input.nilCheck())
		})
	}
}
