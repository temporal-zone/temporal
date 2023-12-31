package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyNumberOfCompoundsPerBlock           = []byte("NumberOfCompoundsPerBlock")
	KeyMinimumCompoundFrequency            = []byte("MinimumCompoundFrequency")
	KeyCompoundModuleEnabled               = []byte("CompoundModuleEnabled")
	DefaultNumberOfCompoundsPerBlock int64 = 100
	DefaultMinimumCompoundFrequency  int64 = 100
	DefaultCompoundModuleEnabled     bool  = true
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(numberOfCompoundsPerBlock int64, minimumCompoundFrequency int64, compoundModuleEnabled bool) Params {
	return Params{
		NumberOfCompoundsPerBlock: numberOfCompoundsPerBlock,
		MinimumCompoundFrequency:  minimumCompoundFrequency,
		CompoundModuleEnabled:     compoundModuleEnabled,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultNumberOfCompoundsPerBlock,
		DefaultMinimumCompoundFrequency,
		DefaultCompoundModuleEnabled,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyNumberOfCompoundsPerBlock, &p.NumberOfCompoundsPerBlock, validateNumberOfCompoundsPerBlock),
		paramtypes.NewParamSetPair(KeyMinimumCompoundFrequency, &p.MinimumCompoundFrequency, validateMinimumCompoundFrequency),
		paramtypes.NewParamSetPair(KeyCompoundModuleEnabled, &p.CompoundModuleEnabled, validateCompoundModuleEnabled),
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
	numberOfCompoundsPerBlock, ok := v.(int64)
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
	minimumCompoundFrequency, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if minimumCompoundFrequency < 1 {
		return fmt.Errorf("MinimumCompoundFrequency can't be less than 1: %d", minimumCompoundFrequency)
	}

	return nil
}

// validateMinimumCompoundFrequency the NumberOfCompoundsPerBlock param
func validateCompoundModuleEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}
