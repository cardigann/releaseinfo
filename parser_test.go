package releaseinfo

import "testing"

/* Fucked-up hall of shame,
 * WWE.Wrestlemania.27.PPV.HDTV.XviD-KYR
 * Unreported.World.Chinas.Lost.Sons.WS.PDTV.XviD-FTP
 * [TestCase("Big Time Rush 1x01 to 10 480i DD2 0 Sianto", "Big Time Rush", 1, new[] { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }, 10)]
 * [TestCase("Desparate Housewives - S07E22 - 7x23 - And Lots of Security.. [HDTV-720p].mkv", "Desparate Housewives", 7, new[] { 22, 23 }, 2)]
 * [TestCase("S07E22 - 7x23 - And Lots of Security.. [HDTV-720p].mkv", "", 7, new[] { 22, 23 }, 2)]
 * (Game of Thrones s03 e - "Game of Thrones Season 3 Episode 10"
 * The.Man.of.Steel.1994-05.33.hybrid.DreamGirl-Novus-HD
 * Superman.-.The.Man.of.Steel.1994-06.34.hybrid.DreamGirl-Novus-HD
 * Superman.-.The.Man.of.Steel.1994-05.33.hybrid.DreamGirl-Novus-HD
 * Constantine S1-E1-WEB-DL-1080p-NZBgeek
 */

func TestParseSeriesName(t *testing.T) {
	tests := []struct {
		postTitle, title string
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
	}

	for _, test := range tests {
		series := ParseSeriesName(test.postTitle)
		cleaned := CleanSeriesTitle(series)
		expected := CleanSeriesTitle(test.title)

		if cleaned != expected {
			t.Fatalf("Expected %s, got %s", expected, cleaned)
		}
	}
}

// [Test]
// public void should_remove_accents_from_title()
// {
//     const string title = "Carniv\u00E0le";

//     title.CleanSeriesTitle().Should().Be("carnivale");
// }

// [TestCase("Discovery TV - Gold Rush : 02 Road From Hell [S04].mp4")]
// public void should_clean_up_invalid_path_characters(string postTitle)
// {
//     Parser.Parser.ParseTitle(postTitle);
// }

// [TestCase("[scnzbefnet][509103] 2.Broke.Girls.S03E18.720p.HDTV.X264-DIMENSION", "2 Broke Girls")]
// public void should_remove_request_info_from_title(string postTitle, string title)
// {
//     Parser.Parser.ParseTitle(postTitle).SeriesTitle.Should().Be(title);
// }
