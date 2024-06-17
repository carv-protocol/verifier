package dcap

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type TcbInfo struct {
	Version                 uint8
	IssueDate               time.Time
	NextUpdate              time.Time
	Fmspc                   []byte
	PceID                   []byte
	TcbType                 uint8
	TcbEvaluationDataNumber uint8
	TcbLevels               []TcbLevelInfo
}

type TcbLevelInfo struct {
	Tcb       Tcb
	TcbDate   time.Time
	TcbStatus string
}

type Tcb struct {
	SgxtcbcompSvn [16]uint8
	Pcesvn        uint16
}

var tcbInfo struct {
	sync.Once
	tcbInfoInstance *TcbInfo
}

// GetTcbInfo initializes or returns the singleton instance of TcbInfo
// GetTcbInfo uses sync.Once to ensure the TcbInfo is loaded only once
func (e *TcbInfo) GetTcbInfo(path string) *TcbInfo {
	tcbInfo.Do(func() {
		var err error
		tcbInfoInstance, err := loadTcbInfoFromUrl(path)
		if err != nil {
			log.Fatalf("Failed to load TCB Info: %v", err)
		}
		tcbInfo.tcbInfoInstance = tcbInfoInstance
	})
	return tcbInfo.tcbInfoInstance
}

func loadTcbInfoFromUrl(url string) (*TcbInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	var root struct {
		TcbInfo struct {
			Version                 uint8  `json:"version"`
			IssueDate               string `json:"issueDate"`
			NextUpdate              string `json:"nextUpdate"`
			Fmspc                   string `json:"fmspc"`
			PceID                   string `json:"pceId"`
			TcbType                 uint8  `json:"tcbType"`
			TcbEvaluationDataNumber uint8  `json:"tcbEvaluationDataNumber"`
			TcbLevels               []struct {
				Tcb struct {
					SgxtcbcompSvn map[string]uint8 `json:""`
					Pcesvn        uint16           `json:"pcesvn"`
				} `json:"tcb"`
				TcbDate   string `json:"tcbDate"`
				TcbStatus string `json:"tcbStatus"`
			} `json:"tcbLevels"`
		} `json:"tcbInfo"`
	}

	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	tcbInfo := &TcbInfo{
		Version:                 root.TcbInfo.Version,
		IssueDate:               parseTime(root.TcbInfo.IssueDate),
		NextUpdate:              parseTime(root.TcbInfo.NextUpdate),
		Fmspc:                   decodeHex(root.TcbInfo.Fmspc),
		PceID:                   decodeHex(root.TcbInfo.PceID),
		TcbType:                 root.TcbInfo.TcbType,
		TcbEvaluationDataNumber: root.TcbInfo.TcbEvaluationDataNumber,
	}

	for _, level := range root.TcbInfo.TcbLevels {
		var compSvn [16]uint8

		for i := 0; i < 16; i++ {
			key := fmt.Sprintf("sgxtcbcomp%02dsvn", i+1)
			if val, ok := level.Tcb.SgxtcbcompSvn[key]; ok {
				compSvn[i] = uint8(val)
			}
		}
		tcbInfo.TcbLevels = append(tcbInfo.TcbLevels, TcbLevelInfo{
			Tcb: Tcb{
				SgxtcbcompSvn: compSvn,
				Pcesvn:        level.Tcb.Pcesvn,
			},
			TcbDate:   parseTime(level.TcbDate),
			TcbStatus: level.TcbStatus,
		})
	}

	return tcbInfo, nil
}

func parseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func decodeHex(s string) []byte {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
