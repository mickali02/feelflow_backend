# FeelFlow Backend - Lab #5 Demo Guide

This document provides a step-by-step guide to run and test the Go backend and connect it with the Flutter frontend application.

**Author:** Mickali Garbutt 
**Project:** `feelflow_backend`

---

## Part 1: Running the Go Backend API

The backend is a Go server that exposes a single healthcheck endpoint.

### Prerequisites

- Go (version 1.21+) installed and configured in the system's PATH.
- A terminal or command prompt (like Git Bash or PowerShell).

### Steps to Run

1.  **Navigate to the Project Directory:**
    Open a terminal and change the directory to the root of this project.
    ```bash
    cd path/to/feelflow_backend
    ```

2.  **Install Dependencies (if not already done):**
    This command will download the necessary router package.
    ```bash
    go mod tidy
    ```

3.  **Run the Server:**
    This command will compile and run the API. The server will start on port 4000.
    ```bash
    go run ./cmd/api
    ```
    You will see a log message indicating the server has started:
    `level=INFO msg="starting server" addr=:4000 env=development`

    **Leave this terminal running.**

### How to Test the Backend Independently

1.  **Open a *new* terminal window.**
2.  **Use `curl` to send a request** to the healthcheck endpoint:
    ```bash
    curl http://localhost:4000/v1/healthcheck
    ```
3.  **Expected Output:**
    You should see a JSON response confirming the server is healthy.
    ```json
    {"status":"available","system_info":{"environment":"development","version":"1.0.0"}}
    ```

---

## Part 2: Connecting with the Flutter Frontend

The Flutter application (`mood_notes` or `feelflow_flutter`) acts as the client.

### Prerequisites

- Flutter SDK installed.
- Chrome browser installed (for web testing).

### Steps to Connect

1.  **Ensure the Go backend from Part 1 is still running.**

2.  **Update the CORS Origin (If Necessary):**
    The Flutter development server runs on a random port. Before running the Flutter app, check its URL in the browser.
    
    If the Flutter app is running on a port other than the one configured, you must update the allowed origin in the Go backend.
    -   **File to edit:** `cmd/api/routes.go`
    -   **Line to change:** `w.Header().Set("Access-Control-Allow-Origin", "http://localhost:PORT_NUMBER")`
    -   Restart the Go server after making the change.

3.  **Run the Flutter Application:**
    Open a new terminal, navigate to the Flutter project's directory, and run:
    ```bash
    flutter run -d chrome
    ```

4.  **Demonstrate the Connection:**
    A Chrome window will open with the Flutter app.
    - Click the **"Check Connection"** button.
    - The app will display a **"Connection Successful!"** message along with the status, environment, and version data fetched directly from the Go API.

This confirms that the Flutter frontend is successfully set up to connect to the backend's healthcheck endpoint.