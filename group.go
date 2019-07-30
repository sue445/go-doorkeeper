package doorkeeper

// A Group represents doorkeeper group
type Group struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CountryCode  string `json:"country_code"`
	Logo         string `json:"logo"`
	Description  string `json:"description"`
	PublicURL    string `json:"public_url"`
	MembersCount int    `json:"members_count"`
}
