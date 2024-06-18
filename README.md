# Curso Basico GoLang

Introducción

![https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg](https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg)

## Qué es Go?

Go, también conocido como GoLang, es un lenguaje de programación creado por Google. Combina la eficiencia en la ejecución de los lenguajes compilados con la facilidad de uso y eficiencia de los lenguajes interpretados. Go se destaca por su simplicidad, eficiencia y capacidad para manejar concurrencia, lo que lo hace ideal para aplicaciones de servidor y microservicios.

## Por que Aprenderlo?

- Gran velocidad de compilación
- Alto rendimiento para tareas pesadas
- Soporte nativo por concurrencia
- Bien pagado
- Obliga a implementar buenas practicas
- Buena Comunidad

## Instalacion

Instalar para windows desde aqui:

[All releases - The Go Programming Language](https://go.dev/dl/)

Podemos crear nuestra app desde cualquier dirección. 

Los paquetes se instalaran en C:/USER/go

Abrimos desde VSCode y instalamos la siguiente extensión:

Nombre: Go
ID: golang.Go
Descripción: Rich Go language support for Visual Studio Code
Versión: 0.41.4
Editor: Go Team at Google
Vínculo de VS Marketplace: [https://marketplace.visualstudio.com/items?itemName=golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Primera App

luego de crear una carpeta en el directorio q queramos. Creamos el archivo main.go.

```go
// Paquete: Contenedor de funciones y variables
package main

// Inclusión de dependencias (paquetes)
import "fmt"

// Función principal
func main() {
	// Código que será ejecutado
	fmt.Println("Hola mundo!")
}

```

corremos el código con ‘go run main.go’.

o con ‘go build main.go’ eso creara un archivo main. escribimos ./main y se ejecutara el programa.

muy exigente con las buenas practicas. da error: espacios en blanco, llaves mal posicionadas.

```go
//permite
func main() {
}
//no permite
func main()
{

}
//en el caso de los espacios la extencion los elimina
func main(){
/
/
	fmt.Println('');

}
```

## Variables, constantes y zero values

- Paso previo: inicializar git. Con git init en la terminal, ubicado en nuestra carpeta raiz.
- git add ‘nombre de la app o carpeta’. en mi caso git add clase1.(en la carpeta clase1 tengo main.go).

### Declarar Constantes

Una constante es una variable que nunca va a cambiar su valor en el tiempo. si se reasigna un valor dará error.

Recuerda que go es un lenguaje estáticamente tipado. Debes indicarle el tipo de dato y variables.

```go

package main

import "fmt"

func main() {
	//Declarar constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	//pi: 3.14
  //pi2: 3.1415
}

```

### Declarar Variables enteras

Hay ciertas diferencias. Si una variables no fue declarada previamente lo indicamos con ‘:’ antes del igual. Por ejemplo.

base := 12

```go
//declarada previamente
base = 12
//no declarada previamente
base := 12
```

El metodo mas común para declarar es con la palabra reservada var y su tipo de dato.

```go
//declaracion e inicializacion
var altura int = 14;
//solo declaracion
var area int
```

### Zero values

Valor por defecto que tienen las variables que no le asignamos un valor.

```go
//Zero values
var a int //variables enteras
var b float64 //variables con decimales
var c string //variables con textos
var d bool // variables true o false
/*
 Esto es lo que retorna
 0 0  false
*/
```

### Ejercicio

Calcular el area de un cuadrado. Sabiendo que su base mide 10.

ignorando la unidad de medida, calcule el area.

Sabemos que la base y la altura de un cuadrado son iguales. Por lo tanto la formula quedaria.

Area = base x base ;

```go
package main

import "fmt"

func main() {
	//Ejercicio Area cuadrado
	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area =", areaCuadrado)
}

```

## Operadores Arigmeticos

```go

x := 10
y := 50

//Suma
resultSuma := x + y
fmt.Println(resultSuma)//60

//Resta
resultResta := x - y
fmt.Println(resultResta)//-40

//Multiplicacion
resultMulti := x * y
fmt.Println(resultMulti)//500

//Division
resultDivision := y / x
fmt.Println(resultDivision)//5

//Resto
resultResto := y % x
fmt.Println(resultResto)

//Incrementar
x++
fmt.Println(x)

//Decrementar
x--
fmt.Println(x)

```

## Tipos de datos primitivos

### Números Enteros

- int = Depende del OS (32 o 64 bits)
- int8 = 8 bits = -128 a 127
- int16 = 16 bits =  -2^15 a 2^15 -1
- int32 = 32 bits = -2^31 a 2^31 -1
- int64 = 64 bits = -2^63 a 2^63 -1

### Números Enteros Positivos

- uint
- uint8
- uint16
- uint32
- uint64

### Números Decimales

- float32
- float64

### Textos y Booleanos

- string = “”
- bool = true o false

### Números Complejos

- Complex64 = Real e imaginario float32
- Complex128 = Real o imaginario float64
- Ejemplo: c:=10 + 8i

## Paquete fmt

### Println

Un print normal con salto de linea al final.

```go
func main() {
	helloMessage := "Hello"
	wordMessage := "Word"
	//Println
	fmt.Println(helloMesagge, wordMessage)//Hello Word
}
```

### Printf

A parte de imprimir agrega una funcion extra.

The verbs:

%s = valor string

%d = valor entero

%v = no se conoce el tipo de dato

\n = salto de linea

Supongamos que declaramos dos variables, una  llamada ‘nombre’ de tipo string, y otra llamada ‘edad’ de tipo number entero (int).

```go
func main(){
	nombre := "Pedro"
	edad := 23
	fmt.Printf("%s tiene %d años", nombre, edad)//Pedro tiene 23 años
}
```

Como ves en el mensaje colocamos donde iran las dos variables con %s y %d. Luego del mensaje, separado de una coma,  indicamos que variables usamos en orden.

### Sprintf

Genera un string pero no lo imprime, solo lo guarda en la variable donde se declare.

```go
func main(){
	nombre := "Pedro"
	edad := 23
	message := fmt.Sprintf("%s tiene %d años", nombre, edad)
	fmt.Println(message)//Pedro tiene 23 años
}
```

### Conocer tipo de dato

se usa el comando ‘%T’

```go
func main(){
	nombre := "Valentin"
	fmt.Printf("la variable nombre es de tipo %T", nombre)
	//la variable nombre es de tipo string
}
```

## Paquetes

Todo programa de Go se compone de paquetes.

Lo arrancamos corriendo en el package main.

Por convención, los nombres de los paquetes es el mismo que el ultimo elemento importado. Por ejemplo si importamos el paquete “math/rand” para usarlo utilizamos al atributo rand.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

//My favorite number is 6
```

### Exportar Nombres

En go, un nombre es exportado si empieza con mayúscula. Por ejemplo Pizza es un nombre exportado, asi como Pi, que este ultimo es exportado desde el paquete math.

pizza y pi como no empiezan con mayúscula no son nombres exportados de un paquete.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)//<-- empieza con minus
}
//./prog.go:9:19: undefined: math.pi
```

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)//<-- empieza con mayus
}
//3.141592653589793
```

## Funciones

Una función puede tener cero o mas argumentos.

En este ejemplo, la función ‘add’ toma dos parámetros de tipo int.

Date cuenta que el tipo d dato de la función va después de su nombre.

Ahora que es el tipo de dato d la función, ES EL TIPO DE DATO QUE VA A RETORNAR LA FUNCION.

La sintaxis es: func <nombre de la función>(<parámetros>)<tipo de dato de la función>{}

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
//55
```

