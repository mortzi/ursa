package data

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UrsaURL something
type UrsaURL struct {
	ID       primitive.ObjectID `json:"_id"`
	URL      string             `json:"url"`
	Category string             `json:"category"`
	Tag      string             `json:"tag"`
}

//String is used when needed to string format
func (url *UrsaURL) String() string {
	return fmt.Sprintf("%s\n\tcategory: %s\n\ttag: %s",
		url.URL,
		url.Category,
		url.Tag)
}

func (url *UrsaURL) toJSON() string {
	j, err := json.Marshal(&url)

	if err != nil {
		return ""
	}

	s := string(j)
	return s
}
