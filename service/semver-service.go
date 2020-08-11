package service

var allowedOperations = []string{"major", "minor", "patch"}
type VersionOps interface {
	Increment(branch string, accessToken string, operation string)
	New(branch string, accessToken string)
}




func IncrementVersion() {
	//TODO: Verify Token Access
}