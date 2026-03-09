import { useState, useCallback } from 'react'
import type { Song } from '../../types/song'

interface Props {
  songs: Song[]
  onBack: () => void
}

function shuffle<T>(arr: T[]): T[] {
  const a = [...arr]
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1))
    ;[a[i], a[j]] = [a[j], a[i]]
  }
  return a
}

interface Round {
  correct: Song
  options: Song[]
}

function generateRound(songs: Song[], usedIds: Set<string>): Round {
  const available = songs.filter((s) => !usedIds.has(s.ID))
  const pool = available.length >= 3 ? available : songs
  const shuffled = shuffle(pool)
  const correct = shuffled[0]
  const wrongs = shuffle(songs.filter((s) => s.ID !== correct.ID)).slice(0, 2)
  return { correct, options: shuffle([correct, ...wrongs]) }
}

const TOTAL_ROUNDS = 5

export default function DescriptionQuiz({ songs, onBack }: Props) {
  const [usedIds] = useState(() => new Set<string>())
  const [round, setRound] = useState<Round>(() => generateRound(songs, new Set()))
  const [selected, setSelected] = useState<string | null>(null)
  const [score, setScore] = useState(0)
  const [roundNum, setRoundNum] = useState(1)
  const [phase, setPhase] = useState<'playing' | 'answered' | 'finished'>('playing')

  const handleSelect = (id: string) => {
    if (phase !== 'playing') return
    setSelected(id)
    setPhase('answered')
    if (id === round.correct.ID) setScore((s) => s + 1)
  }

  const handleNext = useCallback(() => {
    if (roundNum >= TOTAL_ROUNDS) {
      setPhase('finished')
      return
    }
    usedIds.add(round.correct.ID)
    const next = generateRound(songs, usedIds)
    setRound(next)
    setSelected(null)
    setPhase('playing')
    setRoundNum((n) => n + 1)
  }, [roundNum, round, songs, usedIds])

  const handleRestart = () => {
    usedIds.clear()
    const next = generateRound(songs, new Set())
    setRound(next)
    setSelected(null)
    setScore(0)
    setRoundNum(1)
    setPhase('playing')
  }

  if (phase === 'finished') {
    const emoji = score >= 4 ? '🏆' : score >= 2 ? '🎵' : '🎸'
    return (
      <div className="quiz">
        <button className="song-detail__back" onClick={onBack}>
          ← Volver
        </button>
        <div className="quiz__score">
          <div className="quiz__score-emoji">{emoji}</div>
          <h2>¡Juego terminado!</h2>
          <p className="quiz__score-result">
            {score} de {TOTAL_ROUNDS} correctas
          </p>
          <button className="quiz__next" onClick={handleRestart}>
            Jugar de nuevo
          </button>
        </div>
      </div>
    )
  }

  const excerpt =
    round.correct.Meaning.slice(0, 200).trim() +
    (round.correct.Meaning.length > 200 ? '...' : '')

  return (
    <div className="quiz">
      <button className="song-detail__back" onClick={onBack}>
        ← Volver
      </button>
      <div className="quiz__header">
        <span className="quiz__round">
          Ronda {roundNum} / {TOTAL_ROUNDS}
        </span>
        <span className="quiz__score-inline">Puntuación: {score}</span>
      </div>

      <div className="quiz__question">
        <p className="quiz__excerpt">"{excerpt}"</p>
      </div>

      <div className="quiz__options">
        {round.options.map((song) => {
          let cls = 'quiz__option'
          if (selected !== null) {
            if (song.ID === round.correct.ID) cls += ' quiz__option--correct'
            else if (song.ID === selected) cls += ' quiz__option--wrong'
          }
          return (
            <button
              key={song.ID}
              className={cls}
              onClick={() => handleSelect(song.ID)}
              disabled={phase === 'answered'}
            >
              {song.Artist} — {song.Title}
            </button>
          )
        })}
      </div>

      {phase === 'answered' && (
        <div className="quiz__feedback">
          {selected === round.correct.ID
            ? '¡Correcto! 🎉'
            : `Era: ${round.correct.Artist} — ${round.correct.Title}`}
        </div>
      )}

      {phase === 'answered' && (
        <button className="quiz__next" onClick={handleNext}>
          {roundNum < TOTAL_ROUNDS ? 'Siguiente pregunta →' : 'Ver resultado'}
        </button>
      )}
    </div>
  )
}
