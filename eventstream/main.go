package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", events)
	http.ListenAndServe(":8888", nil)
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	tokens := []string{
		"this", "is", "a", "big", "stream", "for", "event", "stream", "test"}

	for _, token := range tokens {
		content := fmt.Sprintf("data:%s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		time.Sleep(time.Millisecond * 420)
	}

}
