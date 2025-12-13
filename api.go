package hltb

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type (
	GameEntry struct {
		Id                  uint64    `json:"id"`
		HltbId              uint64    `json:"hltbId"`
		Title               string    `json:"title"`
		ImageUrl            string    `json:"imageUrl"`
		SteamAppId          uint64    `json:"steamAppId"`
		GogAppId            uint64    `json:"gogAppId"`
		MainStory           float64   `json:"mainStory"`
		MainStoryWithExtras float64   `json:"mainStoryWithExtras"`
		Completionist       float64   `json:"completionist"`
		LastUpdatedAt       time.Time `json:"lastUpdatedAt"`
	}

	TermMatchType int
	Platform      string

	searchByGameTitleRequest struct {
		SearchTerm string        `json:"searchTerm"`
		MatchType  TermMatchType `json:"matchType"`
		Platform   Platform      `json:"platform"`
	}

	SearchByGameTitleOptions struct {
		MatchType TermMatchType
		Platform  Platform
	}
)

var (
	TermMatchTypeExact TermMatchType = 0
	TermMatchTypeFuzzy TermMatchType = 1
)

var (
	All              Platform = ""
	Pc               Platform = "PC"
	Emulated         Platform = "Emulated"
	Nes              Platform = "NES"
	Snes             Platform = "Super Nintendo"
	NintendoDS       Platform = "Nintendo DS"
	Nintendo3DS      Platform = "Nintendo 3DS"
	Nintendo64       Platform = "Nintendo 64"
	NintendoGameCube Platform = "Nintendo GameCube"
	NintendoSwitch   Platform = "Nintendo Switch"
	NintendoSwitch2  Platform = "Nintendo Switch 2"
	GameBoy          Platform = "Game Boy"
	GameBoyColor     Platform = "Game Boy Color"
	GameBoyAdvance   Platform = "Game Boy Advance"
	Playstation      Platform = "PlayStation"
	Playstation2     Platform = "PlayStation 2"
	Playstation3     Platform = "PlayStation 3"
	Playstation4     Platform = "PlayStation 4"
	Playstation5     Platform = "PlayStation 5"
	PlaystationNow   Platform = "PlayStation Now"
	Wii              Platform = "Wii"
	WiiU             Platform = "Wii U"
	Xbox360          Platform = "Xbox 360"
	XboxOne          Platform = "Xbox One"
	XboxSeriesXS     Platform = "Xbox Series X/S"
)

var ErrNotFound = errors.New("hltb: game not found")

func (c *Client) GetByHltbId(ctx context.Context, id uint64) (*GameEntry, error) {
	var game GameEntry
	res, err := c.get(ctx, fmt.Sprintf(c.baseUrl+"/hltb/%d", id), &game)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("hltb.GetByHltbId(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return &game, nil
}

func (c *Client) RefreshByHltbId(ctx context.Context, id uint64) (*GameEntry, error) {
	var game GameEntry
	res, err := c.get(ctx, fmt.Sprintf(c.baseUrl+"/hltb/%d/refresh", id), &game)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("hltb.RefreshByHltbId(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return &game, nil
}

func (c *Client) SearchByGameTitle(ctx context.Context, searchTerm string, options *SearchByGameTitleOptions) ([]GameEntry, error) {
	req := searchByGameTitleRequest{
		SearchTerm: searchTerm,
		MatchType:  TermMatchTypeExact,
		Platform:   All,
	}
	if options != nil {
		if options.MatchType != 0 {
			req.MatchType = options.MatchType
		}
		if options.Platform != "" {
			req.Platform = options.Platform
		}
	}

	var games []GameEntry
	res, err := c.post(ctx, c.baseUrl+"/hltb/search", req, &games)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("hltb.SearchByGameTitle(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return games, nil
}

func (c *Client) GetBySteamAppId(ctx context.Context, id uint64) (*GameEntry, error) {
	var game GameEntry
	res, err := c.get(ctx, fmt.Sprintf(c.baseUrl+"/steam/%d", id), &game)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("hltb.GetBySteamAppId(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return &game, nil
}

func (c *Client) GetByGogAppId(ctx context.Context, id uint64) (*GameEntry, error) {
	var game GameEntry
	res, err := c.get(ctx, fmt.Sprintf(c.baseUrl+"/gog/%d", id), &game)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("hltb.GetByGogAppId(): %s, status code = %d", res.String(), res.StatusCode())
	}

	return &game, nil
}
