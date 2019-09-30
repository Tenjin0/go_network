package main

type Values struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Values, reply *int) error {

}

func (t *Arith) Divide(args *Values, reply *int) error {

}
func ArithServer(port string) {

}
