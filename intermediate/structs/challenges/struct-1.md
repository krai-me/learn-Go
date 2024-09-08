# GoLang Exercise: Circle Struct

## Problem Statement

Define a `Circle` struct that represents a geometric circle. The `Circle` struct should have a single field to store the radius of the circle.

### Tasks:

1. **Define the `Circle` struct**:
   The struct should contain one field:
   - `radius` (of type `float64`) representing the radius of the circle.

2. **Method: `circumference`**
   Define a method `circumference` for the `Circle` struct that calculates and returns the circumference of the circle.
   Use the formula:
   \[
   C = 2 \pi r
   \]
   Where \( r \) is the radius of the circle and \( \pi \) is approximately 3.14159.

3. **Method: `area`**
   Define a method `area` for the `Circle` struct that calculates and returns the area of the circle.
   Use the formula:
   \[
   A = \pi r^2
   \]
   Where \( r \) is the radius of the circle and \( \pi \) is approximately 3.14159.

4. **Take User Input**
   Modify the program to take the radius as **user input** instead of hardcoding the value.

5. **Create an instance of `Circle`**
   Using the user-provided radius, create an instance of `Circle` and use the methods to print both:
   - The **circumference** of the circle.
   - The **area** of the circle.

### Expected Output Example:

If the user enters a radius of `10.0`, the output should be:
Enter the radius of the circle: 10 Circumference of the circle: 62.831853 Area of the circle: 314.159265

### Additional Requirements:
- Use the `fmt.Scan` function to take user input.
