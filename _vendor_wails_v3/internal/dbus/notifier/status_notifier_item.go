// Code generated by dbus-codegen-go DO NOT EDIT.
package notifier

import (
	"context"
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

var (
	// Introspection for org.kde.StatusNotifierItem
	IntrospectDataStatusNotifierItem = introspect.Interface{
		Name: "org.kde.StatusNotifierItem",
		Methods: []introspect.Method{{Name: "ContextMenu", Args: []introspect.Arg{
			{Name: "x", Type: "i", Direction: "in"},
			{Name: "y", Type: "i", Direction: "in"},
		}},
			{Name: "Activate", Args: []introspect.Arg{
				{Name: "x", Type: "i", Direction: "in"},
				{Name: "y", Type: "i", Direction: "in"},
			}},
			{Name: "SecondaryActivate", Args: []introspect.Arg{
				{Name: "x", Type: "i", Direction: "in"},
				{Name: "y", Type: "i", Direction: "in"},
			}},
			{Name: "Scroll", Args: []introspect.Arg{
				{Name: "delta", Type: "i", Direction: "in"},
				{Name: "orientation", Type: "s", Direction: "in"},
			}},
		},
		Signals: []introspect.Signal{{Name: "NewTitle"},
			{Name: "NewIcon"},
			{Name: "NewAttentionIcon"},
			{Name: "NewOverlayIcon"},
			{Name: "NewStatus", Args: []introspect.Arg{
				{Name: "status", Type: "s", Direction: ""},
			}},
			{Name: "NewIconThemePath", Args: []introspect.Arg{
				{Name: "icon_theme_path", Type: "s", Direction: "out"},
			}},
			{Name: "NewMenu"},
		},
		Properties: []introspect.Property{{Name: "Category", Type: "s", Access: "read"},
			{Name: "Id", Type: "s", Access: "read"},
			{Name: "Title", Type: "s", Access: "read"},
			{Name: "Status", Type: "s", Access: "read"},
			{Name: "WindowId", Type: "i", Access: "read"},
			{Name: "IconThemePath", Type: "s", Access: "read"},
			{Name: "Menu", Type: "o", Access: "read"},
			{Name: "ItemIsMenu", Type: "b", Access: "read"},
			{Name: "IconName", Type: "s", Access: "read"},
			{Name: "IconPixmap", Type: "a(iiay)", Access: "read", Annotations: []introspect.Annotation{
				{Name: "org.qtproject.QtDBus.QtTypeName", Value: "KDbusImageVector"},
			}},
			{Name: "OverlayIconName", Type: "s", Access: "read"},
			{Name: "OverlayIconPixmap", Type: "a(iiay)", Access: "read", Annotations: []introspect.Annotation{
				{Name: "org.qtproject.QtDBus.QtTypeName", Value: "KDbusImageVector"},
			}},
			{Name: "AttentionIconName", Type: "s", Access: "read"},
			{Name: "AttentionIconPixmap", Type: "a(iiay)", Access: "read", Annotations: []introspect.Annotation{
				{Name: "org.qtproject.QtDBus.QtTypeName", Value: "KDbusImageVector"},
			}},
			{Name: "AttentionMovieName", Type: "s", Access: "read"},
			{Name: "ToolTip", Type: "(sa(iiay)ss)", Access: "read", Annotations: []introspect.Annotation{
				{Name: "org.qtproject.QtDBus.QtTypeName", Value: "KDbusToolTipStruct"},
			}},
		},
		Annotations: []introspect.Annotation{},
	}
)

// Signal is a common interface for all signals.
type Signal interface {
	Name() string
	Interface() string
	Sender() string

	path() dbus.ObjectPath
	values() []interface{}
}

// Emit sends the given signal to the bus.
func Emit(conn *dbus.Conn, s Signal) error {
	return conn.Emit(s.path(), s.Interface()+"."+s.Name(), s.values()...)
}

// ErrUnknownSignal is returned by LookupSignal when a signal cannot be resolved.
var ErrUnknownSignal = errors.New("unknown signal")

// LookupSignal converts the given raw D-Bus signal with variable body
// into one with typed structured body or returns ErrUnknownSignal error.
func LookupSignal(signal *dbus.Signal) (Signal, error) {
	switch signal.Name {
	case InterfaceStatusNotifierItem + "." + "NewTitle":
		return &StatusNotifierItem_NewTitleSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body:   &StatusNotifierItem_NewTitleSignalBody{},
		}, nil
	case InterfaceStatusNotifierItem + "." + "NewIcon":
		return &StatusNotifierItem_NewIconSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body:   &StatusNotifierItem_NewIconSignalBody{},
		}, nil
	case InterfaceStatusNotifierItem + "." + "NewAttentionIcon":
		return &StatusNotifierItem_NewAttentionIconSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body:   &StatusNotifierItem_NewAttentionIconSignalBody{},
		}, nil
	case InterfaceStatusNotifierItem + "." + "NewOverlayIcon":
		return &StatusNotifierItem_NewOverlayIconSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body:   &StatusNotifierItem_NewOverlayIconSignalBody{},
		}, nil
	case InterfaceStatusNotifierItem + "." + "NewStatus":
		v0, ok := signal.Body[0].(string)
		if !ok {
			return nil, fmt.Errorf("prop .Status is %T, not string", signal.Body[0])
		}
		return &StatusNotifierItem_NewStatusSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body: &StatusNotifierItem_NewStatusSignalBody{
				Status: v0,
			},
		}, nil
	case InterfaceStatusNotifierItem + "." + "NewIconThemePath":
		v0, ok := signal.Body[0].(string)
		if !ok {
			return nil, fmt.Errorf("prop .IconThemePath is %T, not string", signal.Body[0])
		}
		return &StatusNotifierItem_NewIconThemePathSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body: &StatusNotifierItem_NewIconThemePathSignalBody{
				IconThemePath: v0,
			},
		}, nil
	case InterfaceStatusNotifierItem + "." + "NewMenu":
		return &StatusNotifierItem_NewMenuSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body:   &StatusNotifierItem_NewMenuSignalBody{},
		}, nil
	default:
		return nil, ErrUnknownSignal
	}
}

