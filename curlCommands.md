# 🏋️‍♂️ QuickFit API: Example curl Commands

## 📝 Workout Routes

### List all workouts
```sh
curl -X GET http://localhost:3000/workouts
```

### Get a single workout by ID
```sh
curl -X GET http://localhost:3000/workouts/1
```

### Create a new workout
```sh
curl -X POST http://localhost:3000/workouts \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Evening Routine",
    "duration": 30,
    "notes": "Cardio and stretching"
  }'
```

### Update a workout by ID
```sh
curl -X PUT http://localhost:3000/workouts/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Routine",
    "duration": 40,
    "notes": "Updated notes"
  }'
```

### Delete a workout by ID
```sh
curl -X DELETE http://localhost:3000/workouts/1
```

---

## 🏃 Exercise Routes

### List all exercises
```sh
curl -X GET http://localhost:3000/exercises
```

### Get a single exercise by ID
```sh
curl -X GET http://localhost:3000/exercises/1
```

### Create a new exercise
```sh
curl -X POST http://localhost:3000/exercises \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Push Ups",
    "description": "Standard push ups",
    "repetitions": 15,
    "sets": 3
  }'
```

### Update an exercise by ID
```sh
curl -X PUT http://localhost:3000/exercises/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Push Ups",
    "description": "Updated description",
    "repetitions": 20,
    "sets": 4
  }'
```

### Delete an exercise by ID
```sh
curl -X DELETE http://localhost:3000/exercises/1
```
