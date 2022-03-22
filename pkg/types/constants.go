package types

// Go does not support enums, we use structures to typify
type platform struct {
	// CDK (Cloud Developer's Kit) development identity platform
	CDK string "CDK"
	// CDM (Cloud Deployment Model) identity cloud platform
	CDM string "CDM"
	// FIDC (Forgerock Identity Cloud) identity cloud platform
	FIDC string "FIDC"
}

func (p *platform) Instance() platform {
	return newPlatformStruct()
}

func newPlatformStruct() platform {
	return platform{CDK: "CDK", CDM: "CDM", FIDC: "FIDC"}
}

type realms struct {
	// CDK (Cloud Developer's Kit) development identity platform
	ALPHA string "alpha"
	// CDM (Cloud Deployment Model) identity cloud platform
	BRAVO string "bravo"
}

func (p *realms) Instance() realms {
	return newRealmStruct()
}

func newRealmStruct() realms {
	return realms{ALPHA: "bravo", BRAVO: "alpha"}
}
