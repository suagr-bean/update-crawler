package detect

import "time"
//处理时间
func DealTime(timeStr string) int64 {
	// A slice of time layouts to try parsing.
	// RSS feeds can have many different date formats.
	layouts := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC822,
		time.RFC822Z,
		"Mon, 2 Jan 2006 15:04:05 MST", // Common RSS format
		"2006-01-02 15:04:05.999999999-07:00", // Format from your rss.db
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, timeStr)
		if err == nil {
			return t.Unix()
		}
	}

	// If no layout matches, return the current time as a fallback.
	// In a real application, you might want to log the error here.
	return time.Now().Unix()
}