### Functions continued

Cuando tenes dos o mas parámetros continuos con el mismo tipo de dato, podes omitir indicárselo a todos y solo ponérselo al ultimo parámetro.

En este ejemplo cambiaremos

x int, y int

Por

x, y int

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```

### Funcion que retorna mas de un resultado

Una funcion puede retornar cualquier numero de resultados.

La funcion swap retorna dos strings.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

como ves siempre se marca que tipo de dato que va a retornar y cuantos, entre parentesis.

func swap(x, y string) (string, string).

## Ciclos For, for while y for forever

### For

```go
//For condicional
for i := 0; i < 10; i++{
	fmt.Println(i);
}
/*
1
2
...
8
9
*/
```

### For while

```go
// For While
counter := 0
for counter < 10{
	fmt.Println(counter)
	counter++;
}
/*
1
2
...
8
9
*/
```

### For Forever

Itera hasta la eternidad

```go
//For forever
counterForever := 0
for{
	fmt.Println(counterForever)
	counterForever++
}
```

## Operadores lógicos y de comparación

Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son FALSO/FALSE.

Empecemos con los operadores de comparación:

### Operadores de comparación

Son aquellos que retornan TRUE o FALSE en caso de cumplirse o no una expresión. Son los siguientes:

- *valor1 == valor2*: Retorna TRUE si valor1 y valor2 son exactamente iguales.
- *valor1 != valor2*: Retorna TRUE si valor1 es diferente de valor2.
- *valor1 < valor2*: Retorna TRUE si valor1 es menor que valor2
- *valor1 > valor2*: Retorna TRUE si valor1 es mayor que valor2
- *valor1 >= valor2*: Retorna TRUE si valor1 es igual o mayor que valor2
- *valor1 <= valor2*: Retorna TRUE si valor1 es menor o igual que valor2.

### Operadores lógicos

Son aquellos que retorna TRUE o FALSE si cumplen o no una condición utilizando [**puertas lógicas**](https://platzi.com/clases/1050-programacion-basica/15968-que-son-tablas-de-verdad-y-compuertas-logicas/).

### Operador AND:

Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo `&&`.

Ejemplo1: `1>0 && 2 >0` Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

Ejemplo2: `2<0 && 1 > 0` Esto retornará FALSE porque una de las condiciones no es verdadera.

### Operador OR:

Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo `||`.

Ejemplo: `2<0 || 1 > 0` Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

### Operador NOT:

Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:

```go
myBool :=  true
fmt.Println(!myBool) // Esto retornará false
```

## El condicional if

Se utiliza para que dada una condicion se cumpla, se ejecute un codigo.

```go
func main(){
	valor1 := 1
	valor2 := 2
	
	if valor1 == 1 {
		fmt.Println("chi es igual")
	} else {
		fmt.Println("no es igual")
	}
	
	//With and - &&
	if valor1 == 1 && valor2 == 2{
		fmt.Println("Ambas condiciones se cumplen")
	}
	
	//Whith or - ||
	if valor1 == 34 || valor2 == 2{
		fmt.Println("Una condicion o ambas se cumplen")
	}
	
	//Covertir texto a numero	
	//dos retornos value(valor convertido) y err(error)	
	// si no hay error. err sera nil
	value, err := strconv.Atoi("54")		
	if err != nil{
			log.Fatal(err)
	}	
	fmt.Println("Value:", value)
	
	//value, err := strconv.Atoi("safds") -->strconv.Atoi: parsing "safds": invalid syntax
}
```

### Reto 1

Realizar una función que determine si un numero es par o impar.

```go
func esPar(x int) string {
	if x % 2 != o{
		return "es impar"
	}
	return "es par"
}
```

## Switch

Cuando estas ejecutando múltiples condiciones if, una tras otra, se vuelve difícil de leer. Para eso se usa switch y volver un codigo mas legible.

```go
func main(){
	modulo := 5 % 2
	
	switch modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
}
```

En caso de que modulo sea igual a 0. se ejecuta el primer case. El default se ejecuta si no se cumplen los case.

También pude encontrarse así. La variable es declarada después de la palabra reservada switch, separada por un ; de su parseo.

```go
func main(){
	
	
	switch modulo := 5 % 2; modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
}
```

### Switch sin condicion

```go
func main(){
		value := 200
		switch{
			case value > 100:
				fmt.Println("Es mayor a 100")
			case value < 0:
				fmt.Println("Es menor a 0")
			default:
				fmt.Println("Es mayor a 0 y menor a 100")
		}
	

		
}
```

## Keywords defer, break y continue

### Defer

Sintaxis:

defer (linea de codigo a ejecutar).

Con defer, la linea de codigo que viene despues, se ejecutara al final de TODO, antes que ‘muera’ el programa.

```go
func main(){
	//Defer
	defer fmt.Println("me ejecuto a lo ultimo de todo")
	fmt.Println("hola")
	fmt.Println("mundo")
	
}

