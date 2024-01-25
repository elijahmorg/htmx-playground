package lib


// type StateEnum struct {
// 	State   string
// 	Display string
// }

// var (
// 	InUse            = StateEnum{"in_use", "In Use"}
// 	Available        = StateEnum{"available", "Available"}
// 	Offline          = StateEnum{"offline", "Offline"}
// 	UnderMaintenance = StateEnum{"under_maintenance", "Under Maintenance"}
// 	NeedsAttention   = StateEnum{"needs_attention", "Needs Attention"}
// )

//go:generate go-enum  --ptr --marshal --flag --nocase --mustparse --sqlnullstr --sql --names --values --nocomments

// ENUM(in_use, available, offline, under_maintenance, needs_attention)
type StateEnum string


var (
	StateDisplay = map[StateEnum]string{
		StateEnumInUse: "In Use",
		StateEnumAvailable: "Available",
		StateEnumOffline: "Offline",
		StateEnumUnderMaintenance: "Under Maintenance",
		StateEnumNeedsAttention: "Needs Attention",
	}
)
