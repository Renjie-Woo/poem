package demo

type Event struct {
	field1   *string `my_json:"field_1"`
	Field2   *int    `my_json:"field_2"`
	jsonFile *string `json:"json_file"`
}

func (e *Event) SetValue(f1 string, f2 int) {
	e.field1 = &f1
	e.Field2 = &f2
}
