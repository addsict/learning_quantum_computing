package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type Cbit struct {
	bit uint8
}

type Qubit struct {
	initialState Cbit
	alpha        float64 // superposition coefficent of |0>
	beta         float64 // superposition coefficent of |1>
}

func NewQubit(alpha, beta float64) (*Qubit, error) {
	if alpha*alpha+beta*beta != 1.0 {
		return nil, errors.New("coefficients error")
	}
	return &Qubit{
		initialState: Cbit{bit: 0},
		alpha:        alpha,
		beta:         beta,
	}, nil
}

func (q *Qubit) Measure() Cbit {
	f := rand.Float64()
	if f <= q.alpha*q.alpha {
		return Cbit{0}
	} else {
		return Cbit{1}
	}
}

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	q1, err := NewQubit(1.0, 0.0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Measure qubit:%v, result:%v\n", q1, q1.Measure())

	q2, err := NewQubit(0.0, 1.0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Measure qubit:%v, result:%v\n", q2, q2.Measure())

	q3, err := NewQubit(0.8, 0.6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Measure qubit:%v, result:%v\n", q3, q3.Measure())
}
