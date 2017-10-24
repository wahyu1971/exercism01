package gigasecond

// import path for the time package from the standard library
import "time"

const (
	testVersion = 4
	gigasecond  = 1e9 * time.Second // bukan 10^9 , weeeeek
)

func AddGigasecond(t time.Time) (result time.Time) {
	result = t.Add(gigasecond)
	return
}
