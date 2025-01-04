package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_string"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"Creation_Date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("Hasher:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("Data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Hash:", hash)
	return hash[:11]
}

func createURL(OriginalURL string) string {
	shortUrl := generateShortURL(OriginalURL)
	id := shortUrl
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}

func getUrl(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not Forund")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET method")
	fmt.Fprintf(w, "Hello World")
}

func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error in short_url", http.StatusBadRequest)
		panic(err)
	}
	shortUrl := createURL(data.URL)
	// fmt.Fprintf(w, shortUrl)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortUrl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	url, err := getUrl(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	// fmt.Println("URL Shortner...")
	// createURL("a")
	// for key, val := range urlDB {
	// 	fmt.Printf("%s is the Key and %+v\n", key, val.CreationDate.Format("02/01/2006 15:03:04"))
	// }

	//Making of the handler function
	http.HandleFunc("/", handler)
	http.HandleFunc("/shortner", shortURLHandler)
	http.HandleFunc("/redirect/", redirect)

	//Starting Server
	fmt.Println("Server is Running on 3000....")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error in running port 3000:", err)
	}
}
