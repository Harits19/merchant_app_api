package transaction

func OmzetWithoutOutlets(omzets []OmzetPerDayModel) []OmzetPerDayModel {
	for i := range omzets {
		omzets[i].OutletName = ""
	}
	return omzets
}
