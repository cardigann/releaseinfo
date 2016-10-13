package releaseinfo

var mediaFileExtensions = map[string]Quality{
	".webm":   QualityUnknown,
	".m4v":    QualitySDTV,
	".3gp":    QualitySDTV,
	".nsv":    QualitySDTV,
	".ty":     QualitySDTV,
	".strm":   QualitySDTV,
	".rm":     QualitySDTV,
	".rmvb":   QualitySDTV,
	".m3u":    QualitySDTV,
	".ifo":    QualitySDTV,
	".mov":    QualitySDTV,
	".qt":     QualitySDTV,
	".divx":   QualitySDTV,
	".xvid":   QualitySDTV,
	".bivx":   QualitySDTV,
	".nrg":    QualitySDTV,
	".pva":    QualitySDTV,
	".wmv":    QualitySDTV,
	".asf":    QualitySDTV,
	".asx":    QualitySDTV,
	".ogm":    QualitySDTV,
	".ogv":    QualitySDTV,
	".m2v":    QualitySDTV,
	".avi":    QualitySDTV,
	".bin":    QualitySDTV,
	".dat":    QualitySDTV,
	".dvr-ms": QualitySDTV,
	".mpg":    QualitySDTV,
	".mpeg":   QualitySDTV,
	".mp4":    QualitySDTV,
	".avc":    QualitySDTV,
	".vp3":    QualitySDTV,
	".svq3":   QualitySDTV,
	".nuv":    QualitySDTV,
	".viv":    QualitySDTV,
	".dv":     QualitySDTV,
	".fli":    QualitySDTV,
	".flv":    QualitySDTV,
	".wpl":    QualitySDTV,
	".img":    QualityDVD,
	".iso":    QualityDVD,
	".vob":    QualityDVD,
	".mkv":    QualityHDTV720p,
	".ts":     QualityHDTV720p,
	".wtv":    QualityHDTV720p,
	".m2ts":   QualityBluray720p,
}

//         public static HashSet<string> Extensions
//         {
//             get { return new HashSet<string>(_fileExtensions.Keys); }
//         }

//         public static Quality GetQualityForExtension(string extension)
//         {
//             if (_fileExtensions.ContainsKey(extension))
//             {
//                 return _fileExtensions[extension];
//             }

//             return QualityUnknown;
//         }
//     }
// }
