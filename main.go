package main
import (
  "sync"
  "net/http"
  "time"
  "fmt"
)

var (
  total = 0.0
  success = 0.0
  faile = 0.0
  count = 5 
 )
 
 var wg sync.WaitGroup
 
func run(num int) {
  defer wg.Done()
  
  no := 0.0
  ok := 0.0
  
  for i := 0 ; i< num ; i++ {
    //set url here
    resp , err := http.Get("http://www.baidu.com")
    if err := nil {
      fmt.Println("err:",err)
      no += 1
      break
    }
    //Close resp body
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
      no += 1
      break
    }
    ok += 1
  }
  success += ok
  faile += no
  total += float64(num)
}
func main () {
  start_time := time.Now().UnixNano()
  fmt.Println("Start web Test")
  for i := 0 ; i < count;i++ {
    wg.Add(1)
    go run(10)
  }
  
  fmt.Println("Wait web test end")
  wg.Wait()
  end_time := time.Now().UnixNano()
  fmt.Println("End web Test")
  fmt.Println("Total:",total)
  fmt.Println("Success:",success)
  fmt.Println("Faile:",faile)
  fmt.Println("UseTime:,fmt.Sprint("%.4f",float64(end_time-start_time)))
}            
  
  
  
 
