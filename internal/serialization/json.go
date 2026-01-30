package serialization

import (
	"encoding/json"
	"log"
)

func JsonSerialization(response any) []byte {
	marsh, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Can't marshal JSON response int JSON encoding")
	}
	return marsh
}

func JsonUnserialization(buf []byte, structure any) {
	if err := json.Unmarshal(buf, &structure); err != nil {
		log.Fatal("Can't decode structure into JSON data, ", err)
	}
}
