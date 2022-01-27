package main
import "fmt"

func main(){
	var n,u,i,n2 int
	fmt.Printf("Panjang Sisi : ")
	fmt.Scanf("%d",&n)
	n2 = 0
	for i=1;i<=n*2-1;i++ {
		for u=1;u<=n*2+4;u++{
			if u > (n*2)/2+2-n2 && u <= (n*2)/2+2+n2 {
				if(i==n){
					fmt.Printf("-")
				} else {
					fmt.Printf(" ")
				}
			} else {
				fmt.Printf("*")
			}
		}
		if i<n {n2+=1} else {n2-=1}
		fmt.Println()
	}
}