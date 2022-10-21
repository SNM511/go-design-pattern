package factory

import (
	"reflect"
	"testing"
)

func Test_jsonConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		{
			name: "json",
			want: jsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		want ISystemConfigParser
	}{
		{
			name: "json",
			want: jsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAbstractFactory() {
	factory := NewAbstractFactory("HuaWei")
	tv := factory.CreateTelevision()
	air := factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(25)

	factory = NewAbstractFactory("Mi")
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(26)

	// Output:
	// Watch HuaWei TV
	// HuaWei AirConditioner set temperature to 25 ℃
	// Watch Mi TV
	// Mi AirConditioner set temperature to 26 ℃
}
