package main

func main() {

}

// ?Interface
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// ? Base struct
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float32
}

type SpecialProdaction struct {
	Special bool
}

// ? Nesne {Composition}
type Ferrari struct {
	Car
	SpecialProdaction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + x.Model

	return ret
}
