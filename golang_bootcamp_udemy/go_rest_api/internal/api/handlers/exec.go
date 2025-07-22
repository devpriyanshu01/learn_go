package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
	"strings"
	"time"

	"github.com/go-mail/mail/v2"
	"golang.org/x/crypto/argon2"
)

func GetExecsHandler(w http.ResponseWriter, r *http.Request) {

	var execs []models.Exec
	execs, err := sqlconnect.GetExecsDbHandler(execs, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(execs),
		Data:   execs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetOneExecHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	exec, err := sqlconnect.GetExecByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

func AddExecsHandler(w http.ResponseWriter, r *http.Request) {

	var newExecs []models.Exec
	var rawExecs []map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &rawExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Get JSON tag
	fields := GetFieldNames(models.Exec{})

	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, exec := range rawExecs {
		for key := range exec {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found in request. Only use allowed fields.", http.StatusBadRequest)
				return
			}
		}
	}

	err = json.Unmarshal(body, &newExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for _, exec := range newExecs {
		err := CheckBlankFields(exec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecs, err := sqlconnect.AddExecsDBHandler(newExecs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecs),
		Data:   addedExecs,
	}
	json.NewEncoder(w).Encode(response)

}

// PATCH /execs
func PatchExecsHandler(w http.ResponseWriter, r *http.Request) {

	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecs(updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// PATCH /execs/{id}
func PatchOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Exec Id", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	// Apply updates using reflect
	updatedExec, err := sqlconnect.PatchOneExec(id, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExec)

}

func DeleteOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Exec Id", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneExec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// --- Alternate approach
	// w.WriteHeader(http.StatusNoContent)

	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Exec successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Exec
	// Data Validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Search for user if user actually exists
	user, err := sqlconnect.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	// is user active
	if user.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
		return
	}

	// verify password
	err = utils.VerifyPassword(req.Password, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	// Generate JWT Token
	tokenString, err := utils.SignToken(user.ID, req.Username, user.Role)
	if err != nil {
		http.Error(w, "Could not create login token", http.StatusInternalServerError)
		return
	}

	// Send token as a response or as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "testabc",
		Value:    "testing1231",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}
	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Logged out succesfully"}`))
}

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid exec ID", http.StatusBadRequest)
		return
	}

	var req models.UpdatePasswordRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	if req.CurrentPassword == "" || req.NewPassword == "" {
		http.Error(w, "Please enter password", http.StatusBadRequest)
		return
	}

	_, err = sqlconnect.UpdatePasswordInDb(userId, req.CurrentPassword, req.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // Send token as a response or as a cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "Bearer",
	// 	Value:    token,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	Expires:  time.Now().Add(24 * time.Hour),
	// 	SameSite: http.SameSiteStrictMode,
	// })

	// Response Body
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Password updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	err = sqlconnect.ForgotPasswordDbHandler(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// respond with success confirmation
	fmt.Fprintf(w, "Password reset link sent to %s", req.Email)
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("resetcode")

	type request struct {
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid values in request", http.StatusBadRequest)
		return
	}
	// TODO: Data Validation for blank values

	if req.NewPassword != req.ConfirmPassword {
		http.Error(w, "Passwords should match", http.StatusBadRequest)
		return
	}

	// Hash the new password
	err = sqlconnect.ResetPasswordDbHandler(token, req.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Password reset successfully")
}

// my login handler 2
func LoginHandler2(w http.ResponseWriter, r *http.Request) {
	//data validation
	var req models.Exec
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse the body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and Password Required", http.StatusBadRequest)
		return
	}
	//search if the user exists in database or not
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Connection to database failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var currUser models.Exec
	query := "SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?"
	err = db.QueryRow(query, req.Username).Scan(&currUser.ID, &currUser.FirstName, &currUser.LastName, &currUser.Email, &currUser.Username, &currUser.Password, &currUser.InactiveStatus, &currUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No user found", http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to fetch the user", http.StatusInternalServerError)
		return
	}

	//check if user is active
	if currUser.InactiveStatus {
		http.Error(w, "User is inactive", http.StatusForbidden)
		return
	}

	//validate the password
	passwordParts := strings.Split(currUser.Password, ".")
	if len(passwordParts) != 2 {
		http.Error(w, "Invalid Encoded Hash Format", http.StatusForbidden)
		return
	}

	base64Salt := passwordParts[0]
	base64Hash := passwordParts[1]

	decodedSalt, err := base64.StdEncoding.DecodeString(base64Salt)
	if err != nil {
		http.Error(w, "Error decoding the salt", http.StatusInternalServerError)
		return
	}

	decodedHash, err := base64.StdEncoding.DecodeString(base64Hash)
	if err != nil {
		http.Error(w, "Error decoding the hash", http.StatusInternalServerError)
		return
	}

	hash := argon2.IDKey([]byte(req.Password), decodedSalt, 1, 64*1024, 4, 32)

	if len(hash) != len(decodedHash) {
		http.Error(w, "Invalid Password", http.StatusForbidden)
		return
	}

	if subtle.ConstantTimeCompare(hash, decodedHash) == 1 {
		//do nothing
	} else {
		http.Error(w, "Incorrect Password", http.StatusForbidden)
		return
	}

	//generate token
	tokenString, err := utils.SignToken(currUser.ID, currUser.Username, currUser.Role)
	if err != nil {
		http.Error(w, "failed to generate jwt token", http.StatusInternalServerError)
		return
	}

	//send token as a response or a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "testabc",
		Value:    "testing123",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}
	json.NewEncoder(w).Encode(response)
}

