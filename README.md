# Basic golang api!

Basic information for the api
- It serves on port 8000
- localhost:8000/health will let you know if youre up and running!


## Expected input!

```
{
  "id": integer,
  "first_name": string,
  "last_name": string,
  "gender": string,
  "phone_number": string,
  "email": string,
  "address": string,
  "visit_date": string,
  "diagnosis": string,
  "drug_code": string,
  "additional_information": [{
    "notes": string",
    "new_patient": boolean,
    "race": string,
    "ssn": string
  }]
}
```


## Routes!

```
Get Routes:

localhost:8000/health => checks to see if it's running

localhost:8000/patient => returns all patients

localhost:8000/patient/{patientID} => returns patient info with the given id

localhost:8000/search + querystring parameters => checks all patients to find ones that match the querystring parameter. Maximum 1 key=value pair!\



Post routes:

localhost:8000/patient => adds new patient to the database (json file)



Patch Routes: 

localhost:8000/patient/{patientID} => partial update of patient with the given id
```

## Lessons learned from this project

I learned the hard way why generics may be important in certain circumstances. It can make things a lot easier when you can you something like 

`Object.assign()` from javascript and a `for ... in loop`

Another thing I learned is that you can't turn a reflect.Value into a struct and that makes me sad.


