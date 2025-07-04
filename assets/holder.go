package assets

import _ "embed"

//go:embed List.md
var ListMD string

//go:embed Divider.md
var DividerMD string

//go:embed Grid.md
var GridMD string

//go:embed Icon.md
var IconMD string

//go:embed Layout.md
var LayoutMD string

//go:embed Space.md
var SpaceMD string

//go:embed Typography.md
var TypographyMD string

//go:embed AutoComplete.md
var AutoCompleteMD string

//go:embed Cascader.md
var CascaderMD string

//go:embed Button.md
var ButtonMD string

//go:embed CheckBox.md
var CheckBoxMD string

//go:embed DatePicker.md
var DatePickerMD string

var Components = []ComponentName{
	{Name: "List", Content: ListMD},
	{Name: "Divider", Content: DividerMD},
	{Name: "Grid", Content: GridMD},
	{Name: "Icon", Content: IconMD},
	{Name: "Layout", Content: LayoutMD},
	{Name: "Space", Content: SpaceMD},
	{Name: "Typography", Content: TypographyMD},
	{Name: "AutoComplete", Content: AutoCompleteMD},
	{Name: "Cascader", Content: CascaderMD},
	{Name: "Button", Content: ButtonMD},
	{Name: "CheckBox", Content: CheckBoxMD},
	{Name: "DatePicker", Content: DatePickerMD},
}
