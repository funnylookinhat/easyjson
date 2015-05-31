package easyjson

import (
	"./"
	"reflect"
	"testing"
)

var testJson = []string{
	`
{
	"a": "AAA",
	"b": "BBB",
	"c": "CCC",
	"d": {
		"da": "DADADA",
		"db": "DBDBDB",
		"dc": "DCDCDC",
		"dd": {
			"dda": "DDADDADDA",
			"ddb": "DDBDDBDDB",
			"ddc": "DDCDDCDDC"
		}
	},
	"e": false,
	"ee": true,
	"f": 123456,
	"ff": 123.456,
	"g": null,
	"h": [
		{
			"q": "q1q1q1",
			"w": "w1w1w1",
			"e": "e1e1e1"
		},
		{
			"q": "q2q2q2",
			"w": "w2w2w2",
			"e": "e2e2e2"
		},
		{
			"q": "q3q3q3",
			"w": "w3w3w3",
			"e": "e3e3e3"
		}
	],
	"i": [
		[
			1,
			2,
			3
		],
		[
			4,
			5,
			6
		],
		[
			7,
			8,
			9
		]
	]
}
`,

	`
[
	{
		"a": "a1",
		"b": "b1",
		"c": "c1",
		"d": {
			"da": "da1",
			"db": "db1",
			"dc": "dc1"
		}
	},
	{
		"a": "a2",
		"b": "b2",
		"c": "c2",
		"d": {
			"da": "da2",
			"db": "db2",
			"dc": "dc2"
		}
	},
	{
		"a": "a3",
		"b": "b3",
		"c": "c3",
		"d": {
			"da": "da3",
			"db": "db3",
			"dc": "dc3"
		}
	}
]`,
}

type testGetStringCase struct {
	jsonIndex int
	keys      []interface{}
	value     string
	err       bool
}

var testGetStringCases = []testGetStringCase{
	{
		0,
		[]interface{}{"a"},
		"AAA",
		false,
	},
	{
		0,
		[]interface{}{"b"},
		"BBB",
		false,
	},
	{
		0,
		[]interface{}{"d", "da"},
		"DADADA",
		false,
	},
	{
		0,
		[]interface{}{"d", "db"},
		"DBDBDB",
		false,
	},
	{
		0,
		[]interface{}{"d", "dd", "ddc"},
		"DDCDDCDDC",
		false,
	},
	{
		0,
		[]interface{}{"d", "zzdd", "ddc"},
		"",
		true,
	},
	{
		0,
		[]interface{}{"h", 0, "w"},
		"w1w1w1",
		false,
	},
	{
		0,
		[]interface{}{"h", 1, "e"},
		"e2e2e2",
		false,
	},
	{
		0,
		[]interface{}{"f"},
		"",
		true,
	},
	{
		1,
		[]interface{}{0, "c"},
		"c1",
		false,
	},
	{
		1,
		[]interface{}{1, "d", "da"},
		"da2",
		false,
	},
}

func TestGetString(t *testing.T) {
	jsonDatas := make([]easyjson.JSData, 0)

	for _, json := range testJson {
		jsonData, err := easyjson.DecodeJson([]byte(json))

		if err != nil {
			t.Fatalf("Could not decode test json.\n")
		}

		jsonDatas = append(jsonDatas, jsonData)
	}

	for _, test := range testGetStringCases {
		value, err := easyjson.GetString(jsonDatas[test.jsonIndex], test.keys...)

		if err != nil {
			if !test.err {
				t.Errorf("TestGetString Failed: Unexpected error.\n%v\n", test)
				t.Errorf("%s", err)
			}
		}

		if value != test.value {
			t.Errorf("TestGetString Failed: %s != %s\n%v", test.value, value, test)
		}
	}
}

type testGetFloatCase struct {
	jsonIndex int
	keys      []interface{}
	value     float64
	err       bool
}

var testGetFloatCases = []testGetFloatCase{
	{
		0,
		[]interface{}{"f"},
		123456,
		false,
	},
	{
		0,
		[]interface{}{"ff"},
		123.456,
		false,
	},
	{
		0,
		[]interface{}{"a"},
		0,
		true,
	},
	{
		0,
		[]interface{}{"i", 0, 1},
		2,
		false,
	},
}

func TestGetFloat(t *testing.T) {
	jsonDatas := make([]easyjson.JSData, 0)

	for _, json := range testJson {
		jsonData, err := easyjson.DecodeJson([]byte(json))

		if err != nil {
			t.Fatalf("Could not decode test json.\n")
		}

		jsonDatas = append(jsonDatas, jsonData)
	}

	for _, test := range testGetFloatCases {
		value, err := easyjson.GetFloat(jsonDatas[test.jsonIndex], test.keys...)

		if err != nil {
			if !test.err {
				t.Errorf("TestGetFloat Failed: Unexpected error.\n%v\n", test)
				t.Errorf("%s", err)
			}
		}

		if value != test.value {
			t.Errorf("TestGetString Failed: %f != %f\n%v", test.value, value, test)
		}
	}
}

