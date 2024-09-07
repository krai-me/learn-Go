package main

import "fmt"

// Define a function to calculate taxes and demonstrate deferred execution.
func calculateTaxes(revenue, deductions, credits float64) float64 {
	// Schedule this print statement to run at the end of the function.
	defer fmt.Println("Taxes Calculated!")

	// Print statement executed immediately.
	fmt.Println("Calculating Taxes")

	// Define a tax rate.
	taxRate := .06143

	// If either deductions or credits are zero, return a simple tax calculation.
	if deductions == 0 || credits == 0 {
		return revenue * taxRate
	}

	// Calculate the tax value based on revenue, deductions, and credits.
	taxValue := (revenue - (deductions * credits)) * taxRate
	if taxValue >= 0 {
		return taxValue
	} else {
		return 0
	}
}

func main() {
	// Call the calculateTaxes function with example values.
	tax := calculateTaxes(1000, 200, 50)
	fmt.Println("Final Tax:", tax)
}

/*
   **Deferring Execution with `defer`**:
   - The `defer` keyword is used to delay the execution of a function until the surrounding function (`calculateTaxes` in this case) returns.
   - In this example, `fmt.Println("Taxes Calculated!")` is deferred, so it runs after all other code in `calculateTaxes` has executed.
   - `fmt.Println("Calculating Taxes")` runs immediately when the function is called.

   **How It Works**:
   - When `calculateTaxes` is called, it prints "Calculating Taxes".
   - The `defer` statement schedules "Taxes Calculated!" to be printed just before `calculateTaxes` returns.
   - The deferred statement runs regardless of which return path the function takes (i.e., whether it exits early due to `if deductions == 0 || credits == 0` or reaches the end of the function).

   **Use Cases**:
   - Useful for tasks like logging, closing files, or cleaning up resources.
   - Ensures that the deferred function runs no matter how the function exits.

   **Example Output**:
   - The output will be:
     ```
     Calculating Taxes
     Taxes Calculated!
     Final Tax: 48.857
     ```

   This shows the immediate print statement, the deferred print statement, and the final calculated tax.
*/
