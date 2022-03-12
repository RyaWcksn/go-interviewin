package logic

func Register(name, email, password string) map[string]interface{} {
    if len(password) < 4 {
        return map[string]interface{}{"status": "error", "message": "Password must be at least 4 characters long"}
    } else if len(name) < 2{
       return map[string]interface{}{"status": "error", "message": "Name must be at least 2 characters long"}
    }

    return map[string]interface{}{"status": "ok"}
}

func Login(email string, password string) bool {
    return true
}
