package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JackettResponse struct {
	Results []struct {
		Title     string `json:"title"`
		Seeders   int    `json:"seeders"`
		Leechers  int    `json:"leechers"`
		MagnetURI string `json:"magnetURI"`
	} `json:"results"`
}

func main() {
	// Introduce el nombre de la película
	pelicula := "El Señor de los Anillos: La Comunidad del Anillo"

	// URL de la API de Jackett
	url := "https://localhost:9117/api/v2.0/indexers/all/torrents/search?query=" + pelicula

	// Crea la petición HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Establece la cabecera de la API
	req.Header.Set("Authorization", "Bearer YOUR_API_KEY")

	// Envía la petición y recibe la respuesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decodifica la respuesta JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var jackettResponse JackettResponse
	err = json.Unmarshal(body, &jackettResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Muestra los resultados
	for _, torrent := range jackettResponse.Results {
		fmt.Println("**Título:**", torrent.Title)
		fmt.Println("**Seeders:**", torrent.Seeders)
		fmt.Println("**Leechers:**", torrent.Leechers)
		fmt.Println("**MagnetURI:**", torrent.MagnetURI)
		fmt.Println()
	}
}
