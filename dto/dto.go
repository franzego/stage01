package dto

// model for request
type Req struct {
	Value string `json:"value"`
}

// model for response
type Resp struct {
	Id        string `json:"id"`
	Value     string `json:"value"`
	Props     `json:"props"`
	CreatedAt string `json:"created_at"`
}

// model for properties
type Props struct {
	Length                int         `json:"length"`
	IsPalindrome          bool        `json:"is_palindrome"`
	UniqueCharacters      int         `json:"unique_characters"`
	WordCount             int         `json:"word_count"`
	Sha256Hash            string      `json:"sha256_hash"`
	CharacterFrequencyMap interface{} `json:"character_frequnecy_map"`
}
