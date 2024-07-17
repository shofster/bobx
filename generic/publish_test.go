package generic

import (
	"fmt"
	"testing"
	"time"
)

type publishMessage struct {
	op   int
	text string
}

func TestPublish(t *testing.T) {
	testPublisher := NewPublisher[publishMessage](10)

	c1 := 0
	pc1 := make(chan publishMessage)
	testPublisher.Register(pc1)
	go func() {
		for range pc1 {
			c1++
		}
	}()

	c2 := 0
	pc2 := make(chan publishMessage)
	testPublisher.Register(pc2)
	go func() {
		for range pc2 {
			c2++
		}
	}()

	for i := 0; i < 10; i++ {
		testPublisher.Submit(publishMessage{op: i, text: fmt.Sprintf("verbage-%d", i)})
		time.Sleep(1 * time.Millisecond)
	}
	time.Sleep(1000 * time.Millisecond)
	testPublisher.Close()
	if c1 != 10 || c2 != 10 {
		t.Errorf("Expected c1 10: got %d", c1)
		t.Errorf("Expected c2 10: got %d", c2)
	}
}

func TestPublishRetry(t *testing.T) {
	testPublisher := NewPublisher[publishMessage](2)

	wait := make(chan struct{}, 1)

	c1 := 0
	pc1 := make(chan publishMessage)
	testPublisher.Register(pc1)
	go func() {
		for range pc1 {
			c1++
		}
	}()

	c2 := 0
	pc2 := make(chan publishMessage)
	testPublisher.Register(pc2)
	go func() {
		for range pc2 {
			c2++
		}
	}()

	attempts := 200
	failures := 0
	retries := 0
	// send a message every few Microseconds
	r := 200
	rate := 200 * time.Millisecond
	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		for i := 0; i < attempts; i++ {
			if !testPublisher.Submit(publishMessage{op: i, text: fmt.Sprintf("verbage-%d", i)}) {
				failures++
				// sleep time can be adjusted
				time.Sleep(rate)
				retries++
				testPublisher.Submit(publishMessage{op: i, text: fmt.Sprintf("verbage-%d", i)})
			}
		}
	}()
	start := time.Now()
	<-wait
	end := time.Now()

	time.Sleep(3000 * time.Millisecond)
	t.Logf("TestPublishRetry: %d messages @%dMS; failures %d, retries %d, c1 %d, c2 %d",
		attempts, r, failures, retries, c1, c2)
	t.Logf("TestPublishRetry:   elapsed %v", end.Sub(start))
	if failures == 0 || c1 != attempts || c2 != attempts {
		t.Errorf("Expected some Failures : got %d", failures)
		t.Errorf("Expected c1 %d: got %d", attempts, c1)
		t.Errorf("Expected c2 %d: got %d", attempts, c2)
	}
	testPublisher.Close()
}
