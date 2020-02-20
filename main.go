// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.baidu.com/`),
		chromedp.WaitVisible(`#kw`, chromedp.ByQuery), // 等待id=kw渲染成功，成功则说明已经获取到了正确的页面
		chromedp.SendKeys(`#kw`, `美女`, chromedp.ByID), //输入关键词
		chromedp.Click("#su", chromedp.ByID),          // 触发点击事件，
		chromedp.Sleep(1 * time.Second),               //缓一缓
		//chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
		chromedp.Text(`#con-ar`, &res, chromedp.NodeVisible, chromedp.ByID), //获取相关人物节点
	)
	log.Println(strings.TrimSpace(res))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}

//op-img-address-divide-high