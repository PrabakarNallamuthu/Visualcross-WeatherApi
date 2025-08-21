package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/visual-crossing-weather-api/mcp-server/config"
	"github.com/visual-crossing-weather-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_visualcrossingwebservices_rest_services_timeline_location_startdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		locationVal, ok := args["location"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: location"), nil
		}
		location, ok := locationVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: location"), nil
		}
		startdateVal, ok := args["startdate"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: startdate"), nil
		}
		startdate, ok := startdateVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: startdate"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["contentType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contentType=%v", val))
		}
		if val, ok := args["unitGroup"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("unitGroup=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["lang"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang=%v", val))
		}
		if val, ok := args["key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("key=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/VisualCrossingWebServices/rest/services/timeline/%s/%s%s", cfg.BaseURL, location, startdate, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_visualcrossingwebservices_rest_services_timeline_location_startdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_VisualCrossingWebServices_rest_services_timeline_location_startdate",
		mcp.WithDescription("Historical and Forecast Weather API"),
		mcp.WithString("location", mcp.Required(), mcp.Description("")),
		mcp.WithString("startdate", mcp.Required(), mcp.Description("")),
		mcp.WithString("contentType", mcp.Description("data format of the output either json or CSV")),
		mcp.WithString("unitGroup", mcp.Description("")),
		mcp.WithString("include", mcp.Description("data to include in the output (required for CSV format - days,hours,alerts,current,events )")),
		mcp.WithString("lang", mcp.Description("Language to use for weather descriptions")),
		mcp.WithString("key", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_visualcrossingwebservices_rest_services_timeline_location_startdateHandler(cfg),
	}
}
