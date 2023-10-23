# receipt-processor-challenge

For this receipt processor challenge I used GoLang (had never used this language before) with no other other languages.

Since this challenge was done in GoLang and the engineer evaluating the submission has an environment for GoLang setup already, the use of this program is simple.  The server is set to run on localhost:8080, but could be changed by editing the string with the desired input on line 20 in receiptProcessor.go.  To run, open a terminal in the receipt-processor-challenge folder and type in "go run ."

For testing this I was using Postman with headers Key:"Content-Type" Value:"application/json".  The urls I have been using are http://localhost:8080/receipts/process and http://localhost:8080/receipts/{id}/points (where id is the receipt id desired).  The body of the request is input as a raw JSON type of the format specified in the challenge:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

The date and time must be in the same format as above (yyyy/mm/dd and hh/mm in 24 hour time)
Another assumption is that the score is +10 if the time is after 2:00 PM (including 2:00 PM) and before 4:00PM (not including 4:00 PM)

To stop the server, in terminal do CTRL + c (or on mac: command c)

There are many ways to accomplish the goal of this challenge and this is just one I can think of.  I know there are probably other ways to verify the format of date and time such as potentially creating a custom unmarshaller, but for this challenge a simple check seemed to work perfectly fine.  This would replace lines 31 - 46 with a custom unmarshaller.  I am sure with more time I could find a better solution, but this works and isnt terribly complex.  If I knew that this would be implemented with a large database (still using memory for some reason, otherwise the database would likely take care of this) I would potentially organize the receipts in the array by their receipt id to make searching quicker.  This would depend on if we cared more about the processing receipts or the finding of receipt speed.
