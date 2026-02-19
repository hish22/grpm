package serialization

import (
	"encoding/json"

	charmlog "github.com/charmbracelet/log"
)

func JsonSerialization(response any) []byte {
	marsh, err := json.Marshal(response)
	if err != nil {
		charmlog.Error("Failed to marshal JSON response into JSON encoding", "error", err)
	}
	return marsh
}

func JsonUnserialization(buf []byte, structure any) {
	if err := json.Unmarshal(buf, &structure); err != nil {
		charmlog.Error("Failed to decode structure into JSON data", "error", err)
	}
}
