package admin

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/do"
	"blog/internal/model/entity"
	"blog/internal/utility"
	"blog/utility/ujwt"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

var jwtKey = ujwt.JwtKey

type AdminClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

// Create 创建管理员
func Create(ctx context.Context, in *model.AdminInput) (err error) {
	var (
		salt     = genSalt()
		password = encryptPass(in.Password, salt)
	)
	_, err = dao.Admin.Ctx(ctx).Data(do.Admin{
		Username:  in.Username,
		Password:  password,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Salt:      salt,
		LastLogin: gtime.Now(),
		Register:  gtime.Now(),
	}).Insert()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// ValidPass 校验密码
func ValidPass(pass string, admin *entity.Admin) bool {
	return admin.Password == encryptPass(pass, admin.Salt)
}

// GenSalt 生成32位盐值盐值
func genSalt() string {
	return grand.S(32, false)
}

// EncryptPass 加密密码
func encryptPass(pass string, salt string) (encrypt string) {
	encrypt, _ = gmd5.EncryptString(pass + salt)
	return
}

// Login 登录
func Login(ctx context.Context, in *model.LoginInput) (tokenString string, err error) {
	adminOne := dao.Admin.GetAdmin(in.Username)
	// 校验账号和密码
	if adminOne.Id != 0 && ValidPass(in.Password, adminOne) {
		adminClaims := &AdminClaims{
			Id:       adminOne.Id,
			Username: adminOne.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, adminClaims)
		tokenString, _ = token.SignedString(jwtKey)
		err = nil
	} else {
		err = utility.Err.Skip(20100)
	}
	return
}

// Logout 登出
func Logout(ctx context.Context) (err error) {
	// TODO 暂时依靠客户端，后续可以考虑改为Redis黑名单处理
	return nil
}

// Info 获取当前已经登录的管理员信息

// Info 根据上下文中的令牌获取管理员信息。
// 该函数接收一个上下文参数 ctx，用于从请求头中提取令牌信息，并验证令牌。
// 如果令牌有效且属于管理员，函数将返回管理员信息；否则，将返回一个错误。
// 参数:
//
//	ctx - 上下文，用于提取请求头中的令牌信息。
//
// 返回值:
//
//	admin - 如果令牌有效且属于管理员，此参数将返回管理员信息。
//	err - 如果发生错误（例如，令牌无效或不属于管理员），此参数将返回错误信息。
func Info(ctx context.Context) (admin *entity.Admin, err error) {
	// 从请求头中提取令牌信息。
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")

	// 解析令牌并验证其合法性。
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// 检查令牌的声明和有效性。
	if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
		// 如果令牌有效且属于管理员，获取并返回管理员信息。
		admin = dao.Admin.GetAdmin(claims.Username)
	} else {
		// 如果令牌无效或不属于管理员，返回一个错误。
		err = utility.Err.Skip(10100)
	}
	return
}