/*
-Consola:
hola
mundo
me ejecuto a lo ultimo de todo
*/
```

se puede usar por ejemplo cuando abras una conexion con una base de datos, le colocas defer al cierre de esa conexion. Se usa para cerrar cosas.

### Continue y break

Se utilizan dentro del ciclo for

```go
func main(){
	for i := 0; i<10; i++ {
		fmt.Println(i)
		
		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		
		//Break
		
		if i == 7{
			fmt.Println("Es 7")
			break
		}
	}
}
/*
Consola:
0
1
2
Es 2
3
4
5
6
7
Break
*/
```

Básicamente el continue continua, y el break lo detiene(sale del for).

## Arrays y Slices

Los array en go tienen una longitud fija e invariable y deben declarase especificándola

var {nombre} [dimensión/longitud que tendrá]{tipo de dato}

var array [4]int

Ese array al no especificarle ni editarle sus elementos, empezara siendo [0 0 0 0]

```go
func main(){
	//Array
	var array [4]int
	fmt.Println(array) // --> [0 0 0 0]

}
```

```go
func main(){
	//Array
	var array [4]int
	
	//puedo editar inices que quiero
	array[0] = 1
	array[1] = 2
	fmt.Println(array)// --> [1 2 0 0]

}
```

### Metodo len y cap

Muy importantes en cuanto a manejo de arrays en go.

- len me dice cuantos elementos hay en un array.
- cap me indica cual es la capacidad máxima de ese array

```go
func main(){
	//Array
	var array [4]int
	
	//puedo editar inices que quiero
	array[0] = 1
	array[1] = 2
	fmt.Println(len(array), cap(array)) // --> 4 4

}
```

### Slice

No se le indica la longitud o cantidad de valores.

```go
//Slice
slice := []int{0, 1, 2, 3, 4, 5, 6}
fmt.Println(slice, len(slice), cap(slice)) // -->[0 1 2 3 4 5 6] 7 7
```

Investigando: La diferencia principal entre los arrays es que estos tienen una longitud fija e invariable y deben declarase especifiandola

```go
x := [5]int{0, 1 ,2, 3, 4}

```

mientras que los Slices tienen una longitud variable y no hay que especificarla en la declaración

```go
var x [ ]float64

```

en este caso se crea un Slice con una longitud de cero Si queremos crear un slice deberiamos usar la funcion make:

```go
x := make([]float64, 5)

```

esto crea un Slice asociado a un array subjacente de longitud 5. Los Slices siempre están asociados a un array y aunque nunca pueden ser mas largos que el aray, pueden ser mas cortos. La función make también permite un tercer parámetro, que representa la capacidad del array, por lo que

```go
x := make([]float64, 5, 10)

```

representa un Slice de longitud 5 y capacidad de 10

Arrays: Son inmutables, su length y capacity serán siempre las mismas.

Slices: Debido a que no son inmutables y su tamaño es dinámico, puedes tener un tamaño distinto a su capacidad, pero su capacidad siempre va a ser más grande que su tamaño.

Es decir, puedes tener un slice de tamaño 3 y capacidad 5, pero no uno de tamaño 5 y capacidad 3. Cuando usas append y te quedas sin capacidad, ya que los slices usan arrays para almacenar los datos, la función le va a asignar un nuevo array más grande y moverá los datos existentes más los nuevos a éste. El array siempre será del doble de tamaño que el anterior.

En otras palabras, digamos que tienes un slice de 5 elementos (capacidad y tamaño), y usas append para poner otro elemento más, el slice resultante va a tener tamaño 6 (por el nuevo elemento) y la capacidad subirá a 10 (el doble de la capacidad del array anterior).

Si ya estás algo familiarizado con programación, esto funciona porque un slice tiene 3 atributos, un pointer al array donde se almacena la info, el dato del tamaño y el dato de la capacidad. Cuando el slice crece, se crea otro array más grande, se copia la info vieja a ese nuevo array, y el pointer del slice cambia al array nuevo.

### Métodos en el slice

```go
fmt.Println(slice[0])//imprime el elemento 0
fmt.Println(slice[:3])//imprime hasta el elem. 3 sin incluirlo
fmt.Println(slice[2:4])//imprime del elem. 2 al 4 sin incluirlo
fmt.Println(slice[4:])//imprime desde el elem. 4
```

### Agregar elementos la slide

Con metodo append

agrega un elemento a lo ultimo.

```go
slice := []int{0, 1, 2, 3, 4, 5, 6}
slice = append(slice, 7)
fmt.Println(slice, len(slice), cap(slice))// --> [0 1 2 3 4 5 6 7] 8 14
```

Agregando otra lista

```go
newSlice := []int{8, 9, 10}
slice = append(slice, newSlice...)
fmt.Println(slice, len(slice), cap(slice))// --> [0 1 2 3 4 5 6 7 8 9 10] 11 14
```

### Recorriendo slice con range

```go
slice := []string{"hola","que","hace"}

for i , valor := range slice {
	fmt.Println(i, valor)
	/*
	0 hola
	1 que
	2 hace
	*/
}
```

## LLave valor con Mapas

En muchos lenguajes de programacion existen las estructuras ‘llave valor’ esto quiere decir que para acceder a un valor necesitas una llave.

Por ejemplo: Un estacionamiento, en el cual, cada puesto tiene un numero asignado y para poder acceder al auto que tengas ahi guardado, necesitas esa llave, que es el numero de estacionamiento, y el valor seria lo que esta alli seleccionado.

Es como un objeto en javascript.

Como crearlo:

```go
 // Crear un mapa
    mapa := make(map[string]string)

    // Agregar elementos al mapa
    mapa["clave1"] = "valor1"
    mapa["clave2"] = "valor2"
    mapa["clave3"] = "valor3"

    // Acceder a un valor usando una clave
    fmt.Println(mapa["clave2"]) // Imprimirá: valor2

    // Eliminar un elemento del mapa
    delete(mapa, "clave1")

    // Verificar si una clave existe en el mapa
    _, existe := mapa["clave1"]
    fmt.Println(existe) // Imprimirá: false

    // Iterar sobre los elementos del mapa
    for clave, valor := range mapa {
        fmt.Println(clave + " = " + valor)
    }
