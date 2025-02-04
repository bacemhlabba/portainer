package endpoints

import (
	"testing"

	portainer "github.com/portainer/portainer/api"
	"github.com/stretchr/testify/assert"
)

func TestIsEdgeEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *portainer.Endpoint
		expected bool
	}{
		{
			name: "EdgeAgentOnDockerEnvironment",
			endpoint: &portainer.Endpoint{
				Type: portainer.EdgeAgentOnDockerEnvironment,
			},
			expected: true,
		},
		{
			name: "EdgeAgentOnKubernetesEnvironment",
			endpoint: &portainer.Endpoint{
				Type: portainer.EdgeAgentOnKubernetesEnvironment,
			},
			expected: true,
		},
		{
			name: "NonEdgeEnvironment",
			endpoint: &portainer.Endpoint{
				Type: portainer.DockerEnvironment,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEdgeEndpoint(tt.endpoint)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsAssociatedEdgeEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *portainer.Endpoint
		expected bool
	}{
		{
			name: "AssociatedEdgeEndpoint",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: true,
			},
			expected: true,
		},
		{
			name: "NonAssociatedEdgeEndpoint",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "",
				UserTrusted: true,
			},
			expected: false,
		},
		{
			name: "EdgeEndpointInWaitingRoom",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: false,
			},
			expected: false,
		},
		{
			name: "NonEdgeEnvironment",
			endpoint: &portainer.Endpoint{
				Type:        portainer.DockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: true,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAssociatedEdgeEndpoint(tt.endpoint)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHasDirectConnectivity(t *testing.T) {
	tests := []struct {
		name     string
		endpoint *portainer.Endpoint
		expected bool
	}{
		{
			name: "NonEdgeEnvironment",
			endpoint: &portainer.Endpoint{
				Type: portainer.DockerEnvironment,
			},
			expected: true,
		},
		{
			name: "AssociatedEdgeEndpoint",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: true,
				Edge:        portainer.EnvironmentEdgeSettings{AsyncMode: false},
			},
			expected: true,
		},
		{
			name: "AssociatedAsyncEdgeEndpoint",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: true,
				Edge:        portainer.EnvironmentEdgeSettings{AsyncMode: true},
			},
			expected: false,
		},
		{
			name: "EdgeEndpointInWaitingRoom",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: false,
				Edge:        portainer.EnvironmentEdgeSettings{AsyncMode: false},
			},
			expected: false,
		},
		{
			name: "AsyncEdgeEndpointInWaitingRoom",
			endpoint: &portainer.Endpoint{
				Type:        portainer.EdgeAgentOnDockerEnvironment,
				EdgeID:      "some-edge-id",
				UserTrusted: false,
				Edge:        portainer.EnvironmentEdgeSettings{AsyncMode: true},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasDirectConnectivity(tt.endpoint)
			assert.Equal(t, tt.expected, result)
		})
	}
}
