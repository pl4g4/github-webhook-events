# github-webhook-events

Get events from selected items in github webhook

##### Start container: 

`docker-compose up`

##### Testing curl: 

`curl -X POST http://localhost:8080/github-events -d "TestingThisNewStuff"`

##### Review your logs inside container: 

`docker exec -it githubWebhookEvents tail -f githubeventslog.log`