package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Line struct {
	ID    string `json:"id"`
	Start Point  `json:"start"`
	End   Point  `json:"end"`
}

type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type IntersectionResult struct {
	LineID       string `json:"lineId"`
	Intersection Point  `json:"intersection"`
}

func main() {
	http.HandleFunc("/intersections", handleIntersections)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIntersections(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	// Check for the POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check for the header-based auth check
	authHeader := r.Header.Get("Authorization")
	if authHeader != "YOUR_AUTH_TOKEN" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse the request body
	var linestring []Point
	err := json.NewDecoder(r.Body).Decode(&linestring)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the linestring
	if len(linestring) == 0 {
		http.Error(w, "Invalid linestring", http.StatusInternalServerError)
		return
	}

	// Define the set of 50 lines
	lines := generateLines()

	// Find intersecting lines
	intersections := findIntersections(linestring, lines)

	// Return the result as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(intersections)
}

// Function to generate the set of 50 lines (start and end points)
func generateLines() []Line {
	// Generate the lines (start and end points) randomly
	// Replace this with your logic to generate the actual lines
	lines := make([]Line, 50)
	for i := 0; i < 50; i++ {
		start := Point{
			Latitude:  generateRandomCoordinate(),
			Longitude: generateRandomCoordinate(),
		}
		end := Point{
			Latitude:  generateRandomCoordinate(),
			Longitude: generateRandomCoordinate(),
		}
		lines[i] = Line{
			ID:    fmt.Sprintf("L%02d", i+1),
			Start: start,
			End:   end,
		}
	}
	return lines
}

// Function to find intersecting lines
func findIntersections(linestring []Point, lines []Line) []IntersectionResult {
	intersections := make([]IntersectionResult, 0)

	// Implementing  logic to find the intersecting lines based on turfjs and linestring

	return intersections
}

// Helper function to generate random coordinates within a range
func generateRandomCoordinate() float64 {

	return 0.0
}
