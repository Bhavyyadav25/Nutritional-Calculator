# Nutritional-Calculator
The Nutrition Calculator is a Go-based application that calculates the nutritional score and Nutri-Score for various food items. It helps users assess the healthiness of food products based on their energy, sugars, saturated fatty acids, sodium, fruits, fiber, and protein content. The application follows the Nutri-Score guidelines for scoring and categorizing food items into different nutritional levels (A, B, C, D, E).

## Features
- Calculates the nutritional score based on energy, sugars, saturated fatty acids, sodium, fruits, fiber, and protein content of food items.
- Provides Nutri-Score ratings (A, B, C, D, E) for easy assessment of the healthiness of food products.
- Modular code structure with separate files for main application logic (main.go) and Nutri-Score calculation utilities (nutriscore.go).
- Utilizes object-oriented programming principles to enhance code maintainability and readability.

## Prerequisites
Go 1.15 or above

## Getting Started
1. Clone the repository:
git clone https://github.com/your-username/nutrition-calculator.git

2. Navigate to the project directory:
cd nutrition-calculator

3. Run the application:
go run .

4. Modify the input values in main.go to calculate the nutritional score and Nutri-Score for different food items.

## Usage
Modify the main.go file to input the nutritional data for a specific food item. The following attributes are used for calculation:

Energy: Energy density in kJ/100g.
Sugars: Amount of sugars in grams/100g.
SaturatedFattyAcids: Amount of saturated fatty acids in grams/100g.
Sodium: Amount of sodium in mg/100g.
Fruits: Fruits, vegetables, pulses, nuts, and rapeseed, walnut, and olive oils as a percentage of the total.
Fibre: Amount of fiber in grams/100g.
Protein: Amount of protein in grams/100g.
After modifying the input values, run the application using the command go run main.go. The nutritional score and Nutri-Score will be displayed in the console.

## Acknowledgments
The Nutri-Score system: https://en.wikipedia.org/wiki/Nutri-Score
