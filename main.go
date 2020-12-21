package main
 
import(
	"fmt"
	"bufio"
	"os"
)



type Synape struct{

	data int
	connect *Synape

}
type Axon struct{
	synapes *Synape
	polarizeStatus bool
	cellBody *Axon
}
type Dendrites struct{
	a *Axon
	branch *Dendrites
}

func startAnn()int{
	fmt.Printf("Enter 0 (Z) or 1 (O)")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		key := input.Text()
		if key == "Z" || key == "z"{
			return 0
		}else if key == "o" || key == "O"{
			return 1
		}else{
			return -1
		}
	}
	return -1
}

func addAnnDendrites(x int, d *Dendrites)*Dendrites{
	if d == nil{
		sx := &Synape{data:x,connect:nil}
		ax := &Axon{synapes:sx,polarizeStatus:false,cellBody:nil}
		dx := &Dendrites{a:ax,branch:nil}
		return dx
	}else{
		sx := &Synape{data:x,connect:nil}
		ax := &Axon{synapes:sx,polarizeStatus:true,cellBody:nil}
		dx := &Dendrites{a:ax,branch:d}
		return dx
	}
	return nil
}


func main(){
	var epoche *Dendrites = nil
	fmt.Printf("Hello! Do you run?")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		key := input.Text()
		if key == "y" || key== "Y" {
		     			x := startAnn()
		    		 	epoche = addAnnDendrites(x, epoche)
		     			fmt.Printf("connection build---\n%g",epoche)	
		}else if key == "n" || key == "N"{
			os.Exit(0)
		}else{
			fmt.Printf("Press (Q)uit?")
			input := bufio.NewScanner(os.Stdin)
			for input.Scan(){
				key := input.Text()
				if key == "q" || key== "Q" {
					os.Exit(0)
				}else{
					fmt.Print("wrong key")
				}
			}
		}
	}	
}