type testGetBoolCase struct {
	jsonIndex int
	keys      []interface{}
	value     bool
	err       bool
}

var testGetBoolCases = []testGetBoolCase{
	{
		0,
		[]interface{}{"e"},
		false,
		false,
	},
	{
		0,
		[]interface{}{"ee"},
		true,
		false,
	},
	{
		0,
		[]interface{}{"a"},
		false,
		true,
	},
}

func TestGetBool(t *testing.T) {
	jsonDatas := make([]easyjson.JSData, 0)

	for _, json := range testJson {
		jsonData, err := easyjson.DecodeJson([]byte(json))

		if err != nil {
			t.Fatalf("Could not decode test json.\n")
		}

		jsonDatas = append(jsonDatas, jsonData)
	}

	for _, test := range testGetBoolCases {
		value, err := easyjson.GetBool(jsonDatas[test.jsonIndex], test.keys...)

		if err != nil {
			if !test.err {
				t.Errorf("TestGetBool Failed: Unexpected error.\n%v\n", test)
				t.Errorf("%s", err)
			}
		}

		if value != test.value {
			t.Errorf("TestGetString Failed: %t != %t\n%v", test.value, value, test)
		}
	}
}

type testGetSliceCase struct {
	jsonIndex int
	keys      []interface{}
	value     []interface{}
	err       bool
}

var testGetSliceCases = []testGetSliceCase{
	{
		0,
		[]interface{}{"e"},
		nil,
		true,
	},
	{
		0,
		[]interface{}{"h"},
		[]interface{}{
			map[string]interface{}{
				"q": "q1q1q1",
				"w": "w1w1w1",
				"e": "e1e1e1",
			},
			map[string]interface{}{
				"q": "q2q2q2",
				"w": "w2w2w2",
				"e": "e2e2e2",
			},
			map[string]interface{}{
				"q": "q3q3q3",
				"w": "w3w3w3",
				"e": "e3e3e3",
			},
		},
		false,
	},
	{
		0,
		[]interface{}{"i", 0},
		[]interface{}{float64(1), float64(2), float64(3)},
		false,
	},
}

func TestGetSlice(t *testing.T) {
	jsonDatas := make([]easyjson.JSData, 0)

	for _, json := range testJson {
		jsonData, err := easyjson.DecodeJson([]byte(json))

		if err != nil {
			t.Fatalf("Could not decode test json.\n")
		}

		jsonDatas = append(jsonDatas, jsonData)
	}

	for _, test := range testGetSliceCases {
		value, err := easyjson.GetSlice(jsonDatas[test.jsonIndex], test.keys...)

		if err != nil {
			if !test.err {
				t.Errorf("TestGetSlice Failed: Unexpected error.\n%v\n", test)
				t.Errorf("%s", err)
			}
		}

		if !reflect.DeepEqual(value, test.value) {
			t.Errorf("TestGetSlice Failed: %v != %v\n%v", test.value, value, test)
		}
	}
}

type testGetMapCase struct {
	jsonIndex int
	keys      []interface{}
	value     map[string]interface{}
	err       bool
}

var testGetMapCases = []testGetMapCase{
	{
		0,
		[]interface{}{"a"},
		nil,
		true,
	},
	{
		0,
		[]interface{}{"d", "dd"},
		map[string]interface{}{
			"dda": "DDADDADDA",
			"ddb": "DDBDDBDDB",
			"ddc": "DDCDDCDDC",
		},
		false,
	},
	{
		1,
		[]interface{}{0, "d"},
		map[string]interface{}{
			"da": "da1",
			"db": "db1",
			"dc": "dc1",
		},
		false,
	},
}

func TestGetMap(t *testing.T) {
	jsonDatas := make([]easyjson.JSData, 0)

	for _, json := range testJson {
		jsonData, err := easyjson.DecodeJson([]byte(json))

		if err != nil {
			t.Fatalf("Could not decode test json.\n")
		}

		jsonDatas = append(jsonDatas, jsonData)
	}

	for _, test := range testGetMapCases {
		value, err := easyjson.GetMap(jsonDatas[test.jsonIndex], test.keys...)

		if err != nil {
			if !test.err {
				t.Errorf("TestGetMap Failed: Unexpected error.\n%v\n", test)
				t.Errorf("%s", err)
			}
		}

		if !reflect.DeepEqual(value, test.value) {
			t.Errorf("TestGetMap Failed: %v != %v\n%v", test.value, value, test)
		}
	}
}
