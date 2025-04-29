/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateTokens struct {
	Key
	Attributes CreateTokensAttributes `json:"attributes"`
}
type CreateTokensResponse struct {
	Data     CreateTokens `json:"data"`
	Included Included     `json:"included"`
}

type CreateTokensListResponse struct {
	Data     []CreateTokens `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustCreateTokens - returns CreateTokens from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateTokens(key Key) *CreateTokens {
	var createTokens CreateTokens
	if c.tryFindEntry(key, &createTokens) {
		return &createTokens
	}
	return nil
}
