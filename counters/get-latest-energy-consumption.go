package counters

func GetLatestEnergyConsumption(Consumption []int) int {

	latestEnergyConsumption := 0

	for _, val := range Consumption {
		if val > 0 {
			latestEnergyConsumption = val
		}
	}
	return latestEnergyConsumption
}
