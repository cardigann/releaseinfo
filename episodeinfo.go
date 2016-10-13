package releaseinfo

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

type EpisodeInfo struct {
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

type SeriesTitleInfo struct {
	Title            string
	TitleWithoutYear string
	Year             int
}

func (i EpisodeInfo) IsDaily() bool {
	return removeSpace(i.AirDate) != ""
}

func (i EpisodeInfo) IsAbsoluteNumbering() bool {
	return len(i.AbsoluteEpisodeNumbers) > 0
}

func (i EpisodeInfo) IsPossibleSpecialEpisode() bool {
	return removeSpace(i.AirDate) != "" &&
		removeSpace(i.SeriesTitle) != "" &&
		(len(i.EpisodeNumbers) == 0 || i.SeasonNumber == 0) ||
		(removeSpace(i.SeriesTitle) != "" && i.Special)
}

func (i EpisodeInfo) String() string {
	episodeString := "[Unknown Episode]"

	if i.IsDaily() && len(i.EpisodeNumbers) == 0 {
		episodeString = fmt.Sprintf("%s", i.AirDate)
	} else if i.FullSeason {
		episodeString = fmt.Sprintf("S%02d", i.SeasonNumber)
	} else if len(i.EpisodeNumbers) > 0 {
		episodes := []string{}
		for _, episode := range i.EpisodeNumbers {
			episodes = append(episodes, fmt.Sprintf("%02d", episode))
		}
		episodeString = fmt.Sprintf("S%02dE%s", i.SeasonNumber, strings.Join(episodes, "-"))
	} else if len(i.AbsoluteEpisodeNumbers) > 0 {
		episodes := []string{}
		for _, episode := range i.AbsoluteEpisodeNumbers {
			episodes = append(episodes, fmt.Sprintf("%03d", episode))
		}
		episodeString = strings.Join(episodes, "-")
	}

	return fmt.Sprintf("%s - %s (%s)", i.SeriesTitle, episodeString, i.Quality)
}
