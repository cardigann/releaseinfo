package releaseinfo

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func TestParseSeriesName(t *testing.T) {
	for idx, test := range []struct {
		postTitle, expected string
	}{
		{"Chuck - 4x05 - Title", "Chuck"},
		{"Law & Order - 4x05 - Title", "laworder"},
		{"Bad Format", "badformat"},
		{"Mad Men - Season 1 [Bluray720p]", "madmen"},
		{"Mad Men - Season 1 [Bluray1080p]", "madmen"},
		{"The Daily Show With Jon Stewart -", "thedailyshowwithjonstewart"},
		{"The Venture Bros. (2004)", "theventurebros2004"},
		{"Castle (2011)", "castle2011"},
		{"Adventure Time S02 720p HDTV x264 CRON", "adventuretime"},
		{"Hawaii Five 0", "hawaiifive0"},
		{"Match of the Day", "matchday"},
		{"Match of the Day 2", "matchday2"},
		{"[ www.Torrenting.com ] - Revenge.S03E14.720p.HDTV.X264-DIMENSION", "Revenge"},
		{"Seed S02E09 HDTV x264-2HD [eztv]-[rarbg.com]", "Seed"},
		{"Reno.911.S01.DVDRip.DD2.0.x264-DEEP", "Reno 911"},
	} {
		require.Equal(t,
			CleanSeriesTitle(test.expected),
			CleanSeriesTitle(ParseSeriesName(test.postTitle)),
			fmt.Sprintf("Row %d should have correct title", idx+1))
	}
}

func TestRemovingAccentsFromTitle(t *testing.T) {
	require.Equal(t, "carnivale", CleanSeriesTitle("Carniv\u00E0le"))
}

func TestRemovingExtensionsFromTitle(t *testing.T) {
	_, err := ParseTitle("Discovery TV - Gold Rush : 02 Road From Hell [S04].mp4")
	if err != nil {
		t.Fatal(err)
	}
}

func TestParsingYearFromTitle(t *testing.T) {
	for idx, test := range []struct {
		postTitle, expectedTitle, expectedTitleWithoutYear string
		expectedYear                                       int
	}{
		{"House.S01E01.pilot.720p.hdtv", "House", "House", 0},
		{"House.2004.S01E01.pilot.720p.hdtv", "House 2004", "House", 2004},
	} {
		result, err := ParseTitle(test.postTitle)

		require.NoError(t, err)
		require.Equal(t, test.expectedYear, result.SeriesTitleInfo.Year,
			fmt.Sprintf("Row %d should have correct year", idx+1))
		require.Equal(t, test.expectedTitle, result.SeriesTitleInfo.Title,
			fmt.Sprintf("Row %d should have correct title", idx+1))
		require.Equal(t, test.expectedTitleWithoutYear, result.SeriesTitleInfo.TitleWithoutYear,
			fmt.Sprintf("Row %d should have correct title without year)", idx+1))
	}
}

