# GoJira

Gojira is an application that act as middleware to help Gitlab x JIRA Integration


#### Supported Webhook Operation
- Push Request
- Open Merge Request

## Installation and Usage
#### Build
1. Clone the repository using
	`git clone https://github.com/agungcandra/gojira.git`
2. Build GoJira Executable using
`go build -o gojira app/main.go`

#### Run GoJira
1. Setup your environment variable or create `.env` file in the GoJira executable directory that consist
```
USERNAME=jira-email@domain.com
PASSWORD=API-TOKEN
HOST=Your JIRA Host
```

2. Create `workflow.yaml` file, that define the workflow structure of your project, sample can be found in `workflow.sample.yaml`

3. Run GoJira using `./gojira` or `go run app/main.go`

  API Token can be created from using this guilde [How to create JIRA API Token](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/ "How to create JIRA API Token")

## Understanding the Workflow
Workflow is used to define, what kind of operation that we support on the GoJira instance, and also it define steps that need to be taken by the GoJira until reach certain goals
Basic structure for the Workflow itself
```yaml
PROJC:                      # Project Key
  InProgress:               # Target
    Target: In Progress
    Steps:                  # Available Steps
      Backlog:              # State
        - Ready For Sprint  # Steps that in to be taken
        - Start Work        # with current steps
      Ready For Sprint:
        - Start Work
```
### Explanation

1. **PROJECT KEY** -> `PROJC`

That is the project key of the JIRA board that we want to integrate with, we can also define multiple board on those workflow file

2. **Target** -> `InProgress`

Workflow target that represent webhook action, when gojira receive given webhook from gitlab, it will look whether we define an action for those action, currently we have 2 supported target
	- InProgress
		Used to define push webhook action workflow
	- CodeReview
		Used to define merge_request webhook action workflow

3. **Steps**

In the steps we define what kind of transition that we support, from example above we define that for InProgress action, we support transition from `Backlog` and `Ready For Sprint` state, and also we need to define what step that need to be taken from those state until target is achieved.
