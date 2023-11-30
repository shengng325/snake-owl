# Quick test

New Game API
```
curl --location 'http://localhost:8808/new?w=100&h=100'
```

Validate Game API
```
curl 'http://localhost:8808/validate' \
--header 'Content-Type: application/json' \
--data '{
    "gameId": "19715f46-2b0b-45ea-8155-ce90cfbd126a",
    "width": 10,
    "height": 10,
    "score": 99,
    "fruit": {
        "x": 5,
        "y": 5
    },
    "snake": {
        "x": 3,
        "y": 7,
        "velX": 0,
        "velY": 1
    },
    "ticks": [
        {
            "velX": 0,
            "velY": -1
        },
        {
            "velX": 1,
            "velY": 0
        },
        {
            "velX": 0,
            "velY": -1
        },
        {
            "velX": 1,
            "velY": 0
        }
    ]
}'
```