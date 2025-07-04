package assets

import _ "embed"

//go:embed token.txt
var Token string

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

//go:embed Checkbox.md
var CheckBoxMD string

//go:embed DatePicker.md
var DatePickerMD string

//go:embed Form.md
var FormMD string

//go:embed Input.md
var InputMD string

//go:embed InputNumber.md
var InputNumberMD string

//go:embed PinCode.md
var PinCodeMD string

//go:embed Radio.md
var RadioMD string

//go:embed Rating.md
var RatingMD string

//go:embed Select.md
var SelectMD string

//go:embed Slider.md
var SliderMD string

//go:embed Switch.md
var SwitchMD string

//go:embed TagInput.md
var TagInputMD string

//go:embed TimePicker.md
var TimePickerMD string

//go:embed Transfer.md
var TransferMD string

//go:embed TreeSelect.md
var TreeSelectMD string

//go:embed Upload.md
var UploadMD string

//go:embed Anchor.md
var AnchorMD string

//go:embed BackTop.md
var BackTopMD string

//go:embed Breadcrumb.md
var BreadcrumbMD string

//go:embed Navigation.md
var NavigationMD string

//go:embed Pagination.md
var PaginationMD string

//go:embed Steps.md
var StepsMD string

//go:embed Tabs.md
var TabsMD string

//go:embed Tree.md
var TreeMD string

//go:embed Avatar.md
var AvatarMD string

//go:embed Badge.md
var BadgeMD string

//go:embed Card.md
var CardMD string

//go:embed Calendar.md
var CalendarMD string

//go:embed Carousel.md
var CarouselMD string

//go:embed Collapse.md
var CollapseMD string

//go:embed Collapsible.md
var CollapsibleMD string

//go:embed Descriptions.md
var DescriptionsMD string

//go:embed Dropdown.md
var DropdownMD string

//go:embed Empty.md
var EmptyMD string

//go:embed Highlight.md
var HighlightMD string

//go:embed Image.md
var ImageMD string

//go:embed List.md
var ListMD string

//go:embed Cropper.md
var CropperMD string

//go:embed Modal.md
var ModalMD string

//go:embed OverflowList.md
var OverflowListMD string

//go:embed Popover.md
var PopoverMD string

//go:embed ScrollList.md
var ScrollListMD string

//go:embed SideSheet.md
var SideSheetMD string

//go:embed Table.md
var TableMD string

//go:embed Timeline.md
var TimelineMD string

//go:embed Tag.md
var TagMD string

//go:embed Tooltip.md
var TooltipMD string

//go:embed UserGuide.md
var UserGuideMD string

//go:embed Banner.md
var BannerMD string

//go:embed Notification.md
var NotificationMD string

//go:embed Popconfirm.md
var PopconfirmMD string

//go:embed Progress.md
var ProgressMD string

//go:embed Skeleton.md
var SkeletonMD string

//go:embed Spin.md
var SpinMD string

//go:embed Toast.md
var ToastMD string

//go:embed ConfigProvider.md
var ConfigProviderMD string

//go:embed LocaleProvider.md
var LocaleProviderMD string

var Components = []ComponentName{
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
	{Name: "Form", Content: FormMD},
	{Name: "Input", Content: InputMD},
	{Name: "InputNumber", Content: InputNumberMD},
	{Name: "PinCode", Content: PinCodeMD},
	{Name: "Radio", Content: RadioMD},
	{Name: "Rating", Content: RatingMD},
	{Name: "Select", Content: SelectMD},
	{Name: "Slider", Content: SliderMD},
	{Name: "Switch", Content: SwitchMD},
	{Name: "TagInput", Content: TagInputMD},
	{Name: "TimePicker", Content: TimePickerMD},
	{Name: "Transfer", Content: TransferMD},
	{Name: "TreeSelect", Content: TreeSelectMD},
	{Name: "Upload", Content: UploadMD},
	{Name: "Anchor", Content: AnchorMD},
	{Name: "BackTop", Content: BackTopMD},
	{Name: "Breadcrumb", Content: BreadcrumbMD},
	{Name: "Navigation", Content: NavigationMD},
	{Name: "Pagination", Content: PaginationMD},
	{Name: "Steps", Content: StepsMD},
	{Name: "Tabs", Content: TabsMD},
	{Name: "Tree", Content: TreeMD},
	{Name: "Avatar", Content: AvatarMD},
	{Name: "Badge", Content: BadgeMD},
	{Name: "Card", Content: CardMD},
	{Name: "Calendar", Content: CalendarMD},
	{Name: "Carousel", Content: CarouselMD},
	{Name: "Collapse", Content: CollapseMD},
	{Name: "Collapsible", Content: CollapsibleMD},
	{Name: "Descriptions", Content: DescriptionsMD},
	{Name: "Dropdown", Content: DropdownMD},
	{Name: "Empty", Content: EmptyMD},
	{Name: "Highlight", Content: HighlightMD},
	{Name: "Image", Content: ImageMD},
	{Name: "List", Content: ListMD},
	{Name: "Cropper", Content: CropperMD},
	{Name: "Modal", Content: ModalMD},
	{Name: "OverflowList", Content: OverflowListMD},
	{Name: "Popover", Content: PopoverMD},
	{Name: "ScrollList", Content: ScrollListMD},
	{Name: "SideSheet", Content: SideSheetMD},
	{Name: "Table", Content: TableMD},
	{Name: "Timeline", Content: TimelineMD},
	{Name: "Tag", Content: TagMD},
	{Name: "Tooltip", Content: TooltipMD},
	{Name: "UserGuide", Content: UserGuideMD},
	{Name: "Banner", Content: BannerMD},
	{Name: "Notification", Content: NotificationMD},
	{Name: "Popconfirm", Content: PopconfirmMD},
	{Name: "Progress", Content: ProgressMD},
	{Name: "Skeleton", Content: SkeletonMD},
	{Name: "Spin", Content: SpinMD},
	{Name: "Toast", Content: ToastMD},
	{Name: "ConfigProvider", Content: ConfigProviderMD},
	{Name: "LocaleProvider", Content: LocaleProviderMD},
}