// AddMatchSignal registers a match rule for the given signal,
// opts are appended to the automatically generated signal's rules.
func AddMatchSignal(conn *dbus.Conn, s Signal, opts ...dbus.MatchOption) error {
	return conn.AddMatchSignal(append([]dbus.MatchOption{
		dbus.WithMatchInterface(s.Interface()),
		dbus.WithMatchMember(s.Name()),
	}, opts...)...)
}

// RemoveMatchSignal unregisters the previously registered subscription.
func RemoveMatchSignal(conn *dbus.Conn, s Signal, opts ...dbus.MatchOption) error {
	return conn.RemoveMatchSignal(append([]dbus.MatchOption{
		dbus.WithMatchInterface(s.Interface()),
		dbus.WithMatchMember(s.Name()),
	}, opts...)...)
}

// Interface name constants.
const (
	InterfaceStatusNotifierItem = "org.kde.StatusNotifierItem"
)

// StatusNotifierItemer is org.kde.StatusNotifierItem interface.
type StatusNotifierItemer interface {
	// ContextMenu is org.kde.StatusNotifierItem.ContextMenu method.
	ContextMenu(x int32, y int32) (err *dbus.Error)
	// Activate is org.kde.StatusNotifierItem.Activate method.
	Activate(x int32, y int32) (err *dbus.Error)
	// SecondaryActivate is org.kde.StatusNotifierItem.SecondaryActivate method.
	SecondaryActivate(x int32, y int32) (err *dbus.Error)
	// Scroll is org.kde.StatusNotifierItem.Scroll method.
	Scroll(delta int32, orientation string) (err *dbus.Error)
}

// ExportStatusNotifierItem exports the given object that implements org.kde.StatusNotifierItem on the bus.
func ExportStatusNotifierItem(conn *dbus.Conn, path dbus.ObjectPath, v StatusNotifierItemer) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"ContextMenu":       v.ContextMenu,
		"Activate":          v.Activate,
		"SecondaryActivate": v.SecondaryActivate,
		"Scroll":            v.Scroll,
	}, path, InterfaceStatusNotifierItem)
}

// UnexportStatusNotifierItem unexports org.kde.StatusNotifierItem interface on the named path.
func UnexportStatusNotifierItem(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceStatusNotifierItem)
}

// UnimplementedStatusNotifierItem can be embedded to have forward compatible server implementations.
type UnimplementedStatusNotifierItem struct{}

func (*UnimplementedStatusNotifierItem) iface() string {
	return InterfaceStatusNotifierItem
}

func (*UnimplementedStatusNotifierItem) ContextMenu(x int32, y int32) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedStatusNotifierItem) Activate(x int32, y int32) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedStatusNotifierItem) SecondaryActivate(x int32, y int32) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedStatusNotifierItem) Scroll(delta int32, orientation string) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

// NewStatusNotifierItem creates and allocates org.kde.StatusNotifierItem.
func NewStatusNotifierItem(object dbus.BusObject) *StatusNotifierItem {
	return &StatusNotifierItem{object}
}

// StatusNotifierItem implements org.kde.StatusNotifierItem D-Bus interface.
type StatusNotifierItem struct {
	object dbus.BusObject
}

// ContextMenu calls org.kde.StatusNotifierItem.ContextMenu method.
func (o *StatusNotifierItem) ContextMenu(ctx context.Context, x int32, y int32) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceStatusNotifierItem+".ContextMenu", 0, x, y).Store()
	return
}