```

ejemplos

```go
temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	
	infoAlumno := map[string]interface{
		"nombre":            "Juan",
		"apellido":          "Perez",
		"carrera":           "Ingeniería Informática",
    "dni":               "12345678",
    "edad":              25,
    "materiasAprobadas": []string{"Matemáticas", "Programación", "Base de datos"},
		
	}
```

## Structs: La forma de hacer clases en Go

Los structs son tipos de datos compuestos que te permiten agrupar diferentes tipos de datos bajo un mismo nombre. Son similares a las estructuras en otros lenguajes de programación. Los structs son útiles cuando deseas modelar datos con múltiples atributos relacionados.

Para definir un struct en Go, se utiliza la palabra clave **`type`** seguida del nombre del struct y la lista de campos que lo componen, junto con sus tipos. Por ejemplo:

```go
type Persona struct {
    Nombre    string
    Edad      int
    Profesion string
}
```

1. **Crear una instancia de un struct**: Puedes crear una instancia de un struct utilizando la sintaxis de llaves y proporcionando valores para cada campo, como en el siguiente ejemplo:

```go
// Crear una instancia de Persona
persona1 := Persona{
    Nombre:    "Juan",
    Edad:      30,
    Profesion: "Ingeniero",
}
```

2. **Acceder a los campos de un struct**: Puedes acceder a los campos de un struct utilizando el operador de punto (**`.`**), como se muestra aquí:

```go
fmt.Println("Nombre:", persona1.Nombre)
fmt.Println("Edad:", persona1.Edad)
fmt.Println("Profesión:", persona1.Profesion)
```

3. **Modificar los campos de un struct**: Los campos de un struct pueden ser modificados directamente asignándoles nuevos valores, como en el siguiente ejemplo:

```go
// Modificar la edad de la persona
persona1.Edad = 31
```

4. **Embeber structs**: Un struct puede incluir otros structs como campos. Esto se conoce como embebido. Por ejemplo:

```go
type Contacto struct {
    Email   string
    Telefono string
}

type Persona struct {
    Nombre    string
    Edad      int
    Profesion string
    Contacto  Contacto // Embebiendo el struct Contacto
}
```

Los structs en Go son una parte fundamental del lenguaje y se utilizan ampliamente para modelar datos en programas Go. Proporcionan una forma conveniente y eficiente de organizar y manipular datos estructurados.

Ejemplo:

Imaginemos que hemos sido contratados por una empresa llamada "TechCorp" como desarrolladores backend utilizando Go. Una de las tareas que se nos ha asignado es crear un sistema de gestión de empleados. Para esto, utilizaremos structs para modelar la información de los empleados y realizar operaciones básicas como agregar nuevos empleados, actualizar información y generar informes.

```go
package main

import (
    "fmt"
)

// Definición del struct Empleado
type Empleado struct {
    ID        int
    Nombre    string
    Apellido  string
    Departamento string
    Salario   float64
}

// Función para agregar un nuevo empleado
func agregarEmpleado(id int, nombre, apellido, departamento string, salario float64) Empleado {
    nuevoEmpleado := Empleado{
        ID:        id,
        Nombre:    nombre,
        Apellido:  apellido,
        Departamento: departamento,
        Salario:   salario,
    }
    return nuevoEmpleado
}

// Función para actualizar el salario de un empleado
func actualizarSalario(empleado *Empleado, nuevoSalario float64) {
    empleado.Salario = nuevoSalario
}

func main() {
    // Crear un nuevo empleado
    empleado1 := agregarEmpleado(1, "Juan", "Perez", "Desarrollo", 50000.0)
    fmt.Println("Empleado creado:", empleado1)

    // Actualizar el salario del empleado
    actualizarSalario(&empleado1, 55000.0)
    fmt.Println("Salario actualizado:", empleado1.Salario)
}
```

En este ejemplo, hemos definido un struct **`Empleado`** con campos como **`ID`**, **`Nombre`**, **`Apellido`**, **`Departamento`** y **`Salario`**. Luego, hemos creado funciones para agregar un nuevo empleado y actualizar su salario utilizando métodos que reciben un puntero al struct **`Empleado`**. Esto nos permite modificar directamente el struct original.

Este escenario es solo un ejemplo de cómo podríamos utilizar structs en la vida laboral como desarrolladores backend en Go. Los structs nos permiten organizar la información de manera estructurada y realizar operaciones sobre ella de manera eficiente.

**Dato a tener en cuenta**

<aside>
💡 Por Que se uso este simbolo &?

En la línea **`actualizarSalario(&empleado1, 55000.0)`**, estamos pasando la dirección de memoria de la variable **`empleado1`** a la función **`actualizarSalario`**. Esto se hace utilizando el operador **`&`** antes del nombre de la variable (**`&empleado1`**).

En Go, cuando pasas una variable a una función, por defecto se pasa una copia de la variable (por valor). Sin embargo, cuando pasas un puntero a una variable, estás pasando la dirección de memoria donde se almacena esa variable (por referencia). **Esto significa que cualquier modificación realizada dentro de la función utilizando el puntero afectará directamente a la variable original fuera de la función.**

En el contexto de **`actualizarSalario`**, estamos modificando el salario del empleado, que es un campo dentro de la estructura **`Empleado`**. Para poder modificar directamente el valor del salario en la estructura original, necesitamos pasar un puntero a la estructura **`Empleado`**, de lo contrario, solo estaríamos modificando una copia de la estructura dentro de la función, y esos cambios no se reflejarían fuera de ella.

Por lo tanto, usamos **`&empleado1`** para pasar un puntero a la estructura **`Empleado`** a la función **`actualizarSalario`**, lo que nos permite actualizar directamente el salario dentro de la estructura original.

</aside>

## Structs y Punteros

Punteros: Acceso a la memoria, es decir, cuando guardamos una variable, se crea una dirección de memoria y a esa dirección se le guarda el valor de la variable.

En el siguiente ejemplo. `a` y `b` comparten dirección de memoria. Le damos la misma direccion de memoria de `a` a `b`. Por el hecho de declararla e inicializarla asi: **`b := &a`**  . Para acceder al valor al que esta apuntando `b` utilizamos un `*` adelante de la variable, en este caso, `*b` .

```go
func main(){
	a := 50
	b := &a
	
	fmt.Println(b) // --> 0xc00000a0a8
	fmt.Println(*b) // --> 50
	
}
```

Si modificamos el valor que esta apuntando a una direccion de memoria, las demas variables que esten apuntando a la misma direccion de memoria tmb cambiaran de valor.

Por ejemplo: modificaremos `*b` y veamos que pasa con a.

```go
*b = 100
fmt.Println(*b)// --> 100
fmt.Println(a)// --> 100
```

En Go, los punteros permiten referirse directamente a la memoria de una variable. Usar punteros en funciones permite que las funciones modifiquen las variables originales en lugar de trabajar con copias. A continuación, se muestran ejemplos de cómo implementar y utilizar punteros en funciones en Go.

### Ejemplo 1: Modificar el valor de una variable usando punteros

Supongamos que queremos crear una función que modifique el valor de una variable entera.

```go
package main

