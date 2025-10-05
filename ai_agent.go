package model

import (
	"database/sql"
	"time"
)

type AIAgent struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type Prompt struct {
	ID        string    `json:"id"`
	AgentID   string    `json:"agent_id"` // FK → ai_agents.id
	UserID    string    `json:"user_id"`  // FK → users.id
	ErrorID   string    `json:"error_id"` // FK → errors.id (nullable)
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type AIAgentModel struct{ DB *sql.DB }

func NewAIAgentModel(db *sql.DB) *AIAgentModel {
	return &AIAgentModel{DB: db}
}

type PromptModel struct{ DB *sql.DB }

func NewPromptModel(db *sql.DB) *PromptModel {
	return &PromptModel{DB: db}
}
func (m *AIAgentModel) CreateTableIfNotExists() error {
	_, err := m.DB.Exec(`
		CREATE TABLE IF NOT EXISTS ai_agents (
			id UUID PRIMARY KEY,
			name VARCHAR(60) NOT NULL,
			active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP NOT NULL
		);
	`)
	return err
}

func (m *PromptModel) CreateTableIfNotExists() error {
	_, err := m.DB.Exec(`
		CREATE TABLE IF NOT EXISTS prompts (
			id UUID PRIMARY KEY,
			agent_id UUID NOT NULL REFERENCES ai_agents(id),
			user_id  UUID NOT NULL REFERENCES users(id),
			error_id UUID REFERENCES errors(id) ,
			content  TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_prompt_agent ON prompts(agent_id);
	`)
	return err
}
