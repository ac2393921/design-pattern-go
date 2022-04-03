package observer

type Subject interface {
	RegisterObserver(*Observer)
	RemoveObserver(*Observer)
	NotifyObservers(*Observer)
}
