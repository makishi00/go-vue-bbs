import axios from 'axios';
class User {
    async register(email, password) {
        try {
            const response = await axios.post("/api/signup", {email: email, password: password});
            response.message = "会員登録完了しました。";
            return response;
        } catch (e) {
            return {message: "メールアドレスかパスワードが間違えています。"};
        }
    }
}

let user;
export default user = new User();