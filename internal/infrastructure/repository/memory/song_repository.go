package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/junolabsmobile/meaning-of-songs/internal/domain"
)

// SongRepository es un adaptador in-memory para el puerto domain.SongRepository.
type SongRepository struct {
	mu    sync.RWMutex
	songs map[string]*domain.Song
}

func NewSongRepository() *SongRepository {
	return &SongRepository{
		songs: make(map[string]*domain.Song),
	}
}

func (r *SongRepository) FindByID(_ context.Context, id string) (*domain.Song, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	song, ok := r.songs[id]
	if !ok {
		return nil, fmt.Errorf("song not found: %s", id)
	}
	return song, nil
}

func (r *SongRepository) FindAll(_ context.Context) ([]domain.Song, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]domain.Song, 0, len(r.songs))
	for _, s := range r.songs {
		result = append(result, *s)
	}
	return result, nil
}

func (r *SongRepository) Save(_ context.Context, song *domain.Song) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.songs[song.ID] = song
	return nil
}
