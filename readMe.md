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


Additional features to be implemented:
1. recursive 
2. Error handling - error tracking 



   
