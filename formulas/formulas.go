package formulas

import (
	"geo-conv/constants"
	"math"
)

func LatitudeZoneLetter(lat float64) string {
	zone_index := int((lat + 80) / 8)
	if zone_index < 0 {
		zone_index = 0
	}
	if zone_index > len(constants.LatitudeBands)-1 {
		zone_index = len(constants.LatitudeBands) - 1
	}
	return string(constants.LatitudeBands[zone_index])
}

func ConvertToUTM(lat, lon float64) (easting, northing float64, zone int, zoneLetter string) {
	// zone number based on the longitude
	zone = int((lon+180)/6) + 1
	// longitude of the central meridian (origin) of the zone
	lonOrigin := float64(zone-1)*6 - 180 + 3
	// Convert the longitude of the central meridian to radians
	lonOriginRad := lonOrigin * math.Pi / 180
	// Convert the latitude and the longitude to radians
	latRad := lat * math.Pi / 180
	lonRad := lon * math.Pi / 180

	// the radius of curvature in the prime vertical
	N := constants.A / math.Sqrt(1-math.Pow(constants.E*math.Sin(latRad), 2))
	// the square of the tangent of the latitude (used for corrections)
	T := math.Pow(math.Tan(latRad), 2)
	// the second eccentricity squared, used for projection corrections
	C := math.Pow(constants.E*math.Cos(latRad), 2) / (1 - constants.E*constants.E)
	// the meridional difference between the longitude and the central meridian
	A := math.Cos(latRad) * (lonRad - lonOriginRad)
	// the meridional arc length M, which is a complex function involving the latitude
	M := constants.A * ((1-constants.E*constants.E/4-3*math.Pow(constants.E, 4)/64-5*math.Pow(constants.E, 6)/256)*latRad -
		(3*constants.E*constants.E/8+3*math.Pow(constants.E, 4)/32+45*math.Pow(constants.E, 6)/1024)*math.Sin(2*latRad) +
		(15*math.Pow(constants.E, 4)/256+45*math.Pow(constants.E, 6)/1024)*math.Sin(4*latRad) -
		(35*math.Pow(constants.E, 6)/3072)*math.Sin(6*latRad))

	// the easting (X) coordinate in the UTM system, adjusting for false easting (500,000 meters)
	easting = constants.K0*N*(A+(1-T+C)*math.Pow(A, 3)/6+(5-18*T+T*T+72*C-58*math.Pow(constants.E, 2))*math.Pow(A, 5)/120) + 500000.0
	// the northing (Y) coordinate in the UTM system, adjusting for false northing (10,000,000 meters if below equator)
	northing = constants.K0 * (M + N*math.Tan(latRad)*(A*A/2+(5-T+9*C+4*C*C)*math.Pow(A, 4)/24+(61-58*T+T*T+600*C-330*math.Pow(constants.E, 2))*math.Pow(A, 6)/720))
	// // If the latitude is in the southern hemisphere, add a false northing of 10,000,000 meters
	if lat < 0 {
		northing += 10000000.0
	}

	zoneLetter = LatitudeZoneLetter(lat)
	return
}
