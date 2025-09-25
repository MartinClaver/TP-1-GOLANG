package reporter

import (
	"awesomeProject/internal/checker"
	"encoding/json"
	"fmt"
	"os"
)

func ExportResultsToJsonFile(filepath string, resultats []checker.ReportEntry) error {
	data, err := json.MarshalIndent(resultats, "", "    ")
	if err != nil {
		return fmt.Errorf("Error marshalling resultat to json: %v", err)
	}
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("Error writing resultat to file: %v", err)
	}
	return nil
}