import "fmt"

// Funcion que recibe un puntero a un entero y modifica su valor
func incrementarValor(val *int) {
    *val += 10 // Incrementa el valor al que apunta el puntero por 10
}

func main() {
    x := 5
    fmt.Println("Antes de incrementar:", x)

    incrementarValor(&x) // Pasamos la direccion de memoria de x
    fmt.Println("Después de incrementar:", x)
}

```

### Explicación:

1. **Declaración de la función**:

```go
func incrementarValor(val *int) {
    *val += 10
}

```

La función `incrementarValor` toma un puntero a un entero (`*int`). Dentro de la función, usamos el operador de desreferencia (`*`) para acceder y modificar el valor al que apunta el puntero.

2. **Llamada a la función**:

```go
incrementarValor(&x)

```

Usamos el operador de dirección (`&`) para obtener la dirección de memoria de `x` y pasársela a la función. Esto permite que la función modifique directamente la variable `x`.

### Ejemplo 2: Intercambiar dos valores usando punteros

Este ejemplo muestra cómo intercambiar los valores de dos variables usando punteros.

```go
package main

import "fmt"

// Funcion que intercambia los valores de dos enteros usando punteros
func intercambiar(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x := 1
    y := 2
    fmt.Println("Antes del intercambio:", x, y)

    intercambiar(&x, &y) // Pasamos las direcciones de memoria de x e y
    fmt.Println("Después del intercambio:", x, y)
}

```

### Explicación:

1. **Declaración de la función**:

```go
func intercambiar(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

```

La función `intercambiar` toma dos punteros a enteros (`*int`). Intercambia los valores al que apuntan estos punteros usando una variable temporal `temp`.

2. **Llamada a la función**:

```go
intercambiar(&x, &y)

```

Usamos el operador de dirección (`&`) para obtener las direcciones de memoria de `x` e `y` y se las pasamos a la función. Esto permite que la función intercambie directamente los valores de `x` e `y`.

### Ejemplo 3: Modificar una estructura usando punteros

Supongamos que tenemos una estructura y queremos modificar sus campos usando punteros.

```go
package main

import "fmt"

// Definicion de la estructura Persona
type Persona struct {
    nombre string
    edad   int
}

// Funcion que modifica el nombre de una Persona
func cambiarNombre(p *Persona, nuevoNombre string) {
    p.nombre = nuevoNombre
}

func main() {
    persona := Persona{nombre: "Juan", edad: 30}
    fmt.Println("Antes de cambiar el nombre:", persona)

    cambiarNombre(&persona, "Carlos") // Pasamos la direccion de memoria de persona
    fmt.Println("Después de cambiar el nombre:", persona)
}

```

### Explicación:

1. **Declaración de la estructura y la función**:

```go
type Persona struct {
    nombre string
    edad   int
}

func cambiarNombre(p *Persona, nuevoNombre string) {
    p.nombre = nuevoNombre
}

```

La estructura `Persona` tiene dos campos: `nombre` y `edad`. La función `cambiarNombre` toma un puntero a una `Persona` y un nuevo nombre como argumentos, y cambia el nombre del objeto `Persona` al nuevo nombre proporcionado.

1. **Llamada a la función**:

```go
cambiarNombre(&persona, "Carlos")
```

Usamos el operador de dirección (`&`) para obtener la dirección de memoria de `persona` y se la pasamos a la función. Esto permite que la función modifique directamente el campo `nombre` de la estructura `persona`.

## Stringers: personalizar el output de Structs

En Go, la interfaz `fmt.Stringer` es utilizada para definir cómo se debe representar una estructura o un tipo personalizado como una cadena de caracteres. Esto es útil cuando quieres tener un control específico sobre cómo se formatea la salida al usar `fmt.Println` o funciones similares.

### Ejemplo Básico

Supongamos que tienes una estructura `Persona` y quieres definir cómo se debe mostrar la información de una persona.

```go
package main

import (
    "fmt"
)

// Definición de la estructura Persona
type Persona struct {
    Nombre string
    Edad   int
}

// Implementación del método String para la estructura Persona
func (p Persona) String() string {
    return fmt.Sprintf("%s (%d años)", p.Nombre, p.Edad)
}

func main() {
    // Crear una instancia de Persona
    persona := Persona{Nombre: "Juan", Edad: 30}
    
    // Usar fmt.Println que llamará automáticamente al método String
    fmt.Println(persona)
}

```

El método `String` de la estructura `Persona` usa `fmt.Sprintf` para formatear la salida como una cadena que muestra el nombre y la edad de la persona.

Al usar `fmt.Println(persona)`, se llama automáticamente al método `String` de la estructura `Persona`, y se imprime la cadena formateada: `Juan (30 años)`.

Output:

```go
Juan (30 años)
```

## Interfaces y listas de interfaces

En Go, una interfaz es un tipo abstracto que especifica un conjunto de métodos que deben ser implementados por otros tipos. Las interfaces son muy poderosas porque permiten escribir código que es más flexible y reutilizable.

### Ejemplo Avanzado

Supongamos que tenemos múltiples tipos de formas geométricas (círculo y rectángulo) y queremos calcular el área de cada una usando una interfaz.

```go
package main

import (
    "fmt"
    "math"
)

// Definición de la interfaz Forma
type Forma interface {
    Area() float64
}

// Definición de la estructura Círculo
type Circulo struct {
    Radio float64
}

// Implementación del método Area para Círculo
func (c Circulo) Area() float64 {
    return math.Pi * c.Radio * c.Radio
}

// Definición de la estructura Rectángulo
type Rectangulo struct {
    Ancho, Alto float64
}

// Implementación del método Area para Rectángulo
func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Alto
}

