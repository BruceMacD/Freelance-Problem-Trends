package analyzer

// GetCommonWords is used to filter common words from results
func GetCommonWords() map[string]bool {
	return map[string]bool{
		"a":           true,
		"and":         true,
		"the":         true,
		"for":         true,
		"i":           true,
		"in":          true,
		"need":        true,
		"of":          true,
		"with":        true,
		"to":          true,
		"be":          true,
		"is":          true,
		"have":        true,
		"":            true,
		"on":          true,
		"will":        true,
		"you":         true,
		"an":          true,
		"are":         true,
		"my":          true,
		"that":        true,
		"can":         true,
		"we":          true,
		"looking":     true,
		"from":        true,
		"it":          true,
		"our":         true,
		"this":        true,
		"want":        true,
		"app":         true, // this one might be of interest
		"like":        true,
		"or":          true,
		"as":          true,
		"some":        true,
		"am":          true,
		"me":          true,
		"who":         true,
		"de":          true,
		"project":     true,
		"-":           true,
		"build":       true,
		"would":       true,
		"work":        true,
		"create":      true,
		"[login":      true,
		"someone":     true,
		"new":         true,
		"url]":        true,
		"help":        true,
		"if":          true,
		"your":        true,
		"please":      true,
		"should":      true,
		"all":         true,
		"file":        true,
		"not":         true,
		"make":        true,
		"do":          true,
		"company":     true,
		"business":    true,
		"but":         true,
		"site":        true,
		"using":       true,
		"so":          true,
		"&":           true,
		"which":       true,
		"must":        true,
		"one":         true,
		"2":           true,
		"by":          true,
		"view":        true,
		"page":        true,
		"list":        true,
		"email":       true,
		"per":         true,
		"about":       true,
		"at":          true,
		"more":        true,
		"online":      true,
		"attached":    true,
		"up":          true,
		"also":        true,
		"see":         true,
		"write":       true,
		"use":         true,
		"get":         true,
		"good":        true,
		"has":         true,
		",":           true,
		"application": true,
		"marketing":   true,
		"into":        true,
		"name":        true,
		"simple":      true,
		"para":        true,
		"details":     true,
		"provide":     true,
		"product":     true,
		"long":        true,
		"just":        true,
		"time":        true,
		"each":        true,
		"needed":      true,
		"able":        true,
		"...":         true,
		"y":           true,
		"any":         true,
		"/":           true,
		"software":    true,
		"+":           true,
		"en":          true,
		"only":        true,
		"based":       true,
	}
}

// GetSoftSkillWords is used to filter soft skills from results
func GetSoftSkillWords() map[string]bool {
	return map[string]bool{
		"design":     true,
		"designer":   true,
		"data":       true,
		"web":        true,
		"expert":     true,
		"developer":  true,
		"experience": true,
	}
}

// GetSkillWords is used to filter common words from results
func GetSkillWords() map[string]bool {
	return map[string]bool{
		"logo":      true,
		"excel":     true,
		"video":     true,
		"android":   true,
		"content":   true,
		"wordpress": true,
		"entry":     true,
		"mobile":    true,
		"google":    true,
		"english":   true,
	}
}
