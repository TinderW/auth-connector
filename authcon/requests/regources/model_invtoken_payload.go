/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type InvtokenPayload struct {
	Key
	Attributes InvtokenPayloadAttributes `json:"attributes"`
}
type InvtokenPayloadResponse struct {
	Data     InvtokenPayload `json:"data"`
	Included Included        `json:"included"`
}

type InvtokenPayloadListResponse struct {
	Data     []InvtokenPayload `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustInvtokenPayload - returns InvtokenPayload from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustInvtokenPayload(key Key) *InvtokenPayload {
	var invtokenPayload InvtokenPayload
	if c.tryFindEntry(key, &invtokenPayload) {
		return &invtokenPayload
	}
	return nil
}
