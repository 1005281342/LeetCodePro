package main

type Twitter struct {
	// index 唯一ID，增
	Index uint32

	// dd 用户推文列表
	Dd map[int][]uint32

	// users 用户集合
	Users map[int]map[int]bool

	// IdMap
	IdMap map[uint32]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{Dd: make(map[int][]uint32), Users: make(map[int]map[int]bool), IdMap: make(map[uint32]int), Index: 1}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := t.Dd[userId]; !ok {
		t.Dd[userId] = make([]uint32, 0)
	}
	t.Dd[userId] = append(t.Dd[userId], t.Index)
	t.IdMap[t.Index] = tweetId
	//atomic.AddUint32(&t.Index, 1)
	t.Index += 1
	t.Follow(userId, userId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {

	tans := make([]uint32, 0)
	users := t.Users[userId]
	for k, v := range users {
		if !v {
			continue
		}
		tans = merge(tans, t.Dd[k])
	}
	tans = tans[max(len(tans)-10, 0):]
	ans := make([]int, 0, 10)
	for _, v := range tans {
		ans = append(ans, t.IdMap[v])
	}
	return reserve(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reserve(a []int) []int {
	left := 0
	right := len(a) - 1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left += 1
		right -= 1
	}
	return a
}

func merge(a, b []uint32) []uint32 {
	ans := make([]uint32, 0, len(a)+len(b))
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i += 1
		} else {
			ans = append(ans, b[j])
			j += 1
		}
	}

	if i < len(a) {
		ans = append(ans, a[i:]...)
	}

	if j < len(b) {
		ans = append(ans, b[j:]...)
	}
	return ans
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.Users[followerId]; !ok {
		t.Users[followerId] = make(map[int]bool)
		// 初始化：自己关注自己
		t.Users[followerId][followerId] = true
	}
	// 绑定关注
	t.Users[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	if _, ok := t.Users[followerId]; !ok {
		return
	}
	// 解除关注
	t.Users[followerId][followeeId] = false
}

/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);
* obj.Unfollow(followerId,followeeId);
 */
