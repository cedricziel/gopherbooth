package booth

import (
	"appengine"
	"appengine/datastore"
)

// boothKey returns the key used for all booth entries.
func Key(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Booth", "default_booth", 0, nil)
}
