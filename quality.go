package releaseinfo

type Quality struct {
	Id   int
	Name string
}

func (q Quality) String() string {
	return q.Name
}

var (
	QualityUnknown     = Quality{0, "Unknown"}
	QualitySDTV        = Quality{1, "SDTV"}
	QualityDVD         = Quality{2, "DVD"}
	QualityWEBDL1080p  = Quality{3, "WEBDL-1080p"}
	QualityHDTV720p    = Quality{4, "HDTV-720p"}
	QualityWEBDL720p   = Quality{5, "WEBDL-720p"}
	QualityBluray720p  = Quality{6, "Bluray-720p"}
	QualityBluray1080p = Quality{7, "Bluray-1080p"}
	QualityWEBDL480p   = Quality{8, "WEBDL-480p"}
	QualityHDTV1080p   = Quality{9, "HDTV-1080p"}
	QualityRAWHD       = Quality{10, "Raw-HD"}
	QualityHDTV2160p   = Quality{16, "HDTV-2160p"}
	QualityWEBDL2160p  = Quality{18, "WEBDL-2160p"}
	QualityBluray2160p = Quality{19, "Bluray-2160p"}
)
