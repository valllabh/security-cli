/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package test

import (
	"testing"

	"github.com/valllabh/security-cli/cmd"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name   string
		remote string
		branch string
	}{
		{name: "VULN", remote: "tomcat"},
		// {name: "VULN", remote: "https://github.com/laverdet/isolated-vm", branch: "master"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.Scan(tt.remote)
		})
	}
}