func LogoutHandler2(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteStrictMode,
	})

	w.Write([]byte("Logout successful"))
}

func UpdatePassword2(w http.ResponseWriter, r *http.Request) {
	var req models.UpdatePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.CurrentPassword == "" || req.NewPassword == "" {
		http.Error(w, "CurrPassword and NewPassword can't be empty", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "failed to parse id", http.StatusInternalServerError)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "failed connection with database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var username, password, role string
	query := "SELECT username, password, role FROM execs WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&username, &password, &role)
	if err != nil {
		http.Error(w, "failed to fetch curr user", http.StatusInternalServerError)
		return
	}

	fmt.Println("username", username)
	fmt.Println("password", password)

	saltAndHash := strings.Split(password, ".")
	saltBase64Encoded := saltAndHash[0]
	hashBase64Encoded := saltAndHash[1]

	if len(saltAndHash) != 2 {
		http.Error(w, "Invalid Password String", http.StatusInternalServerError)
		return
	}

	//base64 to hash value
	hash, err := base64.StdEncoding.DecodeString(hashBase64Encoded)
	if err != nil {
		http.Error(w, "failed to decode hash", http.StatusInternalServerError)
		return
	}

	//encoded salt to salt
	salt, err := base64.StdEncoding.DecodeString(saltBase64Encoded)
	if err != nil {
		http.Error(w, "failed to decode salt", http.StatusInternalServerError)
		return
	}

	//convert sent_password to hash and compare
	hashPassword := argon2.IDKey([]byte(req.CurrentPassword), salt, 1, 64*1024, 4, 32)

	if len(hashPassword) != len(hash) {
		http.Error(w, "Wrong Password Entered", http.StatusUnauthorized)
		return
	}

	if subtle.ConstantTimeCompare(hash, hashPassword) == 1 {
		//do nothing
	} else {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	//Now we can store new_password to database
	newPassword := req.NewPassword
	newPasswordHashed := argon2.IDKey([]byte(newPassword), salt, 1, 64*1024, 4, 32)

	newSaltEncoded := base64.StdEncoding.EncodeToString(salt)
	newPasswordHashedEncoded := base64.StdEncoding.EncodeToString(newPasswordHashed)

	finalNewPassword := fmt.Sprintf("%s.%s", newSaltEncoded, newPasswordHashedEncoded)

	updateQuery := "UPDATE execs SET password = ? WHERE id = ?"
	_, err = db.Exec(updateQuery, finalNewPassword, id)
	if err != nil {
		http.Error(w, "failed to save new password to db", http.StatusInternalServerError)
		return
	}

	token, err := utils.SignToken(id, username, role)
	if err != nil {
		http.Error(w, "failed to generate token after saving new password", http.StatusInternalServerError)
		return
	}

	//send token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})

	w.Write([]byte("Password Updated Successfully"))
}

func ForgotPasswordHandler2(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Incorrect email sent", http.StatusBadRequest)
		return
	}
	log.Println("received email", req.Email)

	defer r.Body.Close()

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "failed to connect to db", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var exec models.Exec
	err = db.QueryRow("SELECT id FROM execs WHERE email = ?", req.Email).Scan(&exec.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "This email is not registered", http.StatusForbidden)
			return
		}
		http.Error(w, "Faild to retrieve details with provided email", http.StatusInternalServerError)
		return
	}

	duration, err := strconv.Atoi(os.Getenv("RESET_TOKEN_EXP_DURATION"))
	if err != nil {
		http.Error(w, "error retrieving duration from .env", http.StatusInternalServerError)
		return
	}

	mins := time.Duration(duration)
	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		http.Error(w, "error reading bytes", http.StatusInternalServerError)
		return
	}

	log.Println("token bytes:", string(tokenBytes))
	//randBytes to string token conversion
	token := hex.EncodeToString(tokenBytes)
	log.Println("token:", token)

	//randBytes to hash conversion
	hashToken := sha256.Sum256(tokenBytes)
	log.Println("hashToken:", hashToken)

	//hash to string conversion
	hashedTokenString := hex.EncodeToString(hashToken[:])

	_, err = db.Exec("UPDATE execs SET password_reset_token = ?, password_token_expires = ? WHERE id = ?", hashedTokenString, expiry, exec.ID)
	if err != nil {
		http.Error(w, "failed to update fields in db", http.StatusInternalServerError)
		return
	}

	//send the rest email
	resetUrl := fmt.Sprintf("https://localhost:3000/execs/resetpassword/reset2/%s", token)
	message := fmt.Sprintf("Forgot your password? Reset your password using the following link: \n%s\nIf you didn't request a password reset, please ignore this email. This link is only valid for %d mintues.", resetUrl, int(mins))

	m := mail.NewMessage()
	m.SetHeader("From", "schooladmin@school.com")
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", "Your password reset link")
	m.SetBody("text/plain", message)
	d := mail.NewDialer("localhost", 1025, "", "")

	err = d.DialAndSend(m)

	if err != nil {
		http.Error(w, "failed to send password reset email", http.StatusInternalServerError)
		log.Println("error sending email --------->", err)
		return
	}

	fmt.Fprintf(w, "Password reset link to %s", req.Email)
}

