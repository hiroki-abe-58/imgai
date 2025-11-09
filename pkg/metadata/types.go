package metadata

// ExifData holds important EXIF information
type ExifData struct {
	Make         string
	Model        string
	DateTime     string
	Orientation  string
	Width        string
	Height       string
	FocalLength  string
	Aperture     string
	ISO          string
	ShutterSpeed string
	GPS          *GPSData
}

// GPSData holds GPS coordinates
type GPSData struct {
	Latitude  float64
	Longitude float64
}

// StripOptions holds options for stripping metadata
type StripOptions struct {
	Output string
}

// HasGPS returns true if GPS data is available
func (e *ExifData) HasGPS() bool {
	return e.GPS != nil
}

// IsEmpty returns true if no EXIF data is present
func (e *ExifData) IsEmpty() bool {
	return e.Make == "" && e.Model == "" && e.DateTime == "" &&
		e.FocalLength == "" && e.Aperture == "" && e.ISO == "" &&
		e.ShutterSpeed == "" && !e.HasGPS()
}
