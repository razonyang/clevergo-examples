package main

import (
	"net/http"
	"strconv"
	"time"

	"clevergo.tech/clevergo"
)

func main() {
	app := clevergo.New()
	// app.ShutdownTimeout = 5*time.Second // default
	// app.ShutdownSignals = []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM} // default
	app.Get("/sleep", sleep)
	app.Run(":8080")
}

func sleep(c *clevergo.Context) (err error) {
	max := 0
	if duration := c.QueryParam("duration"); duration != "" {
		max, err = strconv.Atoi(duration)
		if err != nil {
			return
		}
	}
	for i := 0; i < max; i++ {
		time.Sleep(time.Second)
	}
	return c.String(http.StatusOK, "done")
}
