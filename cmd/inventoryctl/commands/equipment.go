package commands

type EquipmentCmd struct {
	Get  EquipmentGetCmd  `cmd:"" help:"Get equipment by ID"`
	List EquipmentListCmd `cmd:"" help:"List all equipment"`
}

type EquipmentGetCmd struct {
	ID string `arg:"" help:"Equipment unique ID"`
}

type EquipmentListCmd struct {
	Limit int `default:"50" help:"The maximum number of equipment results to return"`
}
