package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/kr/pretty"
)

const (
	searchBookEndpoint = "/v1/volumes?q="
)

func (r *Requests) QueryBooksByName(searchQuery string) (*models.SearchQueryResponse, error) {

	req, err := http.NewRequest("GET", r.config.GoogleBooksURL+searchBookEndpoint+searchQuery, nil)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[QueryBooksByName] Response error --> %s", err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		println("HESKLWHJSKDHG")
		return nil, bcerrors.NewError(
			"Got invalid response code from API", bcerrors.InternalError,
		).WithExternalMessage(fmt.Sprintf("Expected: 200, Got: %d", resp.StatusCode))
	}
	defer closeResponse(resp)

	pretty.Print(resp.Body)
	var decodedResponse models.SearchQueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&decodedResponse); err != nil {
		if err != nil {
			log.Printf("[QueryBooksByName] Response error #3 --> %s", err.Error())
			return nil, err
		}
	}

	return &decodedResponse, nil
}
