package releaseinfo

import (
	"fmt"

	"golang.org/x/text/language"
)

type ParsedEpisodeInfo struct {
	SeriesTitle            string
	SeriesTitleInfo        SeriesTitleInfo
	Quality                QualityModel
	SeasonNumber           int
	EpisodeNumbers         []int
	AbsoluteEpisodeNumbers []int
	AirDate                string
	Language               language.Tag
	FullSeason             bool
	Special                bool
	ReleaseGroup           string
	ReleaseHash            string
}

func NewParsedEpisodeInfo() ParsedEpisodeInfo {
	return ParsedEpisodeInfo{
		EpisodeNumbers:         []int{},
		AbsoluteEpisodeNumbers: []int{},
	}
}

func (pei ParsedEpisodeInfo) IsDaily() bool {
	return removeSpace(pei.AirDate) != ""
}

func (pei ParsedEpisodeInfo) IsAbsoluteNumbering() bool {
	return len(pei.AbsoluteEpisodeNumbers) > 0
}

func (pei ParsedEpisodeInfo) IsPossibleSpecialEpisode() bool {
	// if we don't have eny episode numbers we are likely a special episode and need to do a search by episode title
	return removeSpace(pei.AirDate) != "" &&
		removeSpace(pei.SeriesTitle) != "" &&
		(len(pei.EpisodeNumbers) == 0 || pei.SeasonNumber == 0) ||
		(removeSpace(pei.SeriesTitle) != "" && pei.Special)
}

func (pei ParsedEpisodeInfo) String() string {
	episodeString := "[Unknown Episode]"

	if pei.IsDaily() && len(pei.EpisodeNumbers) == 0 {
		episodeString = fmt.Sprintf("%s", pei.AirDate)
	} else if pei.FullSeason {
		episodeString = fmt.Sprintf("Season {0:00}", pei.SeasonNumber)
	} else if pei.EpisodeNumbers != nil && len(pei.EpisodeNumbers) > 0 {
		panic("fix me")
		// episodeString = string.Format("S{0:00}E{1}", SeasonNumber, string.Join("-", EpisodeNumbers.Select(c => c.ToString("00"))));
	} else if pei.AbsoluteEpisodeNumbers != nil && len(pei.AbsoluteEpisodeNumbers) > 0 {
		panic("fix me")
		//episodeString = string.Format("{0}", string.Join("-", AbsoluteEpisodeNumbers.Select(c => c.ToString("000"))));
	}

	return fmt.Sprintf("{0} - {1} {2}", pei.SeriesTitle, episodeString, pei.Quality)
}
