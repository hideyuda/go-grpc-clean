package driver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
)

type WebDriver struct {
	Browser *rod.Browser
}

func (w *WebDriver) SignInToUberSuggest(myDomain, compatibleDomain string) error {
	var err error

	l := launcher.New().
		Headless(false).
		Devtools(true)

	defer l.Cleanup() // remove launcher.FlagUserDataDir

	url := l.MustLaunch()

	// Launch a new browser with default options, and connect to it.
	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()

	launcher.Open(browser.ServeMonitor(""))

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://app.neilpatel.com/ja/login").MustWaitLoad()

	// browser.MustHandleAuth(config.UberSuggest.Email, config.UberSuggest.Password)

	// We use css selector to get the search input element and input "git"
	page.MustElement("input[name=email]").
		MustInput(config.UberSuggest.Email)
	// MustType(input.Enter)

	page.MustElement("input[name=password]").
		MustInput(config.UberSuggest.Password)
		// MustType(input.Enter)

	page.MustElement("button.sc-gSYCTC").
		MustClick().
		MustWaitLoad()

	log.Println("Signed in to UberSuggest")

	// Create a new page
	page.MustNavigate("https://app.neilpatel.com/ja/labs/keywords_generator").MustWaitLoad()

	page.MustElement("div.sc-jETAwz.hPEqmO > input.sc-jPnjyd.CFabW").
		MustInput(myDomain).
		MustType(input.Enter).
		MustWaitLoad()

	// page.MustElement("input.sc-ibZcGq.dEABVp").
	// 	MustInput("日本語").
	// 	MustType(input.Enter).
	// 	MustWaitLoad()

	page.MustElement("div.sc-bfbdKt.fnVbQt > button.sc-hxPjim.fZwflK").
		MustClick().
		MustWaitLoad()

	page.MustElement("input.sc-dHvAOr.cFLZpV").
		MustInput(compatibleDomain).
		MustType(input.Enter).
		MustWaitLoad()

	wait := browser.MustWaitDownload()

	page.MustElement("button.sc-khIgXV.sc-hTRkEk.iYFxeN.fgItdR").
		MustClick()

	filePath := fmt.Sprintf("%s-%s-%v.csv", myDomain, compatibleDomain, time.Now().Format("2006-01-02-15-04-05"))

	err = utils.OutputFile(filePath, wait())
	if err != nil {
		log.Fatalf("failed to save file: err=%v", err)
		return err
	}

	return nil

	// We use css selector to get the search input element and input "git"

	// // Wait until css selector get the element then get the text content of it.
	// text := page.MustElement(".codesearch-results p").MustText()

	// fmt.Println(text)

	// // Get all input elements. Rod supports query elements by css selector, xpath, and regex.
	// // For more detailed usage, check the query_test.go file.
	// fmt.Println("Found", len(page.MustElements("input")), "input elements")

	// // Eval js on the page
	// page.MustEval(`() => console.log("hello world")`)

	// // Pass parameters as json objects to the js function. This MustEval will result 3
	// fmt.Println("1 + 2 =", page.MustEval(`(a, b) => a + b`, 1, 2).Int())

	// // When eval on an element, "this" in the js is the current DOM element.
	// fmt.Println(page.MustElement("title").MustEval(`() => this.innerText`).String())

}

func (w *WebDriver) DisableHeadless() {
	// Headless runs the browser on foreground, you can also use flag "-rod=show"
	// Devtools opens the tab in each new tab opened automatically
	l := launcher.New().
		Headless(false).
		Devtools(true)

	defer l.Cleanup() // remove launcher.FlagUserDataDir

	url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// SlowMotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()

	// ServeMonitor plays screenshots of each tab. This feature is extremely
	// useful when debugging with headless mode.
	// You can also enable it with flag "-rod=monitor"
	launcher.Open(browser.ServeMonitor(""))

	defer browser.MustClose()

	// page := browser.MustPage("https://github.com/")

	// page.MustElement("input").MustInput("git").MustType(input.Enter)

	// text := page.MustElement(".codesearch-results p").MustText()

	// fmt.Println(text)

	// utils.Pause() // pause goroutine
}

func (w *WebDriver) HandleError() {
	page := rod.New().MustConnect().MustPage("https://mdn.dev")

	// We use Go's standard way to check error types, no magic.
	check := func(err error) {
		var evalErr *rod.ErrEval
		if errors.Is(err, context.DeadlineExceeded) { // timeout error
			fmt.Println("timeout err")
		} else if errors.As(err, &evalErr) { // eval error
			fmt.Println(evalErr.LineNumber)
		} else if err != nil {
			fmt.Println("can't handle", err)
		}
	}

	// The two code blocks below are doing the same thing in two styles:

	// The block below is better for debugging or quick scripting. We use panic to short-circuit logics.
	// So that we can take advantage of fluent interface (https://en.wikipedia.org/wiki/Fluent_interface)
	// and fail-fast (https://en.wikipedia.org/wiki/Fail-fast).
	// This style will reduce code, but it may also catch extra errors (less consistent and precise).
	{
		err := rod.Try(func() {
			fmt.Println(page.MustElement("a").MustHTML()) // use "Must" prefixed functions
		})
		check(err)
	}

	// The block below is better for production code. It's the standard way to handle errors.
	// Usually, this style is more consistent and precise.
	{
		el, err := page.Element("a")
		if err != nil {
			check(err)
			return
		}
		html, err := el.HTML()
		if err != nil {
			check(err)
			return
		}
		fmt.Println(html)
	}
}

func (w *WebDriver) HandleEvent() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()

	done := make(chan struct{})

	// Listen for all events of console output.
	go page.EachEvent(func(e *proto.RuntimeConsoleAPICalled) {
		fmt.Println(page.MustObjectsToJSON(e.Args))
		close(done)
	})()

	wait := page.WaitEvent(&proto.PageLoadEventFired{})
	page.MustNavigate("https://mdn.dev")
	wait()

	// EachEvent allows us to achieve the same functionality as above.
	if false {
		// Subscribe events before they happen, run the "wait()" to start consuming
		// the events. We can return an optional stop signal to unsubscribe events.
		wait := page.EachEvent(func(e *proto.PageLoadEventFired) (stop bool) {
			return true
		})
		page.MustNavigate("https://mdn.dev")
		wait()
	}

	// Or the for-loop style to handle events to do the same thing above.
	if false {
		page.MustNavigate("https://mdn.dev")

		for msg := range page.Event() {
			e := proto.PageLoadEventFired{}
			if msg.Load(&e) {
				break
			}
		}
	}

	page.MustEval(`() => console.log("hello", "world")`)

	<-done
}

func (w *WebDriver) WaitForAnimation() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://getbootstrap.com/docs/4.0/components/modal/")

	page.MustWaitLoad().MustElement("[data-target='#exampleModalLive']").MustClick()

	saveBtn := page.MustElementR("#exampleModalLive button", "Close")

	// Here, WaitStable will wait until the button's position and size become stable.
	saveBtn.MustWaitStable().MustClick().MustWaitInvisible()

	fmt.Println("done")

}

func (w *WebDriver) WaitForRequest() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.wikipedia.org/").MustWaitLoad()

	// Start to analyze request events
	wait := page.MustWaitRequestIdle()

	// This will trigger the search ajax request
	page.MustElement("#searchInput").MustClick().MustInput("lisp")

	// Wait until there's no active requests
	wait()

	// We want to make sure that after waiting, there are some autocomplete
	// suggestions available.
	fmt.Println(len(page.MustElements(".suggestion-link")) > 0)

}
