package topic

import (
	. "gin_bbs/test/controllers"
	"testing"
)

func TestTopicIndexAPI(t *testing.T) {
	APIGETList(t, "/topics").
		Element(0).
		Object().
		ContainsKey("id").
		ValueEqual("id", 1)

	APIGETList(t, "/topics").
		Element(1).
		Object().
		ContainsKey("id").
		ValueEqual("id", 2)

		// error
	APIGETList(t, "/topics").
		Element(3).
		Object().
		ContainsKey("id").
		ValueEqual("id", 1)
}
