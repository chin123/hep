package fads

import (
	"math"
	"math/rand"
	"reflect"

	"github.com/go-hep/fmom"
	"github.com/go-hep/fwk"
	"github.com/go-hep/random"
)

type EnergySmearing struct {
	fwk.TaskBase

	input  string
	output string

	smear func(eta, ene float64) float64
	seed  int64
	src   rand.Source
}

func (tsk *EnergySmearing) Configure(ctx fwk.Context) fwk.Error {
	var err fwk.Error

	err = tsk.DeclInPort(tsk.input, reflect.TypeOf([]Candidate{}))
	if err != nil {
		return err
	}

	err = tsk.DeclOutPort(tsk.output, reflect.TypeOf([]Candidate{}))
	if err != nil {
		return err
	}

	return err
}

func (tsk *EnergySmearing) StartTask(ctx fwk.Context) fwk.Error {
	var err fwk.Error
	tsk.src = rand.NewSource(tsk.seed)
	return err
}

func (tsk *EnergySmearing) StopTask(ctx fwk.Context) fwk.Error {
	var err fwk.Error

	return err
}

func (tsk *EnergySmearing) Process(ctx fwk.Context) fwk.Error {
	var err fwk.Error
	store := ctx.Store()
	msg := ctx.Msg()

	v, err := store.Get(tsk.input)
	if err != nil {
		return err
	}

	input := v.([]Candidate)
	msg.Debugf(">>> input: %v\n", len(input))

	output := make([]Candidate, 0, len(input))
	defer func() {
		err = store.Put(tsk.output, output)
	}()

	for i := range input {
		cand := &input[i]
		eta := cand.Pos.Eta()
		phi := cand.Pos.Phi()
		ene := cand.Mom.E()

		// apply smearing
		smearEne := random.Gauss(ene, tsk.smear(eta, ene), &tsk.src)
		ene = smearEne()

		if ene <= 0 {
			continue
		}

		mother := cand
		c := cand.Clone()
		eta = cand.Mom.Eta()
		phi = cand.Mom.Phi()
		pt := ene / math.Cosh(eta)

		pxs := pt * math.Cos(phi)
		pys := pt * math.Sin(phi)
		pzs := pt * math.Sinh(eta)

		c.Mom = fmom.NewPxPyPzE(pxs, pys, pzs, ene)
		c.Add(mother)

		output = append(output, *c)
	}

	msg.Debugf(">>> smeared: %v\n", len(output))

	return err
}

func init() {
	fwk.Register(reflect.TypeOf(EnergySmearing{}),
		func(typ, name string, mgr fwk.App) (fwk.Component, fwk.Error) {
			var err fwk.Error
			tsk := &EnergySmearing{
				TaskBase: fwk.NewTask(typ, name, mgr),
				input:    "InputParticles",
				output:   "OutputParticles",
				smear:    func(x, y float64) float64 { return 0 },
				seed:     1234,
			}

			err = tsk.DeclProp("Input", &tsk.input)
			if err != nil {
				return nil, err
			}

			err = tsk.DeclProp("Output", &tsk.output)
			if err != nil {
				return nil, err
			}

			err = tsk.DeclProp("Resolution", &tsk.smear)
			if err != nil {
				return nil, err
			}

			err = tsk.DeclProp("Seed", &tsk.seed)
			if err != nil {
				return nil, err
			}

			return tsk, err
		},
	)
}