func TestParsingSingleEpisodeNumber(t *testing.T) {
	for idx, test := range []struct {
		postTitle, expectedTitle        string
		expectedSeason, expectedEpisode int
	}{
		{"Sonny.With.a.Chance.S02E15", "Sonny With a Chance", 2, 15},
		{"Two.and.a.Half.Me.103.720p.HDTV.X264-DIMENSION", "Two and a Half Me", 1, 3},
		{"Two.and.a.Half.Me.113.720p.HDTV.X264-DIMENSION", "Two and a Half Me", 1, 13},
		{"Two.and.a.Half.Me.1013.720p.HDTV.X264-DIMENSION", "Two and a Half Me", 10, 13},
		{"Chuck.4x05.HDTV.XviD-LOL", "Chuck", 4, 5},
		{"The.Girls.Next.Door.S03E06.DVDRip.XviD-WiDE", "The Girls Next Door", 3, 6},
		{"Degrassi.S10E27.WS.DSR.XviD-2HD", "Degrassi", 10, 27},
		{"Parenthood.2010.S02E14.HDTV.XviD-LOL", "Parenthood 2010", 2, 14},
		{"Hawaii Five 0 S01E19 720p WEB DL DD5 1 H 264 NT", "Hawaii Five 0", 1, 19},
		{"The Event S01E14 A Message Back 720p WEB DL DD5 1 H264 SURFER", "The Event", 1, 14},
		{"Adam Hills In Gordon St Tonight S01E07 WS PDTV XviD FUtV", "Adam Hills In Gordon St Tonight", 1, 7},
		{"Adventure.Inc.S03E19.DVDRip.XviD-OSiTV", "Adventure Inc", 3, 19},
		{"S03E09 WS PDTV XviD FUtV", "", 3, 9},
		{"5x10 WS PDTV XviD FUtV", "", 5, 10},
		{"Castle.2009.S01E14.HDTV.XviD-LOL", "Castle 2009", 1, 14},
		{"Pride.and.Prejudice.1995.S03E20.HDTV.XviD-LOL", "Pride and Prejudice 1995", 3, 20},
		{"The.Office.S03E115.DVDRip.XviD-OSiTV", "The Office", 3, 115},
		{"Parks and Recreation - S02E21 - 94 Meetings - 720p TV.mkv", "Parks and Recreation", 2, 21},
		{"24-7 Penguins-Capitals- Road to the NHL Winter Classic - S01E03 - Episode 3.mkv", "24-7 Penguins-Capitals- Road to the NHL Winter Classic", 1, 3},
		{"Adventure.Inc.S03E19.DVDRip.\"XviD\"-OSiTV", "Adventure Inc", 3, 19},
		{"Hawaii Five-0 (2010) - 1x05 - Nalowale (Forgotten/Missing)", "Hawaii Five-0 (2010)", 1, 5},
		{"Hawaii Five-0 (2010) - 1x05 - Title", "Hawaii Five-0 (2010)", 1, 5},
		{"House - S06E13 - 5 to 9 [DVD]", "House", 6, 13},
		{"The Mentalist - S02E21 - 18-5-4", "The Mentalist", 2, 21},
		{"Breaking.In.S01E07.21.0.Jump.Street.720p.WEB-DL.DD5.1.h.264-KiNGS", "Breaking In", 1, 7},
		{"CSI.525", "CSI", 5, 25},
		{"King of the Hill - 10x12 - 24 Hour Propane People [SDTV]", "King of the Hill", 10, 12},
		{"Brew Masters S01E06 3 Beers For Batali DVDRip XviD SPRiNTER", "Brew Masters", 1, 6},
		{"24 7 Flyers Rangers Road to the NHL Winter Classic Part01 720p HDTV x264 ORENJI", "24 7 Flyers Rangers Road to the NHL Winter Classic", 1, 1},
		{"24 7 Flyers Rangers Road to the NHL Winter Classic Part 02 720p HDTV x264 ORENJI", "24 7 Flyers Rangers Road to the NHL Winter Classic", 1, 2},
		{"24-7 Flyers-Rangers- Road to the NHL Winter Classic - S01E01 - Part 1", "24-7 Flyers-Rangers- Road to the NHL Winter Classic", 1, 1},
		{"S6E02-Unwrapped-(Playing With Food) - [DarkData]", "", 6, 2},
		{"S06E03-Unwrapped-(Number Ones Unwrapped) - [DarkData]", "", 6, 3},
		{"The Mentalist S02E21 18 5 4 720p WEB DL DD5 1 h 264 EbP", "The Mentalist", 2, 21},
		{"01x04 - Halloween, Part 1 - 720p WEB-DL", "", 1, 4},
		{"extras.s03.e05.ws.dvdrip.xvid-m00tv", "extras", 3, 5},
		{"castle.2009.416.hdtv-lol", "castle 2009", 4, 16},
		{"hawaii.five-0.2010.217.hdtv-lol", "hawaii five-0 2010", 2, 17},
		{"Looney Tunes - S1936E18 - I Love to Singa", "Looney Tunes", 1936, 18},
		{"American_Dad!_-_7x6_-_The_Scarlett_Getter_[SDTV]", "American Dad!", 7, 6},
		{"Falling_Skies_-_1x1_-_Live_and_Learn_[HDTV-720p]", "Falling Skies", 1, 1},
		{"Top Gear - 07x03 - 2005.11.70", "Top Gear", 7, 3},
		{"Glee.S04E09.Swan.Song.1080p.WEB-DL.DD5.1.H.264-ECI", "Glee", 4, 9},
		{"S08E20 50-50 Carla [DVD]", "", 8, 20},
		{"Cheers S08E20 50-50 Carla [DVD]", "Cheers", 8, 20},
		{"S02E10 6-50 to SLC [SDTV]", "", 2, 10},
		{"Franklin & Bash S02E10 6-50 to SLC [SDTV]", "Franklin & Bash", 2, 10},
		{"The_Big_Bang_Theory_-_6x12_-_The_Egg_Salad_Equivalency_[HDTV-720p]", "The Big Bang Theory", 6, 12},
		{"Top_Gear.19x06.720p_HDTV_x264-FoV", "Top Gear", 19, 6},
		{"Portlandia.S03E10.Alexandra.720p.WEB-DL.AAC2.0.H.264-CROM.mkv", "Portlandia", 3, 10},
		//{"(Game of Thrones s03 e - \"Game of Thrones Season 3 Episode 10\"", "Game of Thrones", 3, 10},
		{"House.Hunters.International.S05E607.720p.hdtv.x264", "House Hunters International", 5, 607},
		{"Adventure.Time.With.Finn.And.Jake.S01E20.720p.BluRay.x264-DEiMOS", "Adventure Time With Finn And Jake", 1, 20},
		{"Hostages.S01E04.2-45.PM.[HDTV-720p].mkv", "Hostages", 1, 4},
		{"S01E04", "", 1, 4},
		{"1x04", "", 1, 4},
		{"10.Things.You.Dont.Know.About.S02E04.Prohibition.HDTV.XviD-AFG", "10 Things You Dont Know About", 2, 4},
		{"30 Rock - S01E01 - Pilot.avi", "30 Rock", 1, 1},
		{"666 Park Avenue - S01E01", "666 Park Avenue", 1, 1},
		{"Warehouse 13 - S01E01", "Warehouse 13", 1, 1},
		{"Don't Trust The B---- in Apartment 23.S01E01", "Don't Trust The B---- in Apartment 23", 1, 1},
		{"Warehouse.13.S01E01", "Warehouse 13", 1, 1},
		{"Dont.Trust.The.B----.in.Apartment.23.S01E01", "Dont Trust The B---- in Apartment 23", 1, 1},
		{"24 S01E01", "24", 1, 1},
		{"24.S01E01", "24", 1, 1},
		{"Homeland - 2x12 - The Choice [HDTV-1080p].mkv", "Homeland", 2, 12},
		{"Homeland - 2x4 - New Car Smell [HDTV-1080p].mkv", "Homeland", 2, 4},
		{"Top Gear - 06x11 - 2005.08.07", "Top Gear", 6, 11},
		{"The_Voice_US_s06e19_04.28.2014_hdtv.x264.Poke.mp4", "The Voice US", 6, 19},
		{"the.100.110.hdtv-lol", "the 100", 1, 10},
		{"2009x09 [SDTV].avi", "", 2009, 9},
		{"S2009E09 [SDTV].avi", "", 2009, 9},
		{"Shark Week S2009E09 [SDTV].avi", "Shark Week", 2009, 9},
		{"St_Elsewhere_209_Aids_And_Comfort", "St Elsewhere", 2, 9},
		{"[Impatience] Locodol - 0x01 [720p][34073169].mkv", "Locodol", 0, 1},
		{"South.Park.S15.E06.City.Sushi", "South Park", 15, 6},
		{"South Park - S15 E06 - City Sushi", "South Park", 15, 6},
		{"Constantine S1-E1-WEB-DL-1080p-NZBgeek", "Constantine", 1, 1},
		{"Constantine S1E1-WEB-DL-1080p-NZBgeek", "Constantine", 1, 1},
		{"NCIS.S010E16.720p.HDTV.X264-DIMENSION", "NCIS", 10, 16},
		{"[ www.Torrenting.com ] - Revolution.2012.S02E17.720p.HDTV.X264-DIMENSION", "Revolution 2012", 2, 17},
		{"Revolution.2012.S02E18.720p.HDTV.X264-DIMENSION.mkv", "Revolution 2012", 2, 18},
		{"Series - Season 1 - Episode 01 (Resolution).avi", "Series", 1, 1},
		{"5x09 - 100 [720p WEB-DL].mkv", "", 5, 9},
		{"1x03 - 274 [1080p BluRay].mkv", "", 1, 3},
		{"1x03 - The 112th Congress [1080p BluRay].mkv", "", 1, 3},
		{"Revolution.2012.S02E14.720p.HDTV.X264-DIMENSION [PublicHD].mkv", "Revolution 2012", 2, 14},
		{"Castle.2009.S06E03.720p.HDTV.X264-DIMENSION [PublicHD].mkv", "Castle 2009", 6, 3},
		{"19-2.2014.S02E01.720p.HDTV.x264-CROOKS", "19-2 2014", 2, 1},
		{"Community - S01E09 - Debate 109", "Community", 1, 9},
		{"Entourage - S02E02 - My Maserati Does 185", "Entourage", 2, 2},
		{"6x13 - The Family Guy 100th Episode Special", "", 6, 13},
		{"The Young And The Restless - S41 E10478 - 2014-08-15", "The Young And The Restless", 41, 10478},
		{"The Young And The Restless - S42 E10591 - 2015-01-27", "The Young And The Restless", 42, 10591},
		{"Series Title [1x05] Episode Title", "Series Title", 1, 5},
		{"Series Title [S01E05] Episode Title", "Series Title", 1, 5},
		{"Series Title Season 01 Episode 05 720p", "Series Title", 1, 5},
		{"The Young And the Restless - S42 E10713 - 2015-07-20.mp4", "The Young And the Restless", 42, 10713},
		// {"quantico.103.hdtv-lol[ettv].mp4", "quantico", 1, 3},
		{"Fargo - 01x02 - The Rooster Prince - [itz_theo]", "Fargo", 1, 2},
		{"Castle (2009) - [06x16] - Room 147.mp4", "Castle (2009)", 6, 16},
		{"grp-zoos01e11-1080p", "grp-zoo", 1, 11},
		{"grp-zoo-s01e11-1080p", "grp-zoo", 1, 11},
		{"Jeopardy!.S2016E14.2016-01-20.avi", "Jeopardy!", 2016, 14},
		{"Ken.Burns.The.Civil.War.5of9.The.Universe.Of.Battle.1990.DVDRip.x264-HANDJOB", "Ken Burns The Civil War", 1, 5},
		{"Judge Judy 2016 02 25 S20E142", "Judge Judy", 20, 142},
		{"Judge Judy 2016 02 25 S20E143", "Judge Judy", 20, 143},
		{"Red Dwarf - S02 - E06 - Parallel Universe", "Red Dwarf", 2, 6},
		{"O.J.Simpson.Made.in.America.Part.Two.720p.HDTV.x264-2HD", "O J Simpson Made in America", 1, 2},
		{"The.100000.Dollar.Pyramid.2016.S01E05.720p.HDTV.x264-W4F", "The 100000 Dollar Pyramid 2016", 1, 5},
		//[TestCase("Sex And The City S6E15 - Catch-38 [RavyDavy].avi", "Sex And The City", 6, 15)] // -38 is getting treated as abs number
		//[TestCase("Heroes - S01E01 - Genesis 101 [HDTV-720p]", "Heroes", 1, 1)]
		//[TestCase("The 100 S02E01 HDTV x264-KILLERS [eztv]", "The 100", 2, 1)]
	} {
		result, err := ParseTitle(test.postTitle)

		require.NoError(t, err)
		require.Equal(t, test.expectedTitle, result.SeriesTitleInfo.Title,
			fmt.Sprintf("Row %d should have correct title", idx+1))
		require.Len(t, result.EpisodeNumbers, 1,
			fmt.Sprintf("Row %d should have 1 episode number)", idx+1))
		require.Equal(t, test.expectedEpisode, result.EpisodeNumbers[0],
			fmt.Sprintf("Row %d should have correct episode number", idx+1))
		require.Len(t, result.AbsoluteEpisodeNumbers, 0,
			fmt.Sprintf("Row %d should have 0 absolute episode numbers)", idx+1))
		require.Equal(t, test.expectedSeason, result.SeasonNumber,
			fmt.Sprintf("Row %d should have correct season number", idx+1))
		require.False(t, result.FullSeason, 1,
			fmt.Sprintf("Row %d should not be a full season)", idx+1))
	}
}

