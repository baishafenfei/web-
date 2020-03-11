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
 
 var wg sync.WaitGroup
 
 
