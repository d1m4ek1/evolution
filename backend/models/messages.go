package models

import (
	"database/sql"
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
	Messages    sql.NullString `json:"messages"`
	NewMessages sql.NullString `json:"newMessages" db:"new_messages"`
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

func SelectChat(ctx *sqlx.DB, userIDOne int64, userIDTwo int64) (ChatData, error) {
	var chatData ChatData
	var isChatData bool

	chatIDVariantOne := fmt.Sprintf("%d%d", userIDOne, userIDTwo)
	chatIDVariantTwo := fmt.Sprintf("%d%d", userIDTwo, userIDOne)

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

	if !isChatData {
		if _, err := ctx.DB.Exec(`
		INSERT INTO
		chats (user_id_one, user_id_two, chat_id)
		VALUES 
		($1, $2, $3)`, userIDOne, userIDTwo, chatIDVariantOne); err != nil {
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
