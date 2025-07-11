package rediskey

const (
	REDIS_LIST_USER_SESSION = "user:session:lst" // key 為 user_id value為 session_id (用戶最多三個連線資格)
	REDIS_LIST_USER_AGENT   = "user:agent:lst"   // key 為 user_id value 為 agent_id (用戶最多五個代理)
	REDIS_LIST_USER_TODO    = "user:todo:lst"    // key 為 user_id value 為 todo_id (用戶最多五十個任務)
)
