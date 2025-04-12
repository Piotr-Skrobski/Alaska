package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/models"
)

type OMDbService struct {
	apiKey  string
	baseURL string
}

type omdbResponse struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
	BoxOffice  string `json:"BoxOffice"`
	Response   string `json:"Response"`
	Error      string `json:"Error,omitempty"`
}

func NewOMDbService(apiKey string) *OMDbService {
	return &OMDbService{
		apiKey:  apiKey,
		baseURL: "http://www.omdbapi.com/",
	}
}

func (c *OMDbService) GetMovieByTitle(title string) (models.Movie, error) {
	url := fmt.Sprintf("%s?apikey=%s&t=%s", c.baseURL, c.apiKey, title)
	return c.fetchMovie(url)
}

func (c *OMDbService) GetMovieByIMDbID(imdbID string) (models.Movie, error) {
	url := fmt.Sprintf("%s?apikey=%s&i=%s", c.baseURL, c.apiKey, imdbID)
	return c.fetchMovie(url)
}

func (c *OMDbService) fetchMovie(url string) (models.Movie, error) {
	omdbResp, err := c.fetchOMDbResponse(url)
	if err != nil {
		return models.Movie{}, err
	}

	return c.convertToMovie(omdbResp)
}

func (c *OMDbService) fetchOMDbResponse(url string) (omdbResponse, error) {
	var omdbResp omdbResponse

	resp, err := http.Get(url)
	if err != nil {
		return omdbResp, fmt.Errorf("failed to fetch from OMDb: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&omdbResp); err != nil {
		return omdbResp, fmt.Errorf("failed to decode OMDb response: %w", err)
	}

	if omdbResp.Response != "True" {
		return omdbResp, fmt.Errorf("OMDb error: %s", omdbResp.Error)
	}

	return omdbResp, nil
}

func (c *OMDbService) convertToMovie(omdbResp omdbResponse) (models.Movie, error) {
	imdbRating, _ := strconv.ParseFloat(omdbResp.ImdbRating, 64)

	boxOffice := c.parseBoxOffice(omdbResp.BoxOffice)
	actors := c.parseActors(omdbResp.Actors)

	movie := models.Movie{
		Title:      omdbResp.Title,
		Year:       omdbResp.Year,
		Genre:      omdbResp.Genre,
		Director:   omdbResp.Director,
		Writer:     omdbResp.Writer,
		Actors:     actors,
		Plot:       omdbResp.Plot,
		Language:   omdbResp.Language,
		Country:    omdbResp.Country,
		PosterURL:  omdbResp.Poster,
		IMDbRating: imdbRating,
		Runtime:    omdbResp.Runtime,
		BoxOffice:  boxOffice,
		Awards:     omdbResp.Awards,
		Metascore:  omdbResp.Metascore,
		IMDbID:     omdbResp.ImdbID,
	}

	return movie, nil
}

func (c *OMDbService) parseActors(actorsStr string) []string {
	if actorsStr == "N/A" || actorsStr == "" {
		return []string{}
	}

	return strings.Split(actorsStr, ", ")
}

func (c *OMDbService) parseBoxOffice(boxOfficeStr string) models.BoxOffice {
	var boxOffice models.BoxOffice

	if boxOfficeStr == "N/A" || boxOfficeStr == "" {
		return boxOffice
	}

	valueStr := boxOfficeStr
	currencySymbol := ""

	if len(valueStr) > 0 && !('0' <= valueStr[0] && valueStr[0] <= '9') {
		currencySymbol = string(valueStr[0])
		valueStr = valueStr[1:]
	}

	valueStr = strings.ReplaceAll(valueStr, ",", "")
	value, err := strconv.ParseFloat(valueStr, 64)

	if err == nil {
		boxOffice = models.BoxOffice{
			Value:    value,
			Currency: currencySymbol,
		}
	}

	return boxOffice
}
