package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestParkingLotBasicFlow(t *testing.T) {
	pl := NewParkingLot(3)

	pl.Park("KA-01-HH-1234")
	pl.Park("KA-01-HH-9999")
	pl.Park("KA-01-BB-0001")

	if len(pl.Slots) != 3 {
		t.Errorf("Expected 3 cars parked, got %d", len(pl.Slots))
	}

	output := captureOutput(func() {
		pl.Park("KA-01-HH-7777")
	})
	if !strings.Contains(output, "Sorry, parking lot is full") {
		t.Error("Expected full parking lot message")
	}

	output = captureOutput(func() {
		pl.Leave("KA-01-BB-0001", 4)
	})
	if !strings.Contains(output, "Charge $30") {
		t.Error("Expected $30 charge for 4 hours")
	}

	pl.Park("KA-01-HH-2701")
	if len(pl.Slots) != 3 {
		t.Errorf("Expected 3 cars after re-parking, got %d", len(pl.Slots))
	}

	// Test leave with invalid car
	output = captureOutput(func() {
		pl.Leave("XX-00-XX-0000", 2)
	})
	if !strings.Contains(output, "not found") {
		t.Error("Expected not found message for unregistered car")
	}
}

func TestStatusOutput(t *testing.T) {
	pl := NewParkingLot(2)
	pl.Park("CAR-001")
	pl.Park("CAR-002")

	output := captureOutput(func() {
		pl.Status()
	})

	if !strings.Contains(output, "Slot No. Registration No.") {
		t.Error("Expected header in status output")
	}
	if !strings.Contains(output, "1 CAR-001") || !strings.Contains(output, "2 CAR-002") {
		t.Error("Expected correct slot-car mapping in status")
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}
