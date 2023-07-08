package cotation

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Service struct {
	repo   Repository
	client http.Client
}

func NewService(repo Repository) *Service {
	return &Service{
		repo:   repo,
		client: http.Client{Timeout: time.Duration(1) * time.Second},
	}
}

func (s *Service) Get(ctx context.Context) (*Cotation, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c Cotation
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.Store(ctx, &c)
	fmt.Printf("ID: %s\n", id)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (s *Service) WriteFile(cotation *Cotation) error {
	//write file cotation.txt with field bid
	file, err := os.Create("cotation.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", cotation.USDBRL.Bid))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
