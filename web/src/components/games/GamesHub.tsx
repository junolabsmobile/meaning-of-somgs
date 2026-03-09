import type { Song } from '../../types/song'

interface Props {
  onSelectGame: (game: 'description' | 'image') => void
  songs: Song[]
}

export default function GamesHub({ onSelectGame, songs }: Props) {
  const ready = songs.length > 0

  return (
    <div className="games-hub">
      <h2 className="games-hub__title">Juegos Musicales</h2>
      <div className="games-hub__grid">
        <div
          className={`game-card${!ready ? ' game-card--disabled' : ''}`}
          onClick={() => ready && onSelectGame('description')}
          tabIndex={ready ? 0 : -1}
          role="button"
          onKeyDown={(e) => e.key === 'Enter' && ready && onSelectGame('description')}
        >
          <span className="game-card__icon">📝</span>
          <h3 className="game-card__title">Adivina la Canción</h3>
          <p className="game-card__desc">
            Lee el significado de una canción y adivina cuál es entre tres opciones.
          </p>
          {!ready && <p className="game-card__loading">Cargando canciones...</p>}
        </div>

        <div
          className="game-card"
          onClick={() => onSelectGame('image')}
          tabIndex={0}
          role="button"
          onKeyDown={(e) => e.key === 'Enter' && onSelectGame('image')}
        >
          <span className="game-card__icon">🎨</span>
          <h3 className="game-card__title">¿Qué canción es?</h3>
          <p className="game-card__desc">
            Observa la imagen e identifica qué famosa canción representa.
          </p>
        </div>
      </div>
    </div>
  )
}
