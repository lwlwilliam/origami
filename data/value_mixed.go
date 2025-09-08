package data

import "fmt"

func NewMixedValue(v interface{}) Value {
	return &MixedValue{
		Value: v,
	}
}

type MixedValue struct {
	Value interface{}
}

func (m *MixedValue) GetValue(ctx Context) (GetValue, Control) {
	return m, nil
}

func (m *MixedValue) AsString() string {
	return fmt.Sprintf("%v", m.Value)
}

func (m *MixedValue) AsInt() (int, error) {
	switch m.Value.(type) {
	case FloatValue:
		return int(m.Value.(FloatValue).Value), nil
	case IntValue:
		return m.Value.(IntValue).Value, nil
	default:
		return 0, nil
	}
}

func (m *MixedValue) AsFloat() (float64, error) {
	switch m.Value.(type) {
	case FloatValue:
		return m.Value.(FloatValue).Value, nil
	case IntValue:
		return float64(m.Value.(IntValue).Value), nil
	default:
		return 0, nil
	}
}

func (m *MixedValue) AsBool() (bool, error) {
	switch m.Value.(type) {
	case BoolValue:
		return m.Value.(BoolValue).Value, nil
	case StringValue:
		return m.Value.(StringValue).Value != "", nil
	case IntValue:
		return m.Value.(IntValue).Value != 0, nil
	case FloatValue:
		return m.Value.(FloatValue).Value != 0, nil
	case ArrayValue:
		return len(m.Value.(ArrayValue).Value) != 0, nil
	default:
		return false, nil
	}
}
