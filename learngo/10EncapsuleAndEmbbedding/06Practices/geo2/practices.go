// Solved by Github YoonBaek

package geo2

import (
	"errors"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

type Landmark struct {
	name string
	Coordinates
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

func (c *Coordinates) SetLatitude(l float64) error {
	if l < -90 || l > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = l
	return nil
}

func (c *Coordinates) SetLongitude(l float64) error {
	if l < -180 || l > 180 {
		return errors.New("invalid latitude")
	}
	c.longitude = l
	return nil
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		errors.New("invalid name")
	}
	l.name = name
	return nil
}

func (l *Landmark) Name() string {
	return l.name
}
