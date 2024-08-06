# weather_analysis.py

def average_temperatures(weather_data):
    """Finds the average maximum and minimum temperatures from the weather data."""
    total_max_temp = 0
    total_min_temp = 0
    count = len(weather_data)
    
    for data in weather_data:
        total_max_temp += data['max_temp']
        total_min_temp += data['min_temp']
    
    avg_max_temp = total_max_temp / count
    avg_min_temp = total_min_temp / count
    
    return avg_max_temp, avg_min_temp

def average_humidity(weather_data):
    """Calculates the average humidity over the given period."""
    total_humidity = 0
    count = len(weather_data)
    
    for data in weather_data:
        total_humidity += data['humidity']
    
    avg_humidity = total_humidity / count
    
    return avg_humidity

def main():
    # Weather data for Bangalore City from 20th July to 26th July 2024
    weather_data = [
        {'date': '2024-07-20', 'max_temp': 32, 'min_temp': 23, 'humidity': 60},
        {'date': '2024-07-21', 'max_temp': 33, 'min_temp': 24, 'humidity': 65},
        {'date': '2024-07-22', 'max_temp': 31, 'min_temp': 22, 'humidity': 70},
        {'date': '2024-07-23', 'max_temp': 30, 'min_temp': 21, 'humidity': 55},
        {'date': '2024-07-24', 'max_temp': 29, 'min_temp': 20, 'humidity': 50},
        {'date': '2024-07-25', 'max_temp': 34, 'min_temp': 25, 'humidity': 75},
        {'date': '2024-07-26', 'max_temp': 35, 'min_temp': 26, 'humidity': 80}
    ]

    # Calculate average temperatures
    avg_max_temp, avg_min_temp = average_temperatures(weather_data)
    print(f"Average Maximum Temperature: {avg_max_temp:.2f}°C")
    print(f"Average Minimum Temperature: {avg_min_temp:.2f}°C")

    # Calculate average humidity
    avg_humidity = average_humidity(weather_data)
    print(f"Average Humidity: {avg_humidity:.2f}%")

if __name__ == "__main__":
    main()
