//? variable declarations -------------------

var message string = "Hello, world!"
var message2 = "Hello, world!"

var a, b, c, d int = 0, 1, 2, 3
var e, f, g, k = 0, true, 2.0, "hello"

u := 44
t := "hello"
v, n := 0, true

var myFloat32 float32 = 44.
myComplex := complex(3, 4)

var (
	valua  = "Hello"
	valua2 = 2
)

var a [2]string
a[0] = "Hello"
a[1] = "World!"

primes := [6]int{1, 2, 3, 4, 5, 6}
primes := [...]int{1, 2, 3, 4, 5, 6}

var colors []string{"blue", "green", "red", "yellow"}

colors. = append(colors,"blue")




//? ************************************

//? Map ------------------------------------------

type Vertex struct{
	Lat, Long float32
}
m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
	40.56,-74.55
}


myMap := make(map[int]string)

myMap[26] = "Eskişehir"
myMap[34] = "İstanbul"

//? ************************************


//? Enumerator --------------------------------

type Brand string

const (
	FACEBOOK Brand = "FACEBOOK"
	MICROPIC Brand = "MICROPIC"
	GOOGLE   Brand = "GOOGLE"
	DIJIS    Brand = "DIJIS"
)

fmt.Println(FACEBOOK)
//? ************************************

//? Print Operations --------------------------------

fmt.Println("Str", value)
fmt.Println("StrValue:%v\n")
fmt.Println("val1=%T, val2=%T, val3:%T", val1, val2, val3)
//? ************************************



//? Functions --------------------------------
//! (x,y int) == (x int, y int)
func add(x,y int)  {
	
}

//! Ownership
func sayHello(message *string)  {
	*message = "Hi hello"


}
message := "Hi"
sayHello(&message)


// * Go nun sonuç değişkenleri aynı bir değişken gibi adlandırılır

func split(sum int)  (x, y int){
	x = sum *4 / 9
	y sum - x
	
	// !return == return x,y
	return
	
}

// * Sınırsız parametre alan fonksiyonlar

func sayHelloTwo(message ...string) {
	for _, message := range message{
		Println(message)
	}
	
}

//? ************************************