func TestParseFullSeasonReleases(t *testing.T) {
	for idx, test := range []struct {
		postTitle, expectedTitle string
		expectedSeason           int
	}{
		{"30.Rock.Season.04.HDTV.XviD-DIMENSION", "30 Rock", 4},
		{"Parks.and.Recreation.S02.720p.x264-DIMENSION", "Parks and Recreation", 2},
		{"The.Office.US.S03.720p.x264-DIMENSION", "The Office US", 3},
		{"Sons.of.Anarchy.S03.720p.BluRay-CLUE\\REWARD", "Sons of Anarchy", 3},
		{"Adventure Time S02 720p HDTV x264 CRON", "Adventure Time", 2},
		{"Sealab.2021.S04.iNTERNAL.DVDRip.XviD-VCDVaULT", "Sealab 2021", 4},
		{"Hawaii Five 0 S01 720p WEB DL DD5 1 H 264 NT", "Hawaii Five 0", 1},
		{"30 Rock S03 WS PDTV XviD FUtV", "30 Rock", 3},
		{"The Office Season 4 WS PDTV XviD FUtV", "The Office", 4},
		{"Eureka Season 1 720p WEB DL DD 5 1 h264 TjHD", "Eureka", 1},
		{"The Office Season4 WS PDTV XviD FUtV", "The Office", 4},
		{"Eureka S 01 720p WEB DL DD 5 1 h264 TjHD", "Eureka", 1},
		{"Doctor Who Confidential   Season 3", "Doctor Who Confidential", 3},
		{"Fleming.S01.720p.WEBDL.DD5.1.H.264-NTb", "Fleming", 1},
		{"Holmes.Makes.It.Right.S02.720p.HDTV.AAC5.1.x265-NOGRP", "Holmes Makes It Right", 2},
		{"My.Series.S2014.720p.HDTV.x264-ME", "My Series", 2014},
	} {
		result, err := ParseTitle(test.postTitle)

		require.NoError(t, err)
		require.True(t, result.FullSeason, 1,
			fmt.Sprintf("Row %d should be a full season)", idx+1))
		require.Equal(t, test.expectedTitle, result.SeriesTitleInfo.Title,
			fmt.Sprintf("Row %d should have correct title", idx+1))
		require.Len(t, result.EpisodeNumbers, 0,
			fmt.Sprintf("Row %d should have 0 episode numbers)", idx+1))
		require.Len(t, result.AbsoluteEpisodeNumbers, 0,
			fmt.Sprintf("Row %d should have 0 absolute episode numbers)", idx+1))
		require.Equal(t, test.expectedSeason, result.SeasonNumber,
			fmt.Sprintf("Row %d should have correct season number", idx+1))
	}
}

