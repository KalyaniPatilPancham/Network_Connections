package connection

import (
	"demo/models"
	"fmt"
)

// findConnections analyzes the fetched data to determine connections between locations
func FindConnections(data *models.Data) []string {
	connections := make(map[string]struct{})

	routerMap := make(map[int]models.Router)
	locationMap := make(map[int]string)

	for _, location := range data.Locations {
		locationMap[location.ID] = location.Name
	}

	for _, router := range data.Routers {
		routerMap[router.ID] = router
	}

	for _, router := range data.Routers {
		for _, linkID := range router.RouterLinks {
			linkedRouter := routerMap[linkID]
			if router.LocationID != linkedRouter.LocationID {
				pair := fmt.Sprintf("%s <-> %s", locationMap[router.LocationID], locationMap[linkedRouter.LocationID])
				if _, exists := connections[pair]; !exists {
					reversePair := fmt.Sprintf("%s <-> %s", locationMap[linkedRouter.LocationID], locationMap[router.LocationID])
					if _, exists := connections[reversePair]; !exists {
						connections[pair] = struct{}{}
					}
				}
			}
		}
	}

	var result []string
	for conn := range connections {
		result = append(result, conn)
	}

	return result
}
