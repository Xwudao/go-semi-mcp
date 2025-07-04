package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Xwudao/semi-mcp/assets"
	"github.com/Xwudao/semi-mcp/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Semi-Design组件库说明文档，包含API属性和组件列表",
		"0.0.1",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
		server.WithLogging(),
	)

	// Add the calculator handler
	s.AddTool(tools.ListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var result []string
		for _, cmp := range assets.Components {
			result = append(result, cmp.Name)
		}

		return mcp.NewToolResultText(strings.Join(result, "\n\n")), nil
	})
	s.AddTool(tools.GetComponentUsageTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cmpName, err := request.RequireString("component_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get component name: %w", err)
		}

		for _, cmp := range assets.Components {
			if cmp.Name == cmpName {
				return mcp.NewToolResultText(cmp.Content), nil
			}
		}

		return nil, fmt.Errorf("component %s not found", cmpName)
	})

	// sse := server.NewSSEServer(s, server.WithBaseURL("/sse"))

	var eg = new(errgroup.Group)

	// eg.Go(func() error {
	// 	// Start the SSE server
	// 	fmt.Println("Starting SSE server on http://0.0.0.0:3334/sse")
	// 	if err := sse.Start(":3334"); err != nil {
	// 		log.Fatalf("SSE Server error: %v\n", err)
	// 		return err
	// 	}

	// 	return nil
	// })

	eg.Go(func() error {
		// Start the server
		fmt.Println("MCP Server also available on stdio")
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Server error: %v\n", err)
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatalf("Error in server: %v\n", err)
	}
}
