//go:build plan9
// +build plan9

package display

// WatchResize is a no-op on Plan 9; terminal resize detection is not
// currently supported.
func WatchResize(eng *Engine) chan<- bool {
	done := make(chan bool, 1)
	return done
}
