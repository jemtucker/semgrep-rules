package test

//ruleid: naked-return
func foo() (err error) {
	return
}

//ruleid: naked-return
func bar(s int) (err error) {
	return
}

//ok: naked-return
func bash(s int) (err error) {
	return err
}

//ruleid: naked-return
func bazh(s int) (err error) {
	if s == 20 {
		return
	}

	return err
}

//ok: naked-return
func none() {
	return
}

//ruleid: naked-return
func baz() (thing string, err error) {
	return
}

type t struct{}

//ruleid: naked-return
func (t) faz() (err error) {
	return
}
