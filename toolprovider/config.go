package toolprovider

import (
	"fmt"

	"github.com/bitrise-io/bitrise/v2/models"
	"github.com/bitrise-io/bitrise/v2/toolprovider/provider"
)

func getToolRequests(config models.BitriseDataModel) ([]provider.ToolRequest, error) {
	tools := config.Tools

	var toolRequests []provider.ToolRequest
	for toolID, toolVersion := range tools {
		v, strategy, err := ParseVersionString(toolVersion)
		if err != nil {
			return nil, fmt.Errorf("parse %s version: %w", toolID, err)
		}

		var pluginIdentifier *string
		if config.ToolConfig != nil && config.ToolConfig.ExtraPlugins != nil {
			if pluginID, ok := config.ToolConfig.ExtraPlugins[toolID]; ok {
				pluginIdentifier = &pluginID
			}
		}

		toolRequests = append(toolRequests, provider.ToolRequest{
			ToolName:           provider.ToolID(toolID),
			UnparsedVersion:    v,
			ResolutionStrategy: strategy,
			PluginIdentifier: pluginIdentifier,
		})
	}

	return toolRequests, nil
}

func defaultToolConfig() models.ToolConfigModel {
	return models.ToolConfigModel{
		Provider: "asdf",
	}
}
