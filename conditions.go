package geo

import (
	"fmt"
)

// NearCondition stores the point and distance to filter around
type NearCondition struct {
	point Point
	value float64
}

// Near filters by distance around a geographic point
func Near(point Point, value float64) *NearCondition {
	return &NearCondition{point, value}
}

// SQL returns the built SQL to apply the filter.
func (cond *NearCondition) SQL() string {
	const s string = `
		6371 * 2 * ASIN(
			SQRT(
				POWER(
					SIN(
						RADIANS(%v - ABS(Y(position)))/2
					),
					2
				)
				+ COS(RADIANS(%v)) * COS(RADIANS(ABS(Y(position))))
				* POWER(
					SIN(
						RADIANS(%v - X(position))/2
					),
					2
				)
			)
		) < ?
	`
	return fmt.Sprintf(s, cond.point.Lat, cond.point.Lat, cond.point.Lng)
}

// Values returns the SQL placeholders to apply the filter
func (cond *NearCondition) Values() []interface{} {
	return []interface{}{cond.value}
}
