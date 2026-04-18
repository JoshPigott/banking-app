### I like need to make method to a struc and then call them 
- I low key need to lean how to use interfaces

- Try and think more in terms of data

- User pointer when you to update the state
- Or when the size of stru is really big like files
- You can make stru from other stru which give you great control and readablity

const lower case and that 


When doing stru it should be 
session := models.Session{
    ID: id,
    UserID: userID,
    ExpiryTime: expiryTime.Unix()
}
And not 
session := models.Session{ID: id,
UserID: userID, ExpiryTime: expiryTime.Unix()}

Do I need to do internal processing