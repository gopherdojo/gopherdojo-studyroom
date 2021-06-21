module github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima

go 1.16

replace (
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/download => ./download
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/getheader => ./getheader
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options => ./options
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request => ./request
)

require (
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/getheader v0.0.0-00010101000000-000000000000
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options v0.0.0-00010101000000-000000000000
	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request v0.0.0-00010101000000-000000000000
)
