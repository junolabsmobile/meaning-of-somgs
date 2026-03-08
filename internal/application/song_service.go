package application

import (
	"context"

	"github.com/junolabsmobile/meaning-of-songs/internal/domain"
)

// SongService implementa los casos de uso de la aplicación.
type SongService struct {
	repo domain.SongRepository
}

// NewSongService crea el servicio inyectando el repositorio (puerto).
func NewSongService(repo domain.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetSong(ctx context.Context, id string) (*domain.Song, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *SongService) ListSongs(ctx context.Context) ([]domain.Song, error) {
	return s.repo.FindAll(ctx)
}

func (s *SongService) CreateSong(ctx context.Context, song *domain.Song) error {
	return s.repo.Save(ctx, song)
}
