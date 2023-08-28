package constant

import "time"

var SessionCookieDuration = 30 * time.Minute

func GenerateSessionExpiresAt() time.Time {
	return time.Now().Add(SessionCookieDuration)
}
