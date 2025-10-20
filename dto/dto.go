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

// Response to fileter struct
type Respfilter struct {
	Data    Resp      `json:"data"`
	Count   int       `json:"count"`
	Filters Filtersss `json:"filters"`
}

// Filters struct
type Filtersss struct {
	IsPalindrome bool   `json:"is_palindrome"`
	MinLength    int    `json:"min_length"`
	MaxLength    int    `json:"max_length"`
	WordCount    int    `json:"word_count"`
	ContainsChar string `json:"contains_char"`
}
