package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/you/myapp-backend/domain"
	"github.com/you/myapp-backend/infra"
	"github.com/you/myapp-backend/models"
)

// ステップ1：インターフェースを作る（設計図）
// まずは「何ができるか」だけ宣言します。

type UserRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

// 実装用の構造体
// db フィールドでデータベース接続を保持しています
// 小文字で始まるので パッケージ外からは見えません（非公開）

type userRepositoryImpl struct {
	db *infra.DB
}

// func domainToModel(user *domain.User) *models.User {
// 	return &models.User{
// 		ID:        user.ID,
// 		SubID:     user.SubID,
// 		Name:      user.Name,
// 		AvatarURL: user.AvatarURL,
// 	}
// }

func modelToDomain(user *models.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		SubID:     user.SubID,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}
}

// 3. コンストラクタ関数
// 引数: 具体的な *infra.DB を受け取る
// 戻り値: インターフェース型 UserRepository を返す
// 実際の動作: 内部で userRepositoryImpl を作成して返す
func NewUserRepository(db *infra.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user models.User
	if err := r.db.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return modelToDomain(&user), nil
}
