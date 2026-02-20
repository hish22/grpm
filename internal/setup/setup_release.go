package setup

// Create a switch for different file extensions
// based on the extensions, we perfom the setup phase

func SetupAsset(repo string, loaction string, ext string, assetID int, force bool) {
	switch ext {
	case "targz":
		unzipFileTarGz(repo, loaction, assetID, force)
	case "tarzst":
		unzipFileTarZst(repo, loaction, assetID, force)
	case "zip":
		unzipZip(loaction, assetID)
	default:
		MoveBinary(repo, loaction, assetID, force)
	}
}
