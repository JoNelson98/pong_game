package utils

// CalculateEasedSpeed computes the current speed based on how long a key has been held,
// interpolating from maxSpeed to cruiseSpeed over accelerationTime.
func CalculateEasedSpeed(duration, maxSpeed, cruiseSpeed, accelerationTime float64) float64 {
	if duration < accelerationTime {
		// Linear interpolation from maxSpeed to cruiseSpeed
		return maxSpeed - (maxSpeed-cruiseSpeed)*(duration/accelerationTime)
	}
	return cruiseSpeed
}

