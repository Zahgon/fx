package fx

import (
	"context"
	"errors"
	"reflect"

	"go.uber.org/dig"
)

type Annotated struct {
	Name string

	Group string

	Target any
}

func (a Annotated) String() string { _ = "STUB: not implemented"; return "" }

var (
	_inAnnotationField = reflect.StructField{
		Name:      "In",
		Type:      reflect.TypeOf(In{}),
		Anonymous: true,
	}

	_outAnnotationField = reflect.StructField{
		Name:      "Out",
		Type:      reflect.TypeOf(Out{}),
		Anonymous: true,
	}
)

type Annotation interface {
	apply(*annotated) error
	build(*annotated) (any, error)
}

var (
	_typeOfError = reflect.TypeOf((*error)(nil)).Elem()
	_nilError    = reflect.Zero(_typeOfError)
)

type annotationError struct {
	target any
	err    error
}

func (e *annotationError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *annotationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type paramTagsAnnotation struct {
	tags []string
}

var _ Annotation = paramTagsAnnotation{}
var (
	errTagSyntaxSpace            = errors.New(`multiple tags are not separated by space`)
	errTagKeySyntax              = errors.New("tag key is invalid, Use group, name or optional as tag keys")
	errTagValueSyntaxQuote       = errors.New(`tag value should start with double quote. i.e. key:"value" `)
	errTagValueSyntaxEndingQuote = errors.New(`tag value should end in double quote. i.e. key:"value" `)
)

func verifyTagsSpaceSeparated(tagIdx int, tag string) error { _ = "STUB: not implemented"; return nil }

func verifyValueQuote(value string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func verifyAnnotateTag(tag string) error { _ = "STUB: not implemented"; return nil }

func (pt paramTagsAnnotation) apply(ann *annotated) error { _ = "STUB: not implemented"; return nil }

func (pt paramTagsAnnotation) build(ann *annotated) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (pt paramTagsAnnotation) parameters(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParamTags(tags ...string) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

type resultTagsAnnotation struct {
	tags []string
}

var _ Annotation = resultTagsAnnotation{}

func (rt resultTagsAnnotation) apply(ann *annotated) error { _ = "STUB: not implemented"; return nil }

func (rt resultTagsAnnotation) build(ann *annotated) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (rt resultTagsAnnotation) results(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ResultTags(tags ...string) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

type outStructInfo struct {
	Fields  []reflect.StructField
	Offsets []int
}

type _lifecycleHookAnnotationType int

const (
	_unknownHookType _lifecycleHookAnnotationType = iota
	_onStartHookType
	_onStopHookType
)

type lifecycleHookAnnotation struct {
	Type   _lifecycleHookAnnotationType
	Target any
}

var _ Annotation = (*lifecycleHookAnnotation)(nil)

func (la *lifecycleHookAnnotation) String() string { _ = "STUB: not implemented"; return "" }

func (la *lifecycleHookAnnotation) apply(ann *annotated) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *lifecycleHookAnnotation) build(ann *annotated) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var (
	_typeOfLifecycle = reflect.TypeOf((*Lifecycle)(nil)).Elem()
	_typeOfContext   = reflect.TypeOf((*context.Context)(nil)).Elem()
)

func (la *lifecycleHookAnnotation) validateHookDeps(hookParamTypes []reflect.Type, paramTypes []reflect.Type, resultTypes []reflect.Type) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (la *lifecycleHookAnnotation) buildHookInstaller(ann *annotated) (
	hookInstaller reflect.Value,
	paramTypes []reflect.Type,
	remapParams func([]reflect.Value) []reflect.Value,
	err error,
) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil, nil, nil
}

var (
	_nameTag  = "name"
	_groupTag = "group"
)

func makeHookScopeCtor(paramTypes []reflect.Type, resultTypes []reflect.Type, args []reflect.Value) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func injectLifecycle(paramTypes []reflect.Type) ([]reflect.Type, func([]reflect.Value) []reflect.Value) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lifecycleExists(paramTypes []reflect.Type) bool { _ = "STUB: not implemented"; return false }

func (la *lifecycleHookAnnotation) buildHook(fn func(context.Context) error) (hook Hook) {
	_ = "STUB: not implemented"
	return *new(Hook)
}

func OnStart(onStart any) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

func OnStop(onStop any) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

type asAnnotation struct {
	targets []any
	types   []asType
}

type asType struct {
	self bool
	typ  reflect.Type
}

func (a asType) String() string { _ = "STUB: not implemented"; return "" }

func isOut(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isIn(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

var _ Annotation = (*asAnnotation)(nil)

func As(interfaces ...any) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

func Self() any { _ = "STUB: not implemented"; return *new(any) }

type self struct{}

func (at *asAnnotation) apply(ann *annotated) error { _ = "STUB: not implemented"; return nil }

func (at *asAnnotation) build(ann *annotated) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (at *asAnnotation) results(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func extractResultFields(types []reflect.Type) ([]reflect.StructField, func(int, []reflect.Value) reflect.Value) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fromAnnotation struct {
	targets []any
	types   []reflect.Type
}

var _ Annotation = (*fromAnnotation)(nil)

func From(interfaces ...any) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

func (fr *fromAnnotation) apply(ann *annotated) error { _ = "STUB: not implemented"; return nil }

func (fr *fromAnnotation) build(ann *annotated) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (fr *fromAnnotation) parameters(ann *annotated) (
	types []reflect.Type,
	remap func([]reflect.Value) []reflect.Value,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type annotated struct {
	Target      any
	Annotations []Annotation
	ParamTags   []string
	ResultTags  []string
	As          [][]asType
	From        []reflect.Type
	FuncPtr     uintptr
	Hooks       []*lifecycleHookAnnotation

	container *dig.Container
}

func (ann annotated) String() string { _ = "STUB: not implemented"; return "" }

func (ann *annotated) Build() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (ann *annotated) applyOptionalTag() { _ = "STUB: not implemented"; return }

func (ann *annotated) cleanUpAsResults() { _ = "STUB: not implemented"; return }

func (ann *annotated) typeCheckOrigFn() error { _ = "STUB: not implemented"; return nil }

func (ann *annotated) currentResultTypes() (resultTypes []reflect.Type, hasError bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (ann *annotated) currentParamTypes() []reflect.Type { _ = "STUB: not implemented"; return nil }

func Annotate(t any, anns ...Annotation) any { _ = "STUB: not implemented"; return *new(any) }
