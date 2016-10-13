package releaseinfo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPossibleSpecialEpisode(t *testing.T) {
	for idx, test := range []string{
		"Under.the.Dome.S02.Special-Inside.Chesters.Mill.HDTV.x264-BAJSKORV",
		"Under.the.Dome.S02.Special-Inside.Chesters.Mill.720p.HDTV.x264-BAJSKORV",
		"Rookie.Blue.Behind.the.Badge.S05.Special.HDTV.x264-2HD",
	} {
		result, err := ParseTitle(test)

		require.NoError(t, err,
			fmt.Sprintf("Row %d should have no parsing error", idx+1))
		require.True(t, result.IsPossibleSpecialEpisode(),
			fmt.Sprintf("Row %d should be a possible special episode", idx+1))
	}
}