func ResetPasswordHandler2(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("resetcode")
	log.Println("resetcode :----", token)

	type request struct {
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to pasrse the body", http.StatusBadRequest)
		log.Println("error parsing the body :----", err)
		return
	}

	if req.NewPassword == "" || req.ConfirmPassword == "" {
		http.Error(w, "new_password or confirm_password can't be empty", http.StatusBadRequest)
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		http.Error(w, "new_password and confirm_password should match", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user models.Exec
	bytes, err := hex.DecodeString(token)
	if err != nil {
		http.Error(w, "failed decoding sent token", http.StatusInternalServerError)
		return
	}
	log.Println("decoded token string:-----", string(bytes))

	//hash the above []byte
	hashedToken := sha256.Sum256(bytes)
	hashedTokenString := hex.EncodeToString(hashedToken[:])

	query := "SELECT id, email FROM execs WHERE password_reset_token = ? AND password_token_expires > ?"

	err = db.QueryRow(query, hashedTokenString, time.Now().Format(time.RFC3339)).Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "no rows for given query", http.StatusInternalServerError)
			return
		}
		http.Error(w, "failed to receive id and email from db", http.StatusInternalServerError)
		log.Println("error getting id, email from db :------", err)
		return
	}

	//Hash the new_password
	hashedNewPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		http.Error(w, "failed to hash new_password", http.StatusInternalServerError)
		return
	}

	updateQuery := "UPDATE execs SET password = ?, password_reset_token = NULL, password_token_expires = NULL, password_changed_at = ? WHERE id = ?"

	_, err = db.Exec(updateQuery, hashedNewPassword, time.Now().Format(time.RFC3339), user.ID)
	if err != nil {
		http.Error(w, "failed to update password", http.StatusInternalServerError)
		log.Println("error updating password :------", err)
		return
	}

	fmt.Fprintf(w, "Password reset successfully")
}
