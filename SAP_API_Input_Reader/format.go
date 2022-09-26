package sap_api_input_reader

import (
	"sap-api-integrations-bill-of-material-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.BillOfMaterial
	results := make([]requests.Item, 0, len(data.BillOfMaterialItem))

	header := sdc.ConvertToHeader()

	for i := range data.BillOfMaterialItem {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.BillOfMaterial
	return &requests.Header{
		BillOfMaterial:                data.BillOfMaterial,
		BillOfMaterialCategory:        data.BillOfMaterialCategory,
		BillOfMaterialVariant:         data.BillOfMaterialVariant,
		BillOfMaterialVersion:         data.BillOfMaterialVersion,
		EngineeringChangeDocument:     data.EngineeringChangeDocument,
		Material:                      data.Material,
		Plant:                         data.Plant,
		BillOfMaterialHeaderUUID:      data.BillOfMaterialHeaderUUID,
		BillOfMaterialVariantUsage:    data.BillOfMaterialVariantUsage,
		EngineeringChangeDocForEdit:   data.EngineeringChangeDocForEdit,
		IsMultipleBOMAlt:              data.IsMultipleBOMAlt,
		BOMHeaderInternalChangeCount:  data.BOMHeaderInternalChangeCount,
		BOMUsagePriority:              data.BOMUsagePriority,
		BillOfMaterialAuthsnGrp:       data.BillOfMaterialAuthsnGrp,
		BOMVersionStatus:              data.BOMVersionStatus,
		IsVersionBillOfMaterial:       data.IsVersionBillOfMaterial,
		IsLatestBOMVersion:            data.IsLatestBOMVersion,
		IsConfiguredMaterial:          data.IsConfiguredMaterial,
		BOMTechnicalType:              data.BOMTechnicalType,
		BOMGroup:                      data.BOMGroup,
		BOMHeaderText:                 data.BOMHeaderText,
		BOMAlternativeText:            data.BOMAlternativeText,
		BillOfMaterialStatus:          data.BillOfMaterialStatus,
		HeaderValidityStartDate:       data.HeaderValidityStartDate,
		HeaderValidityEndDate:         data.HeaderValidityEndDate,
		ChgToEngineeringChgDocument:   data.ChgToEngineeringChgDocument,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
		IsALE:                         data.IsALE,
		MatFromLotSizeQuantity:        data.MatFromLotSizeQuantity,
		MaterialToLotSizeQuantity:     data.MaterialToLotSizeQuantity,
		BOMHeaderBaseUnit:             data.BOMHeaderBaseUnit,
		BOMHeaderQuantityInBaseUnit:   data.BOMHeaderQuantityInBaseUnit,
		RecordCreationDate:            data.RecordCreationDate,
		LastChangeDate:                data.LastChangeDate,
		BOMIsToBeDeleted:              data.BOMIsToBeDeleted,
		DocumentIsCreatedByCAD:        data.DocumentIsCreatedByCAD,
		LaboratoryOrDesignOffice:      data.LaboratoryOrDesignOffice,
		LastChangeDateTime:            data.LastChangeDateTime,
		ProductDescription:            data.ProductDescription,
		PlantName:                     data.PlantName,
		BillOfMaterialHdrDetailsText:  data.BillOfMaterialHdrDetailsText,
		SelectedBillOfMaterialVersion: data.SelectedBillOfMaterialVersion,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataBillOfMaterial := sdc.BillOfMaterial
	data := sdc.BillOfMaterial.BillOfMaterialItem[num]
	return &requests.Item{
		BillOfMaterial:               dataBillOfMaterial.BillOfMaterial,
		BillOfMaterialCategory:       data.BillOfMaterialCategory,
		BillOfMaterialVersion:        data.BillOfMaterialVersion,
		BillOfMaterialItemNodeNumber: data.BillOfMaterialItemNodeNumber,
		HeaderChangeDocument:         data.HeaderChangeDocument,
		Material:                     data.Material,
		Plant:                        data.Plant,
		ValidityStartDate:            data.ValidityStartDate,
		ValidityEndDate:              data.ValidityEndDate,
		BillOfMaterialComponent:      data.BillOfMaterialComponent,
		ComponentDescription:         data.ComponentDescription,
		BillOfMaterialItemQuantity:   data.BillOfMaterialItemQuantity,
		ComponentScrapInPercent:      data.ComponentScrapInPercent,
	}
}
