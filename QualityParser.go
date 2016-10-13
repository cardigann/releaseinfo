package releaseinfo

import (
	"log"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

var SourceRegex = regexp2.MustCompile(`\b(?:
	(?<bluray>BluRay|Blu-Ray|HDDVD|BD)|
	(?<webdl>WEB[-_. ]DL|WEBDL|WebRip|iTunesHD|WebHD|[. ]WEB[. ](?:[xh]26[45]|DD5[. ]1)|\d+0p[. ]WEB[. ])|
	(?<hdtv>HDTV)|
	(?<bdrip>BDRip)|
	(?<brrip>BRRip)|
	(?<dvd>DVD|DVDRip|NTSC|PAL|xvidvd)|
	(?<dsr>WS[-_. ]DSR|DSR)|
	(?<pdtv>PDTV)|
	(?<sdtv>SDTV)|
	(?<tvrip>TVRip)
	)\b`,
	regexp2.Compiled|regexp2.IgnoreCase|regexp2.IgnorePatternWhitespace)

var RawHDRegex = regexp2.MustCompile(`\b(?<rawhd>RawHD|1080i[-_. ]HDTV|Raw[-_. ]HD|MPEG[-_. ]?2)\b`,
	regexp2.Compiled|regexp2.IgnoreCase)

var ProperRegex = regexp2.MustCompile(`\b(?<proper>proper|repack|rerip)\b`,
	regexp2.Compiled|regexp2.IgnoreCase)

var VersionRegex = regexp2.MustCompile(`\dv(?<version>\d)\b|\[v(?<version>\d)\]`,
	regexp2.Compiled|regexp2.IgnoreCase)

var RealRegex = regexp2.MustCompile(`\b(?<real>REAL)\b`,
	regexp2.Compiled)

var ResolutionRegex = regexp2.MustCompile(`\b(?:(?<_480p>480p|640x480|848x480)|(?<_576p>576p)|(?<_720p>720p|1280x720)|(?<_1080p>1080p|1920x1080)|(?<_2160p>2160p))\b`,
	regexp2.Compiled|regexp2.IgnoreCase)

var CodecRegex = regexp2.MustCompile(`\b(?:(?<x264>x264)|(?<h264>h264)|(?<xvidhd>XvidHD)|(?<xvid>Xvid)|(?<divx>divx))\b`,
	regexp2.Compiled|regexp2.IgnoreCase)

var OtherSourceRegex = regexp2.MustCompile(`(?<hdtv>HD[-_. ]TV)|(?<sdtv>SD[-_. ]TV)`, regexp2.Compiled|regexp2.IgnoreCase)

var AnimeBlurayRegex = regexp2.MustCompile(`bd(?:720|1080)|(?<=[-_. (\[])bd(?=[-_. )\]])`, regexp2.Compiled|regexp2.IgnoreCase)

var HighDefPdtvRegex = regexp2.MustCompile(`hr[-_. ]ws`, regexp2.Compiled|regexp2.IgnoreCase)

func ParseQuality(name string) QualityModel {
	log.Printf("Parsing quality for %q", name)

	normalizedName := removeSpace(name)
	normalizedName = strings.Replace(normalizedName, "_", " ", -1)
	normalizedName = removeSpace(normalizedName)
	normalizedName = strings.ToLower(normalizedName)

	result := parseQualityModifiers(name, normalizedName)

	if match, _ := RawHDRegex.FindStringMatch(normalizedName); match != nil {
		result.Quality = QualityRAWHD
		return result
	}

	sourceMatch, _ := SourceRegex.FindStringMatch(normalizedName)
	resolution := ParseResolution(normalizedName)
	codecRegex, _ := CodecRegex.FindStringMatch(normalizedName)

	if sourceMatch != nil {
		if sourceMatch.GroupByName("bluray") != nil {
			if codecRegex != nil && (codecRegex.GroupByName("xvid") != nil || codecRegex.GroupByName("divx") != nil) {
				result.Quality = QualityDVD
				return result
			}

			if resolution == Resolution2160p {
				result.Quality = QualityBluray2160p
				return result
			}

			if resolution == Resolution1080p {
				result.Quality = QualityBluray1080p
				return result
			}

			if resolution == Resolution480p || resolution == Resolution576p {
				result.Quality = QualityDVD
				return result
			}

			result.Quality = QualityBluray720p
			return result
		}

		if sourceMatch.GroupByName("webdl") != nil {
			if resolution == Resolution2160p {
				result.Quality = QualityWEBDL2160p
				return result
			}

			if resolution == Resolution1080p {
				result.Quality = QualityWEBDL1080p
				return result
			}

			if resolution == Resolution720p {
				result.Quality = QualityWEBDL720p
				return result
			}

			if strings.Contains(name, "[WEBDL]") {
				result.Quality = QualityWEBDL720p
				return result
			}

			result.Quality = QualityWEBDL480p
			return result
		}

		if sourceMatch.GroupByName("hdtv") != nil {
			if resolution == Resolution2160p {
				result.Quality = QualityHDTV2160p
				return result
			}

			if resolution == Resolution1080p {
				result.Quality = QualityHDTV1080p
				return result
			}

			if resolution == Resolution720p {
				result.Quality = QualityHDTV720p
				return result
			}

			if strings.Contains(name, "[HDTV]") {
				result.Quality = QualityHDTV720p
				return result
			}

			result.Quality = QualitySDTV
			return result
		}

		if sourceMatch.GroupByName("bdrip") != nil || sourceMatch.GroupByName("brrip") != nil {
			switch resolution {
			case Resolution720p:
				result.Quality = QualityBluray720p
				return result
			case Resolution1080p:
				result.Quality = QualityBluray1080p
				return result
			default:
				result.Quality = QualityDVD
				return result
			}
		}

		if sourceMatch.GroupByName("dvd") != nil {
			result.Quality = QualityDVD
			return result
		}

		if sourceMatch.GroupByName("pdtv") != nil ||
			sourceMatch.GroupByName("sdtv") != nil ||
			sourceMatch.GroupByName("dsr") != nil ||
			sourceMatch.GroupByName("tvrip") != nil {
			if match, _ := HighDefPdtvRegex.FindStringMatch(normalizedName); match != nil {
				result.Quality = QualityHDTV720p
				return result
			}

			result.Quality = QualitySDTV
			return result
		}
	}

	//Anime Bluray matching
	if match, _ := AnimeBlurayRegex.FindStringMatch(normalizedName); match != nil {
		if resolution == Resolution480p || resolution == Resolution576p || strings.Contains(normalizedName, "480p") {
			result.Quality = QualityDVD
			return result
		}

		if resolution == Resolution1080p || strings.Contains(normalizedName, "1080p") {
			result.Quality = QualityBluray1080p
			return result
		}

		result.Quality = QualityBluray720p
		return result
	}

	if resolution == Resolution2160p {
		result.Quality = QualityHDTV2160p
		return result
	}

	if resolution == Resolution1080p {
		result.Quality = QualityHDTV1080p
		return result
	}

	if resolution == Resolution720p {
		result.Quality = QualityHDTV720p
		return result
	}

	if resolution == Resolution480p {
		result.Quality = QualitySDTV
		return result
	}

	if codecRegex != nil && codecRegex.GroupByName("x264") != nil {
		result.Quality = QualitySDTV
		return result
	}

	if strings.Contains(normalizedName, "848x480") {
		if strings.Contains(normalizedName, "dvd") {
			result.Quality = QualityDVD
		}

		result.Quality = QualitySDTV
	}

	if strings.Contains(normalizedName, "1280x720") {
		if strings.Contains(normalizedName, "bluray") {
			result.Quality = QualityBluray720p
		}

		result.Quality = QualityHDTV720p
	}

	if strings.Contains(normalizedName, "1920x1080") {
		if strings.Contains(normalizedName, "bluray") {
			result.Quality = QualityBluray1080p
		}

		result.Quality = QualityHDTV1080p
	}

	if strings.Contains(normalizedName, "bluray720p") {
		result.Quality = QualityBluray720p
	}

	if strings.Contains(normalizedName, "bluray1080p") {
		result.Quality = QualityBluray1080p
	}

	if otherSourceMatch := otherSourceMatch(normalizedName); otherSourceMatch != QualityUnknown {
		result.Quality = otherSourceMatch
	}

	//Based on extension
	// if result.Quality == QualityUnknown && !name.ContainsInvalidPathChars() {
	// 	// try
	// 	// {
	// 	result.Quality = MediaFileExtensions.GetQualityForExtension(Path.GetExtension(name))
	// 	result.QualitySource = QualitySource.Extension
	// 	// }
	// 	// catch (ArgumentException)
	// 	// {
	// 	//Swallow exception for cases where string contains illegal
	// 	//path characters.
	// 	// }
	// }

	return result
}

func ParseResolution(name string) Resolution {
	match, _ := ResolutionRegex.FindStringMatch(name)

	if match == nil {
		return ResolutionUnknown
	}
	if match.GroupByName("_480p") != nil {
		return Resolution480p
	}

	if match.GroupByName("_576p") != nil {
		return Resolution576p
	}

	if match.GroupByName("_720p") != nil {
		return Resolution720p
	}

	if match.GroupByName("_1080p") != nil {
		return Resolution1080p
	}

	if match.GroupByName("_2160p") != nil {
		return Resolution2160p
	}

	return ResolutionUnknown
}

func otherSourceMatch(name string) Quality {
	match, _ := OtherSourceRegex.FindStringMatch(name)
	if match == nil {
		return QualityUnknown
	}
	if match.GroupByName("sdtv") != nil {
		return QualitySDTV
	}
	if match.GroupByName("hdtv") != nil {
		return QualityHDTV720p
	}

	return QualityUnknown
}

func parseQualityModifiers(name string, normalizedName string) QualityModel {
	result := QualityModel{Quality: QualityUnknown}

	if m, _ := ProperRegex.FindStringMatch(normalizedName); m != nil {
		result.Revision.Version = 2
	}

	versionRegexResult, _ := VersionRegex.FindStringMatch(normalizedName)
	if versionRegexResult != nil {
		result.Revision.Version, _ = strconv.Atoi(versionRegexResult.GroupByName("version").String())
	}

	//TODO: re-enable this when we have a reliable way to determine real
	//TODO: Only treat it as a real if it comes AFTER the season/epsiode number
	// var realRegexResult = RealRegex.Matches(name)
	// if realRegexResult.Count > 0 {
	// 	result.Revision.Real = realRegexResult.Count
	// }

	return result
}

type Resolution string

const (
	Resolution480p    Resolution = "480p"
	Resolution576p    Resolution = "576p"
	Resolution720p    Resolution = "720p"
	Resolution1080p   Resolution = "1080p"
	Resolution2160p   Resolution = "2160p"
	ResolutionUnknown Resolution = "unknown"
)
