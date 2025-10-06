# Label Printer

Go-based HTTP service for generating and printing labels with QR codes for Monash Automation inventory items.

## Features

- Generates labels with item serial numbers, names, and QR codes
- Includes Monash Automation branding
- Prints to Brother QL-700 label printer via CUPS
- REST API for label printing

## Prerequisites

- Go 1.16+
- Brother QL-700 label printer configured in CUPS
- CUPS printing system (Linux/macOS)

## Installation

1. **Clone the repository:**
```bash
git clone https://github.com/yourusername/Label-Printer.git
cd Label-Printer
```

2. **Install dependencies:**
```bash
go mod download  # Downloads dependencies
go install       # Compiles and installs binary
```

3. **Configure printer:**
   - Ensure Brother QL-700 is installed in CUPS as `brother_ql.700`
   - Verify with: `lpstat -p -d`

## Usage

### Starting the Server

```bash
go run .
```

Server starts on `http://localhost:6767`

### API Endpoint

**POST** `/printer`

Generates and prints a label.

**Request Body:**
```json
{
  "name": "Arduino Uno R3",
  "serial": "SN12345",
  "quantity": 2,
  "itemId": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Parameters:**
- `name` (string, required): Item name displayed on label
- `serial` (string, required): Serial number displayed prominently
- `quantity` (int, required): Number of labels to print (must be > 0)
- `itemId` (string, required): Valid UUID encoded in QR code

**Success Response:**
```json
{
  "ok": true,
  "itemId": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Error Response:**
```json
{
  "ok": false,
  "error": "Error description"
}
```

### Example Request

```bash
curl -X POST http://localhost:6767/printer \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Raspberry Pi 4",
    "serial": "RPI-001",
    "quantity": 1,
    "itemId": "123e4567-e89b-12d3-a456-426614174000"
  }'
```

## Label Specifications

- **Dimensions:** 991 x 306 pixels
- **QR Code:** 241 x 241 pixels (High error correction)
- **Margins:** 30 pixels
- **Output Format:** JPEG (95% quality)
- **Font Sizes:** 
  - Serial: 100pt (bold)
  - Name: 40pt (regular)
  - Header: 30pt (regular)

## Project Structure

```
Label-Printer/
├── main.go           # HTTP server and print handler
├── validation.go     # Request validation
├── format.go         # Label generation logic
├── assets/           # Logo images
├── fonts/            # Custom fonts
└── temp/             # Generated label output
```

## Dependencies

- `github.com/golang/freetype` - Font rendering
- `github.com/nfnt/resize` - Image resizing
- `github.com/skip2/go-qrcode` - QR code generation
- `github.com/google/uuid` - UUID validation
- `golang.org/x/image/font/gofont/*` - Embedded fallback fonts

## License

[Add your license here]
