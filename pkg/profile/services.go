package profile

type OscService = string

const (
	OscServiceApi        OscService = "api"
	OscServiceOKS        OscService = "oks"
	OscServiceLBU        OscService = "lbu"
	OscServiceOOS        OscService = "oos"
	OscServiceFCU        OscService = "fcu"
	OscServiceEIM        OscService = "eim"
	OscServiceDirectLink OscService = "direct_link"
)