func TestNotParseSeasonExtrasAndSubpacks(t *testing.T) {
	for idx, postTitle := range []string{
		"Acropolis Now S05 EXTRAS DVDRip XviD RUNNER",
		"Punky Brewster S01 EXTRAS DVDRip XviD RUNNER",
		"Instant Star S03 EXTRAS DVDRip XviD OSiTV",
		"Lie.to.Me.S03.SUBPACK.DVDRip.XviD-REWARD",
		"The.Middle.S02.SUBPACK.DVDRip.XviD-REWARD",
		"CSI.S11.SUBPACK.DVDRip.XviD-REWARD",
	} {
		result, err := ParseTitle(postTitle)
		require.Error(t, err,
			fmt.Sprintf("Row %d should have an error", idx+1))
		require.Nil(t, result,
			fmt.Sprintf("Row %d should have a nil result", idx+1))
	}
}

func TestParsingHashedReleases(t *testing.T) {
	for idx, test := range []struct {
		path, expectedTitle  string
		expectedQuality      Quality
		expectedReleaseGroup string
	}{
		{`C:\Test\Some.Hashed.Release.S01E01.720p.WEB-DL.AAC2.0.H.264-Mercury\0e895c37245186812cb08aab1529cf8ee389dd05.mkv`,
			"Some Hashed Release", QualityWEBDL720p, "Mercury"},
		{`C:\Test\0e895c37245186812cb08aab1529cf8ee389dd05\Some.Hashed.Release.S01E01.720p.WEB-DL.AAC2.0.H.264-Mercury.mkv`,
			"Some Hashed Release", QualityWEBDL720p, "Mercury"},
		{`C:\Test\Fake.Dir.S01E01-Test\yrucreM-462.H.0.2CAA.LD-BEW.p027.10E10S.esaeleR.dehsaH.emoS.mkv`,
			"Some Hashed Release", QualityWEBDL720p, "Mercury"},
		{`C:\Test\Fake.Dir.S01E01-Test\yrucreM-LN 1.5DD LD-BEW P0801 10E10S esaeleR dehsaH emoS.mkv`,
			"Some Hashed Release", QualityWEBDL1080p, "Mercury"},
		{`C:\Test\Weeds.S01E10.DVDRip.XviD-SONARR\AHFMZXGHEWD660.mkv`,
			"Weeds", QualityDVD, "SONARR"},
		{`C:\Test\Deadwood.S02E12.1080p.BluRay.x264-SONARR\Backup_72023S02-12.mkv`,
			"Deadwood", QualityBluray1080p, "SONARR"},
		{`C:\Test\Grimm S04E08 Chupacabra 720p WEB-DL DD5 1 H 264-ECI\123.mkv`,
			"Grimm", QualityWEBDL720p, "ECI"},
		{`C:\Test\Grimm S04E08 Chupacabra 720p WEB-DL DD5 1 H 264-ECI\abc.mkv`,
			"Grimm", QualityWEBDL720p, "ECI"},
		{`C:\Test\Grimm S04E08 Chupacabra 720p WEB-DL DD5 1 H 264-ECI\b00bs.mkv`,
			"Grimm", QualityWEBDL720p, "ECI"},
		{`C:\Test\The.Good.Wife.S02E23.720p.HDTV.x264-NZBgeek/cgajsofuejsa501.mkv`,
			"The Good Wife", QualityHDTV720p, "NZBgeek"},
	} {
		result, err := ParsePath(test.path)

		require.NoError(t, err)
		require.Equal(t, test.expectedTitle, result.SeriesTitleInfo.Title,
			fmt.Sprintf("Row %d should have correct title", idx+1))
		require.Equal(t, test.expectedQuality, result.Quality.Quality,
			fmt.Sprintf("Row %d should have correct quality", idx+1))
		require.Equal(t, test.expectedReleaseGroup, result.ReleaseGroup,
			fmt.Sprintf("Row %d should have correct quality", idx+1))
	}
}

