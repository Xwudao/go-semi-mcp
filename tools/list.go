package tools

import "github.com/mark3labs/mcp-go/mcp"

var ListTool = mcp.NewTool("list components", mcp.WithDescription("List all components Names in the Semi-Design library"))

var GetComponentUsageTool = mcp.NewTool("get component usage",
	mcp.WithDescription("Get the usage of a specific component in the Semi-Design library"),
	mcp.WithString("component_name", mcp.Description("The name of the component to get usage for. The Name must be like ComponentName"), mcp.Required()),
)

var GetTokenTool = mcp.NewTool("get token",
	mcp.WithDescription("Get the token of the Semi-Design library. Include css variables and their values."),
)

var GetCSSVariablesTool = mcp.NewTool("get css variables",
	mcp.WithDescription("Get the CSS variables of the Semi-Design library. Include css variables and their values."),
)
