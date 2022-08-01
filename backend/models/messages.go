package models

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	newerror "iNote/www/backend/pkg/NewError"
	"strings"
)

type UserCardMessagesItems struct {
	UserId    int64  `json:"userId" db:"user_id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Banner    string `json:"banner"`
	NetStatus string `json:"netStatus"`
}

type ChatData struct {
	Id          int64          `json:"id"`
	ChatId      int64          `json:"chatId" db:"chat_id"`
	UserIDOne   int64          `json:"userIDOne" db:"user_id_one"`
	UserIDTwo   int64          `json:"userIDTwo" db:"user_id_two"`
	Messages    pq.StringArray `json:"messages"`
	NewMessages pq.StringArray `json:"newMessages" db:"new_messages"`
	UserDataOne string         `json:"userDataOne" db:"user_data_one"`
	UserDataTwo string         `json:"userDataTwo" db:"user_data_two"`
}

type CompanionData struct {
	UserId    int64  `json:"userId" db:"user_id"`
	Logo      string `json:"logo"`
	Banner    string `json:"banner"`
	Name      string `json:"name"`
	NetStatus string `json:"netStatus" db:"net_status"`
}

func generateTplSubs(ids []int64) (tpl []string) {
	for _, id := range ids {
		tpl = append(tpl, fmt.Sprintf(`u.user_id=%d AND s.settings_id=(SELECT settings_id FROM identifiers ids WHERE ids.user_id=%d)`, id, id))
	}

	return tpl
}

func selectUserCardSubs(ctx *sqlx.DB, queryTpl string) ([]UserCardMessagesItems, error) {
	var isCardItems []UserCardMessagesItems
	rows, err := ctx.Query(fmt.Sprintf(`
	SELECT 
	    u.user_id,
	    u.name,
			u.net_status,
	    s.logo,
	    s.banner
	FROM
	    users u,
	    settings s
	WHERE %s`, queryTpl))
	if err != nil {
		newerror.Wrap("SelectUserCardSubs -> ctx.Query", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i UserCardMessagesItems
		if err := rows.Scan(&i.UserId, &i.Name, &i.NetStatus, &i.Logo, &i.Banner); err != nil {
			newerror.Wrap("rows.Scan", err)
			return nil, err
		}

		isCardItems = append(isCardItems, i)
	}

	return isCardItems, nil
}

func SelectUserCardMessages(ctx *sqlx.DB, userID int64) ([]UserCardMessagesItems, []UserCardMessagesItems, error) {
	var isSubscriptions, isSubscribers []int64
	var isTplQuerySubscriptions, isTplQuerySubscribers string
	var isCardSubscriptions, isCardSubscribers []UserCardMessagesItems
	var err error

	if err := ctx.DB.QueryRow(`
	SELECT
	    subscriptions,
	    subscribers
	FROM
	  	users
	WHERE
	    user_id=$1`, userID).Scan(pq.Array(&isSubscriptions), pq.Array(&isSubscribers)); err != nil {
		newerror.Wrap("ctx.DB.QueryRow", err)
		return nil, nil, err
	}

	if len(isSubscriptions) != 0 {
		isTplQuerySubscriptions = strings.Join(generateTplSubs(isSubscriptions), " OR ")
		isCardSubscriptions, err = selectUserCardSubs(ctx, isTplQuerySubscriptions)
		if err != nil {
			newerror.Wrap("isCardSubscriptions -> selectUserCardSubs", err)
			return nil, nil, err
		}
	}

	if len(isSubscribers) != 0 {
		isTplQuerySubscribers = strings.Join(generateTplSubs(isSubscribers), " OR ")
		isCardSubscribers, err = selectUserCardSubs(ctx, isTplQuerySubscribers)
		if err != nil {
			newerror.Wrap("isCardSubscribers -> selectUserCardSubs", err)
			return nil, nil, err
		}
	}

	return isCardSubscriptions, isCardSubscribers, nil
}

func selectChatByID(ctx *sqlx.DB, chatID int64) (ChatData, error) {
	var chatData ChatData
	if err := ctx.Get(&chatData, `
		SELECT
				*
		FROM
				chats
		WHERE
			chat_id=$1`, chatID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return ChatData{}, err
	}
	return chatData, nil
}

func getUserDataChat(ctx *sqlx.DB, userIDOne, userIDTwo int64) (string, string, error) {
	var companionOne, companionTwo CompanionData
	var jsonCompanionOne, jsonCompanionTwo []byte
	var err error

	if err := ctx.Get(&companionOne, `
	SELECT
	    u.user_id,
	    u.name,
	    s.logo,
	    s.banner,
	    u.net_status
	FROM
	    users AS u,
	    settings AS s
	WHERE
	    u.user_id=$1 AND s.settings_id=(SELECT settings_id FROM identifiers WHERE user_id=$2)`, userIDOne, userIDOne); err != nil {
		newerror.Wrap("ctx.Get", err)
		return "", "", err
	}

	if err := ctx.Get(&companionTwo, `
	SELECT
	    u.user_id,
	    u.name,
	    s.logo,
	    s.banner,
	    u.net_status
	FROM
	    users AS u,
	    settings AS s
	WHERE
	    u.user_id=$1 AND s.settings_id=(SELECT settings_id FROM identifiers WHERE user_id=$2)`, userIDTwo, userIDTwo); err != nil {
		newerror.Wrap("ctx.Get", err)
		return "", "", err
	}

	jsonCompanionOne, err = json.Marshal(companionOne)
	if err != nil {
		newerror.Wrap("json.Marshal", err)
		return "", "", err
	}

	jsonCompanionTwo, err = json.Marshal(companionTwo)
	if err != nil {
		newerror.Wrap("json.Marshal", err)
		return "", "", err
	}

	return string(jsonCompanionOne), string(jsonCompanionTwo), nil
}

func SelectChat(ctx *sqlx.DB, userIDOne, userIDTwo, chatID int64) (ChatData, error) {
	var chatData ChatData
	var isChatData bool
	var chatIDVariantOne, chatIDVariantTwo string

	if chatID != 0 {
		chatData, err := selectChatByID(ctx, chatID)
		if err != nil {
			newerror.Wrap("selectChatByID", err)
			return ChatData{}, err
		}
		return chatData, nil
	}

	chatIDVariantOne = fmt.Sprintf("%d%d", userIDOne, userIDTwo)
	chatIDVariantTwo = fmt.Sprintf("%d%d", userIDTwo, userIDOne)

	if err := ctx.Get(&isChatData, `
	SELECT
	    count(*) = 1
	FROM
			chats
	WHERE
    chat_id=$1 OR chat_id=$2`, chatIDVariantOne, chatIDVariantTwo); err != nil {
		newerror.Wrap("ctx.Get", err)
		return ChatData{}, err
	}

	userDataOne, userDataTwo, err := getUserDataChat(ctx, userIDOne, userIDTwo)
	if err != nil {
		newerror.Wrap("getUserDataChat", err)
		return ChatData{}, err
	}

	if !isChatData {
		if _, err := ctx.DB.Exec(`
		INSERT INTO
		chats (user_id_one, user_id_two, chat_id, user_data_one, user_data_two)
		VALUES 
		($1, $2, $3, $4, $5)`, userIDOne, userIDTwo, chatIDVariantOne, userDataOne, userDataTwo); err != nil {
			newerror.Wrap("ctx.DB.Exec", err)
			return ChatData{}, err
		}
	}

	if err := ctx.Get(&chatData, `
		SELECT
				*
		FROM
				chats
		WHERE
			chat_id=$1 OR chat_id=$2`, chatIDVariantOne, chatIDVariantTwo); err != nil {
		newerror.Wrap("ctx.Get", err)
		return ChatData{}, err
	}

	return chatData, nil
}

func SelectChatItems(ctx *sqlx.DB, userID int64) ([]ChatData, error) {
	var chatDataItems []ChatData

	if err := ctx.Select(&chatDataItems, `
	SELECT
	    *
	FROM
	    chats
	WHERE
	    user_id_one=$1 OR user_id_two=$2`, userID, userID); err != nil {
		newerror.Wrap("ctx.Select", err)
		return nil, err
	}

	return chatDataItems, nil
}

func SetNewMessage(ctx *sqlx.DB, chatID string, message string) error {
	if _, err := ctx.DB.Exec(`
	UPDATE
				chats
	SET
	    new_messages=array_append(new_messages, $1)
	WHERE
	    chat_id=$2`, message, chatID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func SetMessage(ctx *sqlx.DB, chatID string, message string) error {
	if _, err := ctx.DB.Exec(fmt.Sprintf(`
	UPDATE
				chats
	SET
	    messages=messages || '{%s}',
			new_messages=null
	WHERE
	    chat_id=$1`, message), chatID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}
