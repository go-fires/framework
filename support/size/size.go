package size //nolint:typecheck

import "fmt"

// Size is a size type
type Size uint64

const (
	// B is 1 byte
	B Size = 1 << (iota * 10)
	// KB is 1024 bytes
	KB
	// MB is 1024 * 1024 bytes
	MB
	// GB is 1024 * 1024 * 1024 bytes
	GB
	// TB is 1024 * 1024 * 1024 * 1024 bytes
	TB
	// PB is 1024 * 1024 * 1024 * 1024 * 1024 bytes
	PB
	// EB is 1024 * 1024 * 1024 * 1024 * 1024 * 1024 bytes
	EB
)

// Bytes return the size in bytes
func (s Size) Bytes() uint64 {
	return uint64(s)
}

// KB returns the size in kilobytes
func (s Size) KB() float64 {
	return float64(s / KB)
}

// MB returns the size in megabytes
func (s Size) MB() float64 {
	return float64(s / MB)
}

// GB returns the size in gigabytes
func (s Size) GB() float64 {
	return float64(s / GB)
}

// TB returns the size in terabytes
func (s Size) TB() float64 {
	return float64(s / TB)
}

// PB returns the size in petabytes
func (s Size) PB() float64 {
	return float64(s / PB)
}

// EB returns the size in exabytes
func (s Size) EB() float64 {
	return float64(s / EB)
}

// BString returns the size as a string in bytes
func (s Size) BString() string {
	return format(float64(s), "B")
}

// KBString returns the size as a string in kilobytes
func (s Size) KBString() string {
	return format(s.KB(), "KB")
}

// MBString returns the size as a string in megabytes
func (s Size) MBString() string {
	return format(s.MB(), "MB")
}

// GBString returns the size as a string in gigabytes
func (s Size) GBString() string {
	return format(s.GB(), "GB")
}

// TBString returns the size as a string in terabytes
func (s Size) TBString() string {
	return format(s.TB(), "TB")
}

// PBString returns the size as a string in petabytes
func (s Size) PBString() string {
	return format(s.PB(), "PB")
}

// EBString returns the size as a string in exabytes
func (s Size) EBString() string {
	return format(s.EB(), "EB")
}

// String returns the size as a string
func (s Size) String() string {
	switch {
	case s >= EB:
		return s.EBString()
	case s >= PB:
		return s.PBString()
	case s >= TB:
		return s.TBString()
	case s >= GB:
		return s.GBString()
	case s >= MB:
		return s.MBString()
	case s >= KB:
		return s.KBString()
	}

	return s.BString()
}

// Format formats the size
func format(yb float64, s string) string {
	return fmt.Sprintf("%.02f %s", yb, s)
}
