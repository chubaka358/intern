# Observer pattern
Observer pattern example
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
