package main

import (
	"github.com/visual-crossing-weather-api/mcp-server/config"
	"github.com/visual-crossing-weather-api/mcp-server/models"
	tools_timeline_weather_api_single_date_request "github.com/visual-crossing-weather-api/mcp-server/tools/timeline_weather_api_single_date_request"
	tools_timeline_weather_api_date_range_request "github.com/visual-crossing-weather-api/mcp-server/tools/timeline_weather_api_date_range_request"
	tools_weather_forecast "github.com/visual-crossing-weather-api/mcp-server/tools/weather_forecast"
	tools_historical_weather "github.com/visual-crossing-weather-api/mcp-server/tools/historical_weather"
	tools_timeline_weather_api_15_day_forecast_request "github.com/visual-crossing-weather-api/mcp-server/tools/timeline_weather_api_15_day_forecast_request"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_timeline_weather_api_single_date_request.CreateGet_visualcrossingwebservices_rest_services_timeline_location_startdateTool(cfg),
		tools_timeline_weather_api_date_range_request.CreateGet_visualcrossingwebservices_rest_services_timeline_location_startdate_enddateTool(cfg),
		tools_weather_forecast.CreateGet_visualcrossingwebservices_rest_services_weatherdata_forecastTool(cfg),
		tools_historical_weather.CreateGet_visualcrossingwebservices_rest_services_weatherdata_historyTool(cfg),
		tools_timeline_weather_api_15_day_forecast_request.CreateGet_visualcrossingwebservices_rest_services_timeline_locationTool(cfg),
	}
}
