package observer

import "testing"

func TestNewsPublisher(t *testing.T) {
	publisher := NewNewsPublisher()
	sportSubscriber := NewSubscriberSport()
	politicsSubscriber := NewSubscriberPolitics()
	publisher.Attach(sportSubscriber)
	publisher.Attach(politicsSubscriber)

	publisher.SetTitleAndData("Sport", "sport stuff")
	publisher.Notify()
	publisher.SetTitleAndData("Politics", "politics stuff")
	publisher.Notify()

	if data := sportSubscriber.GetData(); data != "sport stuff" {
		t.Errorf("expected %q, got %q", "sport stuff", data)
	}

	if data := politicsSubscriber.GetData(); data != "politics stuff" {
		t.Errorf("expected %q, got %q", "sport stuff", data)
	}

	publisher.Unsubscribe(politicsSubscriber)

	publisher.SetTitleAndData("Politics", "The subscriber should not receive this")
	if politicsSubscriber.GetData() == "The subscriber should not receive this" {
		t.Errorf("politicsSubscriber shouldn't recieve the data")
	}
}
