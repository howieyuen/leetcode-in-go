package _355_design_twitter

type Twitter struct {
	users  map[int][]int
	tweets Tweets
}

type Tweets []Tweet

type Tweet struct {
	tweetId int
	userId  int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users:  make(map[int][]int),
		tweets: make(Tweets, 0),
	}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	tweet := Tweet{
		tweetId: tweetId,
		userId:  userId,
	}
	t.tweets = append(t.tweets, tweet)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	var allTweets []int
	var allFollowerIds = append(t.users[userId], userId)
	for i := len(t.tweets) - 1; i >= 0; i-- {
		var included bool
		for _, followerId := range allFollowerIds {
			if t.tweets[i].userId == followerId {
				included = true
				break
			}
		}
		if included {
			allTweets = append(allTweets, t.tweets[i].tweetId)
		}
		if len(allTweets) == 10 {
			break
		}
	}
	return allTweets
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.users[followerId]; ok {
		t.users[followerId] = append(t.users[followerId], followeeId)
	} else {
		t.users[followerId] = []int{followeeId}
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	for i, followee := range t.users[followerId] {
		if followee == followeeId {
			t.users[followerId] = append(t.users[followerId][:i], t.users[followerId][i+1:]...)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
