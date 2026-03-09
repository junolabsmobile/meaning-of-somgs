import { useState, useEffect } from 'react'
import './App.css'
import type { Song } from './types/song'
import SearchBar from './components/SearchBar'
import SongCard from './components/SongCard'
import SongDetail from './components/SongDetail'
import GamesHub from './components/games/GamesHub'
import DescriptionQuiz from './components/games/DescriptionQuiz'
import ImageQuiz from './components/games/ImageQuiz'

type Section = 'songs' | 'games'
type ActiveGame = 'hub' | 'description' | 'image'

function App() {
  const [songs, setSongs] = useState<Song[]>([])
  const [query, setQuery] = useState('')
  const [selectedSong, setSelectedSong] = useState<Song | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [section, setSection] = useState<Section>('songs')
  const [activeGame, setActiveGame] = useState<ActiveGame>('hub')

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

  const handleSectionChange = (s: Section) => {
    setSection(s)
    setSelectedSong(null)
    if (s === 'games') setActiveGame('hub')
  }

  const renderContent = () => {
    if (section === 'songs') {
      if (selectedSong) {
        return <SongDetail song={selectedSong} onBack={() => setSelectedSong(null)} />
      }
      return (
        <>
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
        </>
      )
    }

    // section === 'games'
    if (activeGame === 'description') {
      return <DescriptionQuiz songs={songs} onBack={() => setActiveGame('hub')} />
    }
    if (activeGame === 'image') {
      return <ImageQuiz onBack={() => setActiveGame('hub')} />
    }
    return <GamesHub onSelectGame={setActiveGame} songs={songs} />
  }

  return (
    <div className="app">
      <header className="app-header">
        <div className="app-header__inner">
          <h1 className="app-header__title">🎵 Meaning of Songs</h1>
          <nav className="app-nav">
            <button
              className={`app-nav__tab${section === 'songs' ? ' app-nav__tab--active' : ''}`}
              onClick={() => handleSectionChange('songs')}
            >
              Canciones
            </button>
            <button
              className={`app-nav__tab${section === 'games' ? ' app-nav__tab--active' : ''}`}
              onClick={() => handleSectionChange('games')}
            >
              Juegos
            </button>
          </nav>
          {section === 'songs' && !selectedSong && (
            <SearchBar value={query} onChange={setQuery} />
          )}
        </div>
      </header>

      <main className="app-main">{renderContent()}</main>
    </div>
  )
}

export default App
