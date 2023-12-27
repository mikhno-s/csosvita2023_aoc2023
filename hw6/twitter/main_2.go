package main

import "fmt"

// The main idea is to use linked list as a user's feed and inserts the folowee's tweets (2) into folower's (1) feed when:
// a) when user 1 follows user 2
// b) when user 1 gets a news feed (if user 1 follows user 2)
// This approach spreads time to computate feed during follow, unfollow and tweets list

type Tweet struct {
	TweetId int
	UserID  int
	prev    *Tweet
}

type Feed struct {
	head *Tweet
}

type User struct {
	ID        int
	Following map[int]bool
	OwnTweets []*Tweet // Saves the only user's tweet ids
	Feed      *Feed    // Leads to the most recent tweet in feed
}

type Twitter struct {
	Users map[int]*User
}

func Constructor() Twitter {
	t := &Twitter{make(map[int]*User)}
	for i := 1; i <= 2; i++ {

		u := &User{
			i,
			make(map[int]bool),
			make([]*Tweet, 0),
			&Feed{},
		}
		t.Users[i] = u
	}
	return *t
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	// get the most recent tweet in feed
	head := this.Users[userId].Feed.head

	// find the place for new tweet in the feed
	for head != nil && tweetId < head.TweetId {
		head = head.prev
	}

	// create a new tweet
	tweet := &Tweet{TweetId: tweetId, UserID: userId, prev: head}

	this.Users[userId].OwnTweets = append(this.Users[userId].OwnTweets, tweet)
	this.Users[userId].Feed.head = tweet
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	feed := make([]int, 0)
	head := this.Users[userId].Feed.head
	// Go over all subscriptions and check if last tweet of any subscrription is
	// more recent than last tweet in user's feed
	// If there is - append the last 10 tweets to user's feed until reach last existed in feed or all 10
	// for _, folowee := range this.Users[userId].Following {
	// 	for i := len(folowee.OwnTweets) - 1; i >= len(folowee.OwnTweets)-10 && i > 0; i-- {

	// 	}
	// }

	for i := 0; i < 10 && head != nil; i++ {
		feed = append(feed, head.TweetId)
		head = head.prev
	}
	return feed
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.Users[followerId].Following, followeeId)
	// remove all folowee tweets from folower's feed
	head := this.Users[followerId].Feed.head

	// starting from the most second recent tweet in feed
	for ; head != nil && head.prev != nil; head = head.prev {
		if head.prev.UserID == followeeId {
			head.prev = head.prev.prev
		}
	}

	// and remove recent tweet if it belongs to the followee user
	if this.Users[followerId].Feed.head != nil && this.Users[followerId].Feed.head.UserID == followeeId {
		this.Users[followerId].Feed.head = this.Users[followerId].Feed.head.prev
	}
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// Add to the following list
	this.Users[followerId].Following[followeeId] = true
	// Get first follower's tweet
	head := this.Users[followerId].Feed.head

	// go over followee own tweets and add TODO max 10 followee's tweets to the followers feed
	for i := len(this.Users[followeeId].OwnTweets) - 1; i >= 0; i-- {
		// creating a copy of a folowee's tweet without prev
		followeeTweet := this.Users[followeeId].OwnTweets[i]
		// search for a place in followers feed for a given folowee tweet
		if head != nil && followeeTweet.TweetId < head.TweetId {
			for head.prev != nil && followeeTweet.TweetId < head.prev.TweetId {
				head = head.prev
			}
		}

		if head == nil {
			followeeTweet.prev = head
			head = followeeTweet
		} else {
			followeeTweet.prev = head.prev
			head.prev = followeeTweet
		}

		// if follower have no tweets
		if this.Users[followerId].Feed.head == nil {
			this.Users[followerId].Feed.head = head
		}
	}

}

func main() {
	twitter := Constructor()

	twitter.PostTweet(1, 0)
	twitter.PostTweet(1, 1)
	twitter.PostTweet(2, 2)
	twitter.PostTweet(1, 3)

	// fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(2, 1)
	fmt.Println(twitter.GetNewsFeed(2))
	twitter.Unfollow(2, 1)
	// fmt.Println(twitter.GetNewsFeed(1))

}
