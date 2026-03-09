import { useState, useEffect } from 'react'
import './App.css'
import type { Song } from './types/song'
import SearchBar from './components/SearchBar'
import SongCard from './components/SongCard'
import SongDetail from './components/SongDetail'

function App() {
  const [songs, setSongs] = useState<Song[]>([])
  const [query, setQuery] = useState('')
  const [selectedSong, setSelectedSong] = useState<Song | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetch('/api/songs')
      .then((res) => {
        if (!res.ok) throw new Error('Error al cargar las canciones')
        return res.json()
      })
      .then((data: Song[]) => {
        const sorted = [...data].sort((a, b) => {
          const artistCmp = a.Artist.localeCompare(b.Artist)
          return artistCmp !== 0 ? artistCmp : a.Title.localeCompare(b.Title)
        })
        setSongs(sorted)
      })
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false))
  }, [])

  const filtered = songs.filter((s) => {
    const q = query.toLowerCase()
    return s.Title.toLowerCase().includes(q) || s.Artist.toLowerCase().includes(q)
  })

  if (selectedSong) {
    return (
      <div className="app">
        <SongDetail song={selectedSong} onBack={() => setSelectedSong(null)} />
      </div>
    )
  }

  return (
    <div className="app">
      <header className="app-header">
        <div className="app-header__inner">
          <h1 className="app-header__title">🎵 Meaning of Songs</h1>
          <SearchBar value={query} onChange={setQuery} />
        </div>
      </header>

      <main className="app-main">
        {loading && <p className="app-status">Cargando canciones...</p>}
        {error && <p className="app-status app-status--error">Error: {error}</p>}
        {!loading && !error && filtered.length === 0 && (
          <p className="app-status">No se encontraron canciones para "{query}".</p>
        )}
        {!loading && !error && (
          <div className="song-grid">
            {filtered.map((song) => (
              <SongCard key={song.ID} song={song} onClick={setSelectedSong} />
            ))}
          </div>
        )}
      </main>
    </div>
  )
}

export default App
