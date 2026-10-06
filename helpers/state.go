package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const ConfigFile = "quickshare.json"

type Record struct {
	IP     string
	Status string
}

var records []Record

func AddRecord(ip string, status string) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if err != nil {
		_, _ = os.Create(path)
		_ = os.WriteFile(path, []byte("[]"), 0o644)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &records)
	if err != nil {
		return err
	}

	for _, record := range records {
		if record.IP == ip {
			return errors.New("ip already exists and it is in " + record.Status + " status")
		}
	}

	record := Record{
		IP:     ip,
		Status: status,
	}

	records = append(records, record)

	json, err := json.Marshal(records)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, json, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func CheckIfIPExists(ip string) (error, bool) {
	path, err := getConfigPath()
	if err != nil {
		return err, false
	}

	_, err = os.Stat(path)
	if err != nil {
		_, _ = os.Create(path)
		_ = os.WriteFile(path, []byte("[]"), 0o644)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err, false
	}

	err = json.Unmarshal(content, &records)
	if err != nil {
		return err, false
	}

	for _, record := range records {
		if record.IP == ip {
			return nil, true
		}
	}

	return errors.New("device with that ip address has not send a pairing request to accept"), false
}

func UpdateRecord(ip string, status string) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if err != nil {
		_, _ = os.Create(path)
		_ = os.WriteFile(path, []byte("[]"), 0o644)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &records)
	if err != nil {
		return err
	}

	recordFound := false
	for _, r := range records {
		if r.IP == ip && r.Status == "pending" {
			recordFound = true
		}
	}

	if !recordFound {
		return errors.New("Record with the given ip address not found")
	}

	updatedRecords := []Record{}
	for _, r := range records {
		if r.IP == ip {
			r.Status = status
		}
		updatedRecords = append(updatedRecords, r)
	}

	json, err := json.Marshal(updatedRecords)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, json, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homeDir, ConfigFile)
	return path, nil
}
