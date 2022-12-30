from typing import Callable, Any

_subscribers: dict[str, list[Callable]] = dict()

def subscribe(event_type: str, function: Callable) -> None:
    """Subscribes to an event."""
    if not event_type in _subscribers:
        _subscribers[event_type] = []
    _subscribers[event_type].append(function)

def post_event(event_type: str, data: Any) -> None:
    """Posts an event with given data."""
    if not event_type in _subscribers:
        return
    for function in _subscribers[event_type]:
        function(data)

def list_events() -> None:
    for idx, event_name in enumerate(_subscribers, start=1):
        print(f"{idx}. {event_name}")