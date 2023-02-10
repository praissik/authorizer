package auth

import (
	"authorizer/pkg/account"
	errors "authorizer/pkg/error"
	"fmt"
	"strings"
)

func Register(email, password string) (string, error) {
	if err := validEmail(email); err != nil {
		return "", err
	}
	if err := validPassword(password); err != nil {
		return "", err
	}

	exists, err := account.IsEmailExists(email)
	if err != nil {
		return "", fmt.Errorf(errors.SomethingWentWrong)
	}
	if exists {
		return "", fmt.Errorf(errors.BusyEmail, email)
	}

	//userEntity.Email = form.Email
	//userEntity.Password = security.GenerateBcryptHash(form.Password)
	//ormService.Flush(userEntity)

	//registerRequestEntity := &entity.RegisterRequestEntity{
	//	UserID:        userEntity,
	//	Token:         generateSHA256TokenForUser(userEntity.ID),
	//	GeneratedDate: time.Now().Unix(),
	//	ExpiresDate:   time.Now().Add(10 * time.Minute).Unix(),
	//}
	//ormService.Flush(registerRequestEntity)

	//emailData := email.Data{
	//	Address:     userEntity.Email,
	//	Token:       registerRequestEntity.Token,
	//	ExpiresDate: registerRequestEntity.ExpiresDate,
	//}
	//err = email.SendEmail(emailData)
	//if err != nil {
	//	return "", err
	//}

	//return security.NewToken(userEntity.Email)
	return "", nil
}

//func (form *Form) ConfirmEmail(in *pb.AuthRequest) error {
//ormService := c.MustGet(service.OrmContextService).(*beeorm.Engine)
//registerRequestEntity := &entity.RegisterRequestEntity{}
//found := ormService.SearchOne(beeorm.NewWhere("Token = ?", c.Query("token")), registerRequestEntity)
//if !found {
//	return fmt.Errorf(errors.RegisterTokenNotExist)
//}
//if registerRequestEntity.UsedDate > 0 {
//	return fmt.Errorf(errors.RegisterTokenUsedUp)
//}
//t := time.Now().Unix()
//registerRequestEntity.UsedDate = t
//if registerRequestEntity.ExpiresDate > t {
//	userEntity := &entity.UserEntity{
//		ID: registerRequestEntity.UserID.ID,
//	}
//	ormService.Load(userEntity)
//	userEntity.Confirmed = true
//	ormService.FlushMany(registerRequestEntity, userEntity)
//	return nil
//} else {
//	ormService.Flush(registerRequestEntity)
//	return fmt.Errorf(errors.RegisterTokenExpired)
//}
//	return nil
//}

//func (form *Form) ResendConfirmEmail(in *pb.AuthRequest) error {
//ormService := c.MustGet(service.OrmContextService).(*beeorm.Engine)
//registerRequestEntity := &entity.RegisterRequestEntity{}
//found := ormService.SearchOne(beeorm.NewWhere("Token = ?", c.Query("token")), registerRequestEntity, "UserID")
//if !found {
//	return fmt.Errorf(errors.RegisterTokenNotExist)
//}
//newRegisterRequestEntity := &entity.RegisterRequestEntity{
//	UserID:        registerRequestEntity.UserID,
//	Token:         generateSHA256TokenForUser(registerRequestEntity.UserID.ID),
//	GeneratedDate: time.Now().Unix(),
//	ExpiresDate:   time.Now().Add(10 * time.Minute).Unix(),
//}
//ormService.Flush(newRegisterRequestEntity)
//emailData := email.Data{
//	Address:     registerRequestEntity.UserID.Email,
//	Token:       newRegisterRequestEntity.Token,
//	ExpiresDate: newRegisterRequestEntity.ExpiresDate,
//}
//return email.SendEmail(emailData)
//	return nil
//}

//func (form *Form) Login(in *pb.AuthRequest) (string, error) {
//session := sessions.Default(c)
//err := form.validLoginForm()
//if err != nil {
//	return "", fmt.Errorf(errors.RegisterTokenNotExist)
//}
//
//userEntity := user.GetUserByEmail(c, form.Email)
//
//if userEntity == nil {
//	return "", &errors.UnauthorizedError{Message: errors.InvalidEmailOrPassword}
//}
//isCompare := security.CompareHashAndText(userEntity.Password, form.Password)
//if !isCompare {
//	return "", &errors.UnauthorizedError{Message: errors.InvalidEmailOrPassword}
//}
//return security.NewToken(userEntity.Email)
//	return "", nil
//}

//func (form Form) Logout() {
//	log.Println("Logout")
//	log.Println(form)
//	session := sessions.Default(c)
//	user := session.Get(userkey)
//	if user == nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
//		return
//	}
//	session.Delete(userkey)
//	if err := session.Save(); err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
//}

//func (form *Form) Valid(in *pb.AuthRequest) (string, error) {
//auth := c.Request.Header.Get("Authorization")
//log.Println(auth)
//auth2 := c.Request.Header
//log.Println(auth2)
//if auth == "" {
//	return "", &errors.UnauthorizedError{Message: "No Authorization header provided"}
//}
//token := strings.TrimPrefix(auth, "Bearer ")
//if token == auth {
//	return "", &errors.UnauthorizedError{Message: "Could not find bearer token in Authorization header"}
//}
//if err := security.ValidToken(token); err != nil {
//	return "", err
//}
//	return "", nil
//}

func validEmail(email string) error {
	if email == "" {
		return fmt.Errorf(errors.EmptyEmail)
	}
	if !strings.ContainsAny(email, "@ & .") {
		return fmt.Errorf(errors.InvalidEmail, email)
	}
	return nil
}

func validPassword(password string) error {
	if password == "" {
		return fmt.Errorf(errors.EmptyPassword)
	}
	if strings.ContainsAny(password, " ") {
		return fmt.Errorf(errors.InvalidPassword)
	}
	if len(password) < 8 {
		return fmt.Errorf(errors.ToShortPassword)
	}
	return nil
}

//func generateSHA256TokenForUser(userID uint64) string {
//	h := sha256.New()
//	h.Write([]byte(time.Now().String() + strconv.FormatUint(userID, 10)))
//	return fmt.Sprintf("%x", h.Sum(nil))
//}
