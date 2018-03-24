import axios from 'axios';

class Auth {
    constructor() {
        this.loginToken = "loginToken";
    }

    async login(email, password) {
        try {
            const response = await axios.post("/api/signin", {email: email, password: password});
            response.message = "ログインしました。";
            return response;
        } catch (e) {
            return {message: "メールアドレスかパスワードが間違えています。"};
        }
    }

    setLoginToken(token) {
        localStorage.setItem(this.loginToken, token);
    }

    getLoginToken() {
        return localStorage.getItem(this.loginToken);
    }

    removeLoginToken() {
        localStorage.removeItem(this.loginToken);
    }
}

let auth;
export default auth = new Auth();