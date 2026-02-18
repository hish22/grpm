package setup

// Create a switch for different file extensions
// based on the extensions, we perfom the setup phase

func SetupAsset(loaction string, ext string, assetID int) {
	switch ext {
	case "targz":
		unzipFileTarGz(loaction, assetID)
	case "tarzst":
		unzipFileTarZst(loaction, assetID)
	case "zip":
		unzipZip(loaction, assetID)
	}
}
