package setup

// Create a switch for different file extensions
// based on the extensions, we perfom the setup phase

func SetupAsset(repo string, loaction string, ext string, assetID int) {
	switch ext {
	case "targz":
		unzipFileTarGz(repo, loaction, assetID)
	case "tarzst":
		unzipFileTarZst(repo, loaction, assetID)
	case "zip":
		unzipZip(loaction, assetID)
	}
}
