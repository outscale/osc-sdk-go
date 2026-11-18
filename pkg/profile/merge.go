package profile

import (
	"errors"
	"fmt"
	"reflect"
)

var errType = reflect.TypeFor[error]()

func planFor(t reflect.Type) []groupPlan {
	groups := make(map[string]groupPlan)
	for field := range t.Fields() {
		fieldType := field.Type

		group := field.Tag.Get("mergegroup")
		if group == "" {
			group = field.Name
		}

		if groupInstance, ok := groups[group]; ok {
			groupInstance.types = append(groupInstance.types, fieldType)
			groupInstance.names = append(groupInstance.names, field.Name)
			groups[group] = groupInstance
		} else {
			groups[group] = groupPlan{
				groupName: group,
				types:     []reflect.Type{fieldType},
				names:     []string{field.Name},
			}
		}
	}

	plans := make([]groupPlan, 0, len(groups))
	for _, plan := range groups {
		plans = append(plans, plan)
	}
	return plans
}

type groupPlan struct {
	groupName string
	types     []reflect.Type
	names     []string
}

func callGetter(src FieldsGetter, name string, want reflect.Type) (reflect.Value, bool, error) {
	f := reflect.ValueOf(src).FieldByName(name)
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return reflect.Value{}, false, nil
	}

	ft := f.Type()
	if ft.NumIn() != 0 || ft.NumOut() != 2 || !ft.Out(0).AssignableTo(want) || !ft.Out(1).Implements(errType) {
		return reflect.Value{}, false, fmt.Errorf("getter is %s, went func() (%s, errror)", ft, want)
	}

	out := f.Call(nil)
	if e := out[1]; !e.IsNil() {
		return reflect.Value{}, false, e.Interface().(error)
	}

	return out[0], true, nil
}

func pick(out reflect.Value, sources *map[string]string, p groupPlan, srcs []FieldsGetter) ([]reflect.Value, error) {
	if len(p.types) != len(p.names) {
		return nil, errors.New("plan is boggus")
	}

	// Maps merge key-by-key and must form their own group.
	if len(p.names) == 1 && p.types[0].Kind() == reflect.Map {
		return pickMap(out.FieldByName(p.names[0]), sources, p, srcs)
	}

	missing := make([]int, 0, len(p.names))
	for i, name := range p.names {
		if !isSet(out.FieldByName(name)) {
			missing = append(missing, i)
		}
	}
	if len(missing) == 0 {
		return nil, nil // out already wins this group
	}

	for _, src := range srcs {
		outGroup := make([]reflect.Value, len(p.types))
		count := 0
		for _, i := range missing {
			val, ok, err := callGetter(src, p.names[i], p.types[i])
			if err != nil {
				return nil, err
			}
			if !ok || !isSet(val) {
				continue
			}

			outGroup[i] = val
			count += 1
		}

		if count == 0 {
			continue // source has nothing for this group
		}
		if count != len(missing) {
			return nil, fmt.Errorf("group fields %s are not fully satisfied by source", p.groupName)
		}

		if sources != nil {
			for _, i := range missing {
				(*sources)[p.names[i]] = src.SourceName
			}
		}

		return outGroup, nil
	}

	return nil, nil
}

func pickMap(dst reflect.Value, sources *map[string]string, p groupPlan, srcs []FieldsGetter) ([]reflect.Value, error) {
	for _, src := range srcs {
		val, ok, err := callGetter(src, p.names[0], p.types[0])
		if err != nil {
			return nil, err
		}
		if !ok || val.IsNil() {
			continue
		}

		iter := val.MapRange()
		for iter.Next() {
			key, v := iter.Key(), iter.Value()
			existing := dst.MapIndex(key)
			if existing.IsValid() && isSet(existing) {
				continue
			}
			if !isSet(v) {
				continue
			}
			if dst.IsNil() {
				dst.Set(reflect.MakeMap(p.types[0]))
			}
			dst.SetMapIndex(key, v)

			if sources != nil {
				(*sources)[p.names[0]+"."+key.String()] = src.SourceName
			}
		}
	}

	// Map merged in place; nothing left to write.
	return nil, nil
}

func isSet(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Slice:
		return !value.IsNil()
	case reflect.Map:
		return !value.IsNil() && value.Len() > 0

	default:
		return !value.IsZero()
	}
}

func Merge(out *Fields, sources *map[string]string, getters ...FieldsGetter) error {
	if out == nil {
		return errors.New("target must be a non-nil pointer")
	}

	outVal := reflect.ValueOf(out).Elem()
	outType := outVal.Type()

	for _, plan := range planFor(outType) {
		values, err := pick(outVal, sources, plan, getters)
		if err != nil {
			return err
		}
		if values == nil {
			continue
		}
		for i, name := range plan.names {
			if !values[i].IsValid() {
				continue
			}
			outVal.FieldByName(name).Set(values[i])
		}
	}
	return nil
}