// Activate calls org.kde.StatusNotifierItem.Activate method.
func (o *StatusNotifierItem) Activate(ctx context.Context, x int32, y int32) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceStatusNotifierItem+".Activate", 0, x, y).Store()
	return
}

// SecondaryActivate calls org.kde.StatusNotifierItem.SecondaryActivate method.
func (o *StatusNotifierItem) SecondaryActivate(ctx context.Context, x int32, y int32) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceStatusNotifierItem+".SecondaryActivate", 0, x, y).Store()
	return
}

// Scroll calls org.kde.StatusNotifierItem.Scroll method.
func (o *StatusNotifierItem) Scroll(ctx context.Context, delta int32, orientation string) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceStatusNotifierItem+".Scroll", 0, delta, orientation).Store()
	return
}

// GetCategory gets org.kde.StatusNotifierItem.Category property.
func (o *StatusNotifierItem) GetCategory(ctx context.Context) (category string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "Category").Store(&category)
	return
}

// GetId gets org.kde.StatusNotifierItem.Id property.
func (o *StatusNotifierItem) GetId(ctx context.Context) (id string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "Id").Store(&id)
	return
}

// GetTitle gets org.kde.StatusNotifierItem.Title property.
func (o *StatusNotifierItem) GetTitle(ctx context.Context) (title string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "Title").Store(&title)
	return
}

// GetStatus gets org.kde.StatusNotifierItem.Status property.
func (o *StatusNotifierItem) GetStatus(ctx context.Context) (status string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "Status").Store(&status)
	return
}

// GetWindowId gets org.kde.StatusNotifierItem.WindowId property.
func (o *StatusNotifierItem) GetWindowId(ctx context.Context) (windowId int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "WindowId").Store(&windowId)
	return
}

// GetIconThemePath gets org.kde.StatusNotifierItem.IconThemePath property.
func (o *StatusNotifierItem) GetIconThemePath(ctx context.Context) (iconThemePath string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "IconThemePath").Store(&iconThemePath)
	return
}

// GetMenu gets org.kde.StatusNotifierItem.Menu property.
func (o *StatusNotifierItem) GetMenu(ctx context.Context) (menu dbus.ObjectPath, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "Menu").Store(&menu)
	return
}

// GetItemIsMenu gets org.kde.StatusNotifierItem.ItemIsMenu property.
func (o *StatusNotifierItem) GetItemIsMenu(ctx context.Context) (itemIsMenu bool, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "ItemIsMenu").Store(&itemIsMenu)
	return
}

// GetIconName gets org.kde.StatusNotifierItem.IconName property.
func (o *StatusNotifierItem) GetIconName(ctx context.Context) (iconName string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "IconName").Store(&iconName)
	return
}

// GetIconPixmap gets org.kde.StatusNotifierItem.IconPixmap property.
//
// Annotations:
//
//	@org.qtproject.QtDBus.QtTypeName = KDbusImageVector
func (o *StatusNotifierItem) GetIconPixmap(ctx context.Context) (iconPixmap []struct {
	V0 int32
	V1 int32
	V2 []byte
}, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "IconPixmap").Store(&iconPixmap)
	return
}

// GetOverlayIconName gets org.kde.StatusNotifierItem.OverlayIconName property.
func (o *StatusNotifierItem) GetOverlayIconName(ctx context.Context) (overlayIconName string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "OverlayIconName").Store(&overlayIconName)
	return
}

// GetOverlayIconPixmap gets org.kde.StatusNotifierItem.OverlayIconPixmap property.
//
// Annotations:
//
//	@org.qtproject.QtDBus.QtTypeName = KDbusImageVector
func (o *StatusNotifierItem) GetOverlayIconPixmap(ctx context.Context) (overlayIconPixmap []struct {
	V0 int32
	V1 int32
	V2 []byte
}, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "OverlayIconPixmap").Store(&overlayIconPixmap)
	return
}

// GetAttentionIconName gets org.kde.StatusNotifierItem.AttentionIconName property.
func (o *StatusNotifierItem) GetAttentionIconName(ctx context.Context) (attentionIconName string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "AttentionIconName").Store(&attentionIconName)
	return
}

// GetAttentionIconPixmap gets org.kde.StatusNotifierItem.AttentionIconPixmap property.
//
// Annotations:
//
//	@org.qtproject.QtDBus.QtTypeName = KDbusImageVector
func (o *StatusNotifierItem) GetAttentionIconPixmap(ctx context.Context) (attentionIconPixmap []struct {
	V0 int32
	V1 int32
	V2 []byte
}, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "AttentionIconPixmap").Store(&attentionIconPixmap)
	return
}

// GetAttentionMovieName gets org.kde.StatusNotifierItem.AttentionMovieName property.
func (o *StatusNotifierItem) GetAttentionMovieName(ctx context.Context) (attentionMovieName string, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "AttentionMovieName").Store(&attentionMovieName)
	return
}

