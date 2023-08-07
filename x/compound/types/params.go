package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyNumberOfCompoundsPerBlock            = []byte("NumberOfCompoundsPerBlock")
	KeyMinimumCompoundFrequency             = []byte("MinimumCompoundFrequency")
	DefaultNumberOfCompoundsPerBlock uint64 = 100
	DefaultMinimumCompoundFrequency  uint64 = 600
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(numberOfCompoundsPerBlock uint64, minimumCompoundFrequency uint64) Params {
	return Params{
		NumberOfCompoundsPerBlock: numberOfCompoundsPerBlock,
		MinimumCompoundFrequency:  minimumCompoundFrequency,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultNumberOfCompoundsPerBlock,
		DefaultMinimumCompoundFrequency,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyNumberOfCompoundsPerBlock, &p.NumberOfCompoundsPerBlock, validateNumberOfCompoundsPerBlock),
		paramtypes.NewParamSetPair(KeyMinimumCompoundFrequency, &p.MinimumCompoundFrequency, validateMinimumCompoundFrequency),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateNumberOfCompoundsPerBlock(p.NumberOfCompoundsPerBlock); err != nil {
		return err
	}

	if err := validateMinimumCompoundFrequency(p.MinimumCompoundFrequency); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateNumberOfCompoundsPerBlock validates the NumberOfCompoundsPerBlock param
func validateNumberOfCompoundsPerBlock(v interface{}) error {
	numberOfCompoundsPerBlock, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if numberOfCompoundsPerBlock < 1 {
		return fmt.Errorf("numberOfCompoundsPerBlock can't be less than 1: %d", numberOfCompoundsPerBlock)
	}

	return nil
}

// validateMinimumCompoundFrequency the NumberOfCompoundsPerBlock param
func validateMinimumCompoundFrequency(v interface{}) error {
	minimumCompoundFrequency, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if minimumCompoundFrequency < 1 {
		return fmt.Errorf("MinimumCompoundFrequency can't be less than 1: %d", minimumCompoundFrequency)
	}

	return nil
}