func main() {
    // Crear instancias de Círculo y Rectángulo
    circulo := Circulo{Radio: 5}
    rectangulo := Rectangulo{Ancho: 4, Alto: 6}

    // Crear una lista de Forma
    formas := []Forma{circulo, rectangulo}

    // Iterar sobre la lista y calcular el área de cada forma
    for _, forma := range formas {
        fmt.Printf("Área: %.2f\n", forma.Area())
    }
}

```

### Explicación:

1. **Definición de la Interfaz**:
    
    ```go
    type Forma interface {
        Area() float64
    }
    ```
    
    La interfaz `Forma` requiere un método `Area` que devuelva un valor de tipo `float64`.
    
2. **Implementación para `Circulo`**:
    
    ```go
    func (c Circulo) Area() float64 {
        return math.Pi * c.Radio * c.Radio
    }
    ```
    
    La estructura `Circulo` implementa el método `Area`.
    
3. **Implementación para `Rectangulo`**:
    
    ```go
    func (r Rectangulo) Area() float64 {
        return r.Ancho * r.Alto
    }
    ```
    
    La estructura `Rectangulo` también implementa el método `Area`.
    
4. **Lista de Interfaces**:
    
    ```go
    formas := []Forma{circulo, rectangulo}
    ```
    
    Se crea una lista de `Forma` que contiene tanto `Circulo` como `Rectangulo`.
    
5. **Iteración y Uso**:

```go
for _, forma := range formas {
    fmt.Printf("Área: %.2f\n", forma
    .Area())
}
```

Se itera sobre la lista y se llama al método `Area` en cada elemento.

## Concurrencia

La concurrencia en Go es una característica poderosa que permite que múltiples tareas se ejecuten de manera simultánea, aprovechando al máximo las capacidades de los procesadores multicore. Go ofrece soporte de concurrencia mediante goroutines y canales, facilitando la escritura de programas concurrentes de manera simple y eficiente.

### Diferencias entre concurrencia y paralelismo

La **concurrencia** es cuando dos o más tareas pueden empezar, ejecutarse y completarse en períodos de tiempos superpuestos. Esto no quiere decir que ambos procesos corran a la misma vez. Mientras que en el **paralelismo** las tareas se ejecutarán literalmente a la misma vez, esto solo es posible cuando tenemos más de un procesador de otra manera es imposible realizar tareas en paralelo y posiblemente hablemos de tareas ejecutándose **concurrentemente**.

### Conceptos Básicos de Concurrencia en Go

1. **Goroutines**: Una goroutine es una función que se ejecuta de manera concurrente con otras goroutines en el mismo espacio de direcciones. Son ligeras y gestionadas por el runtime de Go.
2. **Canales**: Los canales son utilizados para la comunicación segura entre goroutines. Permiten el paso de datos de una goroutine a otra de forma sincronizada.

### Ejemplo Básico

### Goroutines

Para iniciar una goroutine, se utiliza la palabra clave `go` seguida de la llamada a la función.

```go

package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello!")
}

func main() {
    // Iniciar una goroutine
    go sayHello()

    // Esperar un momento para permitir que la goroutine se ejecute
    time.Sleep(time.Second)
    fmt.Println("Done")
}

```

### Explicación:

1. **Función `sayHello`**: Esta función imprime "Hello!".
2. **Goroutine**: `go sayHello()` inicia una nueva goroutine que ejecuta `sayHello`.
3. **Esperar**: `time.Sleep(time.Second)` pausa el programa principal para permitir que la goroutine se complete antes de que el programa termine.

### Otro ejemplo

un simple ejemplo con los empleados de la Estrella de la Muerte, que recordemos si tuvo los problemas de seguridad que tuvo, es porque pasaban más tiempo durmiendo que trabajando

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyStormtrooper(i)
	}
	time.Sleep(4 * time.Second)
}

func sleepyStormtrooper(id int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("The Stormtrooper, number %d, is snoring\n", id)
}
```

Uno cree que el mensaje en la console sera:

"The Stormtrooper, number 0, is snoring

(espera 3s)

"The Stormtrooper, number 1, is snoring

…

Lo que va a pasar es que todos los mensajes se mandaran al mismo tiempo y desordenados, porque son independientes uno de otros.

## Usando canales con nuestras gorrutinas

Obviamente sabemos que la programación no siempre es tan sencilla, y tendremos que realizar programas más complicados con concurrencia, un caso muy común que nos encontraremos es el de comunicar dos **gorrutinas**, para esto en **Go** tenemos lo que se llaman `channels` o `canales`.

Los **canales pueden ser utilizados como variables** como cualquier otro tipo en **Go**, es decir podremos pasarlos a las funciones y ser propiedades de nuestros `structs`.

Para inicializar un canal haremos uso de la función `make` al igual que la utilizamos en los `maps` y `slices`. Además tendremos que especificar el tipo de datos que va a recibir el canal a la hora de crearlos.

```go
c := make(chan string)
```

Con nuestro canal creado ya estamos listo para mandar o recibir mensajes, para ello utilizaremos el operador `<-`, veamos como funciona:

Si queremos mandar datos a nuestro canal:

```go
c <- "Hello, Darth Vader"
```

Y si queremos recibirlos:

```go
m := <- c
```

