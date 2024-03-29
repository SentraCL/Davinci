package util

import "math/rand"

//DavinciTales , tiene puras frases para rellenar la espera del login
type DavinciTales struct {
	says []string
}

//DiLoTuyo , retorna alguna frase dicha por DaVinci
func (d *DavinciTales) DiLoTuyo() string {
	frases := []string{"Admiro aquellos que pueden sonreír en la adversidad.",
		"Estudiar sin deseo arruina la memoria, y no retiene nada de lo que absorbe.",
		"El conocimiento de todas las cosas es posible.",
		"La vida sin amor, no es vida en absoluto.",
		"Cuando no puedes hacer lo que quieres, quieres lo que puedes hacer.",
		"Cuanto más profundo es el sentimiento, mayor es el dolor.",
		"Una pequeña verdad es mejor que una gran mentira.",
		"La sabiduría es la hija de la experiencia.",
		"Cualquiera que defienda un argumento apelando a la autoridad, no está usando su inteligencia, solo está usando su memoria.",
		"Es un alumno pobre el que no va más allá de su maestro.",
		"Fija tu rumbo a una estrella y podrás navegar a través de cualquier tormenta.",
		"He ofendido a Dios y a la humanidad porque mi trabajo no alcanzó la calidad que debería tener.",
		"Aunque la naturaleza comienza con la razón y termina en la experiencia, es necesario que hagamos lo contrario, es decir, comenzar con la experiencia y, a partir de ahí, investigamos la razón.",
		"Nuestra vida está hecha de la muerte de otros.",
		"El que ama la práctica sin la teoría es como el marinero que sube a bordo sin timón ni brújula y nunca sabe dónde acabará.",
		"El matrimonio es como meter la mano en una bolsa de serpientes con la esperanza de sacar una anguila.",
		"Igual que un reino dividido cae, la mente dividida entre muchos temas se confunde y se destruye a sí misma.",
		"El ingenio humano nunca imaginará una invención más hermosa, más simple o más directa que la naturaleza, porque en sus inventos no falta nada, y nada es superfluo.",
		"Antes existirá un cuerpo sin sombra que la virtud sin la envidia.",
		"La experiencia no se equivoca, solo tus juicios se equivocan al esperar de ella lo que no está en su poder.",
		"Las mentiras no resuelven los problemas, solo los empeoran, así que los mentirosos vayan con cuidado.",
		"Es un gran error hablar bien de un hombre inútil, tanto como hablar mal de un buen hombre.",
		"Los hombres luchan en guerras y destruyen todo lo que los rodea. La tierra debería abrirse y tragárselos. El que no valora la vida no se la merece. Nunca destruyas otra vida mediante la ira o la maldad.",
		"Es un hecho reconocido que percibimos errores en el trabajo de otros más fácilmente que en el nuestro.",
		"La mayoría de las personas, si les das un libro, lo husmean un rato y luego intentan comérselo.",
		"Verdaderamente el hombre es el rey de las bestias, porque su brutalidad las supera.",
		"Después de haber recorrido una distancia entre rocas sombrías, llegué a la entrada de una gran caverna. Dos emociones contrarias surgieron en mí: miedo y deseo. Miedo a la amenazante caverna y deseo de ver si había cosas maravillosas en ella.",
		"Quien desee ser rico en un día, será ahorcado en un año.",
		"Llegará un momento en que los hombres consideren el asesinato de animales como ahora ven el asesinato de hombres.",
		"Nada debería ser tan temido como la fama vacía.",
		"El aire se mueve como un río y lleva las nubes con él igual que el agua que corre se lleva todas las cosas que flotan sobre ella.",
		"Debemos dudar de la certeza de todo lo que pasa por los sentidos, pero cuánto más debemos dudar es de las cosas contrarias a los sentidos, como la existencia de Dios y el alma.",
		"Reprueba a tu amigo en secreto y ensálzalo en público.",
		"Si el pintor desea ver bellezas que le encanten, le corresponde a él crearlas, y si desea ver monstruos que espantosos, ridículos o realmente lastimosos, él és el amo y señor de ellos.",
		"Una ola nunca se encuentra sola, sino que se mezcla con las otras olas.",
		"El aprendizaje es lo único que la mente nunca agota, nunca teme y nunca se arrepiente.",
		"Desperté solo para descubrir que el resto del mundo todavía estaba dormido.",
		"El peor mal que se puede hacer un artista es creer que su obra debería parecer buena a sus propios ojos.",
		"Una mentira es algo tan vil que incluso si hablara con justicia de Dios, le quitaría un poco de su divinidad, y tan excelente es la verdad que si elogia las cosas más humildes, son exaltadas.",
		"Ningún consejo es más digno de confianza que el que se da a los barcos que están en peligro.",
		"Una vez que hayas probado a volar, siempre caminarás por la tierra con los ojos mirando al cielo, porque allí has ​​estado, y allí siempre anhelarás regresar.",
		"El aprendizaje nunca agota la mente.",
		"El que realmente sabe no tiene necesidad de gritar.",
		"El que no se opone al mal, ordena que se haga.",
		"La adquisición del conocimiento siempre es útil para el intelecto, ya que puede expulsar cosas inútiles y retener lo bueno.",
		"Debes obrar como el pobre que llega al final de la feria y no puede encontrar otra forma de proveerse a sí mismo más que tomando todas las cosas que ya han visto otros compradores, y han sino rechazadas por su menor valor.",
		"Todo nuestro conocimiento tiene su origen en nuestras percepciones.",
		"El pintor producirá imágenes de poco mérito si toma las obras de otros como su estándar",
		"Pensé que estaba aprendiendo a vivir; Solo estaba aprendiendo a morir.",
		"Igual que un día de provecho trae un sueño feliz, una vida de provecho trae una muerte feliz.",
		"Mi cuerpo no será una tumba para otras criaturas.",
		"Cuando la higuera se quedó sin frutos, nadie la miró. Deseando que la producción de sus frutos fuera alabada por los hombres, fue doblada y rota por ellos.",
		"La ignorancia ciega nos engaña. ¡Oh Desdichados mortales, abrid los ojos!",
		"En los ríos, el agua que tocas es la última de la que ya ha pasado y la primera de la que viene, igual que el presente.",
		"Es mejor imitar el trabajo antiguo que el moderno.",
		"Dios nos vende todas las cosas al precio de la mano de obra.",
		"El hierro se oxida debido al desuso, el agua estancada pierde su pureza, en el clima frío se congela, y la inacción mina los ánimos de la mente.",
		"El conocimiento no es suficiente; debemos poner en práctica.",
		"Un pintor debe comenzar cada lienzo con un baño de negro, porque todas las cosas en la naturaleza son oscuras, excepto cuando están expuestas por la luz.",
		"El mayor engaño que sufren los hombres proviene de sus propias opiniones.",
		"El que piensa poco se equivoca mucho.",
		"Los detalles hacen la perfección, y la perfección no es un detalle.",
		"La percepción actúa más fácilmente que el juicio.",
		"Y tú que deseas representar por medio de palabras la forma del hombre y todos los aspectos de su membrificación, renuncia a esa idea. Porque cuanto más minuciosamente describas, más limitarás la mente del lector, y más lo mantendrás alejado del conocimiento de lo descrito. Y entonces es necesario dibujar y describir.",
		"¿Por qué el ojo ve más claramente cuando estamos dormidos que la imaginación cuando estamos despiertos?",
		"El que está fijo en una estrella, no cambia de opinión.",
		"Aquellos que, en el debate, apelan a sus calificaciones, argumentan desde la memoria, no desde la comprensión.",
		"El que más posee más teme por la pérdida.",
		"El arte es la reina de todas las ciencias, comunicando el conocimiento a todas las generaciones del mundo.",
		"Hace tiempo que me llamó la atención que las personas con logros rara vez se acomodan y esperan que las cosas les sucedan. Salieron y les sucedieron las cosas.",
		"Nada fortalece la autoridad tanto como el silencio.",
		"Date cuenta de que todo se conecta con todo lo demás.",
		"Un humano promedio mira sin ver, escucha sin oír, toca sin sentir, come sin probar, se mueve sin conciencia física, respira sin percibir olor o fragancia y habla sin pensar.",
		"El tiempo es suficientemente extenso para aquellos que lo usan.",
		"Uno no tiene derecho a amar u odiar algo si no ha adquirido un conocimiento profundo de su naturaleza. El gran amor surge del gran conocimiento del objeto amado, y si lo conoces muy poco podrás amarlo solo un poco o nada.",
		"El pintor tiene el Universo en su mente y en sus manos.",
		"Los obstáculos no pueden aplastarme, cada obstáculo cede a una determinación firme.",
		"Un poeta sabe que ha alcanzado la perfección, no cuando no le queda nada por añadir, sino cuando ya no le queda nada por quitar.",
		"Haz que tu trabajo se ajuste a tu propósito.",
		"El artista ve lo que otros solo intuyen.",
		"Un hermoso cuerpo perece, pero una obra de arte no muere.",
		"La paciencia sirve como protección contra los errores igual que lo hace la ropa frente el frío. Porque si te pones más ropa a medida que aumenta el frío, este no tendrá poder para lastimarte. Entonces, de igual forma, debes tener paciencia cuando te encuentres con grandes errores, y entonces serán incapaces de irritar tu mente.",
		"El arte nunca termina, solo se abandona.",
		"Una vez que hayas probado el sabor del cielo, siempre mirarás hacia arriba.",
		"El agua es la fuerza motriz de la naturaleza.",
		"De vez en cuando, aléjate, relájate un poco, porque cuando vuelvas a tu trabajo tu juicio estará más seguro. Vete a cierta distancia porque entonces el trabajo parece más pequeño y se puede ver más de un vistazo y se puede ver más fácilmente la falta de armonía y proporción.",
		"La pasión intelectual expulsa la sensualidad.",
		"Nada se puede amar u odiar a menos que se entienda por primera vez."}
	total := len(frases)
	return frases[rand.Intn(total)]
}
