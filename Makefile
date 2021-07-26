build:
	gofmt -w pkg/crawler/crawler.go
	gofmt -w main.go
	go build -o bin/main main.go

run:
	go run main.go https://www.teamblind.com/post/New-Year-Gift---Curated-List-of-Top-100-LeetCode-Questions-to-Save-Your-Time-OaM1orEU https://www.glassdoor.com/Job/jobs.htm?sc.keyword=Software%20Engineer&locT=C&locId=1147311 
