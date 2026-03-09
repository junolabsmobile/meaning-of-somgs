import { useState, useMemo } from 'react'

interface ImageQuestion {
  image: string
  correct: string
  artist: string
  options: string[]
}

const ALL_QUESTIONS: ImageQuestion[] = [
  {
    image: '/images/stairway.svg',
    correct: 'Stairway to Heaven',
    artist: 'Led Zeppelin',
    options: ['Stairway to Heaven', 'Hotel California', 'Comfortably Numb'],
  },
  {
    image: '/images/hotel-california.svg',
    correct: 'Hotel California',
    artist: 'Eagles',
    options: ['Hotel California', 'Highway to Hell', 'Born to Run'],
  },
  {
    image: '/images/bohemian-rhapsody.svg',
    correct: 'Bohemian Rhapsody',
    artist: 'Queen',
    options: ['Bohemian Rhapsody', 'Paint It Black', 'Paranoid'],
  },
  {
    image: '/images/purple-haze.svg',
    correct: 'Purple Haze',
    artist: 'Jimi Hendrix',
    options: ['Purple Haze', 'Comfortably Numb', 'Smells Like Teen Spirit'],
  },
  {
    image: '/images/comfortably-numb.svg',
    correct: 'Comfortably Numb',
    artist: 'Pink Floyd',
    options: ['Comfortably Numb', 'Paranoid', 'Sailing'],
  },
  {
    image: '/images/highway-to-hell.svg',
    correct: 'Highway to Hell',
    artist: 'AC/DC',
    options: ['Highway to Hell', 'Born to Run', 'Back in Black'],
  },
]

function shuffle<T>(arr: T[]): T[] {
  const a = [...arr]
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1))
    ;[a[i], a[j]] = [a[j], a[i]]
  }
  return a
}

const TOTAL_ROUNDS = 5

interface Props {
  onBack: () => void
}

export default function ImageQuiz({ onBack }: Props) {
  const [questions] = useState<ImageQuestion[]>(() =>
    shuffle(ALL_QUESTIONS).slice(0, TOTAL_ROUNDS),
  )
  const [index, setIndex] = useState(0)
  const [selected, setSelected] = useState<string | null>(null)
  const [score, setScore] = useState(0)
  const [phase, setPhase] = useState<'playing' | 'answered' | 'finished'>('playing')

  const current = questions[index]
  const shuffledOptions = useMemo(() => shuffle(current.options), [current])

  const handleSelect = (option: string) => {
    if (phase !== 'playing') return
    setSelected(option)
    setPhase('answered')
    if (option === current.correct) setScore((s) => s + 1)
  }

  const handleNext = () => {
    if (index + 1 >= TOTAL_ROUNDS) {
      setPhase('finished')
      return
    }
    setIndex((i) => i + 1)
    setSelected(null)
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
          <button className="quiz__next" onClick={onBack}>
            Volver al menú
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="quiz">
      <button className="song-detail__back" onClick={onBack}>
        ← Volver
      </button>
      <div className="quiz__header">
        <span className="quiz__round">
          Pregunta {index + 1} / {TOTAL_ROUNDS}
        </span>
        <span className="quiz__score-inline">Puntuación: {score}</span>
      </div>

      <div className="quiz__question">
        <img
          src={current.image}
          alt="¿Qué canción representa esta imagen?"
          className="quiz__image"
        />
        <p className="quiz__prompt">¿Qué canción representa esta imagen?</p>
      </div>

      <div className="quiz__options">
        {shuffledOptions.map((option) => {
          let cls = 'quiz__option'
          if (selected !== null) {
            if (option === current.correct) cls += ' quiz__option--correct'
            else if (option === selected) cls += ' quiz__option--wrong'
          }
          return (
            <button
              key={option}
              className={cls}
              onClick={() => handleSelect(option)}
              disabled={phase === 'answered'}
            >
              {option}
            </button>
          )
        })}
      </div>

      {phase === 'answered' && (
        <div className="quiz__feedback">
          {selected === current.correct
            ? '¡Correcto! 🎉'
            : `Era: ${current.artist} — ${current.correct}`}
        </div>
      )}

      {phase === 'answered' && (
        <button className="quiz__next" onClick={handleNext}>
          {index + 1 < TOTAL_ROUNDS ? 'Siguiente pregunta →' : 'Ver resultado'}
        </button>
      )}
    </div>
  )
}
