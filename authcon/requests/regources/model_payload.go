/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type Payload struct {
	Key
	Attributes PayloadAttributes `json:"attributes"`
}
type PayloadResponse struct {
	Data     Payload  `json:"data"`
	Included Included `json:"included"`
}

type PayloadListResponse struct {
	Data     []Payload `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustPayload - returns Payload from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPayload(key Key) *Payload {
	var payload Payload
	if c.tryFindEntry(key, &payload) {
		return &payload
	}
	return nil
}
