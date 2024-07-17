package models

// Data represents the structure of the JSON data fetched from the API
type Data struct {
	Routers   []Router   `json:"routers"`
	Locations []Location `json:"locations"`
}

// Router represents a single router
type Router struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LocationID  int    `json:"location_id"`
	RouterLinks []int  `json:"router_links"`
}

// Location represents a single location
type Location struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Postcode string `json:"postcode"`
}
