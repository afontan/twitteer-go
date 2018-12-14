package features

import (
	"errors"
	"github.com/DATA-DOG/godog"
	"github.com/twitteer-go/src/domain"
	"github.com/twitteer-go/src/service"
	"os"
	"testing"
)

var twitterManager *service.TweetManager

func thereAreTwoTweets() error {
	tweetOne := domain.NewTextTweet("test","test")
	tweetTwo := domain.NewTextTweet("test","test")

	twitterManager.PublishTweet(tweetOne)
	twitterManager.PublishTweet(tweetTwo)

	return godog.ErrPending
}

func tryToPublishATweet(cant int) error {
	tweetOne := domain.NewTextTweet("test","grillo62222221312312312312312312312233232323223232232323232323223232323232323232323123123123123123123123131231123123113123231231231212312aaaa")
	twitterManager.PublishTweet(tweetOne)
	return godog.ErrPending
}

func thereShouldBePublishedTweet(cant int) error{
	if len(twitterManager.GetTweets()) == 2 {
		return nil
	}else{
		return errors.New("Not valid")
	}
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are two tweets$`, thereAreTwoTweets)
	s.Step(`^try to publish a (\d+)\+ tweet$`, tryToPublishATweet)
	s.Step(`^there should be (\d+) published tweet$`, thereShouldBePublishedTweet)
	s.BeforeScenario(func(interface{}) {
		twitterManager = service.InitializeService()
	})
}

func TestMain(m *testing.M) {
	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		godog.SuiteContext(s)
	}, godog.Options{
		Format: format,
		Paths:     []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}