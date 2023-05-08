Feature:Home

@home
@text1
Scenario: Search 1
Given I launch the home page
And I enter with text msn
And I hit the search button
Then Verify the page title

@home
@text2
Scenario: Search 2
Given I launch the home page
And I enter with text globo
And I hit the search button
Then Verify the page title