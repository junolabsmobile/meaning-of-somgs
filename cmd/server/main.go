package main

import (
	"context"
	"log"
	"net/http"

	"github.com/junolabsmobile/meaning-of-songs/internal/application"
	"github.com/junolabsmobile/meaning-of-songs/internal/domain"
	httpinfra "github.com/junolabsmobile/meaning-of-songs/internal/infrastructure/http"
	"github.com/junolabsmobile/meaning-of-songs/internal/infrastructure/repository/memory"
)

func main() {
	// Crear adaptador de salida (repositorio in-memory)
	songRepo := memory.NewSongRepository()

	// Crear servicio de aplicación inyectando el puerto
	songService := application.NewSongService(songRepo)

	// Seedear canciones de ejemplo
	seedSongs(context.Background(), songService)

	// Crear router HTTP (adaptador de entrada)
	router := httpinfra.NewRouter(songService)

	// Arrancar servidor
	addr := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func seedSongs(ctx context.Context, svc *application.SongService) {
	songs := []*domain.Song{
		{
			ID:      "1",
			Title:   "Stairway to Heaven",
			Artist:  "Led Zeppelin",
			History: "Grabada en 1971 para el álbum 'Led Zeppelin IV', es considerada una de las mejores canciones de rock de todos los tiempos. Jimmy Page compuso la música y Robert Plant escribió la letra en una sola noche inspirada en una cabaña de Gales. La canción tardó más de ocho minutos en construirse, pasando de una delicada balada acústica a un estallido eléctrico. Su icónico riff de guitarra acústica se convirtió en uno de los más reconocibles de la historia. Fue lanzada sin single para mantener su integridad artística.",
			Meaning: "La letra describe el viaje de una mujer materialista que trata de 'comprar una escalera al cielo', y cómo ese camino lleva al error espiritual. Plant exploró temas de redención, el contraste entre lo material y lo espiritual, y la búsqueda del sentido de la vida. La frase 'there are two paths you can go by' refleja la elección moral que todos enfrentamos. Muchos intérpretes ven referencias a la mitología celta y al ocultismo. Es una canción que invita a la reflexión sobre los valores que guían nuestra existencia.",
		},
		{
			ID:      "2",
			Title:   "Paint It Black",
			Artist:  "The Rolling Stones",
			History: "Lanzada en 1966 como single del álbum 'Aftermath', marcó un giro oscuro y psicodélico para los Stones. Brian Jones introdujo el sitar indio, un sonido inusual para el rock occidental de la época. La canción llegó al número 1 en Reino Unido y Estados Unidos de inmediato. Su producción mezcla el rock crudo con influencias de la música hindú, algo novedoso en el pop de los 60. Ha sido usada en múltiples soundtracks y series, incluyendo 'Full Metal Jacket' y 'Westworld'.",
			Meaning: "La letra narra la perspectiva de alguien sumido en el duelo, deseando que todo a su alrededor se vuelva negro para reflejar su dolor interior. Mick Jagger canta desde el punto de vista de alguien que ha perdido a un ser querido y no puede soportar ver colores alegres. El narrador quiere borrar el rojo de su amada, símbolo de vida y amor ahora ausentes. Es una de las primeras canciones del rock en explorar abiertamente la depresión y el luto. Su oscuridad lírica contrasta con la energía hipnótica del sitar.",
		},
		{
			ID:      "3",
			Title:   "Hotel California",
			Artist:  "Eagles",
			History: "Publicada en 1977 como el tema central del álbum homónimo, es la canción más emblemática de los Eagles. Don Felder compuso el riff de guitarra original mientras conducía por Malibu, y Don Henley escribió la letra. Ganó el Grammy a la mejor grabación del año en 1978. El hotel de la canción no existe realmente; fue inspirado en la cultura hedonista de Los Ángeles de los 70. El duelo de guitarras de Felder y Joe Walsh al final es uno de los solos más celebrados del rock.",
			Meaning: "La canción es una metáfora del sueño americano y sus trampas: la fama, el dinero y el exceso que atrapan sin posibilidad de escape. 'You can check out any time you like, but you can never leave' simboliza la adicción y la vida de los excesos en la industria musical. El hotel representa Hollywood y el estilo de vida de los ricos y famosos que seduce pero destruye. Henley describió la letra como un comentario sobre la decadencia del sueño americano en los años 70. Es una advertencia disfrazada de historia de misterio.",
		},
		{
			ID:      "4",
			Title:   "Bohemian Rhapsody",
			Artist:  "Queen",
			History: "Compuesta por Freddie Mercury y lanzada en 1975 en el álbum 'A Night at the Opera', revolucionó el concepto de lo que podía ser una canción de rock. Dura casi seis minutos y mezcla balada, ópera, hard rock y a cappella en una sola pieza sin precedentes. Las voces del coro fueron grabadas durante semanas con Mercury, May y Taylor cantando las mismas partes repetidamente. EMI se negó a lanzarla como single por su duración, pero el DJ Kenny Everett la tocó en radio y la demanda del público obligó a publicarla. Volvió al número 1 en 1992 tras el fallecimiento de Mercury.",
			Meaning: "Mercury nunca reveló el significado definitivo, pero la letra parece narrar la historia de un joven que ha cometido un crimen y enfrenta las consecuencias ante Dios y el diablo. Muchos intérpretes ven la canción como la expresión de la lucha interna de Mercury con su identidad sexual y los secretos que guardaba. 'Mama, just killed a man' podría ser una metáfora de matar la identidad anterior para convertirse en quien realmente eres. La sección operística con Galileo y Bismillah evoca un juicio divino surrealista. Es una obra que invita a cada oyente a encontrar su propio significado.",
		},
		{
			ID:      "5",
			Title:   "Smells Like Teen Spirit",
			Artist:  "Nirvana",
			History: "Lanzada en 1991 como primer single del álbum 'Nevermind', convirtió a Nirvana en la banda más importante del mundo casi de la noche a la mañana. Kurt Cobain admitió que el riff fue deliberadamente inspirado en Pixies, intentando imitar su dinámica suave-fuerte. El título viene de una marca de desodorante que Kathleen Hanna (de Bikini Kill) le pintó en la pared. El video musical, filmado en un gimnasio de instituto, se convirtió en uno de los más reproducidos de la historia de MTV. Marcó el inicio del grunge como movimiento cultural dominante.",
			Meaning: "Cobain escribió la letra en estado de semiinconsciencia, llenándola de frases crípticas que él mismo describió como 'sin sentido real'. La canción captura la apatía, la rabia y la alienación de la Generación X de los 90. 'Here we are now, entertain us' refleja el cinismo de una generación que se sentía usada y vaciada por la cultura del entretenimiento. Cobain bromeaba diciendo que era un himno para los adolescentes que no tenían himno. Paradójicamente, se convirtió en el anthem masivo que intentaba criticar.",
		},
		{
			ID:      "6",
			Title:   "Purple Haze",
			Artist:  "Jimi Hendrix",
			History: "Publicada en 1967 como el primer single de Jimi Hendrix en Reino Unido, definió instantáneamente su sonido revolucionario. Hendrix compuso la canción en una servilleta tras un sueño en el que caminaba bajo el agua. El riff inicial usa el llamado 'tritono' o 'diablo en la música', un intervalo considerado disonante durante siglos. La grabación tardó solo tres horas en completarse en los estudios de Londres. Es considerada una de las mejores actuaciones de guitarra eléctrica jamás registradas.",
			Meaning: "Aunque muchos la asocian con las drogas psicodélicas de los 60, Hendrix siempre dijo que trataba sobre estar enamorado y desorientado por ese sentimiento. 'Purple haze, all in my brain' describe el estado de confusión y éxtasis que produce el amor intenso. La frase 'excuse me while I kiss the sky' —frecuentemente malinterpretada como 'kiss this guy'— es un vuelo poético sobre la trascendencia. Hendrix también declaró que la canción estaba inspirada en un cuento de ciencia ficción sobre un anillo místico de color púrpura. La ambigüedad deliberada de la letra permite múltiples lecturas.",
		},
		{
			ID:      "7",
			Title:   "Comfortably Numb",
			Artist:  "Pink Floyd",
			History: "Grabada en 1979 para el álbum conceptual 'The Wall', es considerada la canción más grande de Pink Floyd. Roger Waters escribió la letra basándose en una experiencia real de 1977 cuando le administraron un calmante para poder actuar enfermo en Philadelphia. David Gilmour y Waters tuvieron una disputa creativa sobre los acordes, y la versión final combina ambas visiones. Los solos de guitarra de Gilmour son constantemente votados entre los mejores de la historia del rock. La canción aparece en el momento más oscuro del álbum, cuando el protagonista Pink se rinde a la anestesia emocional.",
			Meaning: "La canción narra el diálogo entre un médico (Waters) y su paciente (Gilmour), donde el paciente se rinde a la anestesia emocional para escapar del dolor. Es una metáfora del aislamiento y la disociación que el protagonista experimenta tras acumular trauma y pérdida. 'I have become comfortably numb' representa el alivio peligroso de apagar los sentimientos para sobrevivir. El solo de Gilmour al final es la voz de la emoción que el protagonista ya no puede expresar con palabras. Es un retrato devastadoramente preciso de la depresión y la desconexión emocional.",
		},
		{
			ID:      "8",
			Title:   "Born to Run",
			Artist:  "Bruce Springsteen",
			History: "Publicada en 1975 en el álbum homónimo, fue el disco con el que Springsteen se jugó su carrera: Columbia amenazó con no renovarle el contrato si no producía un hit. La canción tardó seis meses en grabarse, con Springsteen reconstruyendo partes una y otra vez. El sonido fue deliberadamente diseñado para sonar como Phil Spector, con capas de guitarras, saxofón y un wall of sound épico. Springsteen y la E Street Band la tocaron en directo cada noche durante décadas. Es considerada uno de los mejores álbumes debut de la historia del rock.",
			Meaning: "La canción es una épica sobre el deseo de escapar de la mediocridad y encontrar algo más grande, aunque no sepas exactamente qué. 'Baby, we were born to run' no es solo sobre huir, sino sobre buscar significado y conexión en un mundo que parece cerrarse. Wendy, la destinataria de la letra, representa el amor como ancla y como motor del escape simultáneamente. Springsteen describió la canción como 'toda América condensada en cuatro minutos'. Es un himno a la juventud, la esperanza y la necesidad humana de trascender.",
		},
		{
			ID:      "9",
			Title:   "Highway to Hell",
			Artist:  "AC/DC",
			History: "Lanzada en 1979 como single del álbum homónimo, fue el último disco con Bon Scott antes de su muerte en 1980. El título fue propuesto por Scott para describir la gira interminable de la banda por carreteras americanas. La producción de John 'Mutt' Lange dio a AC/DC un sonido más pulido que los acercó al mainstream sin perder agresividad. Fue su primer álbum en entrar al top 10 del Billboard 200 en Estados Unidos. La canción se convirtió en el himno definitivo del hard rock de finales de los 70.",
			Meaning: "Bon Scott escribió la letra como un homenaje irónico y gozoso a la vida de músico de gira: los excesos, la carretera sin fin y el rechazo de las normas sociales. 'Highway to hell' no es una referencia literal al infierno sino a la vida sin restricciones que lleva la banda. Scott celebraba ese camino con orgullo, sin disculpas ni arrepentimientos. La canción es un rechazo del puritanismo y una afirmación del placer y la libertad. Trágicamente, Bon Scott murió solo meses después de su lanzamiento, dándole a la letra una resonancia póstuma inesperada.",
		},
		{
			ID:      "10",
			Title:   "Paranoid",
			Artist:  "Black Sabbath",
			History: "Grabada en 1970 en menos de tres minutos, fue añadida al álbum homónimo casi como relleno cuando el disco quedó corto. Tony Iommi compuso el riff en minutos y Ozzy Osbourne escribió la letra en el tiempo que tardó en grabarse la música. A pesar de su génesis accidental, se convirtió en el single más exitoso de la banda. El álbum 'Paranoid' es considerado el nacimiento del heavy metal como género. La sencillez brutal de la canción la hizo accesible y demoledora al mismo tiempo.",
			Meaning: "La letra describe de manera directa los síntomas de la depresión: incapacidad para disfrutar del presente, sensación de vacío y desconexión. Osbourne la escribió sin pretensiones artísticas, pero capturó con precisión brutal lo que significa sentirse atrapado en la propia mente. 'I tell you to enjoy life, I wish I could' es quizás el verso más honesto sobre la depresión en la historia del rock. La canción normaliza la conversación sobre salud mental décadas antes de que fuera socialmente aceptado hacerlo. Su brevedad y contundencia la hacen más impactante que cualquier balada elaborada.",
		},
		{
			ID:      "11",
			Title:   "Sailing",
			Artist:  "Armin van Buuren ft. Josh Cumbee",
			History: "Lanzada en 2015 en el álbum 'Embrace' de Armin van Buuren, es una de las canciones más emotivas de su discografía. Josh Cumbee aportó una voz etérea que complementó perfectamente la producción atmosférica de van Buuren. La canción fue interpretada en vivo en el festival Tomorrowland ante miles de personas, creando uno de los momentos más memorables del EDM. Su uso de cuerdas orquestales y progresiones lentas la diferencia del trance más enérgico habitual de van Buuren. Forma parte de la transformación artística del DJ holandés hacia un sonido más melódico y emotivo.",
			Meaning: "La letra usa la metáfora de navegar para describir el viaje emocional hacia el amor y la conexión: dejarse llevar por la corriente hacia algo más grande que uno mismo. 'Sailing' evoca la vulnerabilidad de abrirse a otra persona, de soltar el control y confiar en el viaje. Las imágenes del mar y el viento representan las fuerzas externas que nos moldean y los sentimientos que no podemos controlar. Es una canción sobre rendirse al amor no como debilidad sino como acto de valentía. La producción creciente de van Buuren refleja musicalmente ese crescendo emocional.",
		},
	}

	for _, song := range songs {
		if err := svc.CreateSong(ctx, song); err != nil {
			log.Printf("Error seeding song %s: %v", song.Title, err)
		}
	}
}
