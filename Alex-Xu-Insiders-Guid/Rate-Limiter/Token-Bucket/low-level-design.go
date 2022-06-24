package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	UserID     int
	BucketSize int
	RefillRate int
	Interval   int
	Blocked    bool
}

func New(UserID, BucketSize, RefillRate, Interval int) *TokenBucket {
	bucket := new(TokenBucket)
	bucket.UserID = UserID
	bucket.BucketSize = BucketSize
	bucket.RefillRate = RefillRate
	bucket.Interval = Interval
	bucket.Blocked = false
	return bucket
}

func (this *TokenBucket) refill() {
	time.Sleep(time.Duration(this.Interval) * time.Second)
	this.BucketSize = this.RefillRate
	this.Blocked = false
}

func (this *TokenBucket) geteway() bool {
	if this.BucketSize > 0 {
		this.BucketSize--
		return true
	}
	return false
}

func (this *TokenBucket) get(endpoint string) {

	res := this.geteway()

	if res {
		fmt.Printf("https://api.believer/com/%v\n\n", endpoint)
	} else {
		fmt.Println("You have reached the limit")
	}
}

func (this *TokenBucket) run() {

	for true {

		if this.BucketSize == 0 && !this.Blocked {
			go this.refill()
			this.Blocked = true

		} else {

			var endpoint string
			fmt.Print("Enter endpoint: ")
			fmt.Scanln(&endpoint)
			fmt.Println()
			this.get(endpoint)

		}

	}
}

func main() {
	
  /*
  Paramerters:
  1. User ID      <- Different rate limits for different users 
  2. Bucket size  <- Requsts per second
  3. Refill rate  <- Refill rate 
  4. Interval     <- Interval to refill the bucket
  
  Example below:
    User ID: 1       <- User ID  
    Bucket size: 10  <- 10 requests per second
    Refill rate: 5   <- Refill rate 
    Interval: 2      <- Refill the buck in every two seconds
  */
	token := New(1, 10, 5, 2)

	token.run()
}
