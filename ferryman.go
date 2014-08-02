package ferryman

import (
	"log"
	"time"

	"github.com/marconi/rivers"
)

func Run(urgent, delayed rivers.Queue) chan bool {
	// channel to cleanly exit Run
	quit := make(chan bool)

	// we need non-pooled connection since
	// execution happens inside goroutine
	conn := rivers.NewNonPool()

	// watch the delayed queue for jobs ready to
	// be popped and move to urgent queue
	go func() {
		for {
			select {
			case <-quit:
				log.Println("killing Ferryman...")
				conn.Close()
				return
			default:
				jobs, err := delayed.(rivers.DelayedQueue).MultiPop(conn)
				if err != nil {
					log.Println("unable to pop jobs:", err)
					continue
				}

				log.Printf("popped %d jobs\n", len(jobs))
				if len(jobs) > 0 {
					n, err := urgent.(rivers.UrgentQueue).MultiPush(jobs, conn)
					if err != nil {
						log.Println("unable to push jobs:", err)
					} else {
						log.Printf("pushed %d jobs\n", n)
					}
				}

				time.Sleep(1 * time.Second)
			}
		}
	}()

	return quit
}
