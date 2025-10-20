package internal

import "github.com/franzego/stage01/dto"

var HardcodedStrings = []dto.Resp{
	{
		Id:    "b1946ac92492d2347c6235b4d2611184",
		Value: "rotator",
		Props: dto.Props{
			Length:                7,
			IsPalindrome:          true,
			UniqueCharacters:      5,
			WordCount:             1,
			Sha256Hash:            "b1946ac92492d2347c6235b4d2611184",
			CharacterFrequencyMap: map[string]int{"r": 2, "o": 1, "t": 2, "a": 1},
		},
		CreatedAt: "2025-10-20T12:00:00Z",
	},
	{
		Id:    "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		Value: "hello world",
		Props: dto.Props{
			Length:                11,
			IsPalindrome:          false,
			UniqueCharacters:      8,
			WordCount:             2,
			Sha256Hash:            "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
			CharacterFrequencyMap: map[string]int{"h": 1, "e": 1, "l": 3, "o": 2, "w": 1, "r": 1, "d": 1},
		},
		CreatedAt: "2025-10-20T12:00:00Z",
	},
}
