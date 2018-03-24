import axios from 'axios';

class Auth {
    async login(email, password) {
        try {
            const response = await axios.post("/api/signin", {email: email, password: password});
            response.message = "ログインしました。";
            return response;
        } catch (e) {
            return {message: "メールアドレスかパスワードが間違えています。"};
        }
    }

    async () {

    }

    setloginToken(token) {
        localStorage.setItem('loginToken', token);
    }
}

let auth;
export default auth = new Auth();