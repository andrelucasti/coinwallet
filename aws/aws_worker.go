package aws

import (
	"coinwallet/wallet"
	"sync"
)

var results = make(chan Result, 2)

func worker(wg *sync.WaitGroup) {
	awsMessage := wallet.Consumer()

	for _, message := range awsMessage {
		result := Result{
			OutputMessage: OutputMessage{
				MessageId: message.MessageId,
				Body:      message.Body,
			},
		}
		results <- result
	}
	wg.Done()
}

func CreateWorkerPool() {
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 1; i++ {
		go worker(&wg)
	}

	wg.Wait()

	close(results)
}