// GetToolTip gets org.kde.StatusNotifierItem.ToolTip property.
//
// Annotations:
//
//	@org.qtproject.QtDBus.QtTypeName = KDbusToolTipStruct
func (o *StatusNotifierItem) GetToolTip(ctx context.Context) (toolTip struct {
	V0 string
	V1 []struct {
		V0 int32
		V1 int32
		V2 []byte
	}
	V2 string
	V3 string
}, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceStatusNotifierItem, "ToolTip").Store(&toolTip)
	return
}

// StatusNotifierItem_NewTitleSignal represents org.kde.StatusNotifierItem.NewTitle signal.
type StatusNotifierItem_NewTitleSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewTitleSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewTitleSignal) Name() string {
	return "NewTitle"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewTitleSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewTitleSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewTitleSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewTitleSignal) values() []interface{} {
	return []interface{}{}
}

// StatusNotifierItem_NewTitleSignalBody is body container.
type StatusNotifierItem_NewTitleSignalBody struct {
}

// StatusNotifierItem_NewIconSignal represents org.kde.StatusNotifierItem.NewIcon signal.
type StatusNotifierItem_NewIconSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewIconSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewIconSignal) Name() string {
	return "NewIcon"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewIconSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewIconSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewIconSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewIconSignal) values() []interface{} {
	return []interface{}{}
}

// StatusNotifierItem_NewIconSignalBody is body container.
type StatusNotifierItem_NewIconSignalBody struct {
}

// StatusNotifierItem_NewAttentionIconSignal represents org.kde.StatusNotifierItem.NewAttentionIcon signal.
type StatusNotifierItem_NewAttentionIconSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewAttentionIconSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewAttentionIconSignal) Name() string {
	return "NewAttentionIcon"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewAttentionIconSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewAttentionIconSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewAttentionIconSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewAttentionIconSignal) values() []interface{} {
	return []interface{}{}
}

// StatusNotifierItem_NewAttentionIconSignalBody is body container.
type StatusNotifierItem_NewAttentionIconSignalBody struct {
}

// StatusNotifierItem_NewOverlayIconSignal represents org.kde.StatusNotifierItem.NewOverlayIcon signal.
type StatusNotifierItem_NewOverlayIconSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewOverlayIconSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewOverlayIconSignal) Name() string {
	return "NewOverlayIcon"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewOverlayIconSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewOverlayIconSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewOverlayIconSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewOverlayIconSignal) values() []interface{} {
	return []interface{}{}
}

// StatusNotifierItem_NewOverlayIconSignalBody is body container.
type StatusNotifierItem_NewOverlayIconSignalBody struct {
}

// StatusNotifierItem_NewStatusSignal represents org.kde.StatusNotifierItem.NewStatus signal.
type StatusNotifierItem_NewStatusSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewStatusSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewStatusSignal) Name() string {
	return "NewStatus"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewStatusSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewStatusSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewStatusSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewStatusSignal) values() []interface{} {
	return []interface{}{s.Body.Status}
}

// StatusNotifierItem_NewStatusSignalBody is body container.
type StatusNotifierItem_NewStatusSignalBody struct {
	Status string
}

// StatusNotifierItem_NewIconThemePathSignal represents org.kde.StatusNotifierItem.NewIconThemePath signal.
type StatusNotifierItem_NewIconThemePathSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewIconThemePathSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewIconThemePathSignal) Name() string {
	return "NewIconThemePath"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewIconThemePathSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewIconThemePathSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewIconThemePathSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewIconThemePathSignal) values() []interface{} {
	return []interface{}{s.Body.IconThemePath}
}

// StatusNotifierItem_NewIconThemePathSignalBody is body container.
type StatusNotifierItem_NewIconThemePathSignalBody struct {
	IconThemePath string
}

// StatusNotifierItem_NewMenuSignal represents org.kde.StatusNotifierItem.NewMenu signal.
type StatusNotifierItem_NewMenuSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *StatusNotifierItem_NewMenuSignalBody
}

// Name returns the signal's name.
func (s *StatusNotifierItem_NewMenuSignal) Name() string {
	return "NewMenu"
}

// Interface returns the signal's interface.
func (s *StatusNotifierItem_NewMenuSignal) Interface() string {
	return InterfaceStatusNotifierItem
}

// Sender returns the signal's sender unique name.
func (s *StatusNotifierItem_NewMenuSignal) Sender() string {
	return s.sender
}

func (s *StatusNotifierItem_NewMenuSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *StatusNotifierItem_NewMenuSignal) values() []interface{} {
	return []interface{}{}
}

// StatusNotifierItem_NewMenuSignalBody is body container.
type StatusNotifierItem_NewMenuSignalBody struct {
}
