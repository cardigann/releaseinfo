package releaseinfo

// using System
// using System.Collections.Generic
// using System.IO
// using System.Linq
// using System.Text.RegularExpressions
// using NLog
// using NzbDrone.Common.Extensions
// using NzbDrone.Common.Instrumentation
// using NzbDrone.Core.Parser.Model
// using NzbDrone.Core.Tv

var ReportTitleRegex = []regexp2.Regexp{
	//Anime - Absolute Episode Number + Title + Season+Episode
	//Todo: This currently breaks series that start with numbers
	// regexp2.MustCompile(`^(?:(?<absoluteepisode>\d{2,3})(?:_|-|\s|\.)+)+(?<title>.+?)(?:\W|_)+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+)`,
	// regexp2.IgnoreCase | regexp2.Compiled),

	//Multi-Part episodes without a title (S01E05.S01E06)
	regexp2.MustCompile(`^(?:\W*S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:[ex]){1,2}(?<episode>\d{1,3}(?!\d+)))+){2,}`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes without a title, Single (S01E05, 1x05) AND Multi (S01E04E05, 1x04x05, etc)
	regexp2.MustCompile(`^(?:S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:\-|[ex]|\W[ex]|_){1,2}(?<episode>\d{2,3}(?!\d+)))+)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title Absolute Episode Number + Season+Episode
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\](?:_|-|\s|\.)?)(?<title>.+?)(?:(?:[-_\W](?<![()\[!]))+(?<absoluteepisode>\d{2,3}))+(?:_|-|\s|\.)+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+).*?(?<hash>[(\[]\w{8}[)\]])?(?:$|\.)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title Season+Episode + Absolute Episode Number
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\](?:_|-|\s|\.)?)(?<title>.+?)(?:[-_\W](?<![()\[!]))+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+)(?:(?:_|-|\s|\.)+(?<absoluteepisode>(?<!\d+)\d{2,3}(?!\d+)))+.*?(?<hash>\[\w{8}\])?(?:$|\.)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title Season+Episode
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\](?:_|-|\s|\.)?)(?<title>.+?)(?:[-_\W](?<![()\[!]))+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+)(?:\s|\.).*?(?<hash>\[\w{8}\])?(?:$|\.)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title with trailing number Absolute Episode Number
	regexp2.MustCompile(`^\[(?<subgroup>.+?)\][-_. ]?(?<title>[^-]+?\d+?)[-_. ]+(?:[-_. ]?(?<absoluteepisode>\d{3}(?!\d+)))+(?:[-_. ]+(?<special>special|ova|ovd))?.*?(?<hash>\[\w{8}\])?(?:$|\.mkv)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title - Absolute Episode Number
	regexp2.MustCompile(`^\[(?<subgroup>.+?)\][-_. ]?(?<title>.+?)(?:[. ]-[. ](?<absoluteepisode>\d{2,3}(?!\d+|[-])))+(?:[-_. ]+(?<special>special|ova|ovd))?.*?(?<hash>\[\w{8}\])?(?:$|\.mkv)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title Absolute Episode Number
	regexp2.MustCompile(`^\[(?<subgroup>.+?)\][-_. ]?(?<title>.+?)[-_. ]+(?:[-_. ]?(?<absoluteepisode>\d{2,3}(?!\d+)))+(?:[-_. ]+(?<special>special|ova|ovd))?.*?(?<hash>\[\w{8}\])?(?:$|\.mkv)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - Title Season EpisodeNumber + Absolute Episode Number [SubGroup]
	regexp2.MustCompile(`^(?<title>.+?)(?:[-_\W](?<![()\[!]))+(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:[ex]|\W[ex]){1,2}(?<episode>\d{2}(?!\d+)))+).+?(?:[-_. ]?(?<absoluteepisode>\d{3}(?!\d+)))+.+?\[(?<subgroup>.+?)\](?:$|\.mkv)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - Title Absolute Episode Number [SubGroup]
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:_|-|\s|\.)+(?<absoluteepisode>\d{3}(?!\d+)))+(?:.+?)\[(?<subgroup>.+?)\].*?(?<hash>\[\w{8}\])?(?:$|\.)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - Title Absolute Episode Number [Hash]
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:_|-|\s|\.)+(?<absoluteepisode>\d{2,3}(?!\d+)))+(?:[-_. ]+(?<special>special|ova|ovd))?[-_. ]+.*?(?<hash>\[\w{8}\])(?:$|\.)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes with airdate AND season/episode number, capture season/epsiode only
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airdate>\d{4}\W+[0-1][0-9]\W+[0-3][0-9])(?!\W+[0-3][0-9])[-_. ](?:s?(?<season>(?<!\d+)(?:\d{1,2})(?!\d+)))(?:[ex](?<episode>(?<!\d+)(?:\d{1,3})(?!\d+)))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes with airdate AND season/episode number
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airyear>\d{4})\W+(?<airmonth>[0-1][0-9])\W+(?<airday>[0-3][0-9])(?!\W+[0-3][0-9]).+?(?:s?(?<season>(?<!\d+)(?:\d{1,2})(?!\d+)))(?:[ex](?<episode>(?<!\d+)(?:\d{1,3})(?!\d+)))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Multi-episode Repeated (S01E05 - S01E06, 1x05 - 1x06, etc)
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()\[!]))+S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:(?:[ex]|[-_. ]e){1,2}(?<episode>\d{1,3}(?!\d+)))+){2,}`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes with a title, Single episodes (S01E05, 1x05, etc) & Multi-episode (S01E05E06, S01E05-06, S01E05 E06, etc) **
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()\[!]))+S?(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:[ex]|\W[ex]|_){1,2}(?<episode>\d{2,3}(?!\d+))(?:(?:\-|[ex]|\W[ex]|_){1,2}(?<episode>\d{2,3}(?!\d+)))*)\W?(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Mini-Series, treated as season 1, episodes are labelled as Part01, Part 01, Part.1
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+(?:(?:Part\W?|(?<!\d+\W+)e)(?<episode>\d{1,2}(?!\d+)))+)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Mini-Series, treated as season 1, episodes are labelled as Part One/Two/Three/...Nine, Part.One, Part_One
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+(?:Part[-._ ](?<episode>One|Two|Three|Four|Five|Six|Seven|Eight|Nine)(?>[-._ ])))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Mini-Series, treated as season 1, episodes are labelled as XofY
	regexp2.MustCompile(`^(?<title>.+?)(?:\W+(?:(?<episode>(?<!\d+)\d{1,2}(?!\d+))of\d+)+)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Supports Season 01 Episode 03
	regexp2.MustCompile(`(?:.*(?:\""|^))(?<title>.*?)(?:[-_\W](?<![()\[]))+(?:\W?Season\W?)(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:\W|_)+(?:Episode\W)(?:[-_. ]?(?<episode>(?<!\d+)\d{1,2}(?!\d+)))+`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Multi-episode release with no space between series title and season (S01E11E12)
	regexp2.MustCompile(`(?:.*(?:^))(?<title>.*?)(?:\W?|_)S(?<season>(?<!\d+)\d{2}(?!\d+))(?:E(?<episode>(?<!\d+)\d{2}(?!\d+)))+`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Multi-episode with single episode numbers (S6.E1-E2, S6.E1E2, S6E1E2, etc)
	regexp2.MustCompile(`^(?<title>.+?)[-_. ]S(?<season>(?<!\d+)(?:\d{1,2}|\d{4})(?!\d+))(?:[-_. ]?[ex]?(?<episode>(?<!\d+)\d{1,2}(?!\d+)))+`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Single episode season or episode S1E1 or S1-E1
	regexp2.MustCompile(`(?:.*(?:\""|^))(?<title>.*?)(?:\W?|_)S(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:\W|_)?E(?<episode>(?<!\d+)\d{1,2}(?!\d+))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//3 digit season S010E05
	regexp2.MustCompile(`(?:.*(?:\""|^))(?<title>.*?)(?:\W?|_)S(?<season>(?<!\d+)\d{3}(?!\d+))(?:\W|_)?E(?<episode>(?<!\d+)\d{1,2}(?!\d+))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//5 digit episode number with a title
	regexp2.MustCompile(`^(?:(?<title>.+?)(?:_|-|\s|\.)+)(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+)))(?:(?:\-|[ex]|\W[ex]|_){1,2}(?<episode>(?<!\d+)\d{5}(?!\d+)))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//5 digit multi-episode with a title
	regexp2.MustCompile(`^(?:(?<title>.+?)(?:_|-|\s|\.)+)(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+)))(?:(?:[-_. ]{1,3}ep){1,2}(?<episode>(?<!\d+)\d{5}(?!\d+)))+`,
		regexp2.IgnoreCase|regexp2.Compiled),

	// Separated season and episode numbers S01 - E01
	regexp2.MustCompile(`^(?<title>.+?)(?:_|-|\s|\.)+S(?<season>\d{2}(?!\d+))(\W-\W)E(?<episode>(?<!\d+)\d{2}(?!\d+))(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Season only releases
	regexp2.MustCompile(`^(?<title>.+?)\W(?:S|Season)\W?(?<season>\d{1,2}(?!\d+))(\W+|_|$)(?<extras>EXTRAS|SUBPACK)?(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//4 digit season only releases
	regexp2.MustCompile(`^(?<title>.+?)\W(?:S|Season)\W?(?<season>\d{4}(?!\d+))(\W+|_|$)(?<extras>EXTRAS|SUBPACK)?(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes with a title and season/episode in square brackets
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()\[!]))+\[S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]|\W[ex]|_){1,2}(?<episode>(?<!\d+)\d{2}(?!\d+|i|p)))+\])\W?(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Supports 103/113 naming
	regexp2.MustCompile(`^(?<title>.+?)?(?:(?:[-_\W](?<![()\[!]))+(?<season>(?<!\d+)[1-9])(?<episode>[1-9][0-9]|[0][1-9])(?![a-z]|\d+))+`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes with airdate
	regexp2.MustCompile(`^(?<title>.+?)?\W*(?<airyear>\d{4})\W+(?<airmonth>[0-1][0-9])\W+(?<airday>[0-3][0-9])(?!\W+[0-3][0-9])`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Supports 1103/1113 naming
	regexp2.MustCompile(`^(?<title>.+?)?(?:(?:[-_\W](?<![()\[!]))*(?<season>(?<!\d+|\(|\[|e|x)\d{2})(?<episode>(?<!e|x)\d{2}(?!p|i|\d+|\)|\]|\W\d+)))+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//4 digit episode number
	//Episodes without a title, Single (S01E05, 1x05) AND Multi (S01E04E05, 1x04x05, etc)
	regexp2.MustCompile(`^(?:S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]|\W[ex]|_){1,2}(?<episode>\d{4}(?!\d+|i|p)))+)(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//4 digit episode number
	//Episodes with a title, Single episodes (S01E05, 1x05, etc) & Multi-episode (S01E05E06, S01E05-06, S01E05 E06, etc)
	regexp2.MustCompile(`^(?<title>.+?)(?:(?:[-_\W](?<![()\[!]))+S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]|\W[ex]|_){1,2}(?<episode>\d{4}(?!\d+|i|p)))+)\W?(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Episodes with single digit episode number (S01E1, S01E5E6, etc)
	regexp2.MustCompile(`^(?<title>.*?)(?:(?:[-_\W](?<![()\[!]))+S?(?<season>(?<!\d+)\d{1,2}(?!\d+))(?:(?:\-|[ex]){1,2}(?<episode>\d{1}))+)+(\W+|_|$)(?!\\)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//iTunes Season 1\05 Title (Quality).ext
	regexp2.MustCompile(`^(?:Season(?:_|-|\s|\.)(?<season>(?<!\d+)\d{1,2}(?!\d+)))(?:_|-|\s|\.)(?<episode>(?<!\d+)\d{1,2}(?!\d+))`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - Title Absolute Episode Number (e66)
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:(?:_|-|\s|\.)+(?:e|ep)(?<absoluteepisode>\d{2,3}))+.*?(?<hash>\[\w{8}\])?(?:$|\.)`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - [SubGroup] Title Episode Absolute Episode Number ([SubGroup] Series Title Episode 01)
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)[-_. ](?:Episode)(?:[-_. ]+(?<absoluteepisode>(?<!\d+)\d{2,3}(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - Title Absolute Episode Number
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:[-_. ]+(?<absoluteepisode>(?<!\d+)\d{2,3}(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Anime - Title {Absolute Episode Number}
	regexp2.MustCompile(`^(?:\[(?<subgroup>.+?)\][-_. ]?)?(?<title>.+?)(?:(?:[-_\W](?<![()\[!]))+(?<absoluteepisode>(?<!\d+)\d{2,3}(?!\d+)))+(?:_|-|\s|\.)*?(?<hash>\[.{8}\])?(?:$|\.)?`,
		regexp2.IgnoreCase|regexp2.Compiled),

	//Extant, terrible multi-episode naming (extant.10708.hdtv-lol.mp4)
	regexp2.MustCompile(`^(?<title>.+?)[-_. ](?<season>[0]?\d?)(?:(?<episode>\d{2}){2}(?!\d+))[-_. ]`,
		regexp2.IgnoreCase|regexp2.Compiled),
}

var RejectHashedReleasesRegex = []regexp2.Regexp{
	// Generic match for md5 and mixed-case hashes.
	regexp2.MustCompile(`^[0-9a-zA-Z]{32}`, regexp2.Compiled),

	// Generic match for shorter lower-case hashes.
	regexp2.MustCompile(`^[a-z0-9]{24}$`, regexp2.Compiled),

	// Format seen on some NZBGeek releases
	// Be very strict with these coz they are very close to the valid 101 ep numbering.
	regexp2.MustCompile(`^[A-Z]{11}\d{3}$`, regexp2.Compiled),
	regexp2.MustCompile(`^[a-z]{12}\d{3}$`, regexp2.Compiled),

	//Backup filename (Unknown origins)
	regexp2.MustCompile(`^Backup_\d{5,}S\d{2}-\d{2}$"`, regexp2.Compiled),

	//123 - Started appearing December 2014
	regexp2.MustCompile(`^123$"`, regexp2.Compiled),

	//abc - Started appearing January 2015
	regexp2.MustCompile(`^abc$`, regexp2.Compiled|regexp2.IgnoreCase),

	//b00bs - Started appearing January 2015
	regexp2.MustCompile(`^b00bs$`, regexp2.Compiled|regexp2.IgnoreCas),
}

//Regex to detect whether the title was reversed.
var ReversedTitleRegex = regexp2.MustCompile(`[-._ ](p027|p0801|\d{2}E\d{2}S)[-._ ]`, regexp2.Compiled)

var NormalizeRegex = regexp2.MustCompile(`((?:\b|_)(?<!^)(a(?!$)|an|the|and|or|of)(?:\b|_))|\W|_`, regexp2.IgnoreCase|regexp2.Compiled)

var FileExtensionRegex = regexp2.MustCompile(`\.[a-z0-9]{2,4}$`, regexp2.IgnoreCase|regexp2.Compiled)

var SimpleTitleRegex = regexp2.MustCompile(`(?:480[ip]|720[ip]|1080[ip]|[xh][\W_]?26[45]|DD\W?5\W1|[<>?*:|]|848x480|1280x720|1920x1080|(8|10)b(it)?)\s*`, regexp2.IgnoreCase|regexp2.Compiled)

var WebsitePrefixRegex = regexp2.MustCompile(`^\[\s*[a-z]+(\.[a-z]+)+\s*\][- ]*`, regexp2.IgnoreCase|regexp2.Compiled)

var AirDateRegex = regexp2.MustCompile(`^(.*?)(?<!\d)((?<airyear>\d{4})[_.-](?<airmonth>[0-1][0-9])[_.-](?<airday>[0-3][0-9])|(?<airmonth>[0-1][0-9])[_.-](?<airday>[0-3][0-9])[_.-](?<airyear>\d{4}))(?!\d)`, regexp2.IgnoreCase|regexp2.Compiled)

var SixDigitAirDateRegex = regexp2.MustCompile(`(?<=[_.-])(?<airdate>(?<!\d)(?<airyear>[1-9]\d{1})(?<airmonth>[0-1][0-9])(?<airday>[0-3][0-9]))(?=[_.-])`, regexp2.IgnoreCase|regexp2.Compiled)

var CleanReleaseGroupRegex = regexp2.MustCompile(`^(.*?[-._ ](S\d+E\d+)[-._ ])|(-(RP|1|NZBGeek|Obfuscated|sample))+$`, regexp2.IgnoreCase|regexp2.Compiled)

var CleanTorrentSuffixRegex = regexp2.MustCompile(`\[(?:ettv|rartv|rarbg|cttv)\]$`, regexp2.IgnoreCase|regexp2.Compiled)

var ReleaseGroupRegex = regexp2.MustCompile(`-(?<releasegroup>[a-z0-9]+)(?<!WEB-DL|480p|720p|1080p|2160p)(?:\b|[-._ ])`, regexp2.IgnoreCase|regexp2.Compiled)

var AnimeReleaseGroupRegex = regexp2.MustCompile(`^(?:\[(?<subgroup>(?!\s).+?(?<!\s))\](?:_|-|\s|\.)?)`, regexp2.IgnoreCase|regexp2.Compiled)

var YearInTitleRegex = regexp2.MustCompile(`^(?<title>.+?)(?:\W|_)?(?<year>\d{4})`, regexp2.IgnoreCase|regexp2.Compiled)

var WordDelimiterRegex = regexp2.MustCompile(`(\s|\.|,|_|-|=|\|)+`, regexp2.Compiled)
var PunctuationRegex = regexp2.MustCompile(`[^\w\s]`, regexp2.Compiled)
var CommonWordRegex = regexp2.MustCompile(`\b(a|an|the|and|or|of)\b\s?`, regexp2.IgnoreCase|regexp2.Compiled)
var SpecialEpisodeWordRegex = regexp2.MustCompile(`\b(part|special|edition|christmas)\b\s?`, regexp2.IgnoreCase|regexp2.Compiled)
var DuplicateSpacesRegex = regexp2.MustCompile(`\s{2,}`, regexp2.Compiled)

var RequestInfoRegex = regexp2.MustCompile(`\[.+?\]`, regexp2.Compiled)

var Numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ParsePath(path string) ParsedEpisodeInfo {
	var fileInfo = NewFileInfo(path)

	var result = ParseTitle(fileInfo.Name)

	if result == nil {
		Logger.Debug("Attempting to parse episode info using directory and file names. {0}", fileInfo.Directory.Name)
		result = ParseTitle(fileInfo.Directory.Name + " " + fileInfo.Name)
	}

	if result == nil {
		Logger.Debug("Attempting to parse episode info using directory name. {0}", fileInfo.Directory.Name)
		result = ParseTitle(fileInfo.Directory.Name + fileInfo.Extension)
	}

	return result
}

func ParseTitle(title string) ParsedEpisodeInfo {
	// try
	// {
	if !ValidateBeforeParsing(title) {
		return nil
	}

	Logger.Debug("Parsing string '{0}'", title)

	if ReversedTitleRegex.IsMatch(title) {
		var titleWithoutExtension = RemoveFileExtension(title).ToCharArray()
		Array.Reverse(titleWithoutExtension)

		title = string(titleWithoutExtension) + title.Substring(titleWithoutExtension.Length)

		Logger.Debug("Reversed name detected. Converted to '{0}'", title)
	}

	var simpleTitle = SimpleTitleRegex.Replace(title, string.Empty)

	simpleTitle = RemoveFileExtension(simpleTitle)

	// TODO: Quick fix stripping [url] - prefixes.
	simpleTitle = WebsitePrefixRegex.Replace(simpleTitle, string.Empty)

	simpleTitle = CleanTorrentSuffixRegex.Replace(simpleTitle, string.Empty)

	var airDateMatch = AirDateRegex.Match(simpleTitle)
	if airDateMatch.Success {
		simpleTitle = airDateMatch.Groups[1].Value + airDateMatch.Groups["airyear"].Value + "." + airDateMatch.Groups["airmonth"].Value + "." + airDateMatch.Groups["airday"].Value
	}

	var sixDigitAirDateMatch = SixDigitAirDateRegex.Match(simpleTitle)
	if sixDigitAirDateMatch.Success {
		var airYear = sixDigitAirDateMatch.Groups["airyear"].Value
		var airMonth = sixDigitAirDateMatch.Groups["airmonth"].Value
		var airDay = sixDigitAirDateMatch.Groups["airday"].Value

		if airMonth != "00" || airDay != "00" {
			var fixedDate = string.Format("20{0}.{1}.{2}", airYear, airMonth, airDay)

			simpleTitle = simpleTitle.Replace(sixDigitAirDateMatch.Groups["airdate"].Value, fixedDate)
		}
	}

	for _, regex := range ReportTitleRegex {
		var match = regex.Matches(simpleTitle)

		if match.Count != 0 {
			Logger.Trace(regex)
			// try
			// {
			var result = ParseMatchCollection(match)

			if result != nil {
				if result.FullSeason && title.ContainsIgnoreCase("Special") {
					result.FullSeason = false
					result.Special = true
				}

				result.Language = LanguageParser.ParseLanguage(title)
				Logger.Debug("Language parsed: {0}", result.Language)

				result.Quality = QualityParser.ParseQuality(title)
				Logger.Debug("Quality parsed: {0}", result.Quality)

				result.ReleaseGroup = ParseReleaseGroup(title)

				var subGroup = GetSubGroup(match)
				if !subGroup.IsNullOrWhiteSpace() {
					result.ReleaseGroup = subGroup
				}

				Logger.Debug("Release Group parsed: {0}", result.ReleaseGroup)

				result.ReleaseHash = GetReleaseHash(match)
				if !result.ReleaseHash.IsNullOrWhiteSpace() {
					Logger.Debug("Release Hash parsed: {0}", result.ReleaseHash)
				}

				return result
			}
			// }
			// catch (InvalidDateException ex)
			// {
			//     Logger.Debug(ex, ex.Message)
			//     break
			// }
		}
	}
	// }
	// catch (Exception e)
	// {
	//     if (!title.ToLower().Contains("password") && !title.ToLower().Contains("yenc"))
	//         Logger.Error(e, "An error has occurred while trying to parse " + title)
	// }

	Logger.Debug("Unable to parse {0}", title)
	return nil
}

func ParseSeriesName(title string) string {
	Logger.Debug("Parsing string '{0}'", title)

	var parseResult = ParseTitle(title)

	if parseResult == nil {
		return CleanSeriesTitle(title)
	}

	return parseResult.SeriesTitle
}

func CleanSeriesTitle(title string) string {
	var number int64 = 0

	//If Title only contains numbers return it as is.
	if long.TryParse(title, number) {
		return title
	}

	return NormalizeRegex.Replace(title, string.Empty).ToLower().RemoveAccent()
}

func NormalizeEpisodeTitle(title string) string {
	title = SpecialEpisodeWordRegex.Replace(title, string.Empty)
	title = PunctuationRegex.Replace(title, " ")
	title = DuplicateSpacesRegex.Replace(title, " ")

	return title.Trim().ToLower()
}

func NormalizeTitle(title string) string {
	title = WordDelimiterRegex.Replace(title, " ")
	title = PunctuationRegex.Replace(title, string.Empty)
	title = CommonWordRegex.Replace(title, string.Empty)
	title = DuplicateSpacesRegex.Replace(title, " ")

	return title.Trim().ToLower()
}

func ParseReleaseGroup(title string) string {
	title = title.Trim()
	title = RemoveFileExtension(title)
	title = WebsitePrefixRegex.Replace(title, "")

	var animeMatch = AnimeReleaseGroupRegex.Match(title)

	if animeMatch.Success {
		return animeMatch.Groups["subgroup"].Value
	}

	title = CleanReleaseGroupRegex.Replace(title, "")

	var matches = ReleaseGroupRegex.Matches(title)

	if matches.Count != 0 {
		// var group = matches.OfType<Match>().Last().Groups["releasegroup"].Value
		var groupIsNumeric int

		if int.TryParse(group, groupIsNumeric) {
			return nil
		}

		return group
	}

	return nil
}

func RemoveFileExtension(title string) string {
	title = FileExtensionRegex.Replace(title, func(m) {
		var extension = m.Value.ToLower()
		// if (MediaFiles.MediaFileExtensions.Extensions.Contains(extension) || new[] { ".par2", ".nzb" }.Contains(extension)) {
		//     return string.Empty
		// }
		return m.Value
	})

	return title
}

func getSeriesTitleInfo(title string) SeriesTitleInfo {
	var seriesTitleInfo = NewSeriesTitleInfo()
	seriesTitleInfo.Title = title

	var match = YearInTitleRegex.Match(title)

	if !match.Success {
		seriesTitleInfo.TitleWithoutYear = title
	} else {
		seriesTitleInfo.TitleWithoutYear = match.Groups["title"].Value
		seriesTitleInfo.Year = Convert.ToInt32(match.Groups["year"].Value)
	}

	return seriesTitleInfo
}

func parseMatchCollection(MatchCollection matchCollection) ParsedEpisodeInfo {
	var seriesName = matchCollection[0].Groups["title"].Value.Replace('.', ' ').Replace('_', ' ')
	seriesName = RequestInfoRegex.Replace(seriesName, "").Trim(' ')

	var airYear int
	int.TryParse(matchCollection[0].Groups["airyear"].Value, airYear)

	var result ParsedEpisodeInfo

	if airYear < 1900 {
		var seasons = []int{}

		for _, seasonCapture := range matchCollection[0].Groups["season"].Captures {
			var parsedSeason int
			if int.TryParse(seasonCapture.Value, parsedSeason) {
				seasons.Add(parsedSeason)
			}
		}

		//If no season was found it should be treated as a mini series and season 1
		if seasons.Count == 0 {
			seasons.Add(1)
		}

		//If more than 1 season was parsed go to the next REGEX (A multi-season release is unlikely)
		if seasons.Distinct().Count() > 1 {
			return nil
		}

		result = ParsedEpisodeInfo{
			SeasonNumber:           seasons.First(),
			EpisodeNumbers:         []int{},
			AbsoluteEpisodeNumbers: []int{},
		}

		for _, matchGroup := range matchCollection {
			// var episodeCaptures = matchGroup.Groups["episode"].Captures.Cast<Capture>().ToList()
			// var absoluteEpisodeCaptures = matchGroup.Groups["absoluteepisode"].Captures.Cast<Capture>().ToList()

			//Allows use to return a list of 0 episodes (We can handle that as a full season release)
			if episodeCaptures.Any() {
				var first = ParseNumber(episodeCaptures.First().Value)
				var last = ParseNumber(episodeCaptures.Last().Value)

				if first > last {
					return nil
				}

				var count = last - first + 1
				result.EpisodeNumbers = Enumerable.Range(first, count).ToArray()
			}

			if absoluteEpisodeCaptures.Any() {
				var first = Convert.ToInt32(absoluteEpisodeCaptures.First().Value)
				var last = Convert.ToInt32(absoluteEpisodeCaptures.Last().Value)

				if first > last {
					return nil
				}

				var count = last - first + 1
				result.AbsoluteEpisodeNumbers = Enumerable.Range(first, count).ToArray()

				if matchGroup.Groups["special"].Success {
					result.Special = true
				}
			}

			if !episodeCaptures.Any() && !absoluteEpisodeCaptures.Any() {
				//Check to see if this is an "Extras" or "SUBPACK" release, if it is, return NULL
				//Todo: Set a "Extras" flag in EpisodeParseResult if we want to download them ever
				if !matchCollection[0].Groups["extras"].Value.IsNullOrWhiteSpace() {
					return nil
				}

				result.FullSeason = true
			}
		}

		if result.AbsoluteEpisodeNumbers.Any() && !result.EpisodeNumbers.Any() {
			result.SeasonNumber = 0
		}
	} else {
		//Try to Parse as a daily show
		var airmonth = Convert.ToInt32(matchCollection[0].Groups["airmonth"].Value)
		var airday = Convert.ToInt32(matchCollection[0].Groups["airday"].Value)

		//Swap day and month if month is bigger than 12 (scene fail)
		if airmonth > 12 {
			var tempDay = airday
			airday = airmonth
			airmonth = tempDay
		}

		var airDate DateTime

		// try
		// {
		//     airDate = new DateTime(airYear, airmonth, airday)
		// }
		// catch (Exception)
		// {
		//     throw new InvalidDateException("Invalid date found: {0}-{1}-{2}", airYear, airmonth, airday)
		// }

		// //Check if episode is in the future (most likely a parse error)
		// if (airDate > DateTime.Now.AddDays(1).Date || airDate < new DateTime(1970, 1, 1))
		// {
		//     throw new InvalidDateException("Invalid date found: {0}", airDate)
		// }

		// result = new ParsedEpisodeInfo
		// {
		//     AirDate = airDate.ToString(Episode.AIR_DATE_FORMAT),
		// }
	}

	result.SeriesTitle = seriesName
	result.SeriesTitleInfo = GetSeriesTitleInfo(result.SeriesTitle)

	Logger.Debug("Episode Parsed. {0}", result)

	return result
}

func validateBeforeParsing(title string) bool {
	if title.ToLower().Contains("password") && title.ToLower().Contains("yenc") {
		Logger.Debug("")
		return false
	}

	if !title.Any(char.IsLetterOrDigit) {
		return false
	}

	var titleWithoutExtension = RemoveFileExtension(title)

	// if (RejectHashedReleasesRegex.Any(v => v.IsMatch(titleWithoutExtension))) {
	//     Logger.Debug("Rejected Hashed Release Title: " + title)
	//     return false
	// }

	return true
}

func GetSubGroup(matchCollection MatchCollection) string {
	var subGroup = matchCollection[0].Groups["subgroup"]

	if subGroup.Success {
		return subGroup.Value
	}

	return string.Empty
}

func GetReleaseHash(matchCollection MatchCollection) string {
	var hash = matchCollection[0].Groups["hash"]

	if hash.Success {
		var hashValue = hash.Value.Trim('[', ']')

		if hashValue.Equals("1280x720") {
			return string.Empty
		}

		return hashValue
	}

	return string.Empty
}

func parseNumber(value string) int {
	var number int

	if int.TryParse(value, number) {
		return number
	}

	number = Array.IndexOf(Numbers, value.ToLower())

	if number != -1 {
		return number
	}

	// throw new FormatException(string.Format("{0} isn't a number", value))
}
