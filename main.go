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
 )

 var (
   client = flag.Int("clients",1,"Please input client quantity(default: 1 client)")
   times = flag.Int("count",1,"Please input times of one client quantity(default: 1 times)")
   url = flag.String("urlPath","https://www.baidu.com","Please input urlpath you want to test(default:baidu)")
   sleep = flag.Int("sleep",0,"Please input sleep time if you want to sleep duing test(default: close)")
   sleep_count = flag.Int("sleep-count",0,"Please input sleep count(default:0)")
 )

 
 var wg sync.WaitGroup
 
func run(num int) {
  defer wg.Done()
  
  no := 0.0
  ok := 0.0
  count := 0
  
  for i := 0 ; i< num ; i++ {
    //set url here
    resp , err := http.Get(*url)
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
    count += 1
    //如果是刷访问量，这里会sleep
    if *sleep != 0 && *sleep_count != 0 && count == *sleep_count {
      fmt.Println("Need sleep now ",count)
        time.Sleep(time.Duration(*sleep)*time.Millisecond )
        count = 0
    }

  }
  success += ok
  faile += no
  total += float64(num)
}
func main () {
  start_time := time.Now().UnixNano()
  
  flag.Parse()
  if *client == 0 || *times == 0 || *url == "" {
    flag.PrintDefault()
    return
  }
  
  fmt.Println("Start web Test clients:",*client,"times:",*times)
  
  for i := 0 ; i < *client;i++ {
    wg.Add(1)
    go run(*times)
  }
  
  fmt.Println("Wait web test end...")
  wg.Wait()
  end_time := time.Now().UnixNano()
  
  fmt.Println("End web Test")
  fmt.Println("PreTotal:",(*client)*(*times))
  fmt.Println("Total:",total)
  fmt.Println("Success:",success)
  fmt.Println("Faile:",faile)
  fmt.Println("SuccessRate:", fmt.Sprintf("%.2f", ((success/total)*100.0)), "%")  
  fmt.Println("UseTime:", fmt.Sprintf("%.4f", float64(end_time-start_time)/1e9), "s")
}            
  
  
  
 
