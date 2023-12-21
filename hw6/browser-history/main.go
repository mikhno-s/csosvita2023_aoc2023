package main

import "fmt"

type Page struct {
	Name string
	prev *Page
	next *Page
}

type BrowserHistory struct {
	current *Page
}

// Initializes the object with the homepage of the browser.
func Constructor(homepage string) BrowserHistory {
	b := &BrowserHistory{}
	b.current = &Page{Name: homepage}
	return *b
}

// Visits url from the current page. It clears up all the forward history.
func (this *BrowserHistory) Visit(url string) {
	page := &Page{Name: url, prev: this.current}
	this.current.next = page
	this.current = page
}

// Move steps back in history.
// If you can only return x steps in the history and steps > x, you will return only x steps.
// Return the current url after moving back in history at most steps.
func (this *BrowserHistory) Back(steps int) string {
	for ; steps > 0 && this.current.prev != nil; steps-- {
		this.current = this.current.prev
	}
	return this.current.Name
}

// Move steps forward in history.
// If you can only forward x steps in the history and steps > x, you will forward only x steps.
// Return the current url after forwarding in history at most steps.
func (this *BrowserHistory) Forward(steps int) string {
	for ; steps > 0 && this.current.next != nil; steps-- {
		this.current = this.current.next
	}
	return this.current.Name
}

func main() {
	browserHistory := Constructor("leetcode.com")
	browserHistory.Visit("google.com")     // You are in "leetcode.com". Visit "google.com"
	browserHistory.Visit("facebook.com")   // You are in "google.com". Visit "facebook.com"
	browserHistory.Visit("youtube.com")    // You are in "facebook.com". Visit "youtube.com"
	fmt.Println(browserHistory.Back(1))    // You are in "youtube.com", move back to "facebook.com" return "facebook.com"
	fmt.Println(browserHistory.Back(1))    // You are in "facebook.com", move back to "google.com" return "google.com"
	fmt.Println(browserHistory.Forward(1)) // You are in "google.com", move forward to "facebook.com" return "facebook.com"
	browserHistory.Visit("linkedin.com")   // You are in "facebook.com". Visit "linkedin.com"
	fmt.Println(browserHistory.Forward(2)) // You are in "linkedin.com", you cannot move forward any steps.
	fmt.Println(browserHistory.Back(2))    // You are in "linkedin.com", move back two steps to "facebook.com" then to "google.com". return "google.com"
	fmt.Println(browserHistory.Back(7))    // You are in "google.com", you can move back only one step to "leetcode.com". return "leetcode.com"

}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
