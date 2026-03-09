import type { Song } from '../types/song'

interface SongDetailProps {
  song: Song
  onBack: () => void
}

export default function SongDetail({ song, onBack }: SongDetailProps) {
  return (
    <div className="song-detail">
      <button className="song-detail__back" onClick={onBack}>
        ← Volver
      </button>
      <p className="song-detail__artist">{song.Artist.toUpperCase()}</p>
      <h1 className="song-detail__title">{song.Title}</h1>

      <section className="song-detail__section">
        <h2 className="song-detail__section-title">Historia</h2>
        <p className="song-detail__text">{song.History}</p>
      </section>

      <section className="song-detail__section">
        <h2 className="song-detail__section-title">Significado</h2>
        <p className="song-detail__text">{song.Meaning}</p>
      </section>
    </div>
  )
}
