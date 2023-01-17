/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "testing"

func TestScan(t *testing.T) {
	tests := []struct {
		name   string
		remote string
		branch string
	}{
		{name: "VULN", remote: "https://hub.docker.com/_/memcached"},
		// {name: "VULN", remote: "https://github.com/laverdet/isolated-vm", branch: "master"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Scan(tt.remote)
		})
	}
}
