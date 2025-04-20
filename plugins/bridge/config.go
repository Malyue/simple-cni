package main

import (
	"encoding/json"
	"fmt"
	"github.com/containernetworking/cni/pkg/types"
)

type NetConf struct {
	types.NetConf
	MTU    int    `json:"mtu"`
	Bridge string `json:"bridge"`
	Mac    string `json:"mac"`
}

type IPAMConf struct {
	Type   string `json:"type"`
	Subnet string `json:"subnet"`
}

func loadNetConf(bytes []byte, envArgs string) (*NetConf, error) {
	n := &NetConf{}
	if err := json.Unmarshal(bytes, &n); err != nil {
		return nil, fmt.Errorf("failed to load netconf : %v", err)
	}

	if n.Bridge == "" {
		return nil, fmt.Errorf("bridge name is required")
	}

	return n, nil
}
