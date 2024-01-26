


How To Run:
    update config.json with appropriate configuration 
    go run cmd/main.go
    
How To Monitor output file real time:
./__scripts/monitor_json.sh

TODO: 

Format all loggers
add test cases
review delete use case ( right now making size to zero)
review batch process and optimize if necessary
Currently it supports base and subfolders monitoring
Error handling - error tracking 



Notes:
1. identify the polling library,
2. /config/config.json
    {
        "sourceDir":"/tmp",
        "reportDir":"/tmp/out",
        "maxGoroutines": 2
    }

    --
    gather 1: 
    {
        {"file":size},
        
    }

    {
        {
        "key:"filename,
         size:int
        }
    }

    new file, update file
    determine the state - system time , updated ts vs created ts  maintain previous run time , 
    determine previously created value either overwrite or update 
    meta-data : last run time , last status.

   libraries :
   https://github.com/fsnotify/fsnotify/blob/main/README.md

   log library :
   TODO://


   loadConfig() err

   init()

   updateStatus() err

   getSize()int


ref material/links :

ref:
_ttps://go.dev/blog/pipelines
_ttps://stackoverflow.com/questions/52512915/how-to-solve-concurrency-access-of-golang-map
_ttps://stackoverflow.com/questions/25306073/always-have-x-number-of-goroutines-running-at-any-time
_ttps://www.pixelstech.net/article/1677371161-What-is-the-use-of-empty-struct-in-GoLang




   