func TestCleanSeriesTitle(t *testing.T) {
	for idx, test := range []struct {
		postTitle, expectedTitle string
	}{
		{"Conan", "conan"},
		{"Castle (2009)", "castle2009"},
		{"Parenthood.2010", "parenthood2010"},
		{"Law_and_Order_SVU", "lawordersvu"},
		{"CaPitAl", "capital"},
		{"peri.od", "period"},
		{"this.^&%^**$%@#$!That", "thisthat"},
		{"test/test", "testtest"},
		{"90210", "90210"},
		{"24", "24"},
	} {
		require.Equal(t, test.expectedTitle, CleanSeriesTitle(test.postTitle),
			fmt.Sprintf("Row %d should have correct title", idx+1))
	}
}

func TestCleanSeriesTitleRemovesCommonWords(t *testing.T) {
	for idx, word := range []string{
		"the", "and", "or", "an", "of", "a",
	} {
		for _, format := range []string{
			"word.%q.word",
			"word %q word",
			"word-%q-word",
			"word.word.%q",
			"word-word-%q",
			"word-word %q",
		} {
			require.Equal(t, "wordword", CleanSeriesTitle(fmt.Sprintf(format, word)),
				fmt.Sprintf("Row %d should have correct title", idx+1))
		}
	}
}

