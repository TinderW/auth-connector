/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type Tokens struct {
	Key
	Attributes TokensAttributes `json:"attributes"`
}
type TokensResponse struct {
	Data     Tokens   `json:"data"`
	Included Included `json:"included"`
}

type TokensListResponse struct {
	Data     []Tokens `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustTokens - returns Tokens from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTokens(key Key) *Tokens {
	var tokens Tokens
	if c.tryFindEntry(key, &tokens) {
		return &tokens
	}
	return nil
}
