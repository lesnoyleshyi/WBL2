package t1

import (
	"github.com/beevik/ntp"
	"log"
	"time"
)

const ntpHost = "0.beevik-ntp.pool.ntp.org"
const formatStr = time.RFC3339

func getCurrentTime() {
	t, err := ntp.Time(ntpHost)
	if err != nil {
		log.Fatal("error receiving time from ntp server:", err)
	}

	log.Printf("current time is %s", t.Format(formatStr))
}
