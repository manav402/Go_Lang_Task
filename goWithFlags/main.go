package main

import (
	"flag"
	"fmt"
	"os"
)

type greetcmd struct {
	fs    *flag.FlagSet
	name  string
	color bool
}

type calccmd struct{
	fs *flag.FlagSet
	operation string
	num1 float64
	num2 float64
}
func newcalccmd() *calccmd{
	cc := &calccmd{
		fs : flag.NewFlagSet("calc",flag.ContinueOnError),
	}
	cc.fs.StringVar(&cc.operation,"o","+","specify the operations")
	cc.fs.Float64Var(&cc.num1,"n1",1.0,"get first number")
	cc.fs.Float64Var(&cc.num2,"n2",1.0,"get second number")
	return cc
}

func newgreetcmd() *greetcmd {
	gc := &greetcmd{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}
	gc.fs.StringVar(&gc.name, "name", "default", "greet user with name")
	gc.fs.BoolVar(&gc.color, "color", false, "specify the color output")
	return gc
}

func (gc *greetcmd) init(args []string) {
	gc.fs.Parse(args)
}

func (cc *calccmd) init(args []string){
	cc.fs.Parse(args)
}

func (gc *greetcmd) greet() {
	if gc.color {
		fmt.Println("\u001b[34m", "hello ", gc.name, "\u001b[0m")
	} else {
		fmt.Println("hello ", gc.name)
	}
}

func (gc *greetcmd) print(ans float64){
	if gc.color{
		fmt.Println("\u001b[34m", "ans :-", ans, "\u001b[0m")
	}else{
		fmt.Println("and :-",ans)
	}
}

func (cc *calccmd)calculate(args []string) {
	gc := newgreetcmd()
	gc.init(args[1:])
switch cc.operation {
	case "+":
		gc.print(cc.num1 + cc.num2)
	case "-":
		gc.print(cc.num1 - cc.num2)
	case "*":
		gc.print(cc.num1 * cc.num2)
	case "/":
		gc.print(cc.num1 / cc.num2)
	}
}

// create a command wich can print an string greet a user and calculate the equtions as well do colorize output
func run(args []string) {
	for _, v := range args {
		switch v {
		case "greet":
			gc := newgreetcmd()
			gc.init(args[1:])
			gc.greet()
		case "calc":
			cc := newcalccmd()
			cc.init(args[1:])
			cc.calculate(args)
		// default:
		// 	fmt.Println("command not supported")
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("subcommand required")
		return
	}
	run(os.Args[1:])
}

// func colorize(b bool,s string){
// 	if b{
// 		fmt.Println("\u001b[34m",s,"\u001b[0m")
// 	}else{
// 		fmt.Println("hello ",s)
// 	}
// }

// func main(){
// 	var name string
// 	var useColor bool
// 	flag.StringVar(&name,"name","john","this will print the name given by you")
// 	flag.BoolVar(&useColor,"color",false,"change the output color")
// 	flag.Parse()

// 	colorize(useColor,name)
// 	// fmt.Println("hello "+name)
// }
