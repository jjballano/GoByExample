package main 

import ("fmt"
		"time")

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2015, 2, 18, 19, 23, 35, 12345, time.UTC)
	p(then)				

	p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))


    p(" - epoch - ")
    secs := now.Unix()
    nanos := now.UnixNano()

    millis := nanos / 1000

    p(secs)
    p(millis)
    p(nanos)

    p(time.Unix(secs,0)) //time from seconds
    p(time.Unix(0,nanos)) //time from nanos

    p(" - time formatting - ")
    p(now.Format(time.RFC3339))

    t1, _ := time.Parse(time.RFC3339,"2012-11-01T22:08:41+00:00")
    p(t1)

    p("Always the same pattern -> Mon Jan 2 15:04:05 MST 2006")
    p(now.Format("3:04PM"))
    p(now.Format("Mon Jan _2 15:04:05 2006"))
    p(now.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    p(t2)

    fmt.Printf("%d/%02d/%02dT%02d:%02d:%02d-00:00\n",
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e := time.Parse(ansic, "8:41PM")
    p(e)
}		