### I like need to make method to a struc and then call them 
- I low key need to lean how to use interfaces

- Try and think more in terms of data

- User pointer when you to update the state
- Or when the size of stru is really big like files
- You can make stru from other stru which give you great control and readablity

const lower case and that 


When doing stru it should be 
session := domain.Session{
    ID: id,
    UserID: userID,
    ExpiryTime: expiryTime.Unix()
}
And not 
session := domain.Session{ID: id,
UserID: userID, ExpiryTime: expiryTime.Unix()}

Do I need to do internal processing

If you have somthing that does the same thing like swim but does it in different ways that is where you use interfaces

Funcation groupping like this is the right order
func ImportantFuncExported () {}
func ImportnatFunc () {}
func simpleUtil() {}
func main() {}
(same thing with variables but const above varaibales)

I need to try and play around with io.Reader and io.Writer Like what si io