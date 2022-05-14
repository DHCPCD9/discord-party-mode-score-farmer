package main

import (
	"flag"
	"fmt"
	"github.com/micmonay/keybd_event"
	"log"
	"runtime"
	"time"
)

func main() {

	score := flag.Int("score", 0, "Your score")

	flag.Parse()

	kb, err := keybd_event.NewKeyBonding()

	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb.SetKeys(keybd_event.VK_A)

	for i := 0; i < 4; i++ {
		log.Println(fmt.Sprintf("Starting in %d", 4-i))
		time.Sleep(1 * time.Second)
	}

	err = kb.Launching()

	if err != nil {
		panic(err)
	}

	for i := 0; i < *score-1; i++ {
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}

	kb.SetKeys(keybd_event.VK_ENTER)
	kb.Press()

}
