package binocs

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ChannelService struct {
	resty *resty.Client
}

type Channel struct {
	ID        int      `json:"id,omitempty"`
	Ident     string   `json:"ident,omitempty"`
	Type      string   `json:"type,omitempty"`
	Alias     string   `json:"alias"`
	Handle    string   `json:"handle,omitempty"`
	UsedCount int      `json:"used_count,omitempty"`
	LastUsed  string   `json:"last_used,omitempty"`
	Verified  string   `json:"verified,omitempty"`
	Checks    []string `json:"checks,omitempty"`
}

type ChannelAttachment struct {
	NotificationType string `json:"notification_type"`
}

func (s *ChannelService) Create(c Channel) (Channel, error) {
	var result Channel
	resp, err := s.resty.R().
		SetBody(c).
		SetResult(&result).
		Post("/channels")
	if err != nil {
		return result, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return result, fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result, nil
}

func (s *ChannelService) Read(ident string) (Channel, error) {
	var result Channel
	resp, err := s.resty.R().
		SetResult(&result).
		Get("/channels/" + ident)
	if err != nil {
		return Channel{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return Channel{}, fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result, nil
}

func (s *ChannelService) List() ([]Channel, error) {
	var result []Channel
	resp, err := s.resty.R().
		SetResult(&result).
		Get("/channels")
	if err != nil {
		return []Channel{}, err
	}
	if resp.StatusCode() != http.StatusOK {
		return []Channel{}, fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return result, nil
}

func (s *ChannelService) Update(ident string, c Channel) error {
	resp, err := s.resty.R().
		SetBody(c).
		Put("/channels/" + ident)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return nil
}

func (s *ChannelService) Delete(ident string) error {
	resp, err := s.resty.R().
		Delete("/channels/" + ident)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return nil
}

func (s *ChannelService) Attach(channel string, check string) error {
	b := ChannelAttachment{}
	resp, err := s.resty.R().
		SetBody(b).
		Post("/channels/" + channel + "/check/" + check)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return nil
}

func (s *ChannelService) Detach(channel string, check string) error {
	b := ChannelAttachment{}
	resp, err := s.resty.R().
		SetBody(b).
		Delete("/channels/" + channel + "/check/" + check)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("API returned %d", resp.StatusCode())
	}
	return nil
}
