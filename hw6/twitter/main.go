package main

import (
	"container/heap"
	"fmt"
)

type MaxRecentTweetsHeap []int

type User struct {
	ID        int
	Following map[int]bool
	Tweets    []int
}

type Twitter struct {
	Users     map[int]*User
	Tweets    map[int]int
	TweetsNum int
}

func Constructor() Twitter {
	t := &Twitter{make(map[int]*User), make(map[int]int), 0}
	for i := 1; i <= 500; i++ {

		u := &User{
			i,
			make(map[int]bool),
			make([]int, 0),
		}
		t.Users[i] = u
		t.Users[i].Following[i] = true
	}
	return *t
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Users[userId].Tweets = append(this.Users[userId].Tweets, this.TweetsNum)
	this.Tweets[this.TweetsNum] = tweetId
	this.TweetsNum += 1
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.Users[followerId].Following, followeeId)
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.Users[followerId].Following[followeeId] = true
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	feed := make([]int, 0)
	tweets := &MaxRecentTweetsHeap{}
	for followeeId := range this.Users[userId].Following {
		for _, tweet := range this.Users[followeeId].Tweets {
			heap.Push(tweets, tweet)
		}
	}
	for i := 0; i < 10 && tweets.Len() > 0; i++ {
		t := heap.Pop(tweets).(int)
		feed = append(feed, this.Tweets[t])
	}

	return feed
}

func (h MaxRecentTweetsHeap) Len() int           { return len(h) }
func (h MaxRecentTweetsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxRecentTweetsHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *MaxRecentTweetsHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MaxRecentTweetsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	twitter := Constructor()

	twitter.PostTweet(1, 5)
	twitter.PostTweet(1, 3)
	fmt.Println(twitter.GetNewsFeed(1))
}
