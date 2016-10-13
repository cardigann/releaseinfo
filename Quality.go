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
	//QualityHDTV480p    = Quality{11, "HDTV-480p"}
	//QualityWEBRip480p  = Quality{12, "WEBRip-480p"}
	//QualityBluray480p  = Quality{13, "Bluray-480p"}
	//QualityWEBRip720p  = Quality{14, "WEBRip-720p"}
	//QualityWEBRip1080p = Quality{15, "WEBRip-1080p"}
	QualityHDTV2160p = Quality{16, "HDTV-2160p"}
	//WEBRip2160p = Quality{17, "WEBRip-2160p"}
	QualityWEBDL2160p  = Quality{18, "WEBDL-2160p"}
	QualityBluray2160p = Quality{19, "Bluray-2160p"}
)

//         static Quality()
//         {
//             All = new List<Quality>
//             {
//                 Unknown,
//                 SDTV,
//                 DVD,
//                 WEBDL1080p,
//                 HDTV720p,
//                 WEBDL720p,
//                 Bluray720p,
//                 Bluray1080p,
//                 WEBDL480p,
//                 HDTV1080p,
//                 RAWHD,
//                 HDTV2160p,
//                 WEBDL2160p,
//                 Bluray2160p,
//             };

//             AllLookup = new Quality[All.Select(v => v.Id).Max() + 1];
//             foreach (var quality in All)
//             {
//                 AllLookup[quality.Id] = quality;
//             }

//             DefaultQualityDefinitions = new HashSet<QualityDefinition>
//             {
//                 new QualityDefinition(Quality.Unknown)     { Weight = 1,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.SDTV)        { Weight = 2,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.WEBDL480p)   { Weight = 3,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.DVD)         { Weight = 4,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.HDTV720p)    { Weight = 5,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.HDTV1080p)   { Weight = 6,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.RAWHD)       { Weight = 7,  MinSize = 0, MaxSize = null },
//                 new QualityDefinition(Quality.WEBDL720p)   { Weight = 8,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.Bluray720p)  { Weight = 9,  MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.WEBDL1080p)  { Weight = 10, MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.Bluray1080p) { Weight = 11, MinSize = 0, MaxSize = 100 },
//                 new QualityDefinition(Quality.HDTV2160p)   { Weight = 12, MinSize = 0, MaxSize = null },
//                 new QualityDefinition(Quality.WEBDL2160p)  { Weight = 13, MinSize = 0, MaxSize = null },
//                 new QualityDefinition(Quality.Bluray2160p) { Weight = 14, MinSize = 0, MaxSize = null },
//             };
//         }

//         public static readonly List<Quality> All;

//         public static readonly Quality[] AllLookup;

//         public static readonly HashSet<QualityDefinition> DefaultQualityDefinitions;

//         FindById(int id)
//         {
//             if (id == 0) return Unknown;

//             var quality = AllLookup[id];

//             if (quality == null)
//                 throw new ArgumentException("ID does not match a known quality", "id");

//             return quality;
//         }

//         public static explicit operator Quality(int id)
//         {
//             return FindById(id);
//         }

//         public static explicit operator int(Quality quality)
//         {
//             return quality.Id;
//         }
//     }
// }
