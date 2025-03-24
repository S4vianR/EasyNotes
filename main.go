package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Set CORS headers
func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow Content-Type header
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, DELETE") // Allow methods
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setCORSHeaders(w)
	response := map[string]string{"status": "The server is running"}
	json.NewEncoder(w).Encode(response)
}

// Function for fetching the notes from notes.json
func getNotes() []Note {
	notes := []Note{}
	file, err := os.ReadFile("notes.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return notes
	}
	json.Unmarshal(file, &notes)
	return notes
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setCORSHeaders(w)
	notes := getNotes()
	json.NewEncoder(w).Encode(notes)
}

// Function for saving the notes to notes.json
func saveNotes(notes []Note) {
	file, err := json.MarshalIndent(notes, "", " ")
	if err != nil {
		fmt.Println("Error marshalling notes:", err)
		return
	}
	os.WriteFile("notes.json", file, 0644)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setCORSHeaders(w)

	// Handle preflight request
	if r.Method == http.MethodOptions {
		return
	}

	var note Note
	json.NewDecoder(r.Body).Decode(&note)

	// Validate that title and content are not empty
	if note.Title != "" && note.Content != "" {
		notes := getNotes()
		notes = append(notes, note)
		saveNotes(notes)
		json.NewEncoder(w).Encode(notes)
	} else {
		http.Error(w, "Title and Content cannot be empty", http.StatusBadRequest)
	}
}

func main() {
	// Server port
	var port string = "3001"

	// API endpoints
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/notes", getNoteHandler)
	http.HandleFunc("/api/create", createNoteHandler)

	// Run server
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