func TestCleanSeriesTitleShouldntRemoveCommonWordsFromInsideOtherWords(t *testing.T) {
	for idx, word := range []string{
		"the", "and", "or", "an", "of", "a",
	} {
		for _, format := range []string{
			"word.%sword",
			"word %sword",
			"word-%sword",
			"word%s.word",
			"word%s-word",
			"word%s-word",
		} {
			require.Equal(t,
				"word"+strings.ToLower(word)+"word",
				CleanSeriesTitle(fmt.Sprintf(format, word)),
				fmt.Sprintf("Row %d should have correct title", idx+1))
		}
	}
}

func TestCleanSeriesTitleShouldntRemoveCommonWordsFromStartOfString(t *testing.T) {
	for idx, word := range []string{
		"the", "and", "or", "an", "of", "a",
	} {
		for _, format := range []string{
			"%s.word.word",
			"%s-word-word",
			"%s word word",
		} {
			require.Equal(t,
				strings.ToLower(word)+"wordword",
				CleanSeriesTitle(fmt.Sprintf(format, word)),
				fmt.Sprintf("Row %d should have correct title", idx+1))
		}
	}
}

func TestCleanSeriesTitleShouldntRemoveTheFromStartOfTitle(t *testing.T) {
	for idx, test := range []struct {
		postTitle, expectedTitle string
	}{
		{"The Office", "theoffice"},
		{"The Tonight Show With Jay Leno", "thetonightshowwithjayleno"},
		{"The.Daily.Show", "thedailyshow"},
	} {
		require.Equal(t, CleanSeriesTitle(test.postTitle), test.expectedTitle,
			fmt.Sprintf("Row %d should have correct title", idx+1))
	}
}

