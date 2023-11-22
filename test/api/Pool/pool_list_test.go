package api_test

import (
	"testing"

	api_test "github.com/Tinyblargon/proxmox-go-sdk/test/api"
	"github.com/stretchr/testify/require"
)

func Test_Pools_List(t *testing.T) {
	Test := api_test.Test{}
	_ = Test.CreateTest()
	pools, err := Test.GetClient().GetPoolList()
	require.NoError(t, err)
	require.Equal(t, 1, len(pools))
}
