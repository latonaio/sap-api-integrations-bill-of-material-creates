package requests

type HeaderItem struct {
	Header
	ToItem `json:"to_BillOfMaterialItem"`
}

type ToItem struct {
	Results []Item `json:"results"`
}
