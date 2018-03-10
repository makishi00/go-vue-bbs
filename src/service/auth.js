import axios from 'axios';

export default auth;

class Auth {
    constructor(){
        this.token = ""
    }
    async Login(email, password) {
        const hoge = await axios.post("/api/signin")
        console.log(hoge)
    }
    SetToken(t) {
        this.token = t
    }
}

var auth = new Auth()