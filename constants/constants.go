package constants

import "math"

const (
	// Semi-major axis: the equatorial radius of the Earth in meters
	A = 6378137.0

	// flattening factor: how much the Earth is squashed at the poles
	F = 1 / 298.257223563

	// UTM scale factor
	K0 = 0.9996

	// the UTM zone letters
	LatitudeBands = "CDEFGHJKLMNPQRSTUVWX"
)

// the eccentricity of the ellipsoid, which describes how elliptical the Earth is
var E = math.Sqrt(F * (2 - F))
