package services

import (
	"context"
	"fmt"
	"os"

	"github.com/JuanDiego3241/blog/src/models"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// SpotifyService encapsula el cliente Spotify + lógica
type SpotifyService struct {
	client *spotify.Client
}

func NewSpotifyService() (*SpotifyService, error) {
	cfg := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}

	ctx := context.Background()
	httpClient := cfg.Client(ctx)
	client := spotify.New(httpClient)
	return &SpotifyService{client: client}, nil
}

func (s *SpotifyService) FetchPlaylist(ctx context.Context, playlistID string) (*models.Playlist, error) {
	pl, err := s.client.GetPlaylist(ctx, spotify.ID(playlistID))
	if err != nil {
		return nil, fmt.Errorf("playlist no encontrada: %w", err)
	}
	tracksPage, err := s.client.GetPlaylistTracks(ctx, spotify.ID(playlistID))
	if err != nil {
		return nil, fmt.Errorf("error al obtener pistas: %w", err)
	}

	playlist := &models.Playlist{
		ID:          string(pl.ID),
		Name:        pl.Name,
		Description: pl.Description,
		Tracks:      make([]models.Track, len(tracksPage.Tracks)),
	}
	for i, item := range tracksPage.Tracks {
		playlist.Tracks[i] = models.Track{
			ID:     string(item.Track.ID),
			Name:   item.Track.Name,
			Artist: item.Track.Artists[0].Name,
			URI:    string(item.Track.URI),
		}
	}
	return playlist, nil
}