Algo que tenéis que tener muy presente es que las operaciones con canales son bloqueantes, así que mientras un canal este abierto y haya alguna parte de nuestro código esperando recibir valores, no se ejecutará nada por debajo o en el caso de una **gorrutina** la mantendrá bloqueada.

Vamos a verlo mejor con nuestro ejemplo anterior, hagamos que los Stormtroopers se despierten y se pongan a trabajar.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyStormtrooper(i, c)
	}

	for i := 0; i < 5; i++ {
		stormtrooper := <-c
		fmt.Printf("The Stormtrooper, number %d, is awake!\n", stormtrooper)
	}
}

func sleepyStormtrooper(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("The StormTrooper, number %d, is snoring\n", id)
	c <- id
}

```

### Canales

Los canales permiten la comunicación y sincronización entre goroutines.

```go
package main

import (
    "fmt"
)

func sum(a, b int, result chan int) {
    result <- a + b
}

func main() {
    // Crear un canal
    result := make(chan int)

    // Iniciar una goroutine
    go sum(3, 4, result)

    // Recibir el resultado del canal
    res := <-result
    fmt.Println("Sum:", res)
}

```

### Explicación:

1. **Función `sum`**: Esta función calcula la suma de dos enteros y envía el resultado a través del canal `result`.
2. **Crear un canal**: `result := make(chan int)` crea un canal para enteros.
3. **Goroutine**: `go sum(3, 4, result)` inicia una goroutine que ejecuta `sum`.
4. **Recibir del canal**: `res := <-result` recibe el valor enviado a través del canal `result` y lo almacena en `res`.

### Ejemplo Avanzado

### Uso de `sync.WaitGroup` para sincronización

El paquete `sync` proporciona un tipo `WaitGroup` que se puede utilizar para esperar a que un conjunto de goroutines finalicen.

```go

package main

import (
    "fmt"
    "sync"
)

func printNumber(number int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println(number)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go printNumber(i, &wg)
    }

    wg.Wait()
    fmt.Println("All goroutines finished")
}

```

### Explicación:

1. **Función `printNumber`**: Esta función imprime un número y llama a `wg.Done()` cuando termina.
2. **`sync.WaitGroup`**: `var wg sync.WaitGroup` declara un `WaitGroup`.
3. **Iniciar goroutines**: En un bucle, `wg.Add(1)` incrementa el contador del `WaitGroup`, y `go printNumber(i, &wg)` inicia una goroutine.
4. **Esperar a que terminen las goroutines**: `wg.Wait()` bloquea hasta que el contador del `WaitGroup` vuelva a cero, lo que indica que todas las goroutines han llamado a `wg.Done()`.

### Ejemplo Completo: Uso de Goroutines y Canales

Vamos a crear un programa que calcule las sumas de dos listas de números en paralelo y luego combine los resultados.

```go

package main

import (
    "fmt"
    "sync"
)

func sum(numbers []int, result chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    total := 0
    for _, number := range numbers {
        total += number
    }
    result <- total
}

func main() {
    numbers1 := []int{1, 2, 3, 4, 5}
    numbers2 := []int{6, 7, 8, 9, 10}

    result1 := make(chan int)
    result2 := make(chan int)

    var wg sync.WaitGroup
    wg.Add(2)

    go sum(numbers1, result1, &wg)
    go sum(numbers2, result2, &wg)

    go func() {
        wg.Wait()
        close(result1)
        close(result2)
    }()

    total1 := <-result1
    total2 := <-result2

    fmt.Printf("Total1: %d, Total2: %d, Combined: %d\n", total1, total2, total1+total2)
}

```

### Explicación:

1. **Función `sum`**: Calcula la suma de una lista de números y envía el resultado a través de un canal.
2. **Listas de números**: `numbers1` y `numbers2` son las listas que queremos sumar.
3. **Canales**: `result1` y `result2` son canales para recibir los resultados.
4. **`sync.WaitGroup`**: Se usa para esperar a que ambas goroutines terminen.
5. **Goroutines**: `go sum(numbers1, result1, &wg)` y `go sum(numbers2, result2, &wg)` inician las goroutines.
6. **Esperar y cerrar canales**: Una goroutine anónima espera a que todas las goroutines terminen y luego cierra los canales.
7. **Recibir resultados y combinarlos**: `total1 := <-result1` y `total2 := <-result2` reciben los resultados y se combinan.

## Range, Close y Select en channels

En Go, las palabras clave `range`, `close` y `select` tienen usos específicos, principalmente relacionados con el manejo de canales (channels) y estructuras de datos iterables. Aquí te explico cada una:

### `range`

La palabra clave `range` se utiliza para iterar sobre elementos en una variedad de estructuras de datos, como arrays, slices, mapas y cadenas (strings). También se puede usar para iterar sobre valores recibidos de un canal.

### Ejemplos:

**Iterar sobre un canal:**

```go

ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

for n := range ch {
    fmt.Println(n)
}

```

En este ejemplo, `range` itera sobre los valores recibidos del canal hasta que el canal está cerrado.

### `close`

La palabra clave `close` se usa para cerrar un canal, indicando que no se enviarán más valores a ese canal. Esto es útil para evitar el bloqueo de receptores que estén esperando valores.

### Ejemplo:

```go

ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()

for n := range ch {
    fmt.Println(n)
}

```

En este caso, el canal se cierra después de enviar los valores 1, 2 y 3, permitiendo que el bucle `range` termine correctamente.

### `select`

La palabra clave `select` se usa para esperar a que uno de varios canales esté listo para operaciones de envío o recepción. Permite trabajar de manera eficiente con múltiples canales al mismo tiempo.

### Ejemplo:

```go

ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    ch1 <- "hello from ch1"
}()

go func() {
    ch2 <- "hello from ch2"
}()

select {
case msg1 := <-ch1:
    fmt.Println("Received", msg1)
case msg2 := <-ch2:
    fmt.Println("Received", msg2)
}

