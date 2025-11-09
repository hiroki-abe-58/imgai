package metadata

import (
	"fmt"
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

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

// ReadExif reads EXIF data from an image file
func ReadExif(path string) (*ExifData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode EXIF: %w", err)
	}

	data := &ExifData{}

	// Extract common fields
	if make, err := x.Get(exif.Make); err == nil {
		if val, err := make.StringVal(); err == nil {
			data.Make = strings.TrimSpace(val)
		}
	}
	if model, err := x.Get(exif.Model); err == nil {
		if val, err := model.StringVal(); err == nil {
			data.Model = strings.TrimSpace(val)
		}
	}
	if dateTime, err := x.Get(exif.DateTime); err == nil {
		if val, err := dateTime.StringVal(); err == nil {
			data.DateTime = strings.TrimSpace(val)
		}
	}
	if orientation, err := x.Get(exif.Orientation); err == nil {
		if val, err := orientation.Int(0); err == nil {
			data.Orientation = fmt.Sprintf("%d", val)
		}
	}
	if width, err := x.Get(exif.PixelXDimension); err == nil {
		if val, err := width.Int(0); err == nil {
			data.Width = fmt.Sprintf("%d", val)
		}
	}
	if height, err := x.Get(exif.PixelYDimension); err == nil {
		if val, err := height.Int(0); err == nil {
			data.Height = fmt.Sprintf("%d", val)
		}
	}
	if focalLength, err := x.Get(exif.FocalLength); err == nil {
		num, denom, err := focalLength.Rat2(0)
		if err == nil && denom != 0 {
			data.FocalLength = fmt.Sprintf("%.1fmm", float64(num)/float64(denom))
		}
	}
	if aperture, err := x.Get(exif.FNumber); err == nil {
		num, denom, err := aperture.Rat2(0)
		if err == nil && denom != 0 {
			data.Aperture = fmt.Sprintf("f/%.1f", float64(num)/float64(denom))
		}
	}
	if iso, err := x.Get(exif.ISOSpeedRatings); err == nil {
		if val, err := iso.Int(0); err == nil {
			data.ISO = fmt.Sprintf("ISO %d", val)
		}
	}
	if shutterSpeed, err := x.Get(exif.ExposureTime); err == nil {
		num, denom, err := shutterSpeed.Rat2(0)
		if err == nil {
			if num < denom {
				data.ShutterSpeed = fmt.Sprintf("1/%d", denom/num)
			} else {
				data.ShutterSpeed = fmt.Sprintf("%d", num/denom)
			}
		}
	}

	// Extract GPS data
	lat, lon, err := x.LatLong()
	if err == nil {
		data.GPS = &GPSData{
			Latitude:  lat,
			Longitude: lon,
		}
	}

	return data, nil
}

// FormatExif formats EXIF data as a human-readable string
func FormatExif(data *ExifData) string {
	var sb strings.Builder

	if data.Make != "" || data.Model != "" {
		sb.WriteString(fmt.Sprintf("Camera: %s %s\n", data.Make, data.Model))
	}
	if data.DateTime != "" {
		sb.WriteString(fmt.Sprintf("Date: %s\n", data.DateTime))
	}
	if data.Width != "" && data.Height != "" {
		sb.WriteString(fmt.Sprintf("Dimensions: %s x %s\n", data.Width, data.Height))
	}
	if data.FocalLength != "" {
		sb.WriteString(fmt.Sprintf("Focal Length: %s\n", data.FocalLength))
	}
	if data.Aperture != "" {
		sb.WriteString(fmt.Sprintf("Aperture: %s\n", data.Aperture))
	}
	if data.ShutterSpeed != "" {
		sb.WriteString(fmt.Sprintf("Shutter Speed: %ss\n", data.ShutterSpeed))
	}
	if data.ISO != "" {
		sb.WriteString(fmt.Sprintf("ISO: %s\n", data.ISO))
	}
	if data.Orientation != "" {
		sb.WriteString(fmt.Sprintf("Orientation: %s\n", data.Orientation))
	}
	if data.GPS != nil {
		sb.WriteString(fmt.Sprintf("GPS: %.6f, %.6f\n", data.GPS.Latitude, data.GPS.Longitude))
	}

	return sb.String()
}
