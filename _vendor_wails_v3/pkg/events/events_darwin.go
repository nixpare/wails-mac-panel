//go:build darwin

package events

/*
#cgo CFLAGS:  -x objective-c
#cgo LDFLAGS: -framework Cocoa -mmacosx-version-min=10.13

#include "events_darwin.h"
#include <stdlib.h>
#include <stdbool.h>

bool hasListener[MAX_EVENTS] = {false};

void registerListener(unsigned int event) {
	hasListener[event] = true;
}

bool hasListeners(unsigned int event) {
	//return hasListener[event];
	return true;
}

*/
import "C"
