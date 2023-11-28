package mtu_test

import (
	"fmt"
	"mtu"
	"testing"
)

func TestSetAdapterMtu(t *testing.T) {
	err := mtu.SetAdapterMtu("VMware Network Adapter VMnet1", 1500)
	if err != nil {
		fmt.Println(err)
	}
}
