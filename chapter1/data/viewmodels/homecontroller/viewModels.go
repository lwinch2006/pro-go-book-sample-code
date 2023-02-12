package homecontroller

import "partyinvites/data/models"

type ReplyViewModel struct {
	Reply  *models.Rvsp
	Errors []string
}
