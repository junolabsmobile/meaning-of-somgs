import type { Song } from '../types/song'

interface SongCardProps {
  song: Song
  onClick: (song: Song) => void
}

export default function SongCard({ song, onClick }: SongCardProps) {
  const excerpt = song.Meaning.length > 100
    ? song.Meaning.slice(0, 100).trimEnd() + '…'
    : song.Meaning

  return (
    <article className="song-card" onClick={() => onClick(song)} role="button" tabIndex={0}
      onKeyDown={(e) => e.key === 'Enter' && onClick(song)}>
      <p className="song-card__artist">{song.Artist.toUpperCase()}</p>
      <h2 className="song-card__title">{song.Title}</h2>
      <p className="song-card__excerpt">{excerpt}</p>
    </article>
  )
}
