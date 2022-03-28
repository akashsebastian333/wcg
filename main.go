package main

import (
  "io"
  "log"
  "net/http"
  "fmt"
  "io/ioutil"
  "strings"
  "golang.org/x/net/html"
  "os"
  "regexp"
)

func main() {

  if len(os.Args) <= 1 {
        fmt.Println("Usage:", os.Args[0], "domain.com")
        return
    }
        




  domain := os.Args[1]
 
  Url := fmt.Sprintf("http://webcache.googleusercontent.com/search?q=cache:%s&strip=0&vwsrc=0", domain)


// fecting links

  fmt.Println("\n\nLinks\n\n")
  getting_links(Url)


// fecting comments
  fmt.Println("\n\nComments inside\n\n")
  getting_comments(Url)


}


func getting_links(Url string){
  resp, err := http.Get(Url)
    if err != nil {
        log.Fatal(err)
    }
  var d []string
    for _, v := range get_Tags(resp.Body) {
    data := strings.Split(string(v), "\n")
    for i := 0; i < len(data); i++ {
      if strings.Contains(data[i], "support.google.com") || strings.Contains(data[i], "googleusercontent.com") || strings.Contains(data[i], "javascript:void"){
        continue
      } else {
        if (data[i] == "#"){
          continue
        }else{
          d = append(d, data[i])
        }    
      }
   }
}

  rD := removeDuplicate(d)  
  for i := 0; i < len(rD); i++ {
      fmt.Println(rD[i])
    }
}


func getting_comments(Url string) {
  resp, err := http.Get(Url)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error reading HTTP body. ", err)
  }
  re := regexp.MustCompile("<!--(.|\n)*?-->")
  comments := re.FindAllString(string(body), -1)
  if comments == nil {
    fmt.Println("No Comments Found")
  } else {
      for _, comment := range comments {
        fmt.Println(comment)
    } 
  } 
}


func removeDuplicate(intSlice []string) []string  {
  keys := make(map[string]bool)
  list := []string{}
  for _, entry := range intSlice {
    if _, value := keys[entry]; !value {
      keys[entry] = true
      list = append(list, entry)
    }
  }
  return list
}


func get_Tags(body io.Reader) []string {
  var links []string
  x := html.NewTokenizer(body)
  for {
    tags := x.Next()
    switch tags {
    case html.ErrorToken:
      return links
    case html.StartTagToken, html.EndTagToken:
      token := x.Token()
      if "a" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "href" {
            links = append(links, attr.Val)
          }
        }
      }
      if "base" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "href" {
            links = append(links, attr.Val)
          }
        }
      }
      if "link" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "href" {
            links = append(links, attr.Val)
          }
        }
      }
      if "area" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "href" {
            links = append(links, attr.Val)
          }
        }
      }
      if "script" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "img" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "style" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "audio" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "source" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "embed" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "iframe" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "track" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
      if "video" == token.Data {
        for _, attr := range token.Attr {
          if attr.Key == "src" {
            links = append(links, attr.Val)
          }
        }
      }
    }
  }
}
