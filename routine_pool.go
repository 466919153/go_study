package main
import "fmt"
import "time"

func go_worker(name string){
     for i:0;i<10;i++{
          fmt.Println("I,am a goroutine,my name is",name,"-----")
          time.Sleep(1*time.Second)
     
     }
     fmt.Println(name,"run ok!")
}

func main(){
     go go_worker("kobe")
     go go_worker("iverson")
     for{
          time.Sleep(1*time.Second)
     }

}
