
import { On, Off, OffAll, OnMultiple, WailsEvent, dispatchWailsEvent, eventListeners, Once } from './events';

import { expect, describe, it, vi, afterEach, beforeEach } from 'vitest';

afterEach(() => {
  OffAll();
  vi.resetAllMocks();
});

describe('OnMultiple', () => {
  let testEvent = new WailsEvent('a', {});

  it('should stop after a specified number of times', () => {
    const cb = vi.fn();
    OnMultiple('a', cb, 5);
    dispatchWailsEvent(testEvent);
    dispatchWailsEvent(testEvent);
    dispatchWailsEvent(testEvent);
    dispatchWailsEvent(testEvent);
    dispatchWailsEvent(testEvent);
    dispatchWailsEvent(testEvent);
    expect(cb).toBeCalledTimes(5);
  });

  it('should return a cancel fn', () => {
    const cb = vi.fn()
    const cancel = OnMultiple('a', cb, 5)
    dispatchWailsEvent(testEvent)
    dispatchWailsEvent(testEvent)
    cancel()
    dispatchWailsEvent(testEvent)
    dispatchWailsEvent(testEvent)
    expect(cb).toBeCalledTimes(2)
  })
})

describe('On', () => {
  it('should create a listener with a count of -1', () => {
    On('a', () => {})
    expect(eventListeners.get("a")[0].maxCallbacks).toBe(-1)
  })

  it('should return a cancel fn', () => {
    const cancel = On('a', () => {})
    cancel();
  })
})

describe('Once', () => {
  it('should create a listener with a count of 1', () => {
    Once('a', () => {})
    expect(eventListeners.get("a")[0].maxCallbacks).toBe(1)
  })

  it('should return a cancel fn', () => {
    const cancel = EventsOn('a', () => {})
    cancel();
  })
})

describe('Off', () => {
  beforeEach(() => {
    On('a', () => {})
    On('a', () => {})
    On('a', () => {})
    On('b', () => {})
    On('c', () => {})
  })

  it('should cancel all event listeners for a single type', () => {
    Off('a')
    expect(eventListeners.get('a')).toBeUndefined()
    expect(eventListeners.get('b')).not.toBeUndefined()
    expect(eventListeners.get('c')).not.toBeUndefined()
  })

  it('should cancel all event listeners for multiple types', () => {
    Off('a', 'b')
    expect(eventListeners.get('a')).toBeUndefined()
    expect(eventListeners.get('b')).toBeUndefined()
    expect(eventListeners.get('c')).not.toBeUndefined()
  })
})

describe('OffAll', () => {
  it('should cancel all event listeners', () => {
    On('a', () => {})
    On('a', () => {})
    On('a', () => {})
    On('b', () => {})
    On('c', () => {})
    OffAll()
    expect(eventListeners.size).toBe(0)
  })
})
