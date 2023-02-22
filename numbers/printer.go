package numbers

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var ErrNegativeRange = errors.New("from and to values should be positive numbers")

type Stringer func(n int) string

type Printer struct {
	// From is start number, inclusive
	From int
	// To is end number, exclusive
	To        int
	Stringify Stringer

	out io.Writer
}

func (p *Printer) WithStringer(fn Stringer) *Printer {
	p.Stringify = fn
	return p
}

func (p *Printer) Print() {
	reversed := p.From > p.To
	if reversed {
		for i := p.From - 1; i <= p.To; i-- {
			fmt.Fprintln(p.out, p.Stringify(i))
		}
		return
	}

	for i := p.From; i < p.To; i++ {
		fmt.Fprintln(p.out, p.Stringify(i))
	}
}

func NewPrinter(from, to int) (*Printer, error) {
	if from < 0 || to < 0 {
		return nil, ErrNegativeRange
	}

	return &Printer{
		From: from,
		To:   to,
		Stringify: func(n int) string {
			return strconv.Itoa(n)
		},

		out: os.Stdout,
	}, nil
}
