# Observer pattern
Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.
## Usage
```go
package main

import "github.com/chubaka358/intern/pkg/observer"

func main() {
	publisher := observer.NewNewsPublisher()
	sportSubscriber := observer.NewSubscriberSport()
	politicsSubscriber := observer.NewSubscriberPolitics()
	publisher.Attach(sportSubscriber)
	publisher.Attach(politicsSubscriber)

	publisher.SetTitleAndData("Sport", "sport stuff")
	publisher.Notify()
}
```
