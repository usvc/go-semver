package semver

const (
	// RegexpWithCaptureGroup is a string regular expression that
	// also captures the capture groups [
	// 	Prefix,
	// 	Major,
	// 	Minor,
	// 	Patch,
	// 	PreRelease,
	// 	BuildMetadata,
	// ]. this expression is a modified form from https://semver.org
	RegexpWithCaptureGroup = `(?:(?P<Prefix>[vV]))?(?P<Major>0|[1-9]\d*)\.(?P<Minor>0|[1-9]\d*)\.(?P<Patch>0|[1-9]\d*)(?:-(?P<PreRelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<BuildMetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	parseDecimal           = 10
	parse32bit             = 32
)
