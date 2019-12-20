### Colly
---
#### 一、
First , you need learn the "goquery"
* documention : https://github.com/PuerkitoBio/goquery
* install
```go
    go get github.com/PuerkitoBio/goquery
```
#### 二、 
* five Callback function
    * `OnRequest` : Called `before` send the request 
    * `OnError` : if in the request , appear a error .
    * `OnResponse` : Called `after` accpet the response , 
    * `OnHTML` : if the response data is html, Called `after` `OnResponse`
    * `OnScraped` : Called `after` the `OnHTML`  
    
* 根据链接每次准备抓取数据前调用 注册的 OnRequest做每次抓取前的预处理工作
  当抓取数据失败时会调用OnError做错误处理
  抓取到数据后调用OnResponse，做刚抓到数据时的处理工作
  然后分析抓取到的数据会根据页面上的dom节点触发OnHTML回调进行数据分析
  数据分析完毕后会调用 OnScraped函数进行每次抓取后的收尾工作  
  
* colly也提供了部分辅助接口，协助完成数据抓取分析流程, 以下列举一部分主要的支持。
    * queue 用于存放等待抓取的链接
    * proxy 用于代理发起抓取源
    * thread 支持多携程并发处理
    * filter 支持对特殊链接进行过滤
    * depth 可以设置抓取深度控制抓取
    
