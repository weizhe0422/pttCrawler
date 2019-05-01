package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"regexp"
)

func main() {
	var (
		resp *http.Response
		err  error
		contents []byte
	)
	if resp, err = http.Get("https://www.ptt.cc/bbs/index.html"); err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	contents, err = httputil.DumpResponse(resp, true)
	//fmt.Println(string(contents))
	findBoradList(contents)
}

/*
   <div class="b-ent">
       <a class="board" href="/bbs/[0-9a-zA-Z]+/index.html">
           <div class="board-name">[0-9a-zA-Z]+</div>
           <div class="board-nuser"><span class="[^>]*">[0-9]+</span></div>
           <div class="board-class">[^<>]*</div>
           <div class="board-title">[^<>]*</div>
       </a>
   </div>
*/
func findBoradList(contents []byte){
	compile := regexp.MustCompile(`        <div class="b-ent">
            <a class="board" href="(/bbs/[0-9a-zA-Z]+/index.html)">
                <div class="board-name">([0-9a-zA-Z]+)</div>
                <div class="board-nuser"><span class="[^>]*">[0-9]+</span></div>
                <div class="board-class">([^<>]*)</div>
                <div class="board-title">([^<>]*)</div>
            </a>
        </div>`)

	matchs := compile.FindAllSubmatch(contents, -1)
	for _, match := range matchs{
		fmt.Printf("版名:%s, 分類:%s, 標題:%s, URL:%s",match[2],match[3],match[4],match[1])
		fmt.Println()
	}
}
