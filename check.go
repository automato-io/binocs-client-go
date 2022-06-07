package binocs

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type CheckService struct {
	resty *resty.Client
}

type Check struct {
	ID                         int      `json:"id,omitempty"`
	Ident                      string   `json:"ident,omitempty"`
	Name                       string   `json:"name"`
	Protocol                   string   `json:"protocol,omitempty"`
	Resource                   string   `json:"resource,omitempty"`
	Method                     string   `json:"method,omitempty"`
	Interval                   int      `json:"interval,omitempty"`
	Target                     float64  `json:"target,omitempty"`
	Regions                    []string `json:"regions,omitempty"`
	UpCodes                    string   `json:"up_codes,omitempty"`
	UpConfirmationsThreshold   int      `json:"up_confirmations_threshold,omitempty"`
	UpConfirmations            int      `json:"up_confirmations,omitempty"`
	DownConfirmationsThreshold int      `json:"down_confirmations_threshold,omitempty"`
	DownConfirmations          int      `json:"down_confirmations,omitempty"`
	LastStatus                 int      `json:"last_status,omitempty"`
	LastStatusCode             string   `json:"last_status_code,omitempty"`
	LastStatusDuration         string   `json:"last_status_duration,omitempty"`
	Created                    string   `json:"created,omitempty"`
	Updated                    string   `json:"updated,omitempty"`
	Channels                   []string `json:"channels,omitempty"`
}

func (s *CheckService) Create(c Check) (Check, error) {
	var result Check
	resp, err := s.resty.R().
		SetBody(c).
		SetResult(&result).
		Post("/checks")
	if err != nil {
		return result, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return result, fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result, nil
}

func (s *CheckService) Read(ident string) (Check, error) {
	var result Check
	resp, err := s.resty.R().
		SetResult(&result).
		Get("/checks/" + ident)
	if err != nil {
		return Check{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return Check{}, fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result, nil
}

func (s *CheckService) List() ([]Check, error) {
	var result []Check
	resp, err := s.resty.R().
		SetResult(&result).
		Get("/checks")
	if err != nil {
		return []Check{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return []Check{}, fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result, nil
}

func (s *CheckService) Update(ident string, c Check) error {
	resp, err := s.resty.R().
		SetBody(c).
		Put("/checks/" + ident)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return nil
}

func (s *CheckService) Delete(ident string) error {
	resp, err := s.resty.R().
		Delete("/checks/" + ident)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return nil
}
