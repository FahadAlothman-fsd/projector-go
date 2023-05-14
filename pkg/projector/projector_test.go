package projector_test

import (
	"testing"

	"github.com/FahadAlothman-fsd/projector-go/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"baba":  "baz1",
				"femto": "is_supreme_soy",
			},
			"/baba": {
				"baba": "baz2",
			},
			"/baba/baz": {
				"baba": "baz3",
			},
		}}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(&projector.Config{
		Args:      []string{},
		Operation: projector.Print,
		Pwd:       pwd,
		Config:    "gomenasai",
	}, data)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	val, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}
	if val != value {
		t.Errorf("expected value \"%v\", got \"%v\"", value, val)
	}

}
func TestGetValue(t *testing.T) {

	proj := getProjector("/baba/baz", getData())
	test(t, proj, "baba", "baz3")
	test(t, proj, "femto", "is_supreme_soy")

}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/baba/baz", data)
	test(t, proj, "baba", "baz3")
	test(t, proj, "femto", "is_supreme_soy")
	proj.SetValue("baba", "baz4")
	test(t, proj, "baba", "baz4")
	proj.SetValue("femto", "is_not_soy")
	test(t, proj, "femto", "is_not_soy")
	proj = getProjector("/baba", data)
	test(t, proj, "femto", "is_supreme_soy")
	proj = getProjector("/", data)
	test(t, proj, "femto", "is_supreme_soy")

}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/baba/baz", data)
	test(t, proj, "baba", "baz3")
	proj.RemoveValue("baba")
	test(t, proj, "baba", "baz2")
	proj.RemoveValue("femto")
	test(t, proj, "femto", "is_supreme_soy")

}