```

En este ejemplo, `select` esperará hasta que uno de los canales (`ch1` o `ch2`) tenga un valor listo para ser recibido. Solo uno de los `case` se ejecutará, dependiendo de cuál canal esté listo primero.

### Resumen

- `range`: Se utiliza para iterar sobre arrays, slices, mapas, cadenas y canales.
- `close`: Se utiliza para cerrar un canal, indicando que no se enviarán más valores.
- `select`: Se utiliza para esperar en múltiples operaciones de canal, permitiendo manejar canales concurrentemente.

### Ejemplo avanzado usando range, close y select

programa de Go que simula el procesamiento de tareas concurrentes. Imagina que tienes un conjunto de tareas que se deben procesar y quieres distribuir el trabajo entre varios trabajadores concurrentes.

### Ejemplo: Procesamiento Concurrente de Tareas

```go

package main

import (
    "fmt"
    "sync"
    "time"
)

// Función que representa una tarea a procesar
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
        fmt.Printf("Worker %d processing task %d\n", id, task)
        time.Sleep(time.Second) // Simula el tiempo de procesamiento
        results <- task * 2     // Simula el resultado del procesamiento
    }
}

func main() {
    const numWorkers = 3
    const numTasks = 5

    tasks := make(chan int, numTasks)
    results := make(chan int, numTasks)
    var wg sync.WaitGroup

    // Iniciar los trabajadores
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, results, &wg)
    }

    // Enviar tareas a los trabajadores
    for i := 1; i <= numTasks; i++ {
        tasks <- i
    }
    close(tasks) // Cerrar el canal de tareas para indicar que no hay más tareas

    // Esperar a que todos los trabajadores terminen
    go func() {
        wg.Wait()
        close(results) // Cerrar el canal de resultados cuando todos los trabajadores hayan terminado
    }()

    // Recibir y mostrar los resultados
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

```

### Descripción del Código

1. **Definición del Trabajador (`worker`)**:
    - La función `worker` toma un ID de trabajador, un canal de tareas (`tasks`) y un canal de resultados (`results`), y un `WaitGroup`.
    - Cada trabajador procesa tareas del canal de tareas hasta que este se cierra.
    - Simula el procesamiento de cada tarea con un `time.Sleep` y luego envía el resultado al canal de resultados.
2. **Función Principal (`main`)**:
    - Se definen el número de trabajadores (`numWorkers`) y el número de tareas (`numTasks`).
    - Se crean los canales de tareas y resultados con capacidad suficiente para manejar todas las tareas.
    - Se inicia un grupo de trabajadores en goroutines, cada uno ejecutando la función `worker`.
    - Se envían todas las tareas al canal de tareas y luego se cierra el canal para indicar que no hay más tareas.
    - Se inicia una goroutine que espera a que todos los trabajadores terminen usando el `WaitGroup` y luego cierra el canal de resultados.
    - Se reciben y muestran los resultados del procesamiento de tareas hasta que el canal de resultados se cierra.

## Go get: El manejador de paquetes

En Go, `go get` es una herramienta integral del sistema de gestión de paquetes y dependencias. Permite descargar, instalar y gestionar paquetes y módulos de Go desde repositorios remotos. Aquí te explico cómo funciona y cómo se utiliza.

### ¿Qué es `go get`?

`go get` es un comando que se utiliza para:

1. **Descargar paquetes y módulos**: Baja el código fuente de paquetes y módulos desde repositorios remotos como GitHub, Bitbucket, entre otros.
2. **Instalar paquetes**: Compila e instala paquetes en tu espacio de trabajo (`GOPATH`) o en el módulo actual (en el caso de Go Modules).
3. **Actualizar dependencias**: Actualiza los paquetes y módulos ya descargados a sus versiones más recientes o especificadas.

### Uso de `go get`

### Descargar e instalar un paquete

```bash
go get example.com/user/package

```

Este comando descarga el paquete `package` del repositorio en `example.com/user/` y lo instala en el espacio de trabajo.

### Especificar una versión (con Go Modules)

Con Go Modules habilitado, puedes especificar una versión particular de un módulo:

```bash

go get example.com/user/package@v1.2.3

```

Esto descarga e instala la versión `v1.2.3` del módulo.

### Actualizar un paquete

Para actualizar un paquete a su última versión compatible:

```bash

go get -u example.com/user/package

```

Para actualizar a la última versión principal (puede incluir cambios incompatibles):

```bash

go get -u=example.com/user/package

```

### Go Modules y `go get`

Desde Go 1.11, se introdujo el soporte para Go Modules, un sistema de gestión de dependencias que reemplaza al antiguo `GOPATH`. Con Go Modules, el comportamiento de `go get` se ajusta a los módulos.

### Inicializar un módulo

Para empezar a utilizar Go Modules en un proyecto, inicializa el módulo:

```bash

go mod init example.com/user/project

```

Esto crea un archivo `go.mod` en el directorio del proyecto, que mantiene un registro de las dependencias del proyecto.

### Añadir una nueva dependencia

Al usar `go get` en un proyecto con módulos, automáticamente se añade la dependencia al archivo `go.mod`.

```bash

go get example.com/user/package

```

El comando añade la entrada correspondiente en `go.mod` y descarga la dependencia en el caché del módulo.

### Actualizar todas las dependencias

Para actualizar todas las dependencias del proyecto a sus últimas versiones compatibles:

```bash

go get -u ./...

```

### Ejemplo completo

A continuación, un ejemplo de cómo crear un proyecto, inicializar módulos, añadir dependencias y usar `go get`.

1. **Crear el proyecto y inicializar Go Modules**:

```bash

mkdir myproject
cd myproject
go mod init github.com/yourusername/myproject

```

1. **Añadir una dependencia**:

```bash

go get github.com/gorilla/mux

```

Este comando añade `github.com/gorilla/mux` como dependencia en `go.mod`.

1. **Actualizar una dependencia**:

```bash

go get -u github.com/gorilla/mux

```

### Otros comandos relacionados

- **Listar dependencias**:
    
    ```bash
    
    go list -m all
    
    ```
    
    Muestra todas las dependencias del proyecto.
    
- **Descargar todas las dependencias**:
    
    ```bash
    
    go mod download
    
    ```
    
    Descarga todas las dependencias especificadas en `go.mod`.
    
- **Verificar la integridad de las dependencias**:
    
    ```bash
    
    go mod verify
    
    ```
    
    Verifica que las dependencias descargadas coinciden con los hashes en `go.sum`.
