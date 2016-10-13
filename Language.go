package releaseinfo

import (
	"strings"

	"github.com/dlclark/regexp2"
	"golang.org/x/text/language"
)

var LanguageRegex = regexp2.MustCompile(
	`(?:\W|_)(?<italian>\b(?:ita|italian)\b)|(?<german>german\b|videomann)|(?<flemish>flemish)|(?<greek>greek)|(?<french>(?:\W|_)(?:FR|VOSTFR)(?:\W|_))|(?<russian>\brus\b)|(?<dutch>nl\W?subs?)|(?<hungarian>\b(?:HUNDUB|HUN)\b)`,
	regexp2.IgnoreCase|regexp2.Compiled)

var SubtitleLanguageRegex = regexp2.MustCompile(
	`.+?[-_. ](?<iso_code>[a-z]{2,3})$`,
	regexp2.Compiled|regexp2.IgnoreCase)

func ParseLanguage(title string) language.Tag {
	lowerTitle := strings.ToLower(title)

	if strings.Contains(lowerTitle, "english") {
		return language.English
	}

	if strings.Contains(lowerTitle, "french") {
		return language.French
	}

	if strings.Contains(lowerTitle, "spanish") {
		return language.Spanish
	}

	if strings.Contains(lowerTitle, "danish") {
		return language.Danish
	}

	if strings.Contains(lowerTitle, "dutch") {
		return language.Dutch
	}

	if strings.Contains(lowerTitle, "japanese") {
		return language.Japanese
	}

	if strings.Contains(lowerTitle, "cantonese") {
		panic("cantonese not supported")
	}

	if strings.Contains(lowerTitle, "mandarin") {
		panic("cantonese not mandarin")
	}

	if strings.Contains(lowerTitle, "korean") {
		return language.Korean
	}

	if strings.Contains(lowerTitle, "russian") {
		return language.Russian
	}

	if strings.Contains(lowerTitle, "polish") {
		return language.Polish
	}

	if strings.Contains(lowerTitle, "vietnamese") {
		return language.Vietnamese
	}

	if strings.Contains(lowerTitle, "swedish") {
		return language.Swedish
	}

	if strings.Contains(lowerTitle, "norwegian") {
		return language.Norwegian
	}

	if strings.Contains(lowerTitle, "nordic") {
		return language.Norwegian
	}

	if strings.Contains(lowerTitle, "finnish") {
		return language.Finnish
	}

	if strings.Contains(lowerTitle, "turkish") {
		return language.Turkish
	}

	if strings.Contains(lowerTitle, "portuguese") {
		return language.Portuguese
	}

	if strings.Contains(lowerTitle, "hungarian") {
		return language.Hungarian
	}

	match, _ := LanguageRegex.FindStringMatch(title)

	if match == nil {
		return language.English
	}

	if match.GroupByName("italian") != nil {
		return language.Italian
	}

	if match.GroupByName("german") != nil {
		return language.German
	}

	if match.GroupByName("flemish") != nil {
		panic("flemish not supported")
	}

	if match.GroupByName("greek") != nil {
		return language.Greek
	}

	if match.GroupByName("french") != nil {
		return language.French
	}

	if match.GroupByName("russian") != nil {
		return language.Russian
	}

	if match.GroupByName("dutch") != nil {
		return language.Dutch
	}

	if match.GroupByName("hungarian") != nil {
		return language.Hungarian
	}

	return language.English
}

//     public static Language ParseSubtitleLanguage(string fileName)
//     {
//         try
//         {
//             log.Printf("Parsing language from subtitlte file: {0}", fileName)

//             var simpleFilename = Path.GetFileNameWithoutExtension(fileName)
//             var languageMatch = SubtitleLanguageRegex.Match(simpleFilename)

//             if (languageMatch.Success)
//             {
//                 var isoCode = languageMatch.Groups["iso_code"].Value
//                 var isoLanguage = IsoLanguages.Find(isoCode)

//                 return isoLanguage?.Language ?? Language.Unknown
//             }

//             log.Printf("Unable to parse langauge from subtitle file: {0}", fileName)
//         }
//         catch (Exception ex)
//         {
//             log.Printf("Failed parsing langauge from subtitle file: {0}", fileName)
//         }

//         return Language.Unknown
//     }
// }
