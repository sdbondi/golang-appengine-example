package app

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time

	context appengine.Context
}

func (self *Greeting) Save() (*datastore.Key, error) {
	c := self.context
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(self.context))
	_, err := datastore.Put(c, key, self)
	return key, err
}

func NewGreeting(c appengine.Context) *Greeting {
	return &Greeting{"", "", time.Now(), c}
}

func LoadAllGreetings(c appengine.Context, buffer *[]Greeting) error {
	limit := cap(*buffer)
	if limit > 100 {
		limit = 100
	}
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(limit)
	_, err := q.GetAll(c, buffer)
	return err
}

func guestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
