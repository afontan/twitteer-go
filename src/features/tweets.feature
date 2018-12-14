Feature: tweets
  Tweet a more 140 char tweet doesnt modify tweet count

  Scenario: add char exeeded tweet
    Given there are two tweets
    When try to publish a 140+ tweet
    Then there should be 2 published tweet