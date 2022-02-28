# counter
- This project creates seriall unique IDs
- IDs are encoded with three parts
    - Timestamp
    - WorkerID
    - Counter
- Multiple PODs can be spun up and they will generate unique IDs across applicaton (due to WorkerID encoding)
- TODO: Encode goroutine id as 4th part.