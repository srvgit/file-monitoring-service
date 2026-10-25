# File Monitoring Service

A Go application that continuously watches a directory and maintains a JSON report mapping file paths to their byte counts, with configurable concurrent processing.

## Original Interview Problem

Build a Go application with the following requirements:

1. **File monitoring:** Continuously watch a configurable target directory for newly created or modified files.
2. **Data processing:** Read each new or updated file to determine its total number of bytes.
3. **Data storage:** Maintain a JSON object whose keys are file paths and whose values are their byte counts. Persist the report to a configurable storage location.
4. **Concurrency management:** Process files concurrently, with a configurable limit on the number of files processed in parallel.
5. **Configuration:** Accept the target directory, report storage location, and concurrency level through command-line arguments or a configuration file.

### Nonfunctional Requirements

- Follow Go best practices for project structure, naming, error handling, and effective use of interfaces.
- Keep the code well-structured, readable, and maintainable.

### Example

When `/watched-folder/example.txt` contains `hello` (5 bytes, without a trailing newline), the JSON report should include:

```json
{
  "/watched-folder/example.txt": 5
}
```

If its contents change to `hello world`, update the corresponding byte count to `11`.

These are the original assignment requirements; implementation notes and outstanding work are recorded below.

Source: [Original interview problem PDF preserved in repository history](https://github.com/srvgit/file-monitoring-service/blob/5d3d92e095e0ce2e0fe36874b97bebac1a9caefa/Go%20Problem%20Statment.pdf).

---

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
https://go.dev/blog/pipelines
https://stackoverflow.com/questions/52512915/how-to-solve-concurrency-access-of-golang-map
https://stackoverflow.com/questions/25306073/always-have-x-number-of-goroutines-running-at-any-time
https://www.pixelstech.net/article/1677371161-What-is-the-use-of-empty-struct-in-GoLang




   
