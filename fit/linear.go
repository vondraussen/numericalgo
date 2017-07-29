package fit

import (
	"github.com/DzananGanic/numericalgo"
)

type Linear struct {
	Base
}

func NewLinear() *Linear {
	lf := &Linear{}
	return lf
}

func (l *Linear) Fit(x numericalgo.Vector, y numericalgo.Vector) error {
	xMatrix := numericalgo.Matrix{x}
	yMatrix := numericalgo.Matrix{y}

	xT, err := xMatrix.Transpose()

	if err != nil {
		return err
	}

	ones := make(numericalgo.Vector, x.Dim())
	for i := range ones {
		ones[i] = 1
	}

	X, err := xT.AddColumnAt(0, ones)

	if err != nil {
		return err
	}

	Y, err := yMatrix.Transpose()

	if err != nil {
		return err
	}

	coeff, err := X.LeftDivide(Y)

	if err != nil {
		return err
	}

	//l.coeff = numericalgo.Vector{coeff[1][0], coeff[0][0]}
	l.coeff, err = coeff.GetColumnAt(0)

	if err != nil {
		return err
	}

	return nil
}

func (l *Linear) Predict(val float64) float64 {
	c := l.Coef()
	return c[1]*val + c[0]
}
