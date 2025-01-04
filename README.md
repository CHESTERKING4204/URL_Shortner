# URL Shortener

A lightweight and efficient URL Shortener service designed to make long URLs short and manageable. This application allows you to create shortened URLs and redirect users from these short links back to their original URLs.

---

## Features

- **Shorten Long URLs**: Easily create short, unique codes for long URLs.
- **Redirect Functionality**: Use the short code to redirect back to the original URL.
- **Simple Setup**: Easy to install and run locally.

---

## How to Run the Code

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/CHESTERKING4204/URL_Shortner.git
   cd URL_Shortner
   
2. **Open Terminal on main.go file and run the command**
    ```bash
    go run main

3. **Now, go to Postman and follow these steps:**

   - **URL bar:**
     ```bash
     localhost:3000/shortner
     ```

   - **Text Area (Choose JSON format):**
     ```json
     {
        "url": "{URL you want to make short}"
     }
     ```

4. **To use the shortened URL, enter the following format in the browser:**
   ```bash
   localhost:3000/redirect/{SHORT CODE YOU GOT}


