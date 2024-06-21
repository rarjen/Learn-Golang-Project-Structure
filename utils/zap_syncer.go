package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

// Direct writer into syncer
type UdpDirectWriteSyncer struct {
	InfoHost  string
	ErrorHost string
	DebugHost string
	infoConn  net.Conn
	debugConn net.Conn
	errorConn net.Conn
}

func (syncer *UdpDirectWriteSyncer) Sync() error {
	infoConn, err := net.Dial("udp", syncer.InfoHost)
	if err != nil {
		return err
	}

	errorConn, err := net.Dial("udp", syncer.ErrorHost)
	if err != nil {
		return err
	}

	debugConn, err := net.Dial("udp", syncer.DebugHost)
	if err != nil {
		return err
	}

	syncer.infoConn = infoConn
	syncer.errorConn = errorConn
	syncer.debugConn = debugConn

	return nil
}

func (syncer *UdpDirectWriteSyncer) Write(p []byte) (int, error) {
	go func(content []byte) {
		// Parse Data to find De
		var parsedData map[string]any
		if err := json.Unmarshal(content, &parsedData); err != nil {
			fmt.Println("error unmarshal", err)
			return
		}

		// Cheking if message is JSON
		var parseMessage map[string]any
		message, ok := parsedData["message"].(string)
		if ok {
			if err := json.Unmarshal([]byte(message), &parseMessage); err == nil {
				parsedData["message"] = parseMessage
			}
		}

		// Formatted The JSON String
		formattedContent, err := json.MarshalIndent(&parsedData, "", "    ")
		if err != nil {
			fmt.Println("error formatting content")
			return
		}

		logError := parsedData["log.level"].(string)
		if logError == "error" {
			_, err = syncer.errorConn.Write(formattedContent)
		} else if logError == "debug" {
			_, err = syncer.debugConn.Write(formattedContent)
		} else {
			_, err = syncer.infoConn.Write(formattedContent)
		}

		if err != nil {
			fmt.Println("error send log", err)
		}

		fmt.Println(string(formattedContent))
	}(bytes.Clone(p))

	return 0, nil
}

func (syncer *UdpDirectWriteSyncer) Close() error {
	err := syncer.infoConn.Close()
	if err != nil {
		return err
	}

	err = syncer.errorConn.Close()
	if err != nil {
		return err
	}

	return nil
}
