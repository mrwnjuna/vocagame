# 🅿️ Parking Lot System (Go)

A simple command-line Parking Lot system written in Go.  
It simulates an automated parking lot with the following features:
- Allocates parking slots based on proximity to the entrance.
- Calculates parking charges based on time.
- Allows cars to leave and re-enter.
- Prints current parking status.
- Interacts via command file input.

---

## 📦 Features

- Create a parking lot of size `n`
- Park a car by registration number
- Unpark a car and calculate parking charges
- Print current slot status
- Reads commands from a file

---

## 🛠️ How to Run the Project

### 1. Clone the Repository or Copy Files
Ensure the following files are in the same directory:

```
parking-app/
├── main.go
├── parking_lot.go
└── input.txt
```

### 2. Prepare Input File

Example `input.txt`:
```
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
leave KA-01-HH-1234 4
status
```

### 3. Run the Program

Make sure you're in the `parking-app/` directory and run:

```bash
go run . input.txt
```

This reads commands from `input.txt` and prints output to the terminal.

---

## ✅ How to Run the Tests

Unit tests are written using Go’s built-in `testing` package.

### 1. Ensure This File Exists

```
parking_lot_test.go
```

### 2. Run All Tests

```bash
go test
```

To see test coverage:

```bash
go test -cover
```

---

## 🧮 Parking Charges

- First 2 hours: **$10**
- Each additional hour: **+$10/hour**

Example:  
4 hours → $10 (base) + 2×$10 = **$30**

---

## 📄 Example Output

```
Allocated slot number: 1
Allocated slot number: 2
Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30
Slot No. Registration No.
2 KA-01-HH-9999
```

---

## 📌 Notes

- The input file **must be in the same directory** as `main.go`.
- Only the car number is tracked (color is ignored for simplicity).
- Commands are case-sensitive: use lowercase `park`, `leave`, etc.

---