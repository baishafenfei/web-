# web-
可以用于web压力测试和刷访问量
代码参考《用go编写简单的压力测试》

https://blog.csdn.net/oLeiShen/article/details/84106157

使用方法：
测试方法，-client是并发请求协程数量，-times是每个协程请求次数，-url是测试的地址，测试结果如下：
go build main.go
 
main.exe -client=20 -times=3 -url=https://baidu.com  
 
Start web Test clients:20 times:3  
Waiting web test end...  
End web Test  
PreTotal: 60  
Total: 60  
Success: 60  
Failure: 0  
SuccessRate: 100.00 %  
UseTime: 1.4413 s  
