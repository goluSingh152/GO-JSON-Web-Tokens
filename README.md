# GO-JSON-Web-Tokens
This is demo program to show the working of the go-jwt tool, so we can use this library for authorization in out project.<br>
I used <b>"github.com/dgrijalva/jwt-go"</b> library for this project.<br>
This library <b>jwt-go</b> used for json web token methods.<br>
As we know that now a days for each user request, instead of validate the user from database side,<br> we validate the user from cache and by using token
in that case json web token is pretty good for solving those problem and its very secure to use and very fast.

<h2>Steps for Using the Library </h2>
<ol>
<li>Import the "github.com/dgrijalva/jwt-go</li>
<li>Create public key and private key using command : <b>ssh-keygen -t rsa</b></li>
<li>Read this files and store in different PrivateKey and PublicKey Using the ioutill library</li>
<li>Use Jwt.New function for create the token</li>
<li>by using the help of jwt.MapClaims, store the metadata inside the jwt token file</li>
<li>Use the files using "jwt.ParseRSAPrivateKeyFromPEM", most of the peopele do wrong here they try to use  []byte("string") which most of time throw error</li>
<li></li>
<li>By using (*token).SignedString(interface{}) generate the token response</li>
<li>By using jwt.Parse function, parse the function the code and get the metadata from the db</li>
</ol>

Thanks if any changes required please let me know.
