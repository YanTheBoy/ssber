package model

type Flat struct {
	ID string `pq:"id"`
	Name string `pq:"name"`
	Price int
	Description string
	FlatType string `pq:"flat_type"`
	CreatedBy string `pq:"created_by"`
	CreatedAt string
}
