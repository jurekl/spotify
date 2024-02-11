package spotify

import (
	"context"
	//"errors"
	//"fmt"
	"net/http"
	//"net/url"
	"strings"
)

/*
jest w jl_spotifyauth_ext.go
const (
	//ScopeUserReadPlaybackPosition read usersposition in content played
	ScopeUserReadPlaybackPosition = "user-read-playback-position"
)
*/

type SavedEpisodePage struct {
	basePage
	Episodes []SavedEpisode `json:"items"`
}

type SavedEpisode struct {
	// The date and time the show was saved, represented as an ISO
	// 8601 UTC timestamp with a zero offset (YYYY-MM-DDTHH:MM:SSZ).
	// You can use the TimestampLayout constant to convert this to
	// a time.Time value.
	AddedAt     string `json:"added_at"`
	EpisodePage `json:"episode"`
}

// SaveEpisodesForCurrentUser saves one or more episodes to current Spotify user's library.
// API reference: https://developer.spotify.com/documentation/web-api/reference/save-episodes-user
func (c *Client) SaveEpisodesForCurrentUser(ctx context.Context, ids []ID) error {
	spotifyURL := c.baseURL + "me/episodes?ids=" + strings.Join(toStringSlice(ids), ",")
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, spotifyURL, nil)
	if err != nil {
		return err
	}

	return c.execute(req, nil, http.StatusOK)
}

// CurrentUsersEpisodes fetches the user's episodes
// sensible defaults. The default limit is 20 and the default timerange
// is medium_term. This call requires ScopeUserTopRead.
//
// Supported options: Limit, Offset
func (c *Client) CurrentUsersEpisodes(ctx context.Context, opts ...RequestOption) (*SavedEpisodePage, error) {
	spotifyURL := c.baseURL + "me/episodes"
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		spotifyURL += "?" + params
	}

	var result SavedEpisodePage // SimpleEpisodePage

	err := c.get(ctx, spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
