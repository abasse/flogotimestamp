# Timestamp flogo activity
This activity converts a date into a UNIX timestamp.
Supported input format RFC3339 only (at this stage) 



## Installation

```bash
flogo install github.com/abasse/flogotimestamp
```

## Schema
Inputs and Outputs:

```json
  { "inputs":[
        {
          "name": "inputdate",
          "type": "string",
          "required": true
        },
        {
          "name": "dateformat",
          "type": "string",
          "required": false
        }
      ],
      "outputs": [
         {
           "name": "timestamp",
           "type": "integer"
          }
      ]
 }
```
