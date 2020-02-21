// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	cdpruntime "github.com/chromedp/cdproto/runtime"

	"github.com/chromedp/chromedp"
)

var (
	// these are set up in init
	execPath    string
	testdataDir string
	allocOpts   = chromedp.DefaultExecAllocatorOptions[:]
    allocTempDir string
	// allocCtx is initialised in TestMain, to cancel before exiting.
	allocCtx context.Context

	// browserCtx is initialised with allocateOnce
	browserCtx context.Context
	browserOpts []chromedp.ContextOption
)



func main()  {

	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-sync", false),
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", false),
		chromedp.Flag("disable-backgrounding-occluded-windows", false),
		chromedp.Flag("disable-breakpad", false),
		chromedp.Flag("disable-client-side-phishing-detection", false),
		chromedp.Flag("disable-default-apps", false),
		chromedp.Flag("disable-dev-shm-usage", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", false),
		chromedp.Flag("disable-ipc-flooding-protection", false),
		chromedp.Flag("disable-popup-blocking", false),
		chromedp.Flag("disable-prompt-on-repost", false),
		chromedp.Flag("disable-renderer-backgrounding", false),
		chromedp.Flag("disable-sync", false),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", false),
		chromedp.Flag("safebrowsing-disable-auto-update", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("window-size", "800,400"),
	}
	allocOpts = append(allocOpts, options...)


	allocCtx, _ = chromedp.NewExecAllocator(context.Background(), allocOpts...)

	fmt.Println("This is webserver base!")
	http.HandleFunc("/login", login)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bdpw", bdpw)


	//服务器要监听的主机地址和端口号
	http.ListenAndServe(":8080", nil)
}



func login(w http.ResponseWriter, req *http.Request) {


	ua := `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`

	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-sync", false),
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", false),
		chromedp.Flag("disable-backgrounding-occluded-windows", false),
		chromedp.Flag("disable-breakpad", false),
		chromedp.Flag("disable-client-side-phishing-detection", false),
		chromedp.Flag("disable-default-apps", false),
		chromedp.Flag("disable-dev-shm-usage", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", false),
		chromedp.Flag("disable-ipc-flooding-protection", false),
		chromedp.Flag("disable-popup-blocking", false),
		chromedp.Flag("disable-prompt-on-repost", false),
		chromedp.Flag("disable-renderer-backgrounding", false),
		chromedp.Flag("disable-sync", false),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", false),
		chromedp.Flag("safebrowsing-disable-auto-update", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.UserAgent(ua),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("window-size", "800,400"),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()
	// create a timeout
	//ctx, cancel = context.WithTimeout(ctx, 1500*time.Second)
	//defer cancel()

	// navigate to a page, wait for an element, click
	//var example string
	//var res string
	//err := chromedp.Run(ctx,
	//	chromedp.Navigate(`https://pan.baidu.com`),
	//)
	//log.Println("提取内容",res)
	//
	//log.Println(strings.TrimSpace(res))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// navigate to a page, wait for an element, click
	//var cookies = "BIDUPSID=6B2A24B16DA03C01980521FCB0215D63; PSTM=1577971193; BAIDUID=6B2A24B16DA03C01B43D1DD027053138:FG=1; PANWEB=1; BDUSS=Th1Tkoyd3BpN2tWcDBKVXlJUGVZWll2azE4cktLQVNoU1hpc2I3d2lXa2h5RGRlRVFBQUFBJCQAAAAAAAAAAAEAAAC8UjcUaHVtYW4wOTEyMDMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACE7EF4hOxBec; H_WISE_SIDS=139203_122155_139110_128064_135846_141003_139148_138903_141344_137758_138878_137985_141200_140174_131247_132552_137746_138165_107319_138883_140260_141753_140201_140592_138585_141651_138253_140114_136196_131862_140324_140578_133847_140792_140065_134047_131423_140311_140966_136537_110085_127969_140622_140593_139886_140992_138426_138944_140683_141191_140597; SCRC=0c932fcfbcf624a33f7fa23675990408; STOKEN=34e8de60ca22622c9509a0b9ea1774bc63ef688d49cbd06ed863c2ecb2c94d73; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; Hm_lvt_bff7b396e3d9f5901f8eec42a2ad3017=1582095100; BDCLND=FjZ2CFfNnBHcqxJWkFa2Zf%2BceBzbmtSB3TKIezbVAJw%3D; Hm_lvt_7a3960b6f067eb0085b7f96ff5e660b0=1582221536,1582253180,1582256082,1582268001; PSINO=1; delPer=1; H_PS_PSSID=1449_21080_30794_30823_26350; Hm_lpvt_7a3960b6f067eb0085b7f96ff5e660b0=1582275654; cflag=13%3A3; PANPSC=3849945753961645605%3AZdTG9lVgpntuTp0g%2BfOF8lcS2d9ns3O5g0mIZdLHpdQGbqupDlB1glh23woWtLTmJOKKgNI1zw3iLaLUl0KD5De4ZdaW7CoOlL98c8Ccr9ch6uZoP3DwQ9YfJggg9xZJ2d71q7e%2BSo8QMYi%2Bz%2BKjhj7huzw8ix2mdmaUZ5sCjF2zGaJuDNHJ27jq%2FTte52B2jBKIyGsgLB%2BSsVZaqyA8TA%3D%3D"

	chromedp.Run(ctx,

		// 设置cookie
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			network.SetCookie("BAIDUID","6B2A24B16DA03C01B43D1DD027053138:FG=1"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("BDCLND","FjZ2CFfNnBHcqxJWkFa2Zf%2BceBzbmtSB3TKIezbVAJw%3D"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("BDORZ","B490B5EBF6F3CD402E515D22BCDA1598"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("BDUSS","Th1Tkoyd3BpN2tWcDBKVXlJUGVZWll2azE4cktLQVNoU1hpc2I3d2lXa2h5RGRlRVFBQUFBJCQAAAAAAAAAAAEAAAC8UjcUaHVtYW4wOTEyMDMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACE7EF4hOxBec"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("BIDUPSID","6B2A24B16DA03C01980521FCB0215D63"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("HMACCOUNT","890DA65E9DC1226F"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("HMVT","6bcd52f51e9b3dce32bec4a3997715ac|1582277381|7a3960b6f067eb0085b7f96ff5e660b0|1582279157|"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("H_PS_PSSID","1449_21080_30794_30823_26350"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("H_WISE_SIDS","139203_122155_139110_128064_135846_141003_139148_138903_141344_137758_138878_137985_141200_140174_131247_132552_137746_138165_107319_138883_140260_141753_140201_140592_138585_141651_138253_140114_136196_131862_140324_140578_133847_140792_140065_134047_131423_140311_140966_136537_110085_127969_140622_140593_139886_140992_138426_138944_140683_141191_140597"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("Hm_lpvt_7a3960b6f067eb0085b7f96ff5e660b0","1582279160"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("Hm_lvt_7a3960b6f067eb0085b7f96ff5e660b0","1582221536,1582253180,1582256082,1582268001"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("Hm_lvt_bff7b396e3d9f5901f8eec42a2ad3017","1582095100"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("PANPSC","13738774843283734416%3AKkwrx6t0uHBtjxk3WnkKJEYemHfSgA%2FTTFZwkLLhvKtIz%2B7SArqnI0ci2O0YbEjp8K4BmbK5xNkOgcchTM%2FVqKKK%2FFH5WUxXO5sK9zFMuONXLSXztUUmuXZNBoZ8i0CTW8pfgMw2m%2FrQ8rlESRQE3CZTnz2RR6t%2BZFCUk453j8DrbkBUnMiFfw%3D%3D"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("PANWEB","1"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("PSINO","1"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("PSTM","1577971193"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("SCRC","0c932fcfbcf624a33f7fa23675990408"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("STOKEN","34e8de60ca22622c9509a0b9ea1774bc7d355e9a9feac379557f97d7c5ee826f"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("STOKEN","34e8de60ca22622c9509a0b9ea1774bc7d355e9a9feac379557f97d7c5ee826f"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("STOKEN","34e8de60ca22622c9509a0b9ea1774bc63ef688d49cbd06ed863c2ecb2c94d73"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("cflag","13%3A3"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("delPer","1"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			network.SetCookie("pcsett","1582365556-73a522153eea459f8c8abadfe2cef4e5"). //设置cookie
				WithExpires(&expr).
				WithDomain("https://pan.baidu.com/"). //访问网站主体
				WithHTTPOnly(true).
				Do(ctx)
			return nil
		}),
		chromedp.Sleep(2 * time.Second),               //缓一缓
		//chromedp.Navigate(`https://pan.baidu.com/s/1o9M2bGI`),
		//chromedp.WaitVisible(`#fdxf27ax`, chromedp.ByQuery), // 等待id=kw渲染成功，成功则说明已经获取到了正确的页面
		//chromedp.SendKeys(`#fdxf27ax`, `jxg8`, chromedp.ByID), //输入关键词
		//chromedp.Sleep(1 * time.Second),               //缓一缓
		//chromedp.Click("#guarGbx > a > span > span", chromedp.ByQuery),          // 触发点击事件，
		//chromedp.Sleep(2 * time.Second),               //缓一缓
		//chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
		//chromedp.Text(`#layoutAside`, &res, chromedp.NodeVisible, chromedp.ByID), //获取相关人物节点

		// 将chromedp.OuterHTML("body", &res) 替换为下面的代码
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取cookie
			cookies, err := network.GetAllCookies().Do(ctx)
			// 将cookie拼接成header请求中cookie字段的模式
			var c string
			for _, v := range cookies {
				c = c + v.Name + "=" + v.Value + ";"
			}
			log.Println(c)
			if err != nil {
				return err
			}
			return nil
		}),
		//chromedp.Text(`yunData.SHARE_ID`, &res, chromedp.NodeReady, chromedp.BySearch), //获取相关人物节点

		//chromedp.Click("#layoutMain > div.frame-content > div.module-share-header > div > div.slide-show-right > div > div > div.x-button-box > a.g-button.g-button-blue > span > span", chromedp.ByQuery),          // 触发保存
		//chromedp.Click("#fileTreeDialog > div.dialog-body > div > ul > li > div > span > span", chromedp.ByQuery),          // 触发保存到根目录
		//chromedp.Click("#fileTreeDialog > div.dialog-footer.g-clearfix > a.g-button.g-button-blue-large > span > span", chromedp.ByQuery),          // 触发保存确定
		//chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
		chromedp.Sleep(2 * time.Second),
	)




	// run task list
	//err := chromedp.Run(ctx2, chromedp.Navigate(`https://pan.baidu.com/s/1o9M2bGI`),chromedp.Sleep(2000 * time.Second),)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Printf("Go's time.After example:\n%s", example)
	//time.Sleep(200 * time.Second)
}

func hello(w http.ResponseWriter, req *http.Request)  {
	testAllocate()
}

var allocateOnce sync.Once

func testAllocate() (context.Context, context.CancelFunc) {
	// Start the browser exactly once, as needed.
	allocateOnce.Do(func() { browserCtx, _ = testAllocateSeparate() })

	if browserCtx == nil {
		// allocateOnce.Do failed; continuing would result in panics.
		log.Panic(browserCtx)
	}

	// Same browser, new tab; not needing to start new chrome browsers for
	// each test gives a huge speed-up.
	ctx, _ := chromedp.NewContext(browserCtx)
		res := ""
	// Only navigate if we want an html file name, otherwise leave the blank page.
		log.Printf("run...")
		if err := chromedp.Run(ctx, chromedp.Navigate("https://pan.baidu.com/s/1o9M2bGI"),
			//chromedp.OuterHTML("html", &res, chromedp.ByQuery),
			chromedp.WaitVisible(`#fdxf27ax`, chromedp.ByQuery), // 等待id=kw渲染成功，成功则说明已经获取到了正确的页面
			chromedp.SendKeys(`#fdxf27ax`, `jxg8`, chromedp.ByID), //输入关键词
			chromedp.Sleep(1 * time.Second),               //缓一缓
			chromedp.Click("#guarGbx > a > span > span", chromedp.ByQuery),          // 触发点击事件，
			chromedp.Sleep(2 * time.Second),               //缓一缓
			//chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
			//chromedp.Text(`#layoutAside`, &res, chromedp.NodeVisible, chromedp.ByID), //获取相关人物节点
			//chromedp.Click("#layoutMain > div.frame-content > div.module-share-header > div > div.slide-show-right > div > div > div.x-button-box > a.g-button.g-button-blue > span > span", chromedp.ByQuery),          // 触发保存
			//chromedp.Click("#fileTreeDialog > div.dialog-body > div > ul > li > div > span > span", chromedp.ByQuery),          // 触发保存到根目录
			//chromedp.Click("#fileTreeDialog > div.dialog-footer.g-clearfix > a.g-button.g-button-blue-large > span > span", chromedp.ByQuery),          // 触发保存确定
			chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
			); err != nil {
			log.Panic(err)
		}
	log.Printf("result..")
	log.Printf(res)



	cancel := func() {
		if err := chromedp.Cancel(ctx); err != nil {
			log.Panic(err)
		}
	}
	return ctx, cancel
}


func testAllocateSeparate() (context.Context, context.CancelFunc) {
	// Entirely new browser, unlike testAllocate.

	log.Println(browserOpts)
	ctx, _ := chromedp.NewContext( allocCtx, browserOpts...)
	if err := chromedp.Run(ctx); err != nil {
		log.Panic(err)
	}
	chromedp.ListenBrowser(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *cdpruntime.EventExceptionThrown:
			fmt.Errorf("%+v\n", ev.ExceptionDetails)
		}
	})
	cancel := func() {
		if err := chromedp.Cancel(ctx); err != nil {
			log.Panic(err)
		}
	}
	return ctx, cancel
}

func bdpw(w http.ResponseWriter, req *http.Request) {

	ua := `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`

	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(ua),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()


	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 1500*time.Second)
	defer cancel()



	// navigate to a page, wait for an element, click
	var example string
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pan.baidu.com/s/1o9M2bGI`),
		chromedp.WaitVisible(`#fdxf27ax`, chromedp.ByQuery), // 等待id=kw渲染成功，成功则说明已经获取到了正确的页面
		//chromedp.SendKeys(`#fdxf27ax`, `jxg8`, chromedp.ByID), //输入关键词
		//chromedp.Sleep(1 * time.Second),               //缓一缓
		//chromedp.Click("#guarGbx > a > span > span", chromedp.ByQuery),          // 触发点击事件，
		//chromedp.Sleep(2 * time.Second),               //缓一缓
		//chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
		//chromedp.Text(`#layoutAside`, &res, chromedp.NodeVisible, chromedp.ByID), //获取相关人物节点

		// 将chromedp.OuterHTML("body", &res) 替换为下面的代码
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取cookie
			cookies, err := network.GetAllCookies().Do(ctx)
			// 将cookie拼接成header请求中cookie字段的模式
			var c string
			for _, v := range cookies {
				c = c + v.Name + "=" + v.Value + ";"
			}
			log.Println(c)
			if err != nil {
				return err
			}
			return nil
		}),
		//chromedp.Text(`yunData.SHARE_ID`, &res, chromedp.NodeReady, chromedp.BySearch), //获取相关人物节点

		//chromedp.Click("#layoutMain > div.frame-content > div.module-share-header > div > div.slide-show-right > div > div > div.x-button-box > a.g-button.g-button-blue > span > span", chromedp.ByQuery),          // 触发保存
		//chromedp.Click("#fileTreeDialog > div.dialog-body > div > ul > li > div > span > span", chromedp.ByQuery),          // 触发保存到根目录
		//chromedp.Click("#fileTreeDialog > div.dialog-footer.g-clearfix > a.g-button.g-button-blue-large > span > span", chromedp.ByQuery),          // 触发保存确定
		//chromedp.OuterHTML("html", &res, chromedp.ByQuery), //获取html源码
		chromedp.Sleep(2000 * time.Second),
	)
	log.Println("提取内容",res)

	//transferResult := transferFile("https://pan.baidu.com/share/transfer?shareid=2880173538&from=231453366&channel=chunlei&web=1&app_id=250528&bdstoken=29ad595f8b03d26e1409a4fe2b7501da&logid=MTU4MjEzMTI3MTAwMjAuOTI3MTY0NjIwMDIyMzc0&clienttype=0")
	//log.Printf("保存结果",transferResult)

	log.Println(strings.TrimSpace(res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Go's time.After example:\n%s", example)
	time.Sleep(200 * time.Second)
}





func transferFile(url string) (string){
	method := "POST"

	payload := strings.NewReader("fsidlist=%20%5B197430756538058%5D&path=%20/")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", " application/json, text/javascript, */*; q=0.01")
	//req.Header.Add("Accept-Encoding", " gzip, deflate, br")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Content-Length", " 39")
	req.Header.Add("Content-Type", " application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", " BIDUPSID=6B2A24B16DA03C01980521FCB0215D63; PSTM=1577971193; BAIDUID=6B2A24B16DA03C01B43D1DD027053138:FG=1; PANWEB=1; BDUSS=Th1Tkoyd3BpN2tWcDBKVXlJUGVZWll2azE4cktLQVNoU1hpc2I3d2lXa2h5RGRlRVFBQUFBJCQAAAAAAAAAAAEAAAC8UjcUaHVtYW4wOTEyMDMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACE7EF4hOxBec; H_WISE_SIDS=139203_122155_139110_128064_135846_141003_139148_138903_141344_137758_138878_137985_141200_140174_131247_132552_137746_138165_107319_138883_140260_141753_140201_140592_138585_141651_138253_140114_136196_131862_140324_140578_133847_140792_140065_134047_131423_140311_140966_136537_110085_127969_140622_140593_139886_140992_138426_138944_140683_141191_140597; SCRC=0c932fcfbcf624a33f7fa23675990408; STOKEN=34e8de60ca22622c9509a0b9ea1774bc63ef688d49cbd06ed863c2ecb2c94d73; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; Hm_lvt_7a3960b6f067eb0085b7f96ff5e660b0=1582081873,1582081917,1582082080,1582094751; Hm_lvt_bff7b396e3d9f5901f8eec42a2ad3017=1582095100; Hm_lpvt_bff7b396e3d9f5901f8eec42a2ad3017=1582127958; PSINO=1; delPer=1; H_PS_PSSID=1449_21080_30794_30823_26350; BDCLND=doburNTL57tM92nZjEOjthN6IkYg84s7C0Y6Wt1iw9k%3D; Hm_lpvt_7a3960b6f067eb0085b7f96ff5e660b0=1582131246; PANPSC=7354584195918573065%3AKkwrx6t0uHBtjxk3WnkKJEYemHfSgA%2FTTFZwkLLhvKtIz%2B7SArqnI0ci2O0YbEjpWx6VHFFT8q%2FUZOkxKN9%2F2qKK%2FFH5WUxXO5sK9zFMuONXLSXztUUmuXZNBoZ8i0CTW8pfgMw2m%2FrQ8rlESRQE3CZTnz2RR6t%2BZFCUk453j8DrbkBUnMiFfw%3D%3D")
	req.Header.Add("DNT", " 1")
	req.Header.Add("Host", " pan.baidu.com")
	req.Header.Add("Origin", " https://pan.baidu.com")
	req.Header.Add("Referer", " https://pan.baidu.com/s/1zD1KZD_PZZtAoJvWX8g6kw")
	req.Header.Add("Sec-Fetch-Mode", " cors")
	req.Header.Add("Sec-Fetch-Site", " same-origin")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	req.Header.Add("X-Requested-With", " XMLHttpRequest")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body)
}

