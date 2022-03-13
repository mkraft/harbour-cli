module github.com/mkraft/harbourcli

replace github.com/mkraft/mattermost => ./mattermost

replace github.com/mkraft/mattermostjsonl => ./mattermostjsonl

replace github.com/mkraft/slackexport => ./slackexport

go 1.17

require (
	github.com/mkraft/mattermost v0.0.0-00010101000000-000000000000
	github.com/mkraft/mattermostjsonl v0.0.0-00010101000000-000000000000
	github.com/mkraft/slackexport v0.0.0-00010101000000-000000000000
)
