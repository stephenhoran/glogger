# Glogger
#### A Simple, Mostly Compatible, Log Level Logger for Go.

Glogger is a simple and fast log level based logger that is mostly compatible with the standard library logger. If you 
wish to leverage glogger, note that new lines are assumed in a log statement. All previous `log.Println` statements will
not work.

```
package main

import (
	log "glogger"
)

func main() {
	log.SetLogLevel(log.DebugLevel)
	log.SetPrefix("Glogger")

	log.Info("A Info Message!")
	log.Warnf("A Warning %s message", "formatted")
}
```
```
2021-09-28T10:57:03-04:00: [INFO] Glogger A Info Message!
2021-09-28T10:57:03-04:00: [WARN] Glogger A Warning formatted message
```

Glogger supports timestamp formatting as well. Timestamps are defaulted to `RFC3339` time.
