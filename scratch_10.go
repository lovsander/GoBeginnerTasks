package main
import "fmt"

type VarPointer struct {
	val interface{}
	pointer *interface{}
}
// constructor
func NewVarPointer(val interface{}) *VarPointer {
	vp:=new(VarPointer)
	vp.val=val
	vp.pointer = &vp.val
	return vp
}
// check consistency
func (vp *VarPointer) CheckVal() bool {
	return vp.val==*vp.pointer
}

func printInit(args...*VarPointer) {
	for _, arg:= range args{
		fmt.Printf("%T , %T ,%v \n", arg.val, arg.pointer, arg.pointer)
	}
	fmt.Println()
}

func printEnd(args...*VarPointer) {
	for _, arg:= range args{
		fmt.Printf("%v, %v, %v \n", arg.val, *arg.pointer, arg.CheckVal())
	}
	fmt.Println()
}

func main() {
	a := NewVarPointer(5)
	b := NewVarPointer(6)
	c := NewVarPointer(8)
	printInit(a,b,c)
	a.pointer = b.pointer
	b.val = c.val
	printEnd(a,b,c)
}
