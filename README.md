# hng12_stage_1

---

# Classify Number API

## API Specification

This API allows you to classify a number and returns its properties such as whether it's prime, perfect, and Armstrong. It also returns the sum of its digits and a fun fact from the Numbers API.

### **Endpoint**

`GET https://hng12-stage-1-uh7h.onrender.com/api/classify-number?number=371`

### **Query Parameters**

- **number** (required): A valid integer to be classified.

### **Required JSON Response Format**

#### **(200 OK)**

```json
{
  "number": 371,
  "is_prime": false,
  "is_perfect": false,
  "properties": ["armstrong", "odd"],
  "digit_sum": 11, // Sum of its digits
  "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371" // Fun fact from Numbers API
}
```

- **number**: The number being classified.
- **is_prime**: `true` if the number is prime, `false` otherwise.
- **is_perfect**: `true` if the number is perfect, `false` otherwise.
- **properties**: A list of string properties of the number. Possible values are:
  - `["armstrong", "odd"]` if the number is both an Armstrong number and odd.
  - `["armstrong", "even"]` if the number is an Armstrong number and even.
  - `["odd"]` if the number is not an Armstrong number but is odd.
  - `["even"]` if the number is not an Armstrong number but is even.
- **digit_sum**: The sum of the digits of the number.
- **fun_fact**: A fun fact related to the number, fetched from the Numbers API.

### **Possible Error Response Example**

#### **(400 Bad Request)**

If the user submits a non-numeric input:

```json
{
  "number": "alphabet",
  "error": true
}
```

- **number**: The input that was invalid.
- **error**: A boolean indicating that the input was invalid.

### **Acceptance Criteria**

1. **Functionality**:

   - The API accepts **GET requests** with a `number` query parameter.
   - The API returns a **JSON response** with the specified format.
   - The API accepts **all valid integers** as valid inputs, including negative numbers and zero.

2. **Public Access**:
   - The API is **publicly accessible** at `https://hng12-stage-1-uh7h.onrender.com/api/classify-number?number=371`

### **Additional Notes**

- The **Numbers API** is used to fetch the **fun fact** for the number. Specifically, the **math type** is used to get the fun fact (e.g., if the number is an Armstrong number, the fact will mention it).

  Example:

  - For the number 371, the fact might be: `"371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"`.

---
