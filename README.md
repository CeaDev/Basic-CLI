Example of a task saved in the JSON file:
  {
    "id": 1,
    "description": "Wash the dishes",
    "isDone": false
  }

Task Attributes:
  - ID: int value that is automatically assigned when the task is created. This will increment every time a new task is made
  - Description: string value that describes the actual task itself
  - IsDone: boolean value that indicates whether or not a task is completed

Commands:
- list:
  - Command Example: "go run main.go list" 
  - Purpose: prints all of the tasks that are currently saved in the JSON file.
- add:
  - Command Example: "go run main.go add wash the dishes" 
  - Purpose: this command will take all of the arguments that come after the second argument in os.Args and join them together into one string. This string will represent the task itself.
- done:
  - Command Example: "go run main.go done 3"
  - The third argument represents a task ID. This will set the IsDone property of that specific task to TRUE.
  