func TestCleanSeriesTitleShouldntRemoveAFromEndOfString(t *testing.T) {
	require.Equal(t, CleanSeriesTitle("Tokyo Ghoul A"), "tokyoghoula")
}

func TestParsingTitleLanguage(t *testing.T) {
	for idx, test := range []struct {
		postTitle        string
		expectedLanguage language.Tag
	}{
		{"Castle.2009.S01E14.English.HDTV.XviD-LOL", language.English},
		{"Castle.2009.S01E14.French.HDTV.XviD-LOL", language.French},
		{"Castle.2009.S01E14.Spanish.HDTV.XviD-LOL", language.Spanish},
		{"Castle.2009.S01E14.German.HDTV.XviD-LOL", language.German},
		{"Castle.2009.S01E14.Germany.HDTV.XviD-LOL", language.English},
		{"Castle.2009.S01E14.Italian.HDTV.XviD-LOL", language.Italian},
		{"Castle.2009.S01E14.Danish.HDTV.XviD-LOL", language.Danish},
		{"Castle.2009.S01E14.Dutch.HDTV.XviD-LOL", language.Dutch},
		{"Castle.2009.S01E14.Japanese.HDTV.XviD-LOL", language.Japanese},
		{"Castle.2009.S01E14.Cantonese.HDTV.XviD-LOL", language.MustParse("yue")},
		{"Castle.2009.S01E14.Mandarin.HDTV.XviD-LOL", language.MustParse("cmn")},
		{"Castle.2009.S01E14.Korean.HDTV.XviD-LOL", language.Korean},
		{"Castle.2009.S01E14.Russian.HDTV.XviD-LOL", language.Russian},
		{"Castle.2009.S01E14.Polish.HDTV.XviD-LOL", language.Polish},
		{"Castle.2009.S01E14.Vietnamese.HDTV.XviD-LOL", language.Vietnamese},
		{"Castle.2009.S01E14.Swedish.HDTV.XviD-LOL", language.Swedish},
		{"Castle.2009.S01E14.Norwegian.HDTV.XviD-LOL", language.Norwegian},
		{"Castle.2009.S01E14.Finnish.HDTV.XviD-LOL", language.Finnish},
		{"Castle.2009.S01E14.Turkish.HDTV.XviD-LOL", language.Turkish},
		{"Castle.2009.S01E14.Portuguese.HDTV.XviD-LOL", language.Portuguese},
		{"Castle.2009.S01E14.HDTV.XviD-LOL", language.English},
		{"person.of.interest.1x19.ita.720p.bdmux.x264-novarip", language.Italian},
		{"Salamander.S01E01.FLEMISH.HDTV.x264-BRiGAND", language.MustParse("nl-BE")},
		{"H.Polukatoikia.S03E13.Greek.PDTV.XviD-Ouzo", language.Greek},
		{"Burn.Notice.S04E15.Brotherly.Love.GERMAN.DUBBED.WS.WEBRiP.XviD.REPACK-TVP", language.German},
		{"Ray Donovan - S01E01.720p.HDtv.x264-Evolve (NLsub)", language.Dutch},
		{"Shield,.The.1x13.Tueurs.De.Flics.FR.DVDRip.XviD", language.French},
		{"True.Detective.S01E01.1080p.WEB-DL.Rus.Eng.TVKlondike", language.Russian},
		{"The.Trip.To.Italy.S02E01.720p.HDTV.x264-TLA", language.English},
		{"Revolution S01E03 No Quarter 2012 WEB-DL 720p Nordic-philipo mkv", language.Norwegian},
		{"Extant.S01E01.VOSTFR.HDTV.x264-RiDERS", language.French},
		{"Constantine.2014.S01E01.WEBRiP.H264.AAC.5.1-NL.SUBS", language.Dutch},
		{"Elementary - S02E16 - Kampfhaehne - mkv - by Videomann", language.German},
		{"Two.Greedy.Italians.S01E01.The.Family.720p.HDTV.x264-FTP", language.English},
		{"Castle.2009.S01E14.HDTV.XviD.HUNDUB-LOL", language.Hungarian},
		{"Castle.2009.S01E14.HDTV.XviD.ENG.HUN-LOL", language.Hungarian},
		{"Castle.2009.S01E14.HDTV.XviD.HUN-LOL", language.Hungarian},
	} {
		require.Equal(t, ParseLanguage(test.postTitle).String(), test.expectedLanguage.String(),
			fmt.Sprintf("Row %d should have correct language", idx+1))
	}
}

func TestParsingSubtitleTitleLanguage(t *testing.T) {
	for idx, test := range []struct {
		postTitle        string
		expectedLanguage language.Tag
	}{
		{"2 Broke Girls - S01E01 - Pilot.en.sub", language.English},
		{"2 Broke Girls - S01E01 - Pilot.eng.sub", language.English},
		// {"2 Broke Girls - S01E01 - Pilot.sub", language.English},
	} {
		lang, err := ParseSubtitleLanguage(test.postTitle)
		require.NoError(t, err,
			fmt.Sprintf("Row %d should have no error", idx+1))
		require.Equal(t, lang.String(), test.expectedLanguage.String(),
			fmt.Sprintf("Row %d should have correct subtitle language", idx+1))
	}
}

func TestParsingSubtitleTitleLanguageFailsWhenNotPresent(t *testing.T) {
	_, err := ParseSubtitleLanguage("2 Broke Girls - S01E01 - Pilot.sub")
	require.Error(t, err)
}
