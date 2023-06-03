package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

type AuditLogPayload struct {
	Event     string    `json:"event"`
	User      string    `json:"user"`
	Timestamp time.Time `json:"timestamp"`
}

type AuditLogPayloadType struct {
	Event     string    `json:"event"`
	User      string    `json:"user"`
	Timestamp time.Time `json:"timestamp"`
}

func handleAuditLogs(c *gin.Context) {
	// Parse the audit log payload from the request
	var auditLogPayload AuditLogPayloadType
	if err := c.ShouldBindJSON(&auditLogPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Send the audit log message to Slack
	webhookURL := "xoxb-5079192664693-5391519133056-2X35uzRu1tgEiF2TgTuROHa8"
	sendAuditLogToSlack(auditLogPayload, webhookURL)
	c.JSON(http.StatusOK, gin.H{"message": "Audit log received"})
}

func sendAuditLogToSlack(log AuditLogPayloadType, webhookURL string) {
	slackClient := slack.New(webhookURL)
	message := formatAuditLogMessage(log)
	// Send message to Slack
	_, _, err := slackClient.PostMessage("#kubernetes-audit-logs", slack.MsgOptionText(message, false))
	if err != nil {
		panic(err)
	}
}

// Format the audit log message
func formatAuditLogMessage(log AuditLogPayloadType) string {
	return fmt.Sprintf("Audit log received:\nEvent: %s\nUser: %s\nTimestamp: %s", log.Event, log.User, log.Timestamp)
}

func main() {
	router := gin.Default()
	router.POST("/auditLogs", handleAuditLogs)
	router.Run(":8080")
}
