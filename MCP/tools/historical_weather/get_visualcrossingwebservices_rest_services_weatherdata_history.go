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

func Get_visualcrossingwebservices_rest_services_weatherdata_historyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["maxDistance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxDistance=%v", val))
		}
		if val, ok := args["shortColumnNames"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shortColumnNames=%v", val))
		}
		if val, ok := args["endDateTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endDateTime=%v", val))
		}
		if val, ok := args["aggregateHours"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("aggregateHours=%v", val))
		}
		if val, ok := args["collectStationContributions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("collectStationContributions=%v", val))
		}
		if val, ok := args["startDateTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startDateTime=%v", val))
		}
		if val, ok := args["maxStations"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxStations=%v", val))
		}
		if val, ok := args["allowAsynch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allowAsynch=%v", val))
		}
		if val, ok := args["locations"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locations=%v", val))
		}
		if val, ok := args["includeNormals"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeNormals=%v", val))
		}
		if val, ok := args["contentType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contentType=%v", val))
		}
		if val, ok := args["unitGroup"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("unitGroup=%v", val))
		}
		if val, ok := args["key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("key=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/VisualCrossingWebServices/rest/services/weatherdata/history%s", cfg.BaseURL, queryString)
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

func CreateGet_visualcrossingwebservices_rest_services_weatherdata_historyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_VisualCrossingWebServices_rest_services_weatherdata_history",
		mcp.WithDescription("Retrieves hourly or daily historical weather records."),
		mcp.WithString("maxDistance", mcp.Description("")),
		mcp.WithBoolean("shortColumnNames", mcp.Description("")),
		mcp.WithString("endDateTime", mcp.Description("")),
		mcp.WithString("aggregateHours", mcp.Description("")),
		mcp.WithBoolean("collectStationContributions", mcp.Description("")),
		mcp.WithString("startDateTime", mcp.Description("")),
		mcp.WithString("maxStations", mcp.Description("")),
		mcp.WithBoolean("allowAsynch", mcp.Description("")),
		mcp.WithString("locations", mcp.Description("")),
		mcp.WithBoolean("includeNormals", mcp.Description("")),
		mcp.WithString("contentType", mcp.Description("")),
		mcp.WithString("unitGroup", mcp.Description("")),
		mcp.WithString("key", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_visualcrossingwebservices_rest_services_weatherdata_historyHandler(cfg),
	}
}
