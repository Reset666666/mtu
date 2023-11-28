package mtu_test

import (
	"fmt"
	"github.com/Reset666666/mtu"

	"testing"
)

func TestSetAdapterMtu(t *testing.T) {
	err := mtu.SetAdapterMtu("VMware Network Adapter VMnet1", 1500)
	if err != nil {
		fmt.Println(err)
	}
}
