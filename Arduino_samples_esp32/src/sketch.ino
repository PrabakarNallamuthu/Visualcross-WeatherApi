#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>
#include <LiquidCrystal_I2C.h>
#include <ArduinoJson.h>

#define WIFI_SSID "Wokwi-GUEST" // Replace with your Wi-Fi SSID
#define WIFI_PASSWORD "" // Replace with your Wi-Fi password

const char* apiKey = "YOUR_API_KEY"; // Your Visual Crossing Weather API key
const char* location = "New York, NY"; // Your location for weather data
const char* serverName = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/";

// Initialize the LCD
LiquidCrystal_I2C lcd(0x27, 16, 2);

// Custom character for the degree symbol
byte degreeSymbol[8] = {
  B00111,
  B00101,
  B00111,
  B00000,
  B00000,
  B00000,
  B00000,
  B00000
};

void setup() {
  Serial.begin(115200);
  Wire.begin(21, 22); // Initialize I2C with GPIO 21 as SDA and GPIO 22 as SCL
  lcd.begin(16, 2); // Initialize the LCD with 16 columns and 2 rows
  lcd.backlight();
  // Create the custom degree symbol
  lcd.createChar(0, degreeSymbol);

  Serial.println("Connecting to WiFi...");

  // // Pretend successful Wi-Fi connection
  // delay(2000); // Simulate some delay for Wi-Fi connection
  // Serial.println("Connected to WiFi");

  // Actual Wi-Fi code for simulation
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.println("Connected to WiFi");
  fetchWeatherData(); // Fetch initial weather data
}

void loop() {
  // Fetch weather data periodically
  fetchWeatherData();
  delay(60000); // Fetch data every 60 seconds
}

void fetchWeatherData() {
  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    String url = String(serverName) + location + "/today?key=" + apiKey + "&elements=temp,conditions&include=days";

    http.begin(url);
    int httpResponseCode = http.GET();

    if (httpResponseCode > 0) {
      String payload = http.getString();
      Serial.println(payload);

      // Parse JSON
      DynamicJsonDocument doc(2048);
      DeserializationError error = deserializeJson(doc, payload);      
      if (!error) {
        const char* condition = doc["days"][0]["conditions"];
        float temperature = doc["days"][0]["temp"];

        Serial.print("Weather Condition: ");
        Serial.println(condition);

        Serial.print("Temperature: ");
        Serial.println(temperature);

        // Display data on LCD
        lcd.clear();
        lcd.setCursor(0, 0);
        lcd.print("Temp: ");
        lcd.print(temperature, 1);
        lcd.write(byte(0)); // Display custom degree symbol
        lcd.print("F");

        lcd.setCursor(0, 1);
        lcd.print("Cond: ");
        lcd.print(condition);
      } else {
        Serial.println("Failed to parse JSON");
      }
    } else {
      Serial.print("Error on HTTP request: ");
      Serial.println(httpResponseCode);
    }

    http.end();
  } else {
    Serial.println("WiFi Disconnected");
  }
}
