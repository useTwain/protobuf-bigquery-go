package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/example/v1"
	expr "google.golang.org/genproto/googleapis/api/expr/v1beta1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestUnmarshalOptions_Load(t *testing.T) {
	for _, tt := range []struct {
		name          string
		row           []bigquery.Value
		schema        bigquery.Schema
		opt           UnmarshalOptions
		expected      proto.Message
		errorContains string
	}{
		{
			name: "library.Book",
			row: []bigquery.Value{
				"name",
				"author",
				"title",
				true,
			},
			schema: bigquery.Schema{
				{Name: "name", Type: bigquery.StringFieldType},
				{Name: "author", Type: bigquery.StringFieldType},
				{Name: "title", Type: bigquery.StringFieldType},
				{Name: "read", Type: bigquery.BooleanFieldType},
			},
			expected: &library.Book{
				Name:   "name",
				Author: "author",
				Title:  "title",
				Read:   true,
			},
		},

		{
			name: "library.UpdateBookRequest",
			row: []bigquery.Value{
				"name",
				[]bigquery.Value{
					"name",
					"author",
					"title",
					true,
				},
			},
			schema: bigquery.Schema{
				{Name: "name", Type: bigquery.StringFieldType},
				{
					Name: "book",
					Type: bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "name", Type: bigquery.StringFieldType},
						{Name: "author", Type: bigquery.StringFieldType},
						{Name: "title", Type: bigquery.StringFieldType},
						{Name: "read", Type: bigquery.BooleanFieldType},
					},
				},
			},
			expected: &library.UpdateBookRequest{
				Name: "name",
				Book: &library.Book{
					Name:   "name",
					Author: "author",
					Title:  "title",
					Read:   true,
				},
			},
		},

		{
			name: "expr.Value (bool)",
			row: []bigquery.Value{
				true,
				nil,
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
			},
			expected: &expr.Value{
				Kind: &expr.Value_BoolValue{
					BoolValue: true,
				},
			},
		},

		{
			name: "expr.Value (double)",
			row: []bigquery.Value{
				nil,
				float64(42),
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
			},
			expected: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
		},

		{
			name: "error on unknown fields",
			row: []bigquery.Value{
				nil,
				float64(42),
				"bar",
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "foo", Type: bigquery.StringFieldType},
			},
			expected:      &expr.Value{},
			errorContains: "unknown field: foo",
		},

		{
			name: "discard unknown fields",
			row: []bigquery.Value{
				nil,
				float64(42),
				"bar",
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "foo", Type: bigquery.StringFieldType},
			},
			opt: UnmarshalOptions{DiscardUnknown: true},
			expected: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
		},

		{
			name: "enum values",
			row: []bigquery.Value{
				"ENUM_VALUE1",
			},
			schema: bigquery.Schema{
				{Name: "enum_value", Type: bigquery.StringFieldType},
			},
			expected: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
		},

		{
			name: "enum numbers",
			row: []bigquery.Value{
				int64(1),
			},
			schema: bigquery.Schema{
				{Name: "enum_value", Type: bigquery.IntegerFieldType},
			},
			opt: UnmarshalOptions{
				Schema: SchemaOptions{
					UseEnumNumbers: true,
				},
			},
			expected: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
		},

		{
			name: "wrappers",
			row: []bigquery.Value{
				float64(1),
				float64(2),
				"foo",
				[]byte("bar"),
				int64(3),
				int64(4),
				uint64(5),
				uint64(6),
				true,
			},
			schema: bigquery.Schema{
				{Name: "float_value", Type: bigquery.FloatFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "string_value", Type: bigquery.StringFieldType},
				{Name: "bytes_value", Type: bigquery.BytesFieldType},
				{Name: "int32_value", Type: bigquery.IntegerFieldType},
				{Name: "int64_value", Type: bigquery.IntegerFieldType},
				{Name: "uint32_value", Type: bigquery.IntegerFieldType},
				{Name: "uint64_value", Type: bigquery.IntegerFieldType},
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
			},
			expected: &examplev1.ExampleWrappers{
				FloatValue:  wrapperspb.Float(1),
				DoubleValue: wrapperspb.Double(2),
				StringValue: wrapperspb.String("foo"),
				BytesValue:  wrapperspb.Bytes([]byte("bar")),
				Int32Value:  wrapperspb.Int32(3),
				Int64Value:  wrapperspb.Int64(4),
				Uint32Value: wrapperspb.UInt32(5),
				Uint64Value: wrapperspb.UInt64(6),
				BoolValue:   wrapperspb.Bool(true),
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := proto.Clone(tt.expected)
			proto.Reset(actual)
			if err := tt.opt.Load(tt.row, tt.schema, actual); tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tt.expected, actual, protocmp.Transform())
			}
		})
	}
}
