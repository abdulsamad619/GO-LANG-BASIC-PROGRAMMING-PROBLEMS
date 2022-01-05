package main
import ("fmt")

func main() {
	Evenorodd()
	Summingarray()
	TakingInputArray()
}
func Evenorodd(){
	fmt.Println("enter number")
	var a int
	fmt.Scanln(&a)

	if (a%2==0){
		fmt.Println("number is even")
	}else{
		fmt.Println("number is odd")
	}
}
func Summingarray(){
	arr:=[]int{2,2,2,2}
	sum:=0
	for i:=0;i<(len(arr));i++{
		sum=sum+arr[i]
	}
	fmt.Println("sum of an array is ",sum)
}
func TakingInputArray(){
	fmt.Println("enter size of an array")
	size:=0
	fmt.Scanln(&size)
	var arr = make([]int,size)
	for i:=0;i<size;i++{
		fmt.Println("enter element :",i+1)
		fmt.Scanln(&arr[i])
		
	}
	fmt.Println(arr)

	
}