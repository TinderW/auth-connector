/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type GroupInviteToken struct {
	Key
	Attributes GroupInviteTokenAttributes `json:"attributes"`
}
type GroupInviteTokenResponse struct {
	Data     GroupInviteToken `json:"data"`
	Included Included         `json:"included"`
}

type GroupInviteTokenListResponse struct {
	Data     []GroupInviteToken `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustGroupInviteToken - returns GroupInviteToken from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGroupInviteToken(key Key) *GroupInviteToken {
	var groupInviteToken GroupInviteToken
	if c.tryFindEntry(key, &groupInviteToken) {
		return &groupInviteToken
	}
	return nil
}
